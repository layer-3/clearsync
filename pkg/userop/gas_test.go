package userop

import (
	"context"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetGasPricesMiddleware(t *testing.T) {
	t.Run("gas multipliers are applied correctly", func(t *testing.T) {
		mockedMaxFeePerGas := big.NewInt(400200400200)
		mockedMaxPriorityFeePerGas := big.NewInt(42424242)
		mockedGasPricesProvider := NewMockGasPriceProvider(mockedMaxFeePerGas, mockedMaxPriorityFeePerGas)

		tests := []struct {
			name                     string
			gasMultipliers           GasConfig
			wantMaxFeePerGas         *big.Int
			wantMaxPriorityFeePerGas *big.Int
		}{
			{
				name:                     "no multipliers supplied result in 0 gas prices",
				gasMultipliers:           GasConfig{},
				wantMaxFeePerGas:         big.NewInt(0),
				wantMaxPriorityFeePerGas: big.NewInt(0),
			},
			{
				name: "1 multipliers do not change gas prices",
				gasMultipliers: GasConfig{
					MaxPriorityFeePerGasMultiplier: decimal.NewFromFloat(1),
					MaxFeePerGasMultiplier:         decimal.NewFromFloat(1),
				},
				wantMaxFeePerGas:         mockedMaxFeePerGas,
				wantMaxPriorityFeePerGas: mockedMaxPriorityFeePerGas,
			},
			{
				name: "1.5 multipliers are applied correctly",
				gasMultipliers: GasConfig{
					MaxPriorityFeePerGasMultiplier: decimal.NewFromFloat(1.5),
					MaxFeePerGasMultiplier:         decimal.NewFromFloat(1.5),
				},
				wantMaxFeePerGas:         big.NewInt(600300600300),
				wantMaxPriorityFeePerGas: big.NewInt(63636363),
			},
			{
				name: "2.25 multipliers are applied correctly",
				gasMultipliers: GasConfig{
					MaxPriorityFeePerGasMultiplier: decimal.NewFromFloat(2.25),
					MaxFeePerGasMultiplier:         decimal.NewFromFloat(2.25),
				},
				wantMaxFeePerGas:         big.NewInt(900450900450),
				wantMaxPriorityFeePerGas: big.NewInt(95454544),
			},
		}

		ctx := context.Background()

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotMaxFeePerGas, gotMaxPriorityFeePerGas, err := getGasPricesAndApplyMultipliers(ctx, mockedGasPricesProvider, tt.gasMultipliers)
				require.NoError(t, err)

				require.True(t, gotMaxFeePerGas.Cmp(tt.wantMaxFeePerGas) == 0)
				require.True(t, gotMaxPriorityFeePerGas.Cmp(tt.wantMaxPriorityFeePerGas) == 0)
			})
		}
	})
}
