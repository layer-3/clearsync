package quotes

import (
	"fmt"
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
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:      decimal.NewFromInt(3),
			DriverUniswapV3Api: decimal.NewFromInt(0),
		}

		ag := &indexAggregator{
			priceCache: NewPriceCache(weights),
			weights:    weights,
		}

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
				trade.Source = DriverUniswapV3Api
			}
			inputTrades = append(inputTrades, trade)
		}

		var result []float64
		for _, tr := range inputTrades {
			res, ok := ag.indexPrice(tr)
			if ok {
				result = append(result, res.Price.Round(5).InexactFloat64())
			}
		}

		exp := []float64{40000, 40207.54717, 40466.24305, 40530.21098, 40941.96994, 40722.34661, 40788.43173, 40982.49411, 41098.56987, 41104.26589, 41107.33374, 41107.00358, 41277.3339, 41292.39911, 41839.88457, 42682.42649, 43596.62893, 43605.96679, 43637.13205, 43568.57095, 43568.57095}
		require.Equal(t, exp, result)
	})

	inputTrades := []TradeEvent{
		{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(0.3)},
		{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(42500), Amount: decimal.NewFromFloat(0.5)},
		{Source: DriverUniswapV3Api, Market: "btcusdt", Price: decimal.NewFromInt(55000), Amount: decimal.NewFromFloat(0.6)},
		{Source: DriverUniswapV3Api, Market: "btcusdt", Price: decimal.NewFromInt(50000), Amount: decimal.NewFromFloat(0.4)},
		{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1)},
	}

	t.Run("Skip trades with zero price or amount", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:      decimal.NewFromInt(2),
			DriverUniswapV3Api: decimal.NewFromInt(2),
		}

		ag := &indexAggregator{
			priceCache: NewPriceCache(weights),
			weights:    weights,
		}

		inputTrades := []TradeEvent{
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.Zero, Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(44000), Amount: decimal.Zero},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(44000), Amount: decimal.NewFromFloat(1)},
		}

		var result []float64
		for _, tr := range inputTrades {
			res, ok := ag.indexPrice(tr)
			if ok {
				result = append(result, res.Price.Round(5).InexactFloat64())
			}
		}

		exp := []float64{40000, 40190.47619, 40553.28798}
		require.Equal(t, exp, result)
	})

	t.Run("README example 1: equal driver weight", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:      decimal.NewFromInt(2),
			DriverUniswapV3Api: decimal.NewFromInt(2),
		}

		ag := &indexAggregator{
			priceCache: NewPriceCache(weights),
			weights:    weights,
		}

		var result []float64
		for _, tr := range inputTrades {
			res, ok := ag.indexPrice(tr)
			if ok {
				result = append(result, res.Price.Round(5).InexactFloat64())
			}
		}

		exp := []float64{41000, 41223.8806, 42464.61758, 42933.56853, 42503.12993}
		require.Equal(t, exp, result)
	})

	t.Run("README example 2: zero weight for one of the drivers", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:      decimal.NewFromInt(3),
			DriverUniswapV3Api: decimal.NewFromInt(0),
		}

		ag := &indexAggregator{
			priceCache: NewPriceCache(weights),
			weights:    weights,
		}

		var result []float64
		for _, tr := range inputTrades {
			res, ok := ag.indexPrice(tr)
			if ok {
				result = append(result, res.Price.Round(5).InexactFloat64())
			}
		}

		exp := []float64{41000, 41223.8806, 41223.8806, 41223.8806, 40872.3039}
		require.Equal(t, exp, result)
	})

	t.Run("README example 3: trade volume", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:      decimal.NewFromInt(2),
			DriverUniswapV3Api: decimal.NewFromInt(2),
		}

		ag := &indexAggregator{
			priceCache: NewPriceCache(weights),
			weights:    weights,
		}

		inputTrades := []TradeEvent{
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(44000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(46000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(48000), Amount: decimal.NewFromFloat(10)},
		}

		var result []float64
		for _, tr := range inputTrades {
			res, ok := ag.indexPrice(tr)
			if ok {
				result = append(result, res.Price.Round(5).InexactFloat64())
			}
		}

		exp := []float64{40000, 40190.47619, 40553.28798, 41072.02246, 44624.83145}
		require.Equal(t, exp, result)
	})

	t.Run("README example 4: drivers volatility", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:      decimal.NewFromInt(2),
			DriverUniswapV3Api: decimal.NewFromInt(2),
		}

		ag := &indexAggregator{
			priceCache: NewPriceCache(weights),
			weights:    weights,
		}

		ag.priceCache.ActivateDriver(DriverBinance, "btcusdt")
		ag.priceCache.ActivateDriver(DriverUniswapV3Api, "btcusdt")

		// Initial price: 41000
		inputTrades := []TradeEvent{{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(41000), Amount: decimal.NewFromInt(1.0)}}
		// Two equal drivers are sending: 42000 and 40000 prices sequentially.
		inputTrades = append(inputTrades, GenerateTrades([]TradeEvent{
			{Source: DriverUniswapV3Api, Market: "btcusdt", Price: decimal.NewFromInt(44000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(38000), Amount: decimal.NewFromFloat(1.0)}}, 25)...)
		// The drivers start sending the same price: 41000.
		inputTrades = append(inputTrades, GenerateTrades([]TradeEvent{
			{Source: DriverUniswapV3Api, Market: "btcusdt", Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: "btcusdt", Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(1.0)}}, 25)...)

		var result []float64
		for i, tr := range inputTrades {
			res, ok := ag.indexPrice(tr)
			if ok {
				result = append(result, res.Price.Round(0).InexactFloat64())
			}
			fmt.Println(i, res.Price.Round(0))
		}

		exp := []float64{41000, 40950, 41050, 41000}
		// Initial price is 41000.
		require.Equal(t, exp[0], result[0])
		// Price is getting smoothed, and alternates between 41050 and 40950 instead of 40000 and 42000.
		// Another example: if initial prices were 38000 and 44000, smoothed values would be 40850 and 41150.

		require.Equal(t, exp[1], result[48])
		require.Equal(t, exp[2], result[49])
		// Index price when drivers start sending the same price.
		require.Equal(t, exp[3], result[100])
	})
}

func GenerateTrades(tr []TradeEvent, n int) []TradeEvent {
	trades := []TradeEvent{}
	for i := 0; i < n; i++ {
		trades = append(trades, tr...)
	}
	return trades
}
