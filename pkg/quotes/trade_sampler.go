package quotes

import (
	"math/rand"
)

type TradeSampler struct {
	enabled           bool
	defaultPercentage int
}

func NewTradeSampler(conf TradeSamplerConfig) *TradeSampler {
	return &TradeSampler{
		enabled:           conf.Enabled,
		defaultPercentage: conf.DefaultPercentage,
	}
}

func (ts *TradeSampler) Allow(trade TradeEvent) bool {
	if !ts.enabled {
		return true
	}

	if rand.Intn(100) < ts.defaultPercentage {
		return true
	}

	logger.Debugw("skipping trade", "trade", trade)
	return false
}
