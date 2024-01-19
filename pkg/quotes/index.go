package quotes

import (
	"github.com/shopspring/decimal"
)

type indexAggregator struct {
	weights map[DriverType]decimal.Decimal

	drivers    []Driver
	priceCache PriceInterface

	aggregated chan TradeEvent
	outbox     chan<- TradeEvent
}

// NewIndexAggregator creates a new instance of IndexAggregator.
func NewIndexAggregator(driverConfigs []Config, weightsMap map[DriverType]decimal.Decimal, outbox chan<- TradeEvent) Driver {
	aggregated := make(chan TradeEvent, 128)

	drivers := []Driver{}
	for _, d := range driverConfigs {
		if d.Driver == DriverIndex {
			continue
		}
		driver, err := NewDriver(d, aggregated)
		if err != nil {
			continue
		}
		drivers = append(drivers, driver)
	}

	return &indexAggregator{
		priceCache: NewPriceCache(weightsMap),
		weights:    weightsMap,
		drivers:    drivers,
		outbox:     outbox,
		aggregated: aggregated,
	}
}

func newIndex(config Config, outbox chan<- TradeEvent) Driver {
	return NewIndexAggregator(AllDrivers, DefaultWeightsMap, outbox)
}

func (a *indexAggregator) Name() DriverType {
	return DriverIndex
}

func (a *indexAggregator) Start() error {
	logger.Info("starting index quotes service")

	for _, d := range a.drivers {
		go func(d Driver) {
			if err := d.Start(); err != nil {
				logger.Warn(err.Error())
			}
		}(d)
	}
	go func() {
		for event := range a.aggregated {
			indexPrice := a.indexPrice(event)
			a.outbox <- indexPrice
		}
	}()
	return nil
}

func (a *indexAggregator) Subscribe(m Market) error {
	for _, d := range a.drivers {
		err := d.Subscribe(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *indexAggregator) Unsubscribe(m Market) error {
	for _, d := range a.drivers {
		err := d.Unsubscribe(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *indexAggregator) Stop() error {
	for _, d := range a.drivers {
		err := d.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

// indexPrice returns indexPrice based on Weighted Exponential Moving Average of last 20 trades.
func (a *indexAggregator) indexPrice(event TradeEvent) TradeEvent {
	sourceWeight := a.weights[event.Source]
	numEMA, denEMA := a.priceCache.Get(event.Market)

	a.priceCache.ActivateDriver(event.Source, event.Market)
	activeWeights := a.priceCache.ActiveWeights(event.Market)
	if activeWeights == decimal.Zero {
		activeWeights = sourceWeight
	}

	// sourceMultiplier defines how much the trade from a specific market will affect a new price
	sourceMultiplier := sourceWeight.Div(activeWeights)

	// To start the procedure (before we've got trades) we generate the initial values:
	if numEMA == decimal.Zero || denEMA == decimal.Zero {
		numEMA = event.Price.Mul(event.Amount).Mul(sourceMultiplier)
		denEMA = event.Amount.Mul(sourceWeight).Div(activeWeights)
	}

	// Weighted Exponential Moving Average:
	// https://www.financialwisdomforum.org/gummy-stuff/EMA.htm
	newNumEMA := EMA20(numEMA, event.Price.Mul(event.Amount).Mul(sourceMultiplier))
	newDenEMA := EMA20(denEMA, event.Amount.Mul(sourceMultiplier))

	newEMA := newNumEMA.Div(newDenEMA)
	a.priceCache.Update(event.Market, newNumEMA, newDenEMA)

	event.Price = newEMA
	event.Source = DriverIndex

	return event
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
