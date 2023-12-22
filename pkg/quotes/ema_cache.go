package quotes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type EMACache interface {
	Get(market string) (decimal.Decimal, decimal.Decimal)      // Returns priceWeight and weight EMAs for a market
	Update(market string, priceWeight, weight decimal.Decimal) // Updates priceWeight and weight EMAs for a market with a new value
}

type EMAsCache struct {
	ema map[string]emaRecord
	mu  sync.RWMutex
}

type emaRecord struct {
	priceWeight decimal.Decimal
	weight      decimal.Decimal
}

// NewEMAsCacheinitializes new cache for ema prices for markets.
func NewEMAsCache() *EMAsCache {
	cache := new(EMAsCache)
	cache.ema = make(map[string]emaRecord, 0)

	return cache
}

func (p *EMAsCache) Get(market string) (decimal.Decimal, decimal.Decimal) {
	cached, ok := p.ema[market]
	if ok {
		return cached.priceWeight, cached.weight
	}
	return decimal.Zero, decimal.Zero
}

func (p *EMAsCache) Update(market string, priceWeight, weight decimal.Decimal) {
	p.mu.Lock()
	p.ema[market] = emaRecord{priceWeight: priceWeight, weight: weight}
	p.mu.Unlock()
}
