package quotes

import (
	"github.com/layer-3/clearsync/pkg/safe"
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

func (*marketHistory) Comparable() bool {
	return true
}

func newMarketHistory() marketHistory {
	return marketHistory{
		trades:        []trade{},
		activeDrivers: map[DriverType]bool{},
	}
}

type marketKey struct {
	baseUnit  string
	quoteUnit string
}

type PriceCacheVWA struct {
	weights safe.Map[DriverType, decimal.Decimal]
	market  safe.Map[marketKey, marketHistory]
	nTrades int
}

// NewPriceCacheVWA initializes a new cache to store last n trades for each market.
func NewPriceCacheVWA(driversWeights map[DriverType]decimal.Decimal, nTrades int) *PriceCacheVWA {
	cache := new(PriceCacheVWA)
	cache.market = safe.NewMap[marketKey, marketHistory]()
	cache.weights = safe.NewMapWithData(driversWeights)
	cache.nTrades = nTrades

	return cache
}

// AddTrade adds a new trade to the cache for a market.
func (p *PriceCacheVWA) AddTrade(market Market, price, volume, weight decimal.Decimal) {
	key := marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit}
	p.market.UpdateInTx(func(m map[marketKey]marketHistory) {
		history, ok := m[key]
		// Ensure the market exists in the cache
		if !ok {
			history = newMarketHistory()
		}

		// Append the new trade and maintain only the last N trades
		trades := append(history.trades, trade{Price: price, Volume: volume, Weight: weight})
		if len(trades) > p.nTrades {
			trades = trades[len(trades)-p.nTrades:]
		}

		history.trades = trades

		m[key] = history
	})
}

// GetVWA calculates the VWA based on a list of trades.
func (p *PriceCacheVWA) GetVWA(market Market) (decimal.Decimal, bool) {
	record, ok := p.market.Load(marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit})
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

	quotePrice := decimal.NewFromInt(1)
	if market.convertTo != "" {
		quotePrice, ok = p.GetVWA(Market{baseUnit: market.quoteUnit, quoteUnit: market.convertTo})
		if !ok {
			return decimal.Zero, false
		}
	}

	return totalPriceVolume.Div(totalVolume).Mul(quotePrice), true
}

// ActivateDriver makes the driver active for the market.
func (p *PriceCacheVWA) ActivateDriver(driver DriverType, market Market) {
	key := marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit}
	p.market.UpdateInTx(func(m map[marketKey]marketHistory) {
		history, ok := m[key]
		if !ok {
			history = newMarketHistory()
		}

		history.activeDrivers[driver] = true

		m[key] = history
	})
}

// ActiveWeights returns the sum of active driver weights for the market.
// TODO: cache the weights inside the marketHistory
func (p *PriceCacheVWA) ActiveWeights(market Market) decimal.Decimal {
	count := decimal.Zero
	key := marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit}
	// there are not changes in the `market`` map, but we need to read value and `activeDrivers` map transactionally
	p.market.UpdateInTx(func(m map[marketKey]marketHistory) {
		history, ok := m[key]
		if !ok {
			return
		}

		for driver, active := range history.activeDrivers {
			if weight, ok := p.weights.Load(driver); active && ok {
				count = count.Add(weight)
			}
		}
	})

	return count
}
