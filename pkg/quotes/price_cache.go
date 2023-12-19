package quotes

import (
	"time"

	"github.com/shopspring/decimal"
)

type PricesCache struct {
	prices map[string]emaRecord
}

type emaRecord struct {
	amount    decimal.Decimal
	timestamp int64
}

// NewPricesCache initializes new cache for ema prices for markets.
func NewPricesCache() *PricesCache {
	cache := new(PricesCache)
	cache.prices = make(map[string]emaRecord, 0)

	return cache
}

func (p *PricesCache) GetEMA(market string) (decimal.Decimal, int64) {
	cached, ok := p.prices[market]
	if ok {
		return cached.amount, cached.timestamp
	}
	return decimal.Zero, 0
}
func (p *PricesCache) UpdateEMA(market string, newValue decimal.Decimal) {
	p.prices[market] = emaRecord{
		amount:    newValue,
		timestamp: time.Now().Unix(),
	}
}
