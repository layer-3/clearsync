package userop

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point"
	"github.com/layer-3/clearsync/pkg/abi/simple_account/account_abstraction"
)

// UserOperationClient represents a client for creating and posting user operations.
type UserOperationClient interface {
	NewUserOp(ctx context.Context, sender common.Address, calls []Call) (UserOperation, error)
	SendUserOp(ctx context.Context, op UserOperation) (<-chan struct{}, error)
}

type Call struct {
	To       common.Address
	Value    decimal.Decimal // Value is a wei amount to be sent with the call.
	CallData []byte
}

// client represents a user operation client.
type client struct {
	providerRPC *ethclient.Client
	bundlerRPC  *rpc.Client
	chainID     *big.Int

	smartWalletType    SmartWalletType
	entryPoint         common.Address
	isPaymasterEnabled bool
	paymaster          common.Address
	signer             Signer
	middlewares        []middleware
}

// NewClient is a factory that builds a new
// user operation client based on the provided configuration.
func NewClient(config ClientConfig) (UserOperationClient, error) {
	providerRPC, err := ethclient.Dial(config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the blockchain RPC: %w", err)
	}

	bundlerRPC, err := rpc.Dial(config.BundlerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the bundler RPC: %w", err)
	}

	entryPointContract, err := entry_point.NewEntryPoint(config.EntryPoint, providerRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the entry point contract: %w", err)
	}

	isPaymasterEnabled := config.Paymaster.URL != "" && config.Paymaster.Address != common.Address{}
	estimateGas := estimateUserOperationGas(bundlerRPC, config.EntryPoint)
	if isPaymasterEnabled {
		switch typ := config.Paymaster.Type; typ {
		case PaymasterPimlicoERC20:
			estimateGas = getPimlicoERC20PaymasterData(bundlerRPC, config.EntryPoint, config.Paymaster.Address)
		case PaymasterPimlicoVerifying:
		case PaymasterBiconomyERC20:
		case PaymasterBiconomySponsoring:
		default:
			panic(fmt.Errorf("unknown paymaster type: %s", typ))
		}
	}

	var getInitCode middleware
	switch typ := config.SmartWallet.Type; typ {
	case SmartWalletSimpleAccount:
	case SmartWalletBiconomy:
	case SmartWalletKernel:
		getInitCode = getKernelInitCode(
			providerRPC,
			decimal.Zero,
			config.SmartWallet.Factory,
			config.SmartWallet.Logic,
			config.SmartWallet.ECDSAValidator,
			config.SmartWallet.Owner,
		)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", typ)
	}

	return &client{
		providerRPC:        providerRPC,
		bundlerRPC:         bundlerRPC,
		chainID:            config.ChainID,
		smartWalletType:    config.SmartWallet.Type,
		entryPoint:         config.EntryPoint,
		isPaymasterEnabled: isPaymasterEnabled,
		paymaster:          config.Paymaster.Address,
		signer:             config.Signer,
		middlewares: []middleware{ // Middleware order matters - first in, first executed
			getNonce(entryPointContract),
			getInitCode,
			getGasPrice(providerRPC),
			sign(config.Signer, config.EntryPoint, config.ChainID),
			estimateGas,
			sign(config.Signer, config.EntryPoint, config.ChainID), // update signature after gas estimation
		},
	}, nil
}

// NewUserOp builds and fills in a new UserOperation.
func (c *client) NewUserOp(
	ctx context.Context,
	smartWallet common.Address, // TODO: support calculating SW address from SmartWalletType
	calls []Call,
) (UserOperation, error) {
	slog.Info("apply middlewares to user operation")
	op := UserOperation{Sender: smartWallet}

	callData, err := c.buildCallData(calls)
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to build call data: %w", err)
	}
	op.CallData = callData

	for _, fn := range c.middlewares {
		if err := fn(ctx, &op); err != nil {
			return UserOperation{}, fmt.Errorf("failed to apply middleware to user operation: %w", err)
		}
	}

	return op, nil
}

// SendUserOp submits a user operation to a bundler
// and executes the provided callback function.
func (c *client) SendUserOp(ctx context.Context, op UserOperation) (<-chan struct{}, error) {
	slog.Info("sending user operation")

	var userOpHash common.Hash
	if err := c.bundlerRPC.CallContext(ctx, &userOpHash, "eth_sendUserOperation", op, c.entryPoint); err != nil {
		return nil, fmt.Errorf("call to `eth_sendUserOperation` failed: %w", err)
	}

	slog.Info("user operation sent successfully", "hash", userOpHash.Hex())
	waiter := make(chan struct{}, 1)
	go waitForUserOpEvent(c.providerRPC, waiter, c.entryPoint, userOpHash)

	return waiter, nil
}

func (c *client) buildCallData(calls []Call) ([]byte, error) {
	switch c.smartWalletType {
	case SmartWalletSimpleAccount, SmartWalletBiconomy:
		return handleCallSimpleAccount(calls)
	case SmartWalletKernel:
		return handleCallKernel(calls)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", c.smartWalletType)
	}
}

func handleCallSimpleAccount(calls []Call) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(account_abstraction.SimpleAccountMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	addresses := make([]any, 0, len(calls))
	values := make([]any, 0, len(calls))
	calldatas := make([]any, 0, len(calls))
	for _, call := range calls {
		addresses = append(addresses, call.To)
		values = append(values, call.Value.BigInt())
		calldatas = append(calldatas, call.CallData)
	}

	data, err := parsedABI.Pack("executeBatch", addresses, values, calldatas)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for SimpleAccount: %w", err)
	}
	return data, nil
}

type callStructKernel struct {
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`
	Data  []byte         `json:"data"`
}

func handleCallKernel(calls []Call) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(kernelExecuteABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	params := make([]callStructKernel, 0, len(calls))
	for _, call := range calls {
		params = append(params, callStructKernel{
			To:    call.To,
			Value: call.Value.BigInt(),
			Data:  call.CallData,
		})
	}

	data, err := parsedABI.Pack("executeBatch", params)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for Kernel: %w", err)
	}
	return data, nil
}

// waitForUserOpEvent waits for a user operation to be committed on block.
func waitForUserOpEvent(
	client *ethclient.Client,
	done chan<- struct{},
	entryPoint common.Address,
	userOpHash common.Hash,
) {
	waitTimeout := time.Millisecond * 30000
	waitInterval := time.Millisecond * 5000

	ctx, cancel := context.WithTimeout(context.Background(), waitTimeout)
	defer cancel()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{entryPoint},
		Topics:    [][]common.Hash{{}, {userOpHash}},
	}

	ticker := time.NewTicker(waitInterval)
	defer ticker.Stop()
	defer close(done)

	for {
		select {
		case <-ctx.Done():
			slog.Error("timeout waiting for user operation event", "hash", userOpHash.Hex())
			return
		case <-ticker.C:
			logs, err := client.FilterLogs(ctx, query)
			if err != nil {
				slog.Error("failed to filter logs", "error", err)
				continue
			}

			if len(logs) > 0 {
				done <- struct{}{}
				return
			}
		}
	}
}
