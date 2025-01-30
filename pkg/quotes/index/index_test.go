package index

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/common"
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

	defaultWeights = map[common.DriverType]decimal.Decimal{
		common.DriverBinance:   decimal.NewFromInt(2),
		common.DriverUniswapV3: decimal.NewFromInt(2),
	}
)

func Test_IndexAggregatorStrategies(t *testing.T) {
	btcusdt := common.NewMarket("btc", "usdt")

	t.Run("Successful test", func(t *testing.T) {
		weights := map[common.DriverType]decimal.Decimal{
			common.DriverBinance:   decimal.NewFromInt(3),
			common.DriverUniswapV3: decimal.NewFromInt(0),
		}

		trade := TradeEvent{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt}}
		var inputTrades []TradeEvent
		for i, p := range prices {
			decimalPrice := decimal.NewFromInt32(p)
			decimalAmount := decimal.NewFromFloat32(amounts[i])

			trade.Price = decimalPrice
			trade.Amount = decimalAmount
			if i == 20 {
				trade.Source = common.DriverUniswapV3
			}
			inputTrades = append(inputTrades, trade)
		}

		results := testStrategies(inputTrades, newIndexStrategy(WithCustomWeights(weights)))

		// Check strategy
		exp := []float64{40000, 42000, 41500, 44000, 43000, 40000, 41000, 42000, 43000, 42000, 45500, 41000, 41500, 42000, 44000, 46000, 47000, 46000, 44000, 42000}
		require.Equal(t, exp, results[0])
	})

	t.Run("Skip trades with zero price or amount", func(t *testing.T) {
		inputTrades := []TradeEvent{
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)}},
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)}},
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.Zero, Amount: decimal.NewFromFloat(1.0)}},
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(44000), Amount: decimal.Zero}},
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(44000), Amount: decimal.NewFromFloat(1)}},
		}

		results := testStrategies(inputTrades, newIndexStrategy(WithCustomWeights(defaultWeights)))

		// Check strategy
		exp := []float64{40000, 42000, 44000}
		require.Equal(t, exp, results[0])
	})

	inputTrades := []TradeEvent{
		{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(0.3)}},
		{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(42500), Amount: decimal.NewFromFloat(0.5)}},
		{Source: common.DriverUniswapV3, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(55000), Amount: decimal.NewFromFloat(0.6)}},
		{Source: common.DriverUniswapV3, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(50000), Amount: decimal.NewFromFloat(0.4)}},
		{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1)}},
	}

	t.Run("README example 1: equal driver weight", func(t *testing.T) {
		results := testStrategies(inputTrades, newIndexStrategy(WithCustomWeights(defaultWeights)))

		// Check strategy
		exp := []float64{41000, 42500, 48750, 46250, 45000}
		require.Equal(t, exp, results[0])
	})

	t.Run("README example 2: zero weight for one of the drivers", func(t *testing.T) {
		weights := map[common.DriverType]decimal.Decimal{
			common.DriverBinance:   decimal.NewFromInt(3),
			common.DriverUniswapV3: decimal.NewFromInt(0),
		}

		results := testStrategies(inputTrades, newIndexStrategy(WithCustomWeights(weights)))

		// Check strategy
		exp := []float64{41000, 42500, 40000}
		require.Equal(t, exp, results[0])
	})

	t.Run("README example 4: drivers volatility", func(t *testing.T) {
		// Initial price: 41000
		inputTrades := []TradeEvent{{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromInt(1.0)}}}
		// Two equal drivers are sending: 42000 and 40000 prices sequentially.
		inputTrades = append(inputTrades, generateTrades([]TradeEvent{
			{Source: common.DriverUniswapV3, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(42000), Amount: decimal.NewFromFloat(1.0)}},
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(40000), Amount: decimal.NewFromFloat(1.0)}},
		}, 25)...)
		// The drivers start sending the same price: 41000.
		inputTrades = append(inputTrades, generateTrades([]TradeEvent{
			{Source: common.DriverUniswapV3, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(1.0)}},
			{Source: common.DriverBinance, TradeEvent: common.TradeEvent{Market: btcusdt, Price: decimal.NewFromInt(41000), Amount: decimal.NewFromFloat(1.0)}},
		}, 25)...)

		testPriceCache := newPriceCache(defaultWeights, 20, time.Minute)
		results := testStrategies(inputTrades, newIndexStrategy(WithCustomWeights(defaultWeights), withCustomPriceCache(testPriceCache)))

		// Check strategy
		require.Equal(t, float64(41000), results[0][0])
		require.Equal(t, float64(41000), results[0][4])
		require.Equal(t, float64(41000), results[0][50])
	})
}

func generateTrades(tr []TradeEvent, n int) []TradeEvent {
	trades := make([]TradeEvent, 0, len(tr)*n)
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
