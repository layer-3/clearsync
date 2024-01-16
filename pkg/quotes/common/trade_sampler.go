package common

import (
	"math/rand"
)

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

func (ts *tradeSampler) Allow(trade TradeEvent) bool {
	if !ts.enabled {
		return true
	}

	if rand.Intn(100) < ts.defaultPercentage {
		return true
	}

	logger.Debugw("skipping trade", "trade", trade)
	return false
}
