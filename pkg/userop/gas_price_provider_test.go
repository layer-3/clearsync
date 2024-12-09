package userop

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPolygonGasPrices(t *testing.T) {
	tests := []struct {
		name    string
		chainId *big.Int
		wantErr bool
	}{
		{
			name:    "should return gas prices for Polygon",
			chainId: big.NewInt(137),
			wantErr: false,
		},
		{
			name:    "should return gas prices for Amoy",
			chainId: big.NewInt(80002),
			wantErr: false,
		},
		{
			name:    "should return error for other chain",
			chainId: big.NewInt(42),
			wantErr: true,
		},
		{
			name:    "should return error if chainId is nil",
			chainId: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := NewPolygonGasPriceProvider(tt.chainId)

			ctx := context.Background()
			maxFee, maxPriorityFee, err := provider.GetGasPrices(ctx)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, maxFee)
				require.NotNil(t, maxPriorityFee)
			}
		})
	}
}
