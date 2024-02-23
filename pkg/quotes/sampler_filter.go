package quotes

import (
	"crypto/rand"
	"math/big"

	"github.com/ipfs/go-log/v2"
)

var loggerSamplerFilter = log.Logger("sampler_filter")

type samplerFilter struct {
	defaultPercentage *big.Int
}

func newSamplerFilter(conf SamplerFilterConfig) samplerFilter {
	return samplerFilter{
		defaultPercentage: new(big.Int).SetInt64(conf.Percentage),
	}
}

func (ts samplerFilter) Allow(trade TradeEvent) bool {
	percentage, err := rand.Int(rand.Reader, new(big.Int).SetUint64(100))
	if err != nil {
		loggerSamplerFilter.Errorf("failed to generate a random number: %w", err)
		return false
	}

	if percentage.Cmp(ts.defaultPercentage) < 0 {
		return true
	}

	loggerSamplerFilter.Infow("skipping trade", "trade", trade)
	return false
}
