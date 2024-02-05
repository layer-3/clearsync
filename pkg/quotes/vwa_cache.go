package quotes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type trade struct {
	Price  decimal.Decimal
	Volume decimal.Decimal
	Weight decimal.Decimal
}

type marketHistory struct {
	trades        []trade
	activeDrivers map[DriverType]bool
}

type PriceCacheVWA struct {
	weights map[DriverType]decimal.Decimal
	market  map[string]*marketHistory
	mu      sync.RWMutex
	nTrades int
}

// NewPriceCacheVWA initializes a new cache to store last n trades for each market.
func NewPriceCacheVWA(driversWeights map[DriverType]decimal.Decimal, nTrades int) *PriceCacheVWA {
	cache := new(PriceCacheVWA)
	cache.market = make(map[string]*marketHistory)
	cache.weights = driversWeights
	cache.nTrades = nTrades

	return cache
}

// AddTrade adds a new trade to the cache for a market.
func (p *PriceCacheVWA) AddTrade(market string, price, volume, weight decimal.Decimal) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Ensure the market exists in the cache
	if _, ok := p.market[market]; !ok {
		p.market[market] = &marketHistory{
			trades:        []trade{},
			activeDrivers: map[DriverType]bool{},
		}
	}

	// Append the new trade and maintain only the last N trades
	trades := append(p.market[market].trades, trade{Price: price, Volume: volume, Weight: weight})
	if len(trades) > p.nTrades {
		trades = trades[len(trades)-p.nTrades:]
	}
	p.market[market].trades = trades
}

// GetVWA calculates the VWA based on a list of trades.
func (p *PriceCacheVWA) GetVWA(market string) (decimal.Decimal, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	record, ok := p.market[market]
	if !ok || len(record.trades) == 0 {
		return decimal.Zero, false
	}

	var totalPriceVolume, totalVolume decimal.Decimal

	for _, trade := range record.trades {
		totalPriceVolume = totalPriceVolume.Add(trade.Price.Mul(trade.Volume).Mul(trade.Weight))
		totalVolume = totalVolume.Add(trade.Volume.Mul(trade.Weight))
	}

	if totalVolume.IsZero() {
		return decimal.Zero, false
	}

	return totalPriceVolume.Div(totalVolume), true
}

// ActivateDriver makes the driver active for the market.
func (p *PriceCacheVWA) ActivateDriver(driver DriverType, market string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.market[market]
	if ok {
		p.market[market].activeDrivers[driver] = true
		return
	}
	p.market[market] = &marketHistory{trades: []trade{}, activeDrivers: map[DriverType]bool{driver: true}}
}

// ActiveWeights returns the sum of active driver weights for the market.
func (p *PriceCacheVWA) ActiveWeights(market string) decimal.Decimal {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.market[market]
	if ok {
		count := decimal.Zero
		for driver, active := range p.market[market].activeDrivers {
			if active == true {
				weight, ok := p.weights[driver]
				if ok {
					count = count.Add(weight)
				}
			}
		}
		return count
	}
	return decimal.Zero
}
