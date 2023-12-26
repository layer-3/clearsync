package quotes

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	prices = []int32{
		40000, 42000, 41500, 44000,
		43000, 40000, 41000, 42000,
		43000, 42000, 45500, 41000,
		41500, 42000, 44000, 46000,
		47000, 46000, 44000, 42000,
		100000,
	}
	amounts = []float32{
		1, 1.1, 2.4, 0.2,
		2, 3.3, 4.0, 2.9,
		1, 0.1, 0.01, 0.04,
		9, 0.4, 4.4, 5,
		6, 0.1, 2, 1, 1,
	}
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
			emas = append(emas, EMA20(emas[i], price))
		}

		var result []float64
		for _, ema := range emas {
			result = append(result, ema.Round(5).InexactFloat64())
		}

		exp := []float64{40000, 40000, 40190.47619, 40315.19274, 40666.12677, 40888.40041, 40803.79085, 40822.47743, 40934.62244, 41131.32506, 41214.05601, 41622.24115, 41562.98009, 41556.98199, 41599.17418, 41827.82426, 42225.17433, 42679.91963, 42996.11776, 43091.72559, 42987.75173, 48417.48966}
		require.Equal(t, exp, result)
	})
}

func Test_IndexAggregator(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		ag := &IndexAggregator{
			priceCache: NewPriceCache(DefaultWeightsMap),
			weights:    DefaultWeightsMap,
		}
		ag.weights[DriverBinance] = decimal.NewFromInt(3)
		ag.weights[DriverUniswapV3] = decimal.NewFromInt(0)

		var decimalPrices []decimal.Decimal
		var decimalAmounts []decimal.Decimal
		for i, p := range prices {
			decimalPrices = append(decimalPrices, decimal.NewFromInt32(p))
			decimalAmounts = append(decimalAmounts, decimal.NewFromFloat32(amounts[i]))
		}

		var inputTrades []TradeEvent
		trade := TradeEvent{Source: DriverBinance, Market: "btcusdt"}
		for i, p := range decimalPrices {
			trade.Price = p
			trade.Amount = decimalAmounts[i]
			if i == 20 {
				trade.Source = DriverUniswapV3
			}
			inputTrades = append(inputTrades, trade)
		}

		var result []float64
		for _, tr := range inputTrades {
			res := ag.indexPrice(tr)
			result = append(result, res.Price.Round(5).InexactFloat64())
		}

		exp := []float64{40000, 40207.54717, 40466.24305, 40530.21098, 40941.96994, 40722.34661, 40788.43173, 40982.49411, 41098.56987, 41104.26589, 41107.33374, 41107.00358, 41277.3339, 41292.39911, 41839.88457, 42682.42649, 43596.62893, 43605.96679, 43637.13205, 43568.57095, 43568.57095}
		require.Equal(t, exp, result)
	})
}
