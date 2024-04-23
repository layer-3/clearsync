package userop

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/artifacts/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
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
	// NOTE: only `executeBatch` is supported for now.
	//
	// Parameters:
	//   - ctx - is the context of the operation.
	//   - smartWallet - is the address of the smart wallet that will execute the user operation.
	//   - signer - is the signer function that will sign the user operation.
	//   - calls - is the list of calls to be executed in the user operation. Must not be empty.
	//   - walletDeploymentOpts - are the options for the smart wallet deployment. Can be nil if the smart wallet is already deployed.
	// 	 - overrides - are the overrides for the middleware during the user operation creation. Can be nil.
	//
	// Returns:
	//   - UserOperation - user operation with all fields filled in.
	//   - error - if failed to build the user operation.
	NewUserOp(
		ctx context.Context,
		sender common.Address,
		signer Signer,
		calls smart_wallet.Calls,
		walletDeploymentOpts *WalletDeploymentOpts,
		overrides *Overrides,
	) (UserOperation, error)

	// SignUserOp signs the user operation with the provided signer.
	//
	// Parameters:
	//   - ctx - is the context of the operation.
	//   - op - is the user operation to be signed.
	//   - signer - is the signer function that will sign the user operation.
	//
	// Returns:
	//   - UserOperation - user operation with modified signature.
	//   - error - if failed to sign the user operation
	SignUserOp(
		ctx context.Context,
		op UserOperation,
		signer Signer,
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

// WalletDeploymentOpts represents data required
// 1. to deploy a new smart wallet
// 2. to get the address of the already deployed wallet.
type WalletDeploymentOpts struct {
	Owner common.Address
	Index decimal.Decimal
}

// Each field overrides the corresponding middleware during the user operation creation.
type Overrides struct {
	Nonce     *big.Int
	InitCode  []byte
	GasPrices *GasPriceOverrides
	GasLimits *GasLimitOverrides
}

// These override provider's estimation. NOTE: if all are supplied, provider's estimation is NOT performed.
type GasPriceOverrides struct {
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

// These override the bundler's estimation. NOTE: if all are supplied, bundler's estimation is NOT performed.
type GasLimitOverrides struct {
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
}

// backend represents a user operation client.
type backend struct {
	provider EthBackend
	bundler  RPCBackend
	chainID  *big.Int

	smartWallet smart_wallet.Config
	entryPoint  common.Address
	paymaster   common.Address

	getNonce     middleware
	getInitCode  middleware
	getGasPrices middleware
	getGasLimits middleware
	sign         middleware
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

	providerURL, err := url.Parse(config.ProviderURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the blockchain RPC URL: %w", err)
	}

	providerRPC, err := NewEthBackend(*providerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the blockchain RPC: %w", err)
	}

	chainID, err := providerRPC.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}
	slog.Debug("fetched chain ID", "chainID", chainID.String())

	bundlerURL, err := url.Parse(config.BundlerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the blockchain RPC URL: %w", err)
	}

	bundlerRPC, err := NewRPCBackend(*bundlerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the bundler RPC: %w", err)
	}

	entryPointContract, err := entry_point_v0_6_0.NewEntryPoint(config.EntryPoint, providerRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the entry point contract: %w", err)
	}

	getInitCode, err := getInitCodeMiddleware(providerRPC, config.SmartWallet)
	if err != nil {
		return nil, fmt.Errorf("failed to build initCode middleware: %w", err)
	}

	getGasLimits, err := getGasLimitsMiddleware(bundlerRPC, config)
	if err != nil {
		return nil, fmt.Errorf("failed to build gas limits estimation middleware: %w", err)
	}

	return &backend{
		provider: providerRPC,
		bundler:  bundlerRPC,
		chainID:  chainID,

		smartWallet: config.SmartWallet,
		entryPoint:  config.EntryPoint,
		paymaster:   config.Paymaster.Address,

		getNonce:     getNonceMiddleware(entryPointContract),
		getInitCode:  getInitCode,
		getGasPrices: getGasPricesMiddleware(providerRPC, config.Gas),
		getGasLimits: getGasLimits,
		sign:         getSignMiddleware(config.EntryPoint, chainID),
	}, nil
}

func (c *backend) IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error) {
	accountAddress, err := c.GetAccountAddress(ctx, owner, index)
	if err != nil {
		return false, fmt.Errorf("failed to get account address: %w", err)
	}

	return smart_wallet.IsAccountDeployed(ctx, c.provider, accountAddress)
}

func (c *backend) GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error) {
	return smart_wallet.GetAccountAddress(ctx, c.provider, c.smartWallet, c.entryPoint, owner, index)
}

func (c *backend) NewUserOp(
	ctx context.Context,
	smartWallet common.Address,
	signer Signer,
	calls smart_wallet.Calls,
	walletDeploymentOpts *WalletDeploymentOpts,
	overrides *Overrides,
) (UserOperation, error) {
	if signer == nil {
		return UserOperation{}, ErrNoSigner
	}

	if len(calls) == 0 {
		return UserOperation{}, ErrNoCalls
	}

	ctx = context.WithValue(ctx, ctxKeySigner, signer)

	callData, err := smart_wallet.BuildCallData(*c.smartWallet.Type, calls)
	if err != nil {
		return UserOperation{}, fmt.Errorf("failed to build call data: %w", err)
	}

	op := UserOperation{Sender: smartWallet, CallData: callData}

	slog.Debug("applying middlewares to user operation")

	overridesPresent := overrides != nil

	// getNonce
	if overridesPresent && overrides.Nonce != nil {
		op.Nonce = decimal.NewFromBigInt(overrides.Nonce, 0)
	} else {
		if err := c.getNonce(ctx, &op); err != nil {
			return UserOperation{}, err
		}
	}

	// getInitCode
	if overridesPresent && overrides.InitCode != nil {
		op.InitCode = overrides.InitCode
	} else {
		isDeployed, err := smart_wallet.IsAccountDeployed(ctx, c.provider, smartWallet)
		if err != nil {
			return UserOperation{}, fmt.Errorf("failed to check if smart account is already deployed: %w", err)
		}

		if !isDeployed {
			if walletDeploymentOpts == nil {
				return UserOperation{}, ErrNoWalletDeploymentOpts
			}

			if walletDeploymentOpts.Owner == (common.Address{}) {
				return UserOperation{}, ErrNoWalletOwnerInWDO
			}

			ctx = context.WithValue(ctx, ctxKeyOwner, walletDeploymentOpts.Owner)
			ctx = context.WithValue(ctx, ctxKeyIndex, walletDeploymentOpts.Index)

			if err := c.getInitCode(ctx, &op); err != nil {
				return UserOperation{}, err
			}
		}
	}

	// getGasPrices
	if overridesPresent && overrides.GasPrices != nil {
		if overrides.GasPrices.MaxFeePerGas != nil {
			op.MaxFeePerGas = decimal.NewFromBigInt(overrides.GasPrices.MaxFeePerGas, 0)
		}
		if overrides.GasPrices.MaxPriorityFeePerGas != nil {
			op.MaxPriorityFeePerGas = decimal.NewFromBigInt(overrides.GasPrices.MaxPriorityFeePerGas, 0)
		}
	}
	err = c.getGasPrices(ctx, &op)
	if err != nil {
		return UserOperation{}, err
	}

	// sign before estimating gas limits, so that signature is well-formed.
	// If signature is corrupted, this can cause SmartWallet estimation to fail,
	// and the bundler will return an error.
	err = c.sign(ctx, &op)
	if err != nil {
		return UserOperation{}, err
	}

	// getGasLimits
	if overridesPresent && overrides.GasLimits != nil {
		if overrides.GasLimits.CallGasLimit != nil {
			op.CallGasLimit = decimal.NewFromBigInt(overrides.GasLimits.CallGasLimit, 0)
		}
		if overrides.GasLimits.VerificationGasLimit != nil {
			op.VerificationGasLimit = decimal.NewFromBigInt(overrides.GasLimits.VerificationGasLimit, 0)
		}
		if overrides.GasLimits.PreVerificationGas != nil {
			op.PreVerificationGas = decimal.NewFromBigInt(overrides.GasLimits.PreVerificationGas, 0)
		}
	}
	err = c.getGasLimits(ctx, &op)
	if err != nil {
		return UserOperation{}, err
	}

	// sign
	err = c.sign(ctx, &op)
	if err != nil {
		return UserOperation{}, err
	}

	slog.Debug("middlewares applied successfully", "userop", op)
	return op, nil
}

func (c *backend) SignUserOp(ctx context.Context, op UserOperation, signer Signer) (UserOperation, error) {
	if signer == nil {
		return UserOperation{}, ErrNoSigner
	}

	ctx = context.WithValue(ctx, ctxKeySigner, signer)
	if err := c.sign(ctx, &op); err != nil {
		return UserOperation{}, err
	}

	return op, nil
}

func (c *backend) SendUserOp(ctx context.Context, op UserOperation) (<-chan Receipt, error) {
	ctx, cancel := context.WithCancel(ctx)
	userOpHash, err := op.UserOpHash(c.entryPoint, c.chainID)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to calculate user operation hash: %w", err)
	}
	recCh := make(chan Receipt, 1)

	go subscribeUserOpEvent(ctx, cancel, c.provider, recCh, c.entryPoint, userOpHash)

	// ERC4337-standardized call to the bundler
	slog.Debug("sending user operation")
	if err := c.bundler.CallContext(ctx, &userOpHash, "eth_sendUserOperation", op, c.entryPoint); err != nil {
		return nil, fmt.Errorf("call to `eth_sendUserOperation` failed: %w", err)
	}

	slog.Info("user operation sent successfully", "userOpHash", userOpHash.Hex())
	return recCh, nil
}

// subscribeUserOpEvent waits for a user operation to be committed on block.
func subscribeUserOpEvent(
	ctx context.Context,
	cancel context.CancelFunc,
	client EthBackend,
	recCh chan<- Receipt,
	entryPoint common.Address,
	userOpHash common.Hash,
) {
	defer close(recCh)

	entryPointContract, err := entry_point_v0_6_0.NewEntryPoint(entryPoint, client)
	if err != nil {
		slog.Error("failed to connect to the entry point contract", "error", err)
		cancel()
		return
	}

	eventCh := make(chan *entry_point_v0_6_0.EntryPointUserOperationEvent)

	sub, err := entryPointContract.WatchUserOperationEvent(nil, eventCh, [][32]byte{userOpHash}, nil, nil)
	if err != nil {
		slog.Error("failed to subscribe to UserOperationEvent", "error", err)
		cancel()
		return
	}

	select {
	case event := <-eventCh:
		receipt := processUserOpEvent(client, event)
		recCh <- *receipt
	case err := <-sub.Err():
		slog.Error("subscription error", "error", err)
		cancel()
	case <-ctx.Done():
		err := ctx.Err()
		if err == nil || errors.Is(err, context.Canceled) {
			slog.Debug("context done for userOp event", "reason", "context canceled", "hash", userOpHash.Hex())
		} else if errors.Is(err, context.DeadlineExceeded) {
			slog.Error("context done for userOp event", "reason", "deadline exceeded", "hash", userOpHash.Hex())
		} else {
			slog.Error("context done for userOp event", "reason", "context error", "error", err, "hash", userOpHash.Hex())
		}
	}

	sub.Unsubscribe()
}

func processUserOpEvent(client EthBackend, event *entry_point_v0_6_0.EntryPointUserOperationEvent) *Receipt {
	receipt := Receipt{
		UserOpHash:    event.UserOpHash,
		TxHash:        event.Raw.TxHash,
		Sender:        event.Sender,
		Nonce:         decimal.NewFromBigInt(event.Nonce, 0),
		Success:       event.Success,
		ActualGasCost: decimal.NewFromBigInt(event.ActualGasCost, 0),
		ActualGasUsed: decimal.NewFromBigInt(event.ActualGasUsed, 0),
	}

	if !receipt.Success { // Try to fetch revert reason.
		txReceipt, err := client.TransactionReceipt(context.Background(), receipt.TxHash)
		if err != nil {
			slog.Error("failed to fetch userop transaction receipt by hash", "error", err)
			return &receipt
		}

		if userOpRevertReasonLog := filterLogsByEventID(txReceipt.Logs, userOpRevertReasonID); userOpRevertReasonLog != nil {
			unpackedRevertReasonParams, err := entryPointUserOpEventsABI.Unpack("UserOperationRevertReason", userOpRevertReasonLog.Data)
			if err != nil {
				slog.Error("Error decoding UserOperationRevertReason params:", err)
				return &receipt
			}

			if len(unpackedRevertReasonParams) == 2 {
				slog.Debug("parsed userOperationRevertReason logs", "data", hexutil.Encode(userOpRevertReasonLog.Data), "parsedParams", unpackedRevertReasonParams)
				revertData, ok := unpackedRevertReasonParams[1].([]byte)
				if !ok {
					slog.Error("failed to parse revert reason", "unpackedRevertReasonParams", unpackedRevertReasonParams)
					return &receipt
				}
				receipt.RevertData = revertData
			} else {
				slog.Warn("unexpected number of unpackedRevertReasonParams", "unpackedRevertReasonParams", unpackedRevertReasonParams)
			}
		}
	}

	return &receipt
}

// Return only one log for simplicity, although several logs
// with the same event signature can be emitted during one tx.
func filterLogsByEventID(logs []*types.Log, eventID common.Hash) *types.Log {
	for _, log := range logs {
		if log.Topics[0] == eventID {
			return log
		}
	}
	return nil
}
