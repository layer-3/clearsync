package userop

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/entry_point_v0_6_0"
	"github.com/layer-3/clearsync/pkg/smart_wallet"
)

type ctxKey int

const (
	ctxKeySigner ctxKey = iota
	ctxKeyOwner
	ctxKeyIndex
)

// middleware is a function that modifies a user operation.
// It is used to create a pipeline of operations to be executed
// to fill in the user operation with all the necessary data.
type middleware func(ctx context.Context, op *UserOperation) error

// TODO: possible improvement: when there is a userOp for this SW already in the mempool, we should return incremented nonce
func getNonceMiddleware(entryPoint *entry_point_v0_6_0.EntryPoint) middleware {
	return func(_ context.Context, op *UserOperation) error {
		slog.Debug("getting nonce")
		key := new(big.Int)
		nonce, err := entryPoint.GetNonce(nil, op.Sender, key)
		if err != nil {
			return err
		}

		op.Nonce = decimal.NewFromBigInt(nonce, 0)
		return nil
	}
}

func getInitCodeMiddleware(provider EthBackend, smartWalletConfig smart_wallet.Config) (middleware, error) {
	return func(ctx context.Context, op *UserOperation) error {
		owner, ok := ctx.Value(ctxKeyOwner).(common.Address)
		if !ok {
			return fmt.Errorf("`owner` not found, but required in context to get init code")
		}

		index, ok := ctx.Value(ctxKeyIndex).(decimal.Decimal)
		if !ok {
			return fmt.Errorf("`index` not found, but required in context to get init code")
		}

		// if Smart Wallet is already deployed - return empty init code
		if isDeployed, err := smart_wallet.IsAccountDeployed(ctx, provider, op.Sender); err != nil {
			return fmt.Errorf("failed to check if smart account is already deployed: %w", err)
		} else if isDeployed {
			return nil
		}

		initCode, err := smart_wallet.GetInitCode(smartWalletConfig, owner, index)
		if err != nil {
			return fmt.Errorf("failed to get init code: %w", err)
		}

		op.InitCode = initCode
		return nil
	}, nil
}

func getGasPricesMiddleware(provider EthBackend, gasConfig GasConfig) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		if !op.MaxFeePerGas.IsZero() && !op.MaxPriorityFeePerGas.IsZero() {
			slog.Debug("skipping gas price estimation, using provided gas prices")
			return nil
		}

		slog.Debug("getting gas prices")

		// Calculate maxPriorityFeePerGas
		var maxPriorityFeePerGas *big.Int
		if !op.MaxPriorityFeePerGas.IsZero() {
			maxPriorityFeePerGas = op.MaxPriorityFeePerGas.BigInt()
		} else {
			var maxPriorityFeePerGasStr string
			if err := provider.RPC().CallContext(ctx, &maxPriorityFeePerGasStr, "eth_maxPriorityFeePerGas"); err != nil {
				return err
			}

			var ok bool
			maxPriorityFeePerGas, ok = new(big.Int).SetString(maxPriorityFeePerGasStr, 0)
			if !ok {
				return fmt.Errorf("failed to parse maxPriorityFeePerGas: %s", maxPriorityFeePerGasStr)
			}

			// Increase maxPriorityFeePerGas to give user more
			// flexibility in setting the gas price.
			maxPriorityFeePerGas.Mul(
				maxPriorityFeePerGas,
				gasConfig.MaxPriorityFeePerGasMultiplier.BigInt())
		}

		// Calculate maxFeePerGas
		var maxFeePerGas *big.Int
		if !op.MaxFeePerGas.IsZero() {
			maxFeePerGas = op.MaxFeePerGas.BigInt()
		} else {
			// Get the latest block to read its base fee
			block, err := provider.BlockByNumber(ctx, nil)
			if err != nil {
				return err
			}
			blockBaseFee := block.BaseFee()
			slog.Debug("block base fee", "baseFee", blockBaseFee.String())

			maxFeePerGas = new(big.Int).Mul(blockBaseFee, gasConfig.MaxFeePerGasMultiplier.BigInt())
			maxFeePerGas.Add(maxFeePerGas, maxPriorityFeePerGas)
		}

		slog.Debug("calculated gas price", "maxFeePerGas", maxFeePerGas, "maxPriorityFeePerGas", maxPriorityFeePerGas)

		op.MaxFeePerGas = decimal.NewFromBigInt(maxFeePerGas, 0)
		op.MaxPriorityFeePerGas = decimal.NewFromBigInt(maxPriorityFeePerGas, 0)

		return nil
	}
}

func getBiconomyPaymasterAndData(
	bundler RPCBackend,
	paymasterCtx map[string]any,
	entryPoint common.Address,
) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		opModified := struct {
			Sender               common.Address  `json:"sender"`
			Nonce                decimal.Decimal `json:"nonce"`
			InitCode             string          `json:"initCode"`
			CallData             string          `json:"callData"`
			CallGasLimit         decimal.Decimal `json:"callGasLimit"`
			VerificationGasLimit decimal.Decimal `json:"verificationGasLimit"`
			PreVerificationGas   decimal.Decimal `json:"preVerificationGas"`
			MaxFeePerGas         decimal.Decimal `json:"maxFeePerGas"`
			MaxPriorityFeePerGas decimal.Decimal `json:"maxPriorityFeePerGas"`
			PaymasterAndData     string          `json:"paymasterAndData"`
			Signature            string          `json:"signature,omitempty"`
		}{
			Sender:               op.Sender,
			Nonce:                op.Nonce,
			InitCode:             hexutil.Encode(op.InitCode),
			CallData:             hexutil.Encode(op.CallData),
			CallGasLimit:         op.CallGasLimit,
			VerificationGasLimit: op.VerificationGasLimit,
			PreVerificationGas:   op.PreVerificationGas,
			MaxFeePerGas:         op.MaxFeePerGas,
			MaxPriorityFeePerGas: op.MaxPriorityFeePerGas,
			PaymasterAndData:     hexutil.Encode(op.PaymasterAndData),
			Signature:            hexutil.Encode(op.Signature),
		}

		// Biconomy-standardized gas estimation with paymaster
		// see https://docs.biconomy.io/Paymaster/api/sponsor-useroperation#sponsorship-paymaster
		var est gasEstimate
		if err := bundler.CallContext(ctx, &est, "pm_sponsorUserOperation", opModified, entryPoint, paymasterCtx); err != nil {
			return fmt.Errorf("failed to call pm_sponsorUserOperation: %w", err)
		}

		callGasLimit, verificationGasLimit, preVerificationGas, err := est.convert()
		if err != nil {
			return fmt.Errorf("failed to convert gas estimates: %w", err)
		}

		op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0)
		op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
		op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)

		paymasterAndData, err := hexutil.Decode(est.PaymasterAndData)
		if err != nil {
			return fmt.Errorf("failed to decode paymasterAndData: %w", err)
		}
		op.PaymasterAndData = paymasterAndData

		return nil
	}
}

func getGasLimitsMiddleware(bundler RPCBackend, config ClientConfig) (middleware, error) {
	estimateGas := estimateUserOperationGas(bundler, config.EntryPoint)

	if config.Paymaster.Type != nil && *config.Paymaster.Type != PaymasterDisabled {
		switch typ := *config.Paymaster.Type; typ {
		case PaymasterPimlicoERC20:
			estimateGas = getPimlicoERC20PaymasterData(
				bundler,
				config.EntryPoint,
				config.Paymaster.Address,
				config.Paymaster.PimlicoERC20.VerificationGasOverhead,
			)
		case PaymasterPimlicoVerifying:
			// NOTE: PimlicoVerifying is the easiest to implement
			return nil, ErrPaymasterNotSupported
		case PaymasterBiconomyERC20:
			return nil, ErrPaymasterNotSupported
		case PaymasterBiconomySponsoring:
			// NOTE: tried to add BiconomySponsoring, but it is not responding correctly
			return nil, ErrPaymasterNotSupported
		default:
			return nil, fmt.Errorf("unknown paymaster type: %s", typ)
		}
	}

	return estimateGas, nil
}

func estimateUserOperationGas(bundler RPCBackend, entryPoint common.Address) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		if !op.CallGasLimit.IsZero() && !op.VerificationGasLimit.IsZero() && !op.PreVerificationGas.IsZero() {
			slog.Debug("skipping gas estimation, using provided gas limits")
			return nil
		}

		slog.Debug("estimating userOp gas limits", "userOp", op)

		// ERC4337-standardized gas estimation
		var est gasEstimate
		if err := bundler.CallContext(
			ctx,
			&est,
			"eth_estimateUserOperationGas",
			op,
			entryPoint,
		); err != nil {
			return fmt.Errorf("error estimating gas: %w", err)
		}

		preVerificationGas, verificationGasLimit, callGasLimit, err := est.convert()
		if err != nil {
			return fmt.Errorf("failed to convert gas estimates: %w", err)
		}

		if !op.CallGasLimit.IsZero() {
			callGasLimit = op.CallGasLimit.BigInt()
		}
		if !op.VerificationGasLimit.IsZero() {
			verificationGasLimit = op.VerificationGasLimit.BigInt()
		}
		if !op.PreVerificationGas.IsZero() {
			preVerificationGas = op.PreVerificationGas.BigInt()
		}

		slog.Debug("estimated userOp gas", "callGasLimit", callGasLimit, "verificationGasLimit", verificationGasLimit, "preVerificationGas", preVerificationGas)

		op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)
		op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
		op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0)

		return nil
	}
}

func getPimlicoERC20PaymasterData(
	bundler RPCBackend,
	entryPoint common.Address,
	paymaster common.Address,
	verificationGasOverhead decimal.Decimal,
) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		estimate := estimateUserOperationGas(bundler, entryPoint)
		if err := estimate(ctx, op); err != nil {
			return err
		}

		op.VerificationGasLimit = op.VerificationGasLimit.Add(verificationGasOverhead)
		op.PaymasterAndData = paymaster.Bytes()

		return nil
	}
}

// gasEstimate holds gas estimates for a user operation.
type gasEstimate struct {
	// depending on provider, any of the following types can be received here: string, int
	CallGasLimit         any `json:"callGasLimit"`
	VerificationGasLimit any `json:"verificationGasLimit"`
	PreVerificationGas   any `json:"preVerificationGas"`

	PaymasterAndData string `json:"paymasterAndData,omitempty"`
}

func (est gasEstimate) convert() (
	preVerificationGas *big.Int,
	verificationGasLimit *big.Int,
	callGasLimit *big.Int,
	err error,
) {
	preVerificationGas, err = est.fromAny(est.PreVerificationGas)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("preVerificationGas: %w", err)
	}

	verificationGasLimit, err = est.fromAny(est.VerificationGasLimit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("verificationGasLimit: %w", err)
	}

	callGasLimit, err = est.fromAny(est.CallGasLimit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("callGasLimit: %w", err)
	}

	return preVerificationGas,
		verificationGasLimit,
		callGasLimit,
		nil
}

func (est gasEstimate) fromAny(a any) (*big.Int, error) {
	switch v := a.(type) {
	case string:
		n, err := strconv.ParseInt(v, 16, 64)
		if err != nil { // it IS hexadecimal
			nBig, err := hexutil.DecodeBig(v)
			if err != nil {
				return nil, fmt.Errorf("failed to parse: %w (got '%s')", err, v)
			}
			return nBig, nil
		}

		return new(big.Int).SetInt64(n), nil
	case int64:
		return new(big.Int).SetInt64(v), nil
	case float64:
		return new(big.Int).SetInt64(int64(v)), nil
	default:
		return nil, fmt.Errorf("unexpected type: %T", v)
	}
}

func getSignMiddleware(entryPoint common.Address, chainID *big.Int) middleware {
	return func(ctx context.Context, op *UserOperation) error {
		signer, ok := ctx.Value(ctxKeySigner).(Signer)
		if !ok {
			return fmt.Errorf("signer not found in context")
		}

		signature, err := signer(*op, entryPoint, chainID)
		if err != nil {
			return fmt.Errorf("failed to sign user operation: %w", err)
		}

		op.Signature = signature

		b, err := op.MarshalJSON()
		if err != nil {
			return fmt.Errorf("failed to marshal user operation: %w", err)
		}
		slog.Debug("userop signed",
			"hash", op.UserOpHash(entryPoint, chainID).String(),
			"json", string(b))

		return nil
	}
}
