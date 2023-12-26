package quotes

import (
	"sync"

	"github.com/shopspring/decimal"
)

type PriceInterface interface {
	Get(market string) (decimal.Decimal, decimal.Decimal)
	Update(driver DriverType, market string, priceWeight, weight decimal.Decimal)
	ActiveDrivers(market string) decimal.Decimal
}

type PriceCache struct {
	weightsMap map[DriverType]decimal.Decimal
	market     map[string]*price
	mu         sync.RWMutex
}

// price contains priceWeight and weight EMAs, and list of active drivers for the market.
type price struct {
	priceWeight   decimal.Decimal
	weight        decimal.Decimal
	activeDrivers map[DriverType]bool
}

// NewPriceCache initializes new cache for ema prices for markets.
func NewPriceCache(weightsMap map[DriverType]decimal.Decimal) *PriceCache {
	cache := new(PriceCache)
	cache.market = make(map[string]*price, 0)
	cache.market = make(map[string]*price, 0)
	cache.weightsMap = weightsMap

	return cache
}

// Get returns the price record for the market from cache.
func (p *PriceCache) Get(market string) (decimal.Decimal, decimal.Decimal) {
	cached, ok := p.market[market]
	if ok {
		return cached.priceWeight, cached.weight
	}
	return decimal.Zero, decimal.Zero
}

// Update updates or creates a price record in cache. It also makes the driver active in the drivers map.
func (p *PriceCache) Update(driver DriverType, market string, priceWeight, weight decimal.Decimal) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.market[market]
	if ok {
		p.market[market].priceWeight = priceWeight
		p.market[market].weight = weight
		p.market[market].activeDrivers[driver] = true
		return
	}
	p.market[market] = &price{priceWeight: priceWeight, weight: weight, activeDrivers: map[DriverType]bool{driver: true}}
}

// ActiveDrivers returns the sum of active driver weights for the market).
func (p *PriceCache) ActiveDrivers(market string) decimal.Decimal {
	_, ok := p.market[market]
	if ok {
		count := decimal.Zero
		for driver, active := range p.market[market].activeDrivers {
			if active == true {
				weight, ok := p.weightsMap[driver]
				if ok {
					count.Add(weight)
				}
			}
		}
		return count
	}
	return decimal.Zero
}
