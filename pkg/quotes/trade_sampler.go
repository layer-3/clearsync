package quotes

import (
	"math/rand"

	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/config"
)

type TradeSampler struct {
	enabled           bool
	defaultPercentage int
}

func NewTradeSampler(conf config.TradeSampler) *TradeSampler {
	return &TradeSampler{
		enabled:           conf.Enabled,
		defaultPercentage: conf.DefaultPercentage,
	}
}

func (ts *TradeSampler) Allow(trade trade.Event) bool {
	if !ts.enabled {
		return true
	}

	if rand.Intn(100) < ts.defaultPercentage {
		return true
	}

	logger.Debugf("skipping trade: %v", trade)

	return false
}
