package quotes

import (
	"time"

	"github.com/shopspring/decimal"
)

type PriceCache interface {
	GetEMA(market string) (decimal.Decimal, int64)     // Returns last EMA for a market
	UpdateEMA(market string, newValue decimal.Decimal) // Replaces the last EMA for a market with a new value
}

type PricesCache struct {
	prices map[string]emaRecord
}

type emaRecord struct {
	amount    decimal.Decimal // Last EMA value
	timestamp int64           // Timestamp is used to update the last ema value in a specific period of time (1 minute)
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
