package filter

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

func TestNewSamplerFilter(t *testing.T) {
	t.Parallel()

	defaultPercentage := int64(rand.Int())
	conf := SamplerConfig{
		Percentage: defaultPercentage,
	}

	tradeSampler := newSamplerFilter(conf)
	require.Equal(t, tradeSampler.defaultPercentage, new(big.Int).SetInt64(defaultPercentage))
}

func TestSamplerFilter_Allow(t *testing.T) {
	t.Parallel()

	t.Run("DefaultPercentage is in specified range", func(t *testing.T) {
		t.Parallel()

		conf := SamplerConfig{
			Percentage: 200, // should be greater than rand.Intn(100)
		}
		ts := newSamplerFilter(conf)

		require.True(t, ts.Allow(common.TradeEvent{}))
	})

	t.Run("Should return false", func(t *testing.T) {
		t.Parallel()

		conf := SamplerConfig{
			Percentage: 0,
		}
		ts := newSamplerFilter(conf)

		require.False(t, ts.Allow(common.TradeEvent{}))
	})
}
