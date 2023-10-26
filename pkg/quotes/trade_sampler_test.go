package quotes

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/config"
)

func TestNewTradeSampler(t *testing.T) {
	defaultPercentage := rand.Int()
	conf := config.TradeSampler{
		Enabled:           false,
		DefaultPercentage: defaultPercentage,
	}

	tradeSampler := NewTradeSampler(conf)
	require.Equal(t, tradeSampler.defaultPercentage, defaultPercentage)
	require.False(t, tradeSampler.enabled)
}

func TestAllow(t *testing.T) {
	t.Run("TradeSampler is not enabled", func(t *testing.T) {
		conf := config.TradeSampler{
			Enabled:           false,
			DefaultPercentage: 0,
		}
		ts := NewTradeSampler(conf)

		require.True(t, ts.Allow(trade.Event{}))
	})

	t.Run("DefaultPercentage is in specified range", func(t *testing.T) {
		conf := config.TradeSampler{
			Enabled:           true,
			DefaultPercentage: 200, // should be greater than rand.Intn(100)
		}
		ts := NewTradeSampler(conf)

		require.True(t, ts.Allow(trade.Event{}))
	})

	t.Run("Should return false", func(t *testing.T) {
		conf := config.TradeSampler{
			Enabled:           true,
			DefaultPercentage: 0,
		}
		ts := NewTradeSampler(conf)

		require.False(t, ts.Allow(trade.Event{}))
	})
}
