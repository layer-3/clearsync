package quotes

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTradeSampler(t *testing.T) {
	defaultPercentage := rand.Int()
	conf := TradeSamplerConfig{
		Enabled:           false,
		DefaultPercentage: defaultPercentage,
	}

	tradeSampler := newTradeSampler(conf)
	require.Equal(t, tradeSampler.defaultPercentage, defaultPercentage)
	require.False(t, tradeSampler.enabled)
}

func TestTradeSampler_Allow(t *testing.T) {
	t.Run("TradeSampler is not enabled", func(t *testing.T) {
		conf := TradeSamplerConfig{
			Enabled:           false,
			DefaultPercentage: 0,
		}
		ts := newTradeSampler(conf)

		require.True(t, ts.Allow(TradeEvent{}))
	})

	t.Run("DefaultPercentage is in specified range", func(t *testing.T) {
		conf := TradeSamplerConfig{
			Enabled:           true,
			DefaultPercentage: 200, // should be greater than rand.Intn(100)
		}
		ts := newTradeSampler(conf)

		require.True(t, ts.Allow(TradeEvent{}))
	})

	t.Run("Should return false", func(t *testing.T) {
		conf := TradeSamplerConfig{
			Enabled:           true,
			DefaultPercentage: 0,
		}
		ts := newTradeSampler(conf)

		require.False(t, ts.Allow(TradeEvent{}))
	})
}
