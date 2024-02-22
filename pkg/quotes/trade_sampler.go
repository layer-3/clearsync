package quotes

import (
	"crypto/rand"
	"math/big"

	"github.com/ipfs/go-log/v2"
)

var loggerTradeSampler = log.Logger("trade_sampler")

type tradeSampler struct {
	enabled           bool
	defaultPercentage *big.Int
}

func newTradeSampler(conf TradeSamplerConfig) *tradeSampler {
	return &tradeSampler{
		enabled:           conf.Enabled,
		defaultPercentage: new(big.Int).SetInt64(conf.DefaultPercentage),
	}
}

func (ts *tradeSampler) allow(trade TradeEvent) bool {
	if !ts.enabled {
		return true
	}

	percentage, err := rand.Int(rand.Reader, new(big.Int).SetUint64(100))
	if err != nil {
		loggerTradeSampler.Errorf("failed to generate a random number: %w", err)
		return false
	}

	if percentage.Cmp(ts.defaultPercentage) == -1 {
		return true
	}

	loggerTradeSampler.Infow("skipping trade", "trade", trade)
	return false
}
