package quotes

import (
	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

type ConfFuncVWA func(*strategyVWA)

type strategyVWA struct {
	weights    map[DriverType]decimal.Decimal
	priceCache *PriceCacheVWA
}

// NewStrategyVWA creates a new instance of Volume-Weighted Average Price index price calculator.
func NewStrategyVWA(configs ...ConfFuncVWA) priceCalculator {
	s := strategyVWA{
		priceCache: NewPriceCacheVWA(defaultWeightsMap, 20),
		weights:    defaultWeightsMap,
	}
	for _, conf := range configs {
		conf(&s)
	}
	return s
}

// WithCustomWeightsVWA configures custom drivers weights. Should be passed as an argument to the NewStrategyVWA() constructor.
func WithCustomWeightsVWA(driversWeights map[DriverType]decimal.Decimal) ConfFuncVWA {
	return func(strategy *strategyVWA) {
		strategy.weights = driversWeights
		strategy.priceCache.weights = safe.NewMapWithData[DriverType, decimal.Decimal](driversWeights)
	}
}

// WithCustomPriceCacheVWA configures price cache. Should be passed as an argument to the NewStrategyVWA() constructor.
func WithCustomPriceCacheVWA(priceCache *PriceCacheVWA) ConfFuncVWA {
	return func(strategy *strategyVWA) {
		strategy.priceCache = priceCache
	}
}

// calculateIndexPrice returns indexPrice based on Volume Weighted Average Price of last 20 trades.
func (a strategyVWA) calculateIndexPrice(event TradeEvent) (decimal.Decimal, bool) {
	sourceWeight := a.weights[event.Source]
	if event.Market.IsEmpty() || event.Price.IsZero() || event.Amount.IsZero() || sourceWeight.IsZero() {
		return decimal.Decimal{}, false
	}

	a.priceCache.ActivateDriver(event.Source, event.Market)
	activeWeights := a.priceCache.ActiveWeights(event.Market)

	// sourceMultiplier defines how much the trade from a specific market will affect a new price
	sourceMultiplier := sourceWeight
	if activeWeights != decimal.Zero {
		sourceMultiplier = sourceWeight.Div(activeWeights)
	}

	// Add the current trade to the cache
	a.priceCache.AddTrade(event.Market, event.Price, event.Amount, sourceMultiplier)

	return a.priceCache.GetVWA(event.Market)
}
