package quotes

import (
	"testing"
	"time"

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
		6, 0.1, 2, 1,
		1,
	}

	defaultWeights = map[DriverType]decimal.Decimal{
		DriverBinance:       decimal.NewFromInt(2),
		DriverUniswapV3: decimal.NewFromInt(2),
	}
)

func Test_IndexAggregatorStrategies(t *testing.T) {
	btcusdt := NewMarket("btc", "usdt")

	t.Run("Successful test", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:       decimal.NewFromInt(3),
			DriverUniswapV3: decimal.NewFromInt(0),
		}

		trade := TradeEvent{Source: DriverBinance, Market: btcusdt}
		var inputTrades []TradeEvent
		for i, p := range prices {
			decimalPrice := decimal.NewFromInt32(p)
			decimalAmount := decimal.NewFromFloat32(amounts[i])

			trade.Price = decimalPrice
			trade.Amount = decimalAmount
			if i == 20 {
				trade.Source = DriverUniswapV3
			}
			inputTrades = append(inputTrades, trade)
		}

		results := testStrategies(inputTrades, newStrategyVWA(WithCustomWeightsVWA(weights)))

		// Check VWA strategy
		expVWA := []float64{40000, 41047.61905, 41288.88889, 41404.25532, 41880.59701, 41260, 41185.71429, 41325.44379, 41418.99441, 41422.22222, 41424.4864, 41423.54571, 41448.98336, 41457.01275, 41808.32025, 42377.0692, 43024.3874, 43031.31548, 43074.41602, 43051.03373}
		require.Equal(t, expVWA, results[0])
	})

	t.Run("Skip trades with zero price or amount", func(t *testing.T) {
		inputTrades := []TradeEvent{
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.Zero, Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(44000), Amount: decimal.Zero},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(44000), Amount: decimal.NewFromFloat(1)},
		}

		results := testStrategies(inputTrades, newStrategyVWA(WithCustomWeightsVWA(defaultWeights)))

		// Check VWA strategy
		expVWA := []float64{40000, 41000, 42000}
		require.Equal(t, expVWA, results[0])
	})

	inputTrades := []TradeEvent{
		{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(0.3)},
		{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(42500), Amount: decimal.NewFromFloat(0.5)},
		{Source: DriverUniswapV3, Market: btcusdt, Price: decimal.NewFromInt(55000), Amount: decimal.NewFromFloat(0.6)},
		{Source: DriverUniswapV3, Market: btcusdt, Price: decimal.NewFromInt(50000), Amount: decimal.NewFromFloat(0.4)},
		{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1)},
	}

	t.Run("README example 1: equal driver weight", func(t *testing.T) {
		results := testStrategies(inputTrades, newStrategyVWA(WithCustomWeightsVWA(defaultWeights)))

		// Check VWA strategy
		expVWA := []float64{41000, 41937.5, 45500, 46192.30769, 44472.22222}
		require.Equal(t, expVWA, results[0])
	})

	t.Run("README example 2: zero weight for one of the drivers", func(t *testing.T) {
		weights := map[DriverType]decimal.Decimal{
			DriverBinance:       decimal.NewFromInt(3),
			DriverUniswapV3: decimal.NewFromInt(0),
		}

		results := testStrategies(inputTrades, newStrategyVWA(WithCustomWeightsVWA(weights)))

		// Check VWA strategy
		expVWA := []float64{41000, 41937.5, 40861.11111}
		require.Equal(t, expVWA, results[0])
	})

	t.Run("README example 3: trade volume", func(t *testing.T) {
		inputTrades := []TradeEvent{
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(44000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(46000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(48000), Amount: decimal.NewFromFloat(10)},
		}

		results := testStrategies(inputTrades, newStrategyVWA(WithCustomWeightsVWA(defaultWeights)))

		// Check VWA strategy
		expVWA := []float64{40000, 41000, 42000, 43000, 46571.42857}
		require.Equal(t, expVWA, results[0])
	})

	t.Run("README example 4: drivers volatility", func(t *testing.T) {
		// Initial price: 41000
		inputTrades := []TradeEvent{{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromInt(1.0)}}
		// Two equal drivers are sending: 42000 and 40000 prices sequentially.
		inputTrades = append(inputTrades, generateTrades([]TradeEvent{
			{Source: DriverUniswapV3, Market: btcusdt, Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)},
		}, 25)...)
		// The drivers start sending the same price: 41000.
		inputTrades = append(inputTrades, generateTrades([]TradeEvent{
			{Source: DriverUniswapV3, Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(1.0)},
			{Source: DriverBinance, Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(1.0)},
		}, 25)...)

		testPriceCacheVWA := newPriceCacheVWA(defaultWeights, 20, time.Minute)
		testPriceCacheVWA.ActivateDriver(DriverBinance, btcusdt)
		testPriceCacheVWA.ActivateDriver(DriverUniswapV3, btcusdt)

		results := testStrategies(inputTrades, newStrategyVWA(WithCustomWeightsVWA(defaultWeights), withCustomPriceCacheVWA(testPriceCacheVWA)))

		// Check VWA strategy
		require.Equal(t, float64(41000), results[0][0])
		require.Equal(t, float64(41000), results[0][25])
		require.Equal(t, float64(41000), results[0][50])
	})
}

func generateTrades(tr []TradeEvent, n int) []TradeEvent {
	trades := []TradeEvent{}
	for i := 0; i < n; i++ {
		trades = append(trades, tr...)
	}
	return trades
}

func testStrategies(inputTrades []TradeEvent, priceCalculators ...priceCalculator) [][]float64 {
	results := make([][]float64, len(priceCalculators))

	for i, pc := range priceCalculators {
		for _, tr := range inputTrades {
			res, ok := pc.calculateIndexPrice(tr)
			if ok {
				results[i] = append(results[i], res.Round(5).InexactFloat64())
			}
		}
	}

	return results
}

func TestIsPriceOutOfRange(t *testing.T) {
	testCases := []struct {
		name         string
		eventPrice   decimal.Decimal
		lastPrice    decimal.Decimal
		maxPriceDiff decimal.Decimal
		isOutOfRange bool
	}{
		// incoming event price | last price | max price diff | expected
		{"Price exactly at upper bound", decimal.NewFromFloat(120), decimal.NewFromFloat(100), decimal.NewFromFloat(0.20), false},
		{"Price exactly at lower bound", decimal.NewFromFloat(80), decimal.NewFromFloat(100), decimal.NewFromFloat(0.20), false},
		{"Price above upper bound", decimal.NewFromFloat(121), decimal.NewFromFloat(100), decimal.NewFromFloat(0.20), true},
		{"Price below lower bound", decimal.NewFromFloat(79), decimal.NewFromFloat(100), decimal.NewFromFloat(0.20), true},
		{"Price within range", decimal.NewFromFloat(110), decimal.NewFromFloat(100), decimal.NewFromFloat(0.20), false},
		{"Price within negative range", decimal.NewFromFloat(90), decimal.NewFromFloat(100), decimal.NewFromFloat(0.20), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPriceOutOfRange(tc.eventPrice, tc.lastPrice, tc.maxPriceDiff)
			if result != tc.isOutOfRange {
				t.Errorf("Test %s failed. Expected %v, got %v", tc.name, tc.isOutOfRange, result)
			}
		})
	}
}
