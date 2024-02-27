package userop

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point"
)

// Client represents a client for creating and posting user operations.
type Client interface {
	// IsAccountDeployed checks whether the smart wallet for the specified owner EOA and index is deployed.
	//
	// Parameters:
	//   - owner - is the EOA address of the smart wallet owner.
	//   - index - is the index of the smart wallet, 0 by default. SW index allows to deploy multiple smart wallets for the same owner.
	//
	// Returns:
	//   - bool - true if the smart wallet is deployed, false if not
	//   - error - if failed to check.
	IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error)

	// GetAccountAddress returns the address of the smart wallet for the specified owner EOA and index.
	//
	// Parameters:
	//   - owner - is the EOA address of the smart wallet owner.
	//   - index - is the index of the smart wallet, 0 by default. SW index allows to deploy multiple smart wallets for the same owner.
	//
	// Returns:
	//   - common.Address - an address of the smart wallet
	//   - error - if failed to calculate it.
	GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error)

	// NewUserOp builds a new UserOperation and fills all the fields.
	//
	// Parameters:
	//   - ctx - is the context of the operation.
	//   - smartWallet - is the address of the smart wallet that will execute the user operation.
	//   - signer - is the signer function that will sign the user operation.
	//   - calls - is the list of calls to be executed in the user operation.
	//   - walletDeploymentOpts - are the options for the smart wallet deployment. Can be nil if the smart wallet is already deployed.
	//
	// Returns:
	//   - UserOperation - user operation with all fields filled in.
	//   - error - if failed to build the user operation.
	NewUserOp(
		ctx context.Context,
		sender common.Address,
		signer Signer,
		calls []Call,
		walletDeploymentOpts *WalletDeploymentOpts,
	) (UserOperation, error)

	// SendUserOp submits a user operation to a bundler and returns a channel to await for the userOp receipt.
	//
	// Parameters:
	//   - ctx - is the context of the operation.
	//   - op - is the user operation to be sent.
	//
	// Returns:
	//   - <-chan Receipt - a channel to await for the userOp receipt.
	//   - error - if failed to send the user operation
	SendUserOp(ctx context.Context, op UserOperation) (done <-chan Receipt, err error)
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

// backend represents a user operation client.
type backend struct {
	provider EthBackend
	bundler  RPCBackend
	chainID  *big.Int

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
	RevertData    []byte // non-empty if Success is false and EntryPoint was able to catch revert reason.
}

// NewClient is a factory that builds a new
// user operation client based on the provided configuration.
func NewClient(config ClientConfig) (Client, error) {
	err := config.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	providerRPC, err := NewEthBackend(config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the blockchain RPC: %w", err)
	}

	chainID, err := providerRPC.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}
	slog.Debug("fetched chain ID", "chainID", chainID.String())

	bundlerRPC, err := NewRPCBackend(config.BundlerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the bundler RPC: %w", err)
	}

	entryPointContract, err := entry_point.NewEntryPoint(config.EntryPoint, providerRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the entry point contract: %w", err)
	}

	getGasEstimation, err := getGasEstimation(bundlerRPC, config)
	if err != nil {
		return nil, fmt.Errorf("failed to build gas estimation middleware: %w", err)
	}

	getInitCode, err := getInitCode(providerRPC, config.SmartWallet)
	if err != nil {
		return nil, fmt.Errorf("failed to build initCode middleware: %w", err)
	}

	return &backend{
		provider:    providerRPC,
		bundler:     bundlerRPC,
		chainID:     chainID,
		smartWallet: config.SmartWallet,
		entryPoint:  config.EntryPoint,
		paymaster:   config.Paymaster.Address,
		middlewares: []middleware{ // Middleware order matters - first in, first executed.
			getNonce(entryPointContract),
			getInitCode,
			getGasPrice(providerRPC, config.Gas),
			sign(config.EntryPoint, chainID),
			getGasEstimation,
			sign(config.EntryPoint, chainID), // update signature after gas estimation
		},
	}, nil
}

func (c *backend) IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error) {
	accountAddress, err := c.GetAccountAddress(ctx, owner, index)
	if err != nil {
		return false, fmt.Errorf("failed to get account address: %w", err)
	}

	return isAccountDeployed(c.provider, accountAddress)
}

func (c *backend) GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error) {
	getInitCode, err := getInitCode(c.provider, c.smartWallet)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to build initCode middleware: %w", err)
	}

	op := UserOperation{Sender: common.Address{}}
	ctx = context.WithValue(ctx, ctxKeyOwner, owner)
	ctx = context.WithValue(ctx, ctxKeyIndex, index)

	if err := getInitCode(ctx, &op); err != nil {
		return common.Address{}, fmt.Errorf("failed to get init code: %w", err)
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

	_, err = c.provider.CallContract(ctx, msg, nil)
	if err == nil {
		panic(fmt.Errorf("'getSenderAddress' call returned no error, but expected one"))
	}

	var scError rpc.DataError
	if ok := errors.As(err, &scError); !ok {
		panic(fmt.Errorf("unexpected error type '%T' containing message %w)", err, err))
	}
	errorData := scError.ErrorData().(string)

	senderAddressResultError, ok := entryPointABI.Errors["SenderAddressResult"]
	if !ok {
		panic(fmt.Errorf("ABI does not contain 'SenderAddressResult' error"))
	}

	// check if the error signature is correct
	if id := senderAddressResultError.ID.String(); errorData[0:10] != id[0:10] {
		panic(fmt.Errorf("'getSenderAddress' unexpected error signature: %s", errorData[0:10]))
	}

	// check if the error data has the correct length
	if len(errorData) < 74 {
		panic(fmt.Errorf("'getSenderAddress' revert data expected to have lenght of 74, but got: %d", len(errorData)))
	}

	return common.HexToAddress(errorData[34:]), nil
}

func (c *backend) NewUserOp(
	ctx context.Context,
	smartWallet common.Address,
	signer Signer,
	calls []Call,
	walletDeploymentOpts *WalletDeploymentOpts,
) (UserOperation, error) {
	slog.Debug("applying middlewares to user operation")
	if signer == nil {
		return UserOperation{}, ErrNoSigner
	}
	ctx = context.WithValue(ctx, ctxKeySigner, signer)

	isDeployed, err := isAccountDeployed(c.provider, smartWallet)
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

	callData, err := c.buildCallData(calls)
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to build call data: %w", err)
	}

	op := UserOperation{Sender: smartWallet, CallData: callData}
	for _, fn := range c.middlewares {
		if err := fn(ctx, &op); err != nil {
			return UserOperation{}, fmt.Errorf("failed to apply middleware to user operation: %w", err)
		}
	}

	slog.Debug("middlewares applied successfully", "userop", op)
	return op, nil
}

func (c *backend) SendUserOp(ctx context.Context, op UserOperation) (<-chan Receipt, error) {
	ctx, cancel := context.WithCancel(ctx)
	userOpHash := op.UserOpHash(c.entryPoint, c.chainID)
	done := make(chan Receipt, 1)

	go waitForUserOpEvent(ctx, cancel, c.provider, done, c.entryPoint, userOpHash)

	// ERC4337-standardized call to the bundler
	slog.Debug("sending user operation")
	if err := c.bundler.CallContext(ctx, &userOpHash, "eth_sendUserOperation", op, c.entryPoint); err != nil {
		return nil, fmt.Errorf("call to `eth_sendUserOperation` failed: %w", err)
	}

	slog.Info("user operation sent successfully", "userOpHash", userOpHash.Hex())
	return done, nil
}

func isAccountDeployed(provider EthBackend, swAddress common.Address) (bool, error) {
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
	return !(byteCode == "" || byteCode == "0x"), nil
}

func (c *backend) buildCallData(calls []Call) ([]byte, error) {
	switch *c.smartWallet.Type {
	case SmartWalletSimpleAccount, SmartWalletBiconomy:
		return handleCallSimpleAccount(calls)
	case SmartWalletKernel:
		return handleCallKernel(calls)
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", c.smartWallet.Type)
	}
}

// handleCallSimpleAccount packs CallData for SimpleAccount smart wallet.
func handleCallSimpleAccount(calls []Call) ([]byte, error) {
	addresses := make([]any, 0, len(calls))
	calldatas := make([]any, 0, len(calls))
	for _, call := range calls {
		addresses = append(addresses, call.To)
		calldatas = append(calldatas, call.CallData)
	}

	// Pack the data for the `executeBatch` smart account function.
	// Biconomy v2.0: https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol#L128
	// NOTE: you can NOT send native token with SimpleAccount v0.6.0 because of `executeBatch` signature
	data, err := simpleAccountABI.Pack("executeBatch", addresses, calldatas)
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
	data, err := kernelExecuteABI.Pack("executeBatch", params)
	if err != nil {
		return nil, fmt.Errorf("failed to pack executeBatch data for Kernel: %w", err)
	}
	return data, nil
}

// waitForUserOpEvent waits for a user operation to be committed on block.
func waitForUserOpEvent(
	ctx context.Context,
	cancel context.CancelFunc,
	client EthBackend,
	done chan<- Receipt,
	entryPoint common.Address,
	userOpHash common.Hash,
) {
	ticker := time.NewTicker(time.Millisecond * 5000)
	defer ticker.Stop()
	defer close(done)

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

			receipt, err := processLogs(logs, userOpHash)
			if err != nil {
				return
			}
			if receipt != nil {
				done <- *receipt
				return
			}
		}
	}
}

func processLogs(logs []types.Log, userOpHash common.Hash) (*Receipt, error) {
	// There are several events where userOpHash is used as a topic
	// Namely, UserOperationEvent, AccountDeployed and UserOperationRevertReason
	// see https://github.com/eth-infinitism/account-abstraction/blob/v0.6.0/contracts/interfaces/IEntryPoint.sol#L19-L47
	userOpEventLog := filterLogsByEventID(logs, userOpEventID)
	if userOpEventLog == nil {
		return nil, nil
	}

	// Decode the ABI-encoded message
	unpackedUserOpParams, err := entryPointUserOpEventsABI.Unpack("UserOperationEvent", userOpEventLog.Data)
	if err != nil {
		slog.Error("Error decoding UserOperationEvent params:", err)
		return nil, err
	}

	receipt := Receipt{
		UserOpHash: userOpHash,
		TxHash:     userOpEventLog.TxHash,
		Sender:     common.BytesToAddress(userOpEventLog.Topics[2].Bytes()),
	}

	if len(unpackedUserOpParams) == 4 {
		slog.Debug("parsed userOperationEvent logs", "data", hexutil.Encode(userOpEventLog.Data), "parsedParams", unpackedUserOpParams)
		receipt.Nonce = decimal.NewFromBigInt(unpackedUserOpParams[0].(*big.Int), 0)
		receipt.Success = unpackedUserOpParams[1].(bool)
		receipt.ActualGasCost = decimal.NewFromBigInt(unpackedUserOpParams[2].(*big.Int), 0)
		receipt.ActualGasUsed = decimal.NewFromBigInt(unpackedUserOpParams[3].(*big.Int), 0)
	} else {
		slog.Warn("unexpected number of unpackedUserOpParams", "unpackedUserOpParams", unpackedUserOpParams)
	}

	if !receipt.Success { // Try to fetch revert reason.
		if userOpRevertReasonLog := filterLogsByEventID(logs, userOpRevertReasonID); userOpRevertReasonLog != nil {
			unpackedRevertReasonParams, err := entryPointUserOpEventsABI.Unpack("UserOperationRevertReason", userOpRevertReasonLog.Data)
			if err != nil {
				slog.Error("Error decoding UserOperationRevertReason params:", err)
				return nil, err
			}

			if len(unpackedRevertReasonParams) == 2 {
				slog.Debug("parsed userOperationRevertReason logs", "data", hexutil.Encode(userOpRevertReasonLog.Data), "parsedParams", unpackedRevertReasonParams)
				receipt.RevertData = unpackedRevertReasonParams[1].([]byte)
			} else {
				slog.Warn("unexpected number of unpackedRevertReasonParams", "unpackedRevertReasonParams", unpackedRevertReasonParams)
			}
		}
	}

	return &receipt, nil
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
