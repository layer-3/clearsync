package quotes

import (
	"math/rand"

	"github.com/ipfs/go-log/v2"
)

var loggerTradeSampler = log.Logger("trade_sampler")

type tradeSampler struct {
	enabled           bool
	defaultPercentage int
}

func newTradeSampler(conf TradeSamplerConfig) *tradeSampler {
	return &tradeSampler{
		enabled:           conf.Enabled,
		defaultPercentage: conf.DefaultPercentage,
	}
}

func (ts *tradeSampler) allow(trade TradeEvent) bool {
	if !ts.enabled {
		return true
	}

	if rand.Intn(100) < ts.defaultPercentage {
		return true
	}

	loggerBinance.Debugw("skipping trade", "trade", trade)
	return false
}
