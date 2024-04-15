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
	weights      map[DriverType]decimal.Decimal
	priceCache   *PriceCacheVWA
	maxPriceDiff decimal.Decimal
}

// newStrategyVWA creates a new instance of Volume-Weighted Average Price index price calculator.
func newStrategyVWA(configs ...ConfFuncVWA) priceCalculator {
	s := strategyVWA{
		priceCache:   newPriceCacheVWA(defaultWeightsMap, 20, 15*time.Minute),
		weights:      defaultWeightsMap,
		maxPriceDiff: decimal.NewFromFloat(1.2),
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

func withCustomMaxPriceDiff(priceDiff decimal.Decimal) ConfFuncVWA {
	return func(strategy *strategyVWA) {
		strategy.maxPriceDiff = priceDiff
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

	timeEmpty := time.Time{}
	if event.CreatedAt == timeEmpty {
		event.CreatedAt = time.Now()
	}

	a.priceCache.AddTrade(event.Market, event.Price, event.Amount, sourceMultiplier, event.CreatedAt)

	price, ok := a.priceCache.GetVWA(event.Market)

	lastPrice := a.priceCache.getLastPrice(event.Market)
	if lastPrice != decimal.Zero {
		if isPriceOutOfRange(price, lastPrice, a.maxPriceDiff) {
			loggerIndex.Warn("skipping outcoming outlier trade", "driver", event)
			return decimal.Zero, false
		}
	}
	a.priceCache.SetLastPrice(event.Market, price)
	return price, ok
}

func (a strategyVWA) getLastPrice(market Market) decimal.Decimal {
	return a.priceCache.getLastPrice(market)
}
