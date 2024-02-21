package userop

import (
	"context"
	"errors"
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

type WalletDeploymentOpts struct {
	Owner common.Address
	Index decimal.Decimal
}

// UserOperationClient represents a client for creating and posting user operations.
type UserOperationClient interface {
	IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error)
	GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error)
	NewUserOp(
		ctx context.Context,
		sender common.Address,
		signer Signer,
		calls []Call,
		walletDeploymentOpts *WalletDeploymentOpts,
	) (UserOperation, error)
	SendUserOp(ctx context.Context, op UserOperation) (<-chan struct{}, error)
}

// Call represents sufficient data to build a single transaction,
// which is a part of a user operation
// to be executed in a batch with other ones.
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

	smartWalletConfig  SmartWallet
	entryPoint         common.Address
	isPaymasterEnabled bool
	paymaster          common.Address
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
			// NOTE: PimlicoVerifying is the easiest to add
			return nil, fmt.Errorf("%w: %s", ErrPaymasterNotSupported, typ)
		case PaymasterBiconomyERC20:
			return nil, ErrPaymasterNotSupported
		case PaymasterBiconomySponsoring:
			// NOTE: tried to add BiconomySponsoring, but it is not responding correctly
			return nil, ErrPaymasterNotSupported
		default:
			return nil, fmt.Errorf("unknown paymaster type: %s", typ)
		}
	}

	getInitCode, err := getInitCode(providerRPC, config.SmartWallet)
	if err != nil {
		return nil, fmt.Errorf("failed to build initCode middleware: %w", err)
	}

	return &client{
		providerRPC:        providerRPC,
		bundlerRPC:         bundlerRPC,
		chainID:            config.ChainID,
		smartWalletConfig:  config.SmartWallet,
		entryPoint:         config.EntryPoint,
		isPaymasterEnabled: isPaymasterEnabled,
		paymaster:          config.Paymaster.Address,
		middlewares: []middleware{ // Middleware order matters - first in, first executed.
			getNonce(entryPointContract),
			getInitCode,
			getGasPrice(providerRPC, config.Gas),
			sign(config.EntryPoint, config.ChainID),
			estimateGas,
			sign(config.EntryPoint, config.ChainID), // update signature after gas estimation
		},
	}, nil
}

func (c *client) IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error) {
	swAddress, err := c.GetAccountAddress(ctx, owner, index)
	if err != nil {
		return false, fmt.Errorf("failed to get account address: %w", err)
	}

	return isAccountDeployed(c.providerRPC, swAddress)
}

func (c *client) GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error) {
	getInitCode, err := getInitCode(c.providerRPC, c.smartWalletConfig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to build initCode middleware: %w", err)
	}

	ctx = context.WithValue(ctx, ctxKeyOwner, owner)
	ctx = context.WithValue(ctx, ctxKeyIndex, index)

	op := UserOperation{
		Sender: common.Address{},
	}

	if err := getInitCode(ctx, &op); err != nil {
		return common.Address{}, fmt.Errorf("failed to get init code: %w", err)
	}

	entryPointABI, err := abi.JSON(strings.NewReader(entry_point.EntryPointMetaData.ABI))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// calculate the smart wallet address that will be generated by the entry point
	// See https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/core/EntryPoint.sol#L356
	getSenderAddressData, err := entryPointABI.Pack("getSenderAddress", op.InitCode)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to pack getSenderAddress data: %w", err)
	}

	msg := ethereum.CallMsg{
		To:   &c.entryPoint,
		Data: getSenderAddressData,
	}

	_, err = c.providerRPC.CallContract(ctx, msg, nil)
	if err == nil {
		panic(fmt.Errorf("'getSenderAddress' call returned no error, but expected one"))
	}

	var scError rpc.DataError
	ok := errors.As(err, &scError)
	if !ok {
		panic(fmt.Errorf("unexpected error type '%T' containing message %w)", err, err))
	}
	errorData := scError.ErrorData().(string)

	senderAddressResultError, ok := entryPointABI.Errors["SenderAddressResult"]
	if !ok {
		panic(fmt.Errorf("ABI does not contain 'SenderAddressResult' error"))
	}

	// check if the error signature is correct
	if errorData[0:10] != senderAddressResultError.ID.String()[0:10] {
		panic(fmt.Errorf("'getSenderAddress' unexpected error signature: %s", errorData[0:10]))
	}

	// check if the error data has the correct length
	if len(errorData) < 74 {
		panic(fmt.Errorf("'getSenderAddress' revert data expected to have lenght of 74, but got: %d", len(errorData)))
	}

	return common.HexToAddress(errorData[34:]), nil
}

// NewUserOp builds and fills in a new UserOperation.
func (c *client) NewUserOp(
	ctx context.Context,
	smartWallet common.Address,
	signer Signer,
	calls []Call,
	walletDeploymentOpts *WalletDeploymentOpts,
) (UserOperation, error) {
	slog.Info("apply middlewares to user operation")

	isDeployed, err := isAccountDeployed(c.providerRPC, smartWallet)
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	}

	if !isDeployed {
		if walletDeploymentOpts == nil {
			return UserOperation{}, ErrNoWalletDeploymentOpts
		}
		ctx = context.WithValue(ctx, ctxKeyOwner, walletDeploymentOpts.Owner)
		ctx = context.WithValue(ctx, ctxKeyIndex, walletDeploymentOpts.Index)
	}

	if signer == nil {
		return UserOperation{}, ErrNoSigner
	}
	ctx = context.WithValue(ctx, ctxKeySigner, signer)
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
	// ERC4337-standardized call to the bundler
	if err := c.bundlerRPC.CallContext(ctx, &userOpHash, "eth_sendUserOperation", op, c.entryPoint); err != nil {
		return nil, fmt.Errorf("call to `eth_sendUserOperation` failed: %w", err)
	}

	slog.Info("user operation sent successfully", "hash", userOpHash.Hex())
	waiter := make(chan struct{}, 1)
	go waitForUserOpEvent(c.providerRPC, waiter, c.entryPoint, userOpHash)

	return waiter, nil
}

func isAccountDeployed(provider *ethclient.Client, swAddress common.Address) (bool, error) {
	var result any
	if err := provider.Client().CallContext(
		context.Background(),
		&result,
		"eth_getCode",
		swAddress,
		"latest",
	); err != nil {
		return false, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	}

	byteCode, ok := result.(string)
	if !ok {
		return false, fmt.Errorf("unexpected type: %T", result)
	}

	// assume that the smart account is deployed if it has non-zero byte code
	if byteCode == "" || byteCode == "0x" {
		return false, nil
	}

	return true, nil
}

func (c *client) buildCallData(calls []Call) ([]byte, error) {
	switch c.smartWalletConfig.Type {
	case SmartWalletSimpleAccount, SmartWalletBiconomy:
		return handleCallSimpleAccount(calls)
	case SmartWalletKernel:
		return handleCallKernel(calls)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", c.smartWalletConfig.Type)
	}
}

// handleCallSimpleAccount packs calldata for SimpleAccount smart wallet.
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

	// pack the data for the `executeBatch` smart account function
	// Biconomy v2.0: https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol#L128
	// eth-infinitism v0.6.0: TODO: downgrade to v0.6.0, update ABI
	data, err := parsedABI.Pack("executeBatch", addresses, values, calldatas)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for SimpleAccount: %w", err)
	}
	return data, nil
}

// callStructKernel represents a call to the Zerodev Kernel smart wallet.
// The idea is the same as in Call type,
// but tailed specifically to the Zerodev Kernel ABI.
type callStructKernel struct {
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`
	Data  []byte         `json:"data"`
}

// handleCallKernel packs calldata for Zerodev Kernel smart wallet.
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

	// pack the data for the `executeBatch` smart account function
	// Zerodev Kernel v2.3: https://github.com/zerodevapp/kernel/blob/v2.3/src/Kernel.sol#L98
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
