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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point"
	"github.com/layer-3/clearsync/pkg/abi/simple_account"
)

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
	SendUserOp(ctx context.Context, op UserOperation) (<-chan Receipt, error)
}

// Call represents sufficient data to build a single transaction,
// which is a part of a user operation
// to be executed in a batch with other ones.
type Call struct {
	To       common.Address
	Value    decimal.Decimal // Value is a wei amount to be sent with the call.
	CallData []byte
}

// WalletDeploymentOpts represents data required
// 1. to deploy a new smart wallet
// 2. to get the address of the already deployed wallet.
type WalletDeploymentOpts struct {
	Owner common.Address
	Index decimal.Decimal
}

// client represents a user operation client.
type client struct {
	providerRPC *ethclient.Client
	bundlerRPC  *rpc.Client
	chainID     *big.Int

	smartWallet SmartWalletConfig
	entryPoint  common.Address
	paymaster   common.Address
	middlewares []middleware
}

type Receipt struct {
	UserOpHash    common.Hash
	TxHash        common.Hash
	Sender        common.Address
	Nonce         decimal.Decimal
	Success       bool
	ActualGasCost decimal.Decimal
	ActualGasUsed decimal.Decimal
	RevertData    []byte // non-empty if Success is false
}

// NewClient is a factory that builds a new
// user operation client based on the provided configuration.
func NewClient(config ClientConfig) (UserOperationClient, error) {
	providerRPC, err := ethclient.Dial(config.ProviderURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the blockchain RPC: %w", err)
	}

	chainID, err := providerRPC.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}
	slog.Debug("fetched chain ID", "chainID", chainID.String())

	bundlerRPC, err := rpc.Dial(config.BundlerURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the bundler RPC: %w", err)
	}

	entryPointContract, err := entry_point.NewEntryPoint(config.EntryPoint, providerRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the entry point contract: %w", err)
	}

	estimateGas := estimateUserOperationGas(bundlerRPC, config.EntryPoint)
	if config.Paymaster.Type != nil && *config.Paymaster.Type != PaymasterDisabled {
		switch typ := config.Paymaster.Type; *typ {
		case PaymasterPimlicoERC20:
			estimateGas = getPimlicoERC20PaymasterData(
				bundlerRPC,
				config.EntryPoint,
				config.Paymaster.Address,
				config.Paymaster.PimlicoERC20.VerificationGasOverhead,
			)
		case PaymasterPimlicoVerifying:
			// NOTE: PimlicoVerifying is the easiest to implement
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
		providerRPC: providerRPC,
		bundlerRPC:  bundlerRPC,
		chainID:     chainID,
		smartWallet: config.SmartWallet,
		entryPoint:  config.EntryPoint,
		paymaster:   config.Paymaster.Address,
		middlewares: []middleware{ // Middleware order matters - first in, first executed.
			getNonce(entryPointContract),
			getInitCode,
			getGasPrice(providerRPC, config.Gas),
			sign(config.EntryPoint, chainID),
			estimateGas,
			sign(config.EntryPoint, chainID), // update signature after gas estimation
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
	getInitCode, err := getInitCode(c.providerRPC, c.smartWallet)
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
	slog.Debug("apply middlewares to user operation")

	isDeployed, err := isAccountDeployed(c.providerRPC, smartWallet)
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
	}

	if !isDeployed {
		if walletDeploymentOpts == nil {
			return UserOperation{}, ErrNoWalletDeploymentOpts
		}

		if walletDeploymentOpts.Owner == (common.Address{}) {
			return UserOperation{}, ErrNoWalletOwner
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
func (c *client) SendUserOp(ctx context.Context, op UserOperation) (<-chan Receipt, error) {
	slog.Debug("sending user operation")

	ctx, cancel := context.WithCancel(ctx)

	userOpHash := op.UserOpHash(c.entryPoint, c.chainID)

	waiter := make(chan Receipt, 1)
	go waitForUserOpEvent(ctx, cancel, c.providerRPC, waiter, c.entryPoint, userOpHash)

	// ERC4337-standardized call to the bundler
	if err := c.bundlerRPC.CallContext(ctx, &userOpHash, "eth_sendUserOperation", op, c.entryPoint); err != nil {
		return nil, fmt.Errorf("call to `eth_sendUserOperation` failed: %w", err)
	}

	slog.Info("user operation sent successfully", "userOpHash", userOpHash.Hex())

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
	switch *c.smartWallet.Type {
	case SmartWalletSimpleAccount, SmartWalletBiconomy:
		return handleCallSimpleAccount(calls)
	case SmartWalletKernel:
		return handleCallKernel(calls)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", c.smartWallet.Type)
	}
}

// handleCallSimpleAccount packs calldata for SimpleAccount smart wallet.
func handleCallSimpleAccount(calls []Call) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(simple_account.SimpleAccountMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	addresses := make([]any, 0, len(calls))
	calldatas := make([]any, 0, len(calls))
	for _, call := range calls {
		addresses = append(addresses, call.To)
		calldatas = append(calldatas, call.CallData)
	}

	// pack the data for the `executeBatch` smart account function
	// Biconomy v2.0: https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol#L128
	// NOTE: you can NOT send native token with SimpleAccount v0.6.0 because of `executeBatch` signature
	data, err := parsedABI.Pack("executeBatch", addresses, calldatas)
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
	// Zerodev Kernel v2.2: https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L82
	data, err := parsedABI.Pack("executeBatch", params)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for Kernel: %w", err)
	}
	return data, nil
}

// waitForUserOpEvent waits for a user operation to be committed on block.
func waitForUserOpEvent(
	ctx context.Context,
	cancel context.CancelFunc,
	client *ethclient.Client,
	done chan<- Receipt,
	entryPoint common.Address,
	userOpHash common.Hash,
) {
	defer close(done)

	waitInterval := time.Millisecond * 5000

	fromBlock, err := client.BlockNumber(ctx)
	if err != nil {
		slog.Error("failed to get block number", "error", err)
		cancel()
		return
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{entryPoint},
		Topics:    [][]common.Hash{{}, {userOpHash}},
		FromBlock: big.NewInt(int64(fromBlock)),
	}

	ticker := time.NewTicker(waitInterval)
	defer ticker.Stop()

	userOpEvents, err := abi.JSON(strings.NewReader(entrypointUserOpEventsABI))
	if err != nil {
		slog.Error("error parsing ABI", "err", err)
		return
	}

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

			// keccak256("UserOperationEvent(bytes32,address,address,uint256,bool,uint256,uint256)")
			userOpEventID := common.HexToHash("0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f")

			// There are several events where userOpHash is used as a topic
			// Namely, UserOperationEvent, AccountDeployed and UserOperationRevertReason
			// see https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/interfaces/IEntryPoint.sol#L19-L47
			userOpEventLog := filterLogsByEventID(logs, userOpEventID)
			if userOpEventLog == nil {
				continue
			}

			topics := userOpEventLog.Topics
			sender := common.BytesToAddress(topics[2].Bytes())
			data := userOpEventLog.Data

			userOpReceipt := Receipt{
				UserOpHash: userOpHash,
				TxHash:     userOpEventLog.TxHash,
				Sender:     sender,
			}

			// Decode the ABI-encoded message
			unpackedUserOpParams, err := userOpEvents.Unpack("UserOperationEvent", data)
			if err != nil {
				slog.Error("Error decoding UserOperationEvent params:", err)
				return
			}

			if len(unpackedUserOpParams) == 4 {
				slog.Debug("parsed userOperationEvent logs", "data", hexutil.Encode(data), "parsedParams", unpackedUserOpParams)

				userOpReceipt.Nonce = decimal.NewFromBigInt(unpackedUserOpParams[0].(*big.Int), 0)
				userOpReceipt.Success = unpackedUserOpParams[1].(bool)
				userOpReceipt.ActualGasCost = decimal.NewFromBigInt(
					unpackedUserOpParams[2].(*big.Int),
					0,
				)
				userOpReceipt.ActualGasUsed = decimal.NewFromBigInt(
					unpackedUserOpParams[3].(*big.Int),
					0,
				)
			} else {
				slog.Warn("unexpected number of unpackedUserOpParams", "unpackedUserOpParams", unpackedUserOpParams)
			}

			if !userOpReceipt.Success {
				// try fetching revert reason
				// keccak256("UserOperationRevertReason(bytes32,address,uint256,bytes)")
				userOpRevertReasonID := common.HexToHash("0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201")
				userOpRevertReasonLog := filterLogsByEventID(logs, userOpRevertReasonID)

				if userOpRevertReasonLog != nil {
					unpackedRevertReasonParams, err := userOpEvents.Unpack("UserOperationRevertReason", userOpRevertReasonLog.Data)
					if err != nil {
						slog.Error("Error decoding UserOperationRevertReason params:", err)
						return
					}

					if len(unpackedRevertReasonParams) == 2 {
						slog.Debug("parsed userOperationRevertReason logs", "data", hexutil.Encode(userOpRevertReasonLog.Data), "parsedParams", unpackedRevertReasonParams)
						userOpReceipt.RevertData = unpackedRevertReasonParams[1].([]byte)
					} else {
						slog.Warn("unexpected number of unpackedRevertReasonParams", "unpackedRevertReasonParams", unpackedRevertReasonParams)
					}
				}
			}

			done <- userOpReceipt
			return
		}
	}
}

// Return only one log for simplicity, although several logs
// with the same event signature can be emitted during one tx.
func filterLogsByEventID(logs []types.Log, eventID common.Hash) *types.Log {
	for _, log := range logs {
		if log.Topics[0] == eventID {
			return &log
		}
	}
	return nil
}
