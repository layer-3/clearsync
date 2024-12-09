package userop

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

func getGasPricesAndApplyMultipliers(ctx context.Context, provider GasPriceProvider, gasConfig GasConfig) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	logger.Debug("getting gas prices")

	maxFeePerGas, maxPriorityFeePerGas, err = provider.GetGasPrices(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get gas prices: %w", err)
	}

	logger.Debug("fetched gas price", "maxFeePerGas", maxFeePerGas, "maxPriorityFeePerGas", maxPriorityFeePerGas)

	maxFeePerGas = decimal.NewFromBigInt(maxFeePerGas, 0).Mul(gasConfig.MaxFeePerGasMultiplier).BigInt()
	maxPriorityFeePerGas = decimal.NewFromBigInt(maxPriorityFeePerGas, 0).Mul(gasConfig.MaxPriorityFeePerGasMultiplier).BigInt()

	logger.Debug("calculated gas price", "maxFeePerGas", maxFeePerGas, "maxPriorityFeePerGas", maxPriorityFeePerGas)

	return maxFeePerGas, maxPriorityFeePerGas, nil
}

// GasEstimate holds gas estimates for a user operation.
type GasEstimate struct {
	// depending on provider, any of the following types can be received here: string, int
	CallGasLimit         any `json:"callGasLimit"`
	VerificationGasLimit any `json:"verificationGasLimit"`
	PreVerificationGas   any `json:"preVerificationGas"`

	PaymasterAndData string `json:"paymasterAndData,omitempty"`
}

func (est GasEstimate) convert() (
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

func (est GasEstimate) fromAny(a any) (*big.Int, error) {
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

// overwriteGasLimitsIfUnset applies gas limits if they are not already set.
func overwriteGasLimitsIfUnset(
	est GasEstimate,
	op *UserOperation,
) error {
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

	logger.Debug("estimated userOp gas", "callGasLimit", callGasLimit, "verificationGasLimit", verificationGasLimit, "preVerificationGas", preVerificationGas)

	op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0)
	op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
	op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)

	return nil
}

// overwriteGasLimits overwrites gas limits with the ones from the estimate.
func overwriteGasLimits(
	est GasEstimate,
	op *UserOperation,
) error {
	preVerificationGas, verificationGasLimit, callGasLimit, err := est.convert()
	if err != nil {
		return fmt.Errorf("failed to convert gas estimates: %w", err)
	}

	logger.Debug("estimated userOp gas", "callGasLimit", callGasLimit, "verificationGasLimit", verificationGasLimit, "preVerificationGas", preVerificationGas)

	op.CallGasLimit = decimal.NewFromBigInt(callGasLimit, 0)
	op.VerificationGasLimit = decimal.NewFromBigInt(verificationGasLimit, 0)
	op.PreVerificationGas = decimal.NewFromBigInt(preVerificationGas, 0)

	return nil
}
