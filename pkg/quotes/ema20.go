package quotes

import "github.com/shopspring/decimal"

type ConfFuncEMA func(*strategyEMA)

type strategyEMA struct {
	weights    map[DriverType]decimal.Decimal
	priceCache *PriceCacheEMA
	nTrades    int32
}

// NewStrategyEMA creates a new instance of EMA index price calculator.
func NewStrategyEMA(configs ...ConfFuncEMA) priceCalculator {
	s := strategyEMA{
		priceCache: NewPriceCacheEMA(DefaultWeightsMap),
		weights:    DefaultWeightsMap,
		nTrades:    20,
	}
	for _, conf := range configs {
		conf(&s)
	}
	return s
}

// WithCustomWeightsEMA configures custom drivers weights. Should be passed as an argument to the NewStrategyEMA() constructor.
func WithCustomWeightsEMA(driversWeights map[DriverType]decimal.Decimal) ConfFuncEMA {
	return func(strategy *strategyEMA) {
		strategy.weights = driversWeights
		strategy.priceCache.weights = driversWeights
	}
}

// WithCustomPriceCacheEMA configures price cache. Should be passed as an argument to the NewStrategyEMA() constructor.
func WithCustomPriceCacheEMA(priceCache *PriceCacheEMA) ConfFuncEMA {
	return func(strategy *strategyEMA) {
		strategy.priceCache = priceCache
	}
}

// calculateIndex returns indexPrice based on Weighted Exponential Moving Average of last 20 trades.
func (a strategyEMA) calculateIndex(event TradeEvent) (decimal.Decimal, bool) {
	sourceWeight := a.weights[event.Source]
	if event.Market == "" || event.Price.String() == "0" || event.Amount.String() == "0" || sourceWeight.IsZero() {
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
	newNumEMA := EMA(numEMA, event.Price.Mul(event.Amount).Mul(sourceMultiplier), a.nTrades)
	newDenEMA := EMA(denEMA, event.Amount.Mul(sourceMultiplier), a.nTrades)

	newEMA := newNumEMA.Div(newDenEMA)
	a.priceCache.Update(event.Market, newNumEMA, newDenEMA)

	return newEMA, true
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
