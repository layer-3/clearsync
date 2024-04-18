package quotes

import (
	"time"

	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

var defaultWeightsMap = map[DriverType]decimal.Decimal{
	DriverKraken:    decimal.NewFromInt(15),
	DriverBinance:   decimal.NewFromInt(20),
	DriverUniswapV3: decimal.NewFromInt(50),
	DriverSyncswap:  decimal.NewFromInt(50),
	DriverQuickswap: decimal.NewFromInt(50),
	DriverInternal:  decimal.NewFromInt(75),
}

type ConfFunc func(*indexStrategy)

type indexStrategy struct {
	weights    map[DriverType]decimal.Decimal
	priceCache *PriceCache
}

// newIndexStrategy creates a new instance of Weighted Price index price calculator.
func newIndexStrategy(configs ...ConfFunc) priceCalculator {
	s := indexStrategy{
		priceCache: newPriceCache(defaultWeightsMap, 20, 15*time.Minute),
		weights:    defaultWeightsMap,
	}
	for _, conf := range configs {
		conf(&s)
	}
	return s
}

// WithCustomWeights configures custom drivers weights. Should be passed as an argument to the NewStrategy() constructor.
func WithCustomWeights(driversWeights map[DriverType]decimal.Decimal) ConfFunc {
	return func(strategy *indexStrategy) {
		strategy.weights = driversWeights
		strategy.priceCache.weights = safe.NewMapWithData(driversWeights)
	}
}

// withCustomPriceCache configures price cache. Should be passed as an argument to the NewStrategy() constructor.
func withCustomPriceCache(priceCache *PriceCache) ConfFunc {
	return func(strategy *indexStrategy) {
		strategy.priceCache = priceCache
	}
}

// calculateIndexPrice returns indexPrice based on Volume Weighted Average Price of last 20 trades.
func (a indexStrategy) calculateIndexPrice(event TradeEvent) (decimal.Decimal, bool) {
	sourceWeight := a.weights[event.Source]
	if event.Market.IsEmpty() || event.Price.IsZero() || event.Amount.IsZero() || sourceWeight.IsZero() {
		return decimal.Decimal{}, false
	}

	timeEmpty := time.Time{}
	if event.CreatedAt == timeEmpty {
		event.CreatedAt = time.Now()
	}

	a.priceCache.AddTrade(event.Market, event.Price, event.Amount, event.CreatedAt, event.Source)

	return a.priceCache.GetIndexPrice(&event)
}

func (a indexStrategy) getLastPrice(market Market) decimal.Decimal {
	return a.priceCache.getLastPrice(market)
}

func (a indexStrategy) setLastPrice(market Market, price decimal.Decimal) {
	a.priceCache.setLastPrice(market, price)
}
