package quotes

import (
	"time"

	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

var defaultWeightsMap = map[DriverType]decimal.Decimal{
	DriverKraken:        decimal.NewFromInt(15),
	DriverBinance:       decimal.NewFromInt(20),
	DriverUniswapV3Api:  decimal.NewFromInt(50),
	DriverUniswapV3Geth: decimal.NewFromInt(50),
	DriverSyncswap:      decimal.NewFromInt(50),
	DriverQuickswap:     decimal.NewFromInt(50),
}

type ConfFuncVWA func(*strategyVWA)

type strategyVWA struct {
	weights    map[DriverType]decimal.Decimal
	priceCache *PriceCacheVWA
}

// newStrategyVWA creates a new instance of Volume-Weighted Average Price index price calculator.
func newStrategyVWA(configs ...ConfFuncVWA) priceCalculator {
	s := strategyVWA{
		priceCache: newPriceCacheVWA(defaultWeightsMap, 20, 15*time.Minute),
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
		strategy.priceCache.weights = safe.NewMapWithData(driversWeights)
	}
}

// withCustomPriceCacheVWA configures price cache. Should be passed as an argument to the NewStrategyVWA() constructor.
func withCustomPriceCacheVWA(priceCache *PriceCacheVWA) ConfFuncVWA {
	return func(strategy *strategyVWA) {
		strategy.priceCache = priceCache
	}
}

// calculateIndexPrice returns indexPrice based on Volume Weighted Average Price of last 20 trades.
func (a strategyVWA) calculateIndexPrice(event TradeEvent) (decimal.Decimal, bool) {
	if event.Market.IsEmpty() || event.Price.IsZero() || event.Amount.IsZero() {
		return decimal.Decimal{}, false
	}

	// Add the current trade to the cache
	timeEmpty := time.Time{}
	if event.CreatedAt == timeEmpty {
		event.CreatedAt = time.Now()
	}

	a.priceCache.AddTrade(event.Market, event.Price, event.Amount, event.CreatedAt)

	return a.priceCache.GetVWA(event.Market)
}
