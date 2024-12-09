package userop

import (
	"context"
	"math/big"
	"net/url"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetGasPricesMiddleware(t *testing.T) {
	t.Run("gas multipliers are applied correctly", func(t *testing.T) {
		t.Skip("manual test")
		providerUrl := "provider_url"

		providerURL, err := url.Parse(providerUrl)
		require.NoError(t, err)
		providerRPC, err := NewEthBackend(*providerURL)
		require.NoError(t, err)

		tests := []struct {
			name           string
			gasMultipliers GasConfig
		}{
			{
				name:           "no multipliers supplied result in 0 gas prices",
				gasMultipliers: GasConfig{},
			},
			{
				name: "1 multipliers do not change gas prices",
				gasMultipliers: GasConfig{
					MaxPriorityFeePerGasMultiplier: decimal.NewFromFloat(1),
					MaxFeePerGasMultiplier:         decimal.NewFromFloat(1),
				},
			},
			{
				name: "1.5 multipliers are applied correctly",
				gasMultipliers: GasConfig{
					MaxPriorityFeePerGasMultiplier: decimal.NewFromFloat(1.5),
					MaxFeePerGasMultiplier:         decimal.NewFromFloat(1.5),
				},
			},
			{
				name: "2.25 multipliers are applied correctly",
				gasMultipliers: GasConfig{
					MaxPriorityFeePerGasMultiplier: decimal.NewFromFloat(2.25),
					MaxFeePerGasMultiplier:         decimal.NewFromFloat(2.25),
				},
			},
		}

		ctx := context.Background()
		maxFeePerGas, maxPriorityFeePerGas, err := getGasPrices(ctx, providerRPC)
		require.NoError(t, err)

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotMaxFeePerGas, gotMaxPriorityFeePerGas, err := getGasPricesAndApplyMultipliers(ctx, providerRPC, tt.gasMultipliers)
				require.NoError(t, err)

				calculatedMaxFeePerGas := decimal.NewFromBigInt(maxFeePerGas, 0).Mul(tt.gasMultipliers.MaxFeePerGasMultiplier).BigInt()
				require.True(t, gotMaxFeePerGas.Cmp(calculatedMaxFeePerGas) == 0)
				calculatedMaxPriorityFeePerGas := decimal.NewFromBigInt(maxPriorityFeePerGas, 0).Mul(tt.gasMultipliers.MaxPriorityFeePerGasMultiplier).BigInt()
				require.True(t, gotMaxPriorityFeePerGas.Cmp(calculatedMaxPriorityFeePerGas) == 0)
			})
		}
	})
}

func TestGetPolygonGasPrices(t *testing.T) {
	t.Run("Should return gas prices for Polygon", func(t *testing.T) {
		chainId := big.NewInt(137)
		maxFee, maxPriorityFee, err := getPolygonGasPrices(chainId)
		require.NoError(t, err)
		require.NotNil(t, maxFee)
		require.NotNil(t, maxPriorityFee)
		// Gas station states that for Polygon Mainnet the maxPriorityFee is at least 30 Gwei
		require.True(t, maxPriorityFee.Cmp(big.NewInt(30).Exp(big.NewInt(10), big.NewInt(9), nil)) >= 0)
	})

	t.Run("Should return gas prices for Amoy", func(t *testing.T) {
		chainId := big.NewInt(80002)
		maxFee, maxPriorityFee, err := getPolygonGasPrices(chainId)
		require.NoError(t, err)
		require.NotNil(t, maxFee)
		require.NotNil(t, maxPriorityFee)
	})

	t.Run("Should return error for other chain", func(t *testing.T) {
		chainId := big.NewInt(42)
		_, _, err := getPolygonGasPrices(chainId)
		require.Error(t, err)
	})

	t.Run("Should return error if chainId is nil", func(t *testing.T) {
		_, _, err := getPolygonGasPrices(nil)
		require.Error(t, err)
	})
}
