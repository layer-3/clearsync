package quotes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type PriceCache interface {
	GetEMA(market string) decimal.Decimal              // Returns last EMA for a market
	UpdateEMA(market string, newValue decimal.Decimal) // Replaces the last EMA for a market with a new value
}

type PricesCache struct {
	prices map[string]decimal.Decimal
	mu     sync.RWMutex
}

// NewPricesCache initializes new cache for ema prices for markets.
func NewPricesCache() *PricesCache {
	cache := new(PricesCache)
	cache.prices = make(map[string]decimal.Decimal, 0)

	return cache
}

func (p *PricesCache) GetEMA(market string) decimal.Decimal {
	cached, ok := p.prices[market]
	if ok {
		return cached
	}
	return decimal.Zero
}
func (p *PricesCache) UpdateEMA(market string, newValue decimal.Decimal) {
	p.mu.Lock()
	p.prices[market] = newValue
	p.mu.Unlock()
}
