package userop

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point"
	"github.com/layer-3/clearsync/pkg/abi/itoken"
	"github.com/layer-3/clearsync/pkg/abi/simple_account/account_abstraction"
)

// UserOperationClient represents a client for creating and posting user operations.
type UserOperationClient interface {
	NewUserOp(
		ctx context.Context,
		sender common.Address,
		receiver common.Address,
		token common.Address,
		amount decimal.Decimal,
	) (UserOperation, error)
	SendUserOp(ctx context.Context, op UserOperation, callback func()) error
}

// Client represents a user operation client.
type Client struct {
	providerRPC *ethclient.Client
	bundlerRPC  *rpc.Client
	chainID     *big.Int

	entryPoint         common.Address
	isPaymasterEnabled bool
	paymaster          common.Address
	signer             func(userOperation UserOperation, entryPoint common.Address, chainId *big.Int) common.Hash
	middlewares        []middleware
}

// NewClient is a factory that builds a new
// user operation client based on the provided configuration.
func NewClient(config ClientConfig) (UserOperationClient, error) {
	providerClient, err := ethclient.Dial(config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}

	bundlerRPC, err := rpc.Dial(config.BundlerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum RPC: %w", err)
	}

	entryPointContract, err := entry_point.NewEntryPoint(config.EntryPoint, providerClient)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the entry point contract: %w", err)
	}

	return &Client{
		providerRPC:        providerClient,
		bundlerRPC:         bundlerRPC,
		chainID:            config.ChainID,
		entryPoint:         config.EntryPoint,
		isPaymasterEnabled: config.Paymaster.URL != "",
		paymaster:          config.Paymaster.Address,
		signer:             config.Signer,
		middlewares: []middleware{
			// Middleware order matters - first in, first executed
			getNonce(entryPointContract),
			// getInitCode(providerRPC), // TODO: add smart wallet deployment support
			getGasPrice(providerClient),
			estimateUserOperationGas(bundlerRPC, config.EntryPoint),
		},
	}, nil
}

// NewUserOp builds and fills in a new UserOperation.
func (c *Client) NewUserOp(
	ctx context.Context,
	sender common.Address,
	receiver common.Address,
	token common.Address,
	amount decimal.Decimal,
) (UserOperation, error) {
	op := UserOperation{Sender: sender}

	erc20, err := abi.JSON(strings.NewReader(itoken.IERC20MetaData.ABI))
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to parse IERC20 ABI: %w", err)
	}

	var addresses []common.Address
	var values []*big.Int
	var callData [][]byte
	if c.isPaymasterEnabled {
		slog.Info("paymaster is enabled")
		approveData, err := erc20.Pack("approve", c.paymaster, amount.BigInt())
		if err != nil {
			return UserOperation{}, fmt.Errorf("failed to pack approve data: %w", err)
		}

		addresses = append(addresses, token)
		values = append(values, new(big.Int))
		callData = append(callData, approveData)

		op.PaymasterAndData = c.paymaster[:]
	}

	transferData, err := erc20.Pack("transfer", receiver, amount.BigInt())
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to pack transfer data: %w", err)
	}
	addresses = append(addresses, token)
	values = append(values, new(big.Int))
	callData = append(callData, transferData)

	// Pack calldata for SimpleAccount / Biconomy contract
	parsedABI, err := abi.JSON(strings.NewReader(account_abstraction.SimpleAccountMetaData.ABI))
	data, err := parsedABI.Pack("executeBatch", addresses, values, callData)
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to pack executeBatch data: %w", err)
	}

	// Pack calldata for Zerodev Kernel contract
	// parsedABI, err := abi.JSON(strings.NewReader(kernel.KernelMetaData.ABI))
	// params := []callStructKernel{
	// 	{
	// 		To:    token,
	// 		Value: new(big.Int),
	// 		Data:  callData[0],
	// 	},
	// 	{
	// 		To:    token,
	// 		Value: new(big.Int),
	// 		Data:  callData[1],
	// 	},
	// }
	// data, err := parsedABI.Pack("executeBatch", params)
	// if err != nil {
	// 	return UserOperation{}, err
	// }

	op.CallData = data

	slog.Info("apply middlewares to user operation")
	for _, fn := range c.middlewares {
		if err := fn(ctx, &op); err != nil {
			return UserOperation{}, fmt.Errorf("failed to apply middleware to user operation: %w", err)
		}
	}

	op.Signature = c.signer(op, c.entryPoint, c.chainID) // should come last

	return op, nil
}

// type callStructKernel struct {
// 	To    common.Address
// 	Value *big.Int
// 	Data  []byte
// }

// SendUserOp submits a user operation to a bundler
// and executes the provided callback function.
func (c *Client) SendUserOp(ctx context.Context, op UserOperation, callback func()) error {
	slog.Info("sending user operation", "json", op.Marshal())

	var userOpHash common.Hash
	err := c.bundlerRPC.CallContext(ctx, &userOpHash, "eth_sendUserOperation", op.Marshal())
	if err != nil {
		return fmt.Errorf("failed to send user operation: %w", err)
	}

	slog.Info("user operation sent successfully", "hash", userOpHash.Hex())
	callback()
	return nil
}
