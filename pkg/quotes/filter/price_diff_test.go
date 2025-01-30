package filter

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

func TestNewPriceDiffFilter(t *testing.T) {
	t.Parallel()

	t.Run("Correct treshold", func(t *testing.T) {
		t.Parallel()

		conf := PriceDiffConfig{
			Threshold: "0.1",
		}
		priceDiffFilter := newPriceDiffFilter(conf)

		require.Equal(t, priceDiffFilter.threshold.String(), decimal.NewFromFloat(0.1).String())
	})

	t.Run("Incorrect treshold, that should be set to default", func(t *testing.T) {
		t.Parallel()

		conf := PriceDiffConfig{
			Threshold: "0,1",
		}
		priceDiffFilter := newPriceDiffFilter(conf)

		require.Equal(t, priceDiffFilter.threshold.String(), decimal.NewFromFloat(0.05).String())
	})
}

func TestPriceDiffFilter_Allow(t *testing.T) {
	t.Parallel()

	conf := PriceDiffConfig{
		Threshold: "0.1",
	}
	priceDiffFilter := newPriceDiffFilter(conf)

	t.Run("First trade event should be accepted by default", func(t *testing.T) {
		trade := common.TradeEvent{
			Market: common.NewMarket("btc", "usd"),
			Price:  decimal.NewFromFloat(50_000),
		}

		require.True(t, priceDiffFilter.Allow(trade))
	})

	t.Run("Trade event should be declined if price diff is too small", func(t *testing.T) {
		trade := common.TradeEvent{
			Market: common.NewMarket("btc", "usd"),
			// Price diff is 1000, which is less than 0.1 * 50_000 = 5000
			Price: decimal.NewFromFloat(49_000),
		}

		require.False(t, priceDiffFilter.Allow(trade))
	})

	t.Run("Trade event should be accepted if price diff is big enough", func(t *testing.T) {
		trade := common.TradeEvent{
			Market: common.NewMarket("btc", "usd"),
			// Price diff is 20_000, which is more than 0.1 * 50_000 = 5000
			Price: decimal.NewFromFloat(30_000),
		}

		require.True(t, priceDiffFilter.Allow(trade))
	})

	t.Run("Trade event should be compared to the last accepted price", func(t *testing.T) {
		trade := common.TradeEvent{
			Market: common.NewMarket("btc", "usd"),
			// Price diff is 21K with current price (30K) and 1K with previous price (50K)
			// It should be accepted because it's more than 0.1 * 30K = 3K
			// but declined if compared to the previous price 0.1 * 50K = 5K
			Price: decimal.NewFromFloat(51_000),
		}

		require.True(t, priceDiffFilter.Allow(trade))
	})
}
