package quotes

import "github.com/shopspring/decimal"

type ConfFunc func(*strategyEMA20)

type strategyEMA20 struct {
	weights    map[DriverType]decimal.Decimal
	priceCache PriceInterface
}

// NewStrategyEMA20 creates a new instance of EMA20 index price calculator.
func NewStrategyEMA20(configs ...ConfFunc) priceCalculator {
	s := strategyEMA20{
		priceCache: NewPriceCache(DefaultWeightsMap),
		weights:    DefaultWeightsMap,
	}
	for _, conf := range configs {
		conf(&s)
	}
	return s
}

// WithCustomWeightsEMA20 configures custom drivers weights. Should be passed as an argument to the NewStrategyEMA20() constructor.
func WithCustomWeightsEMA20(driversWeights map[DriverType]decimal.Decimal) ConfFunc {
	return func(strategy *strategyEMA20) {
		strategy.weights = driversWeights
		strategy.priceCache = NewPriceCache(driversWeights)
	}
}

// WithCustomPriceCacheEMA20 configures price cache. Should be passed as an argument to the NewStrategyEMA20() constructor.
func WithCustomPriceCacheEMA20(priceCache PriceInterface) ConfFunc {
	return func(strategy *strategyEMA20) {
		strategy.priceCache = priceCache
	}
}

// calculateIndex returns indexPrice based on Weighted Exponential Moving Average of last 20 trades.
func (a strategyEMA20) calculateIndex(event TradeEvent) (decimal.Decimal, bool) {
	sourceWeight := a.weights[event.Source]
	if event.Market == "" || event.Price.String() == "0" || event.Amount.String() == "0" || sourceWeight == decimal.Zero {
		return decimal.Decimal{}, false
	}

	numEMA, denEMA := a.priceCache.Get(event.Market)

	a.priceCache.ActivateDriver(event.Source, event.Market)
	activeWeights := a.priceCache.ActiveWeights(event.Market)

	// sourceMultiplier defines how much the trade from a specific market will affect a new price
	sourceMultiplier := sourceWeight
	if activeWeights != decimal.Zero {
		sourceMultiplier = sourceWeight.Div(activeWeights)
	}

	// To start the procedure (before we've got trades) we generate the initial values:
	if numEMA == decimal.Zero || denEMA == decimal.Zero {
		numEMA = event.Price.Mul(event.Amount).Mul(sourceMultiplier)
		denEMA = event.Amount.Mul(sourceMultiplier)
	}

	// Weighted Exponential Moving Average:
	// https://www.financialwisdomforum.org/gummy-stuff/EMA.htm
	newNumEMA := EMA20(numEMA, event.Price.Mul(event.Amount).Mul(sourceMultiplier))
	newDenEMA := EMA20(denEMA, event.Amount.Mul(sourceMultiplier))

	newEMA := newNumEMA.Div(newDenEMA)
	a.priceCache.Update(event.Market, newNumEMA, newDenEMA)

	return newEMA, true
}

// EMA20 returns Exponential Moving Average for 20 intervals based on previous EMA, and current price.
func EMA20(lastEMA, newPrice decimal.Decimal) decimal.Decimal {
	return EMA(lastEMA, newPrice, 20)
}

// EMA returns Exponential Moving Average based on previous EMA, current price and the number of intervals.
func EMA(previousEMA, newValue decimal.Decimal, intervals int32) decimal.Decimal {
	if intervals <= 0 {
		return decimal.Zero
	}

	smoothing := smoothing(intervals)
	// EMA = ((newValue − previous EMA) × smoothing constant) + (previous EMA).
	return smoothing.Mul(newValue.Sub(previousEMA)).Add(previousEMA)
}

// smoothing returns a smothing constant which equals 2 ÷ (number of periods + 1).
func smoothing(intervals int32) decimal.Decimal {
	smoothingFactor := decimal.NewFromInt(2)
	alpha := smoothingFactor.Div(decimal.NewFromInt32(intervals).Add(decimal.NewFromInt(1)))
	return alpha
}
