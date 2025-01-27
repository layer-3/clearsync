package filter

import (
	"crypto/rand"
	"math/big"

	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

var loggerSamplerFilter = log.Logger("sampler_filter")

type SamplerConfig struct {
	Percentage int64 `yaml:"percentage" env:"PERCENTAGE" env-default:"5"`
}

type samplerFilter struct {
	defaultPercentage *big.Int
}

func newSamplerFilter(conf SamplerConfig) samplerFilter {
	return samplerFilter{
		defaultPercentage: new(big.Int).SetInt64(conf.Percentage),
	}
}

func (ts samplerFilter) Allow(trade common.TradeEvent) bool {
	percentage, err := rand.Int(rand.Reader, new(big.Int).SetUint64(100))
	if err != nil {
		loggerSamplerFilter.Errorf("failed to generate a random number: %w", err)
		return false
	}

	if percentage.Cmp(ts.defaultPercentage) < 0 {
		return true
	}

	// loggerSamplerFilter.Infow("skipping trade", "trade", trade)
	return false
}
