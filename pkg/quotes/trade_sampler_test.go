package quotes

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTradeSampler(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	t.Run("TradeSampler is not enabled", func(t *testing.T) {
		t.Parallel()

		conf := TradeSamplerConfig{
			Enabled:           false,
			DefaultPercentage: 0,
		}
		ts := newTradeSampler(conf)

		require.True(t, ts.allow(TradeEvent{}))
	})

	t.Run("DefaultPercentage is in specified range", func(t *testing.T) {
		t.Parallel()

		conf := TradeSamplerConfig{
			Enabled:           true,
			DefaultPercentage: 200, // should be greater than rand.Intn(100)
		}
		ts := newTradeSampler(conf)

		require.True(t, ts.allow(TradeEvent{}))
	})

	t.Run("Should return false", func(t *testing.T) {
		t.Parallel()

		conf := TradeSamplerConfig{
			Enabled:           true,
			DefaultPercentage: 0,
		}
		ts := newTradeSampler(conf)

		require.False(t, ts.allow(TradeEvent{}))
	})
}
