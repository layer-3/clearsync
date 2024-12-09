package userop

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/shopspring/decimal"
)

type GasPriceProvider interface {
	GetGasPrices(ctx context.Context) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error)
}

type EVMGasPriceProvider struct {
	provider EthBackend
}

func NewEVMGasPriceProvider(provider EthBackend) *EVMGasPriceProvider {
	return &EVMGasPriceProvider{provider: provider}
}

func (p *EVMGasPriceProvider) GetGasPrices(ctx context.Context) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	var maxPriorityFeePerGasStr string
	if err := p.provider.RPC().CallContext(ctx, &maxPriorityFeePerGasStr, "eth_maxPriorityFeePerGas"); err != nil {
		return nil, nil, err
	}

	maxPriorityFeePerGas, ok := new(big.Int).SetString(maxPriorityFeePerGasStr, 0)
	if !ok {
		return nil, nil, fmt.Errorf("failed to parse maxPriorityFeePerGas: %s", maxPriorityFeePerGasStr)
	}
	logger.Debug("fetched maxPriorityFeePerGas", "maxPriorityFeePerGas", maxPriorityFeePerGas.String())

	// Get the latest block to read its base fee
	block, err := p.provider.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	blockBaseFee := block.BaseFee()
	logger.Debug("fetched block base fee", "baseFee", blockBaseFee.String())

	maxFeePerGas = blockBaseFee.Add(blockBaseFee, maxPriorityFeePerGas)

	return maxFeePerGas, maxPriorityFeePerGas, nil
}

type PolygonGasPriceProvider struct {
	chainId *big.Int
}

func NewPolygonGasPriceProvider(chainId *big.Int) *PolygonGasPriceProvider {
	return &PolygonGasPriceProvider{chainId: chainId}
}

func (p *PolygonGasPriceProvider) GetGasPrices(ctx context.Context) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	var resp *http.Response

	if p.chainId == nil {
		return nil, nil, fmt.Errorf("chain ID is nil")
	}

	switch {
	case p.chainId.Uint64() == 137:
		resp, err = http.Get("https://gasstation.polygon.technology/v2")
	case p.chainId.Uint64() == 80002:
		resp, err = http.Get("https://gasstation.polygon.technology/amoy")
	default:
		return nil, nil, fmt.Errorf("unsupported chain ID: %v", p.chainId)
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

	maxFeePerGas = gasData.Fast.MaxFee.Mul(gweiMult).BigInt()
	maxPriorityFeePerGas = gasData.Fast.MaxPriorityFee.Mul(gweiMult).BigInt()

	return maxFeePerGas, maxPriorityFeePerGas, nil
}

type MockGasPriceProvider struct {
	maxFeePerGas         *big.Int
	maxPriorityFeePerGas *big.Int
}

func NewMockGasPriceProvider(maxFeePerGas, maxPriorityFeePerGas *big.Int) *MockGasPriceProvider {
	return &MockGasPriceProvider{maxFeePerGas: maxFeePerGas, maxPriorityFeePerGas: maxPriorityFeePerGas}
}

func (p *MockGasPriceProvider) GetGasPrices(ctx context.Context) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	return p.maxFeePerGas, p.maxPriorityFeePerGas, nil
}
