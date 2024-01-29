package quotes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type PriceCacheEMA struct {
	weights map[DriverType]decimal.Decimal
	market  map[string]*price
	mu      sync.RWMutex
}

// price contains numEMA, denEMA, and list of active drivers for the market.
type price struct {
	numEMA        decimal.Decimal
	denEMA        decimal.Decimal
	driversActive map[DriverType]bool
}

// NewPriceCacheEMA initializes new cache for ema prices for markets.
func NewPriceCacheEMA(weightsMap map[DriverType]decimal.Decimal) *PriceCacheEMA {
	cache := new(PriceCacheEMA)
	cache.market = make(map[string]*price, 0)
	cache.market = make(map[string]*price, 0)
	cache.weights = weightsMap

	return cache
}

// Get returns the price record for the market from cache.
func (p *PriceCacheEMA) Get(market string) (decimal.Decimal, decimal.Decimal) {
	p.mu.Lock()
	defer p.mu.Unlock()

	cached, ok := p.market[market]
	if ok {
		return cached.numEMA, cached.denEMA
	}
	return decimal.Zero, decimal.Zero
}

// Update updates or creates a price record in cache.
func (p *PriceCacheEMA) Update(market string, numEMA, denEMA decimal.Decimal) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.market[market]
	if ok {
		p.market[market].numEMA = numEMA
		p.market[market].denEMA = denEMA
		return
	}
	p.market[market] = &price{numEMA: numEMA, denEMA: denEMA, driversActive: map[DriverType]bool{}}
}

// ActivateDriver makes the driver active for the market.
func (p *PriceCacheEMA) ActivateDriver(driver DriverType, market string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.market[market]
	if ok {
		p.market[market].driversActive[driver] = true
		return
	}
	p.market[market] = &price{numEMA: decimal.Zero, denEMA: decimal.Zero, driversActive: map[DriverType]bool{driver: true}}
}

// ActiveWeights returns the sum of active driver weights for the market.
func (p *PriceCacheEMA) ActiveWeights(market string) decimal.Decimal {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.market[market]
	if ok {
		count := decimal.Zero
		for driver, active := range p.market[market].driversActive {
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
