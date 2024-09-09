package userop

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

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
}
