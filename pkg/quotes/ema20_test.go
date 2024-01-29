package quotes

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func Test_EMA20(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		var decimalPrices []decimal.Decimal
		for _, p := range prices {
			decimalPrices = append(decimalPrices, decimal.NewFromInt32(p))
		}

		var emas []decimal.Decimal
		emas = append(emas, decimalPrices[0])

		for i, price := range decimalPrices {
			emas = append(emas, EMA(emas[i], price, 20))
		}

		var result []float64
		for _, ema := range emas {
			result = append(result, ema.Round(5).InexactFloat64())
		}

		exp := []float64{40000, 40000, 40190.47619, 40315.19274, 40666.12677, 40888.40041, 40803.79085, 40822.47743, 40934.62244, 41131.32506, 41214.05601, 41622.24115, 41562.98009, 41556.98199, 41599.17418, 41827.82426, 42225.17433, 42679.91963, 42996.11776, 43091.72559, 42987.75173, 48417.48966}
		require.Equal(t, exp, result)
	})
}
