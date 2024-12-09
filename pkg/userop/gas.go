package userop

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

func getGasPricesAndApplyMultipliers(ctx context.Context, provider EthBackend, gasConfig GasConfig) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	logger.Debug("getting gas prices")

	chainId, err := provider.ChainID(ctx)
	if err != nil || chainId == nil {
		return nil, nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	isPolygon := chainId.Uint64() == 137 || chainId.Uint64() == 80002

	// for Polygon and Amoy, fetch from polygon gas station
	if isPolygon {
		maxFeePerGas, maxPriorityFeePerGas, err = getPolygonGasPrices(chainId)
		if err != nil {
			logger.Error("failed to get gas prices from polygon gas station", "error", err)
		}
	}

	// for other chains, or in case gas station is down, fetch from provider
	if !isPolygon || err != nil {
		maxFeePerGas, maxPriorityFeePerGas, err = getGasPrices(ctx, provider)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get gas prices: %w", err)
		}
	}

	logger.Debug("fetched gas price", "maxFeePerGas", maxFeePerGas, "maxPriorityFeePerGas", maxPriorityFeePerGas)

	maxFeePerGas = decimal.NewFromBigInt(maxFeePerGas, 0).Mul(gasConfig.MaxFeePerGasMultiplier).BigInt()
	maxPriorityFeePerGas = decimal.NewFromBigInt(maxPriorityFeePerGas, 0).Mul(gasConfig.MaxPriorityFeePerGasMultiplier).BigInt()

	logger.Debug("calculated gas price", "maxFeePerGas", maxFeePerGas, "maxPriorityFeePerGas", maxPriorityFeePerGas)

	return maxFeePerGas, maxPriorityFeePerGas, nil
}

func getPolygonGasPrices(chainId *big.Int) (*big.Int, *big.Int, error) {
	var resp *http.Response
	var err error

	if chainId == nil {
		return nil, nil, fmt.Errorf("chain ID is nil")
	}

	switch {
	case chainId.Uint64() == 137:
		resp, err = http.Get("https://gasstation.polygon.technology/v2")
	case chainId.Uint64() == 80002:
		resp, err = http.Get("https://gasstation.polygon.technology/amoy")
	default:
		return nil, nil, fmt.Errorf("unsupported chain ID: %v", chainId)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %v", err)
	}

	var gasData struct {
		Fast struct {
			MaxPriorityFee decimal.Decimal `json:"maxPriorityFee"`
			MaxFee         decimal.Decimal `json:"maxFee"`
		} `json:"fast"`
	}

	err = json.Unmarshal(body, &gasData)
	if err != nil {
		return nil, nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	gweiMult := decimal.NewFromInt(1e9)

	maxFeePerGas := gasData.Fast.MaxFee.Mul(gweiMult).BigInt()
	maxPriorityFeePerGas := gasData.Fast.MaxPriorityFee.Mul(gweiMult).BigInt()

	return maxFeePerGas, maxPriorityFeePerGas, nil
}

func getGasPrices(ctx context.Context, provider EthBackend) (*big.Int, *big.Int, error) {
	var maxPriorityFeePerGasStr string
	if err := provider.RPC().CallContext(ctx, &maxPriorityFeePerGasStr, "eth_maxPriorityFeePerGas"); err != nil {
		return nil, nil, err
	}

	maxPriorityFeePerGas, ok := new(big.Int).SetString(maxPriorityFeePerGasStr, 0)
	if !ok {
		return nil, nil, fmt.Errorf("failed to parse maxPriorityFeePerGas: %s", maxPriorityFeePerGasStr)
	}
	logger.Debug("fetched maxPriorityFeePerGas", "maxPriorityFeePerGas", maxPriorityFeePerGas.String())

	// Get the latest block to read its base fee
	block, err := provider.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	blockBaseFee := block.BaseFee()
	logger.Debug("fetched block base fee", "baseFee", blockBaseFee.String())

	maxFeePerGas := blockBaseFee.Add(blockBaseFee, maxPriorityFeePerGas)

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
