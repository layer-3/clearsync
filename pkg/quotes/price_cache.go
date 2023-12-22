package quotes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type PriceCache interface {
	GetEMA(market string) (decimal.Decimal, decimal.Decimal)      // Returns priceWeight and weight EMAs for a market
	UpdateEMA(market string, priceWeight, weight decimal.Decimal) // Replaces priceWeight and weight EMAs for a market with a new value
}

type PricesCache struct {
	prices map[string]emaRecord
	mu     sync.RWMutex
}

type emaRecord struct {
	priceWeight decimal.Decimal
	weight      decimal.Decimal
}

// NewPricesCache initializes new cache for ema prices for markets.
func NewPricesCache() *PricesCache {
	cache := new(PricesCache)
	cache.prices = make(map[string]emaRecord, 0)

	return cache
}

func (p *PricesCache) GetEMA(market string) (decimal.Decimal, decimal.Decimal) {
	cached, ok := p.prices[market]
	if ok {
		return cached.priceWeight, cached.weight
	}
	return decimal.Zero, decimal.Zero
}
func (p *PricesCache) UpdateEMA(market string, priceWeight, weight decimal.Decimal) {
	p.mu.Lock()
	p.prices[market] = emaRecord{priceWeight: priceWeight, weight: weight}
	p.mu.Unlock()
}
