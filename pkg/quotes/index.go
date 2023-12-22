package quotes

import (
	"github.com/shopspring/decimal"
)

type IndexAggregator struct {
	weightsMap    map[DriverType]decimal.Decimal
	activeWeights map[DriverType]decimal.Decimal

	drivers  []Driver
	emaCache EMACache

	aggregated chan TradeEvent
	outbox     chan<- TradeEvent
}

func NewIndexAggregator(driverConfigs []Config, weightsMap map[DriverType]decimal.Decimal, outbox chan<- TradeEvent) (*IndexAggregator, error) {
	aggregated := make(chan TradeEvent, 128)

	drivers := []Driver{}
	for _, d := range driverConfigs {
		driver, err := NewDriver(d, aggregated)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	return &IndexAggregator{
		emaCache:      NewEMAsCache(),
		weightsMap:    weightsMap,
		activeWeights: make(map[DriverType]decimal.Decimal),
		drivers:       drivers,
		outbox:        outbox,
		aggregated:    aggregated,
	}, nil
}

func (a *IndexAggregator) Start(markets []Market) error {
	logger.Info("starting index quotes service")

	for _, d := range a.drivers {
		go func(d Driver) {
			markets := make([]Market, 0, len(markets))
			for _, m := range markets {
				market := Market{BaseUnit: m.BaseUnit, QuoteUnit: m.QuoteUnit}
				markets = append(markets, market)
			}

			if err := d.Start(markets); err != nil {
				logger.Warn(err.Error())
			}
			a.activeWeights[d.Name()] = a.weightsMap[d.Name()]
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

func (a *IndexAggregator) Subscribe(m Market) error {
	for _, d := range a.drivers {
		err := d.Subscribe(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *IndexAggregator) Stop() error {
	for _, d := range a.drivers {
		err := d.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *IndexAggregator) indexPrice(event TradeEvent) TradeEvent {
	driverWeight := a.activeWeights[event.Source]
	priceWeightEMA, weightEMA := a.emaCache.Get(event.Market)

	// To start the procedure (before we've got trades) we generate the initial values:
	if priceWeightEMA == decimal.Zero || weightEMA == decimal.Zero {
		priceWeightEMA = event.Price.Mul(event.Amount).Mul(driverWeight).Div(a.totalWeights())
		weightEMA = event.Amount.Mul(driverWeight).Div(a.totalWeights())
	}

	// Volume-weighted Exponential Moving Average (V-EMA)
	// https://www.financialwisdomforum.org/gummy-stuff/EMA.htm
	newPriceWeightEMA := EMA20(priceWeightEMA, event.Price.Mul(event.Amount).Mul(driverWeight).Div(a.totalWeights()))
	newWeightEMA := EMA20(weightEMA, event.Amount.Mul(driverWeight).Div(a.totalWeights()))

	newEMA := newPriceWeightEMA.Div(newWeightEMA)
	a.emaCache.Update(event.Market, newPriceWeightEMA, newWeightEMA)

	event.Price = newEMA
	event.Source = DriverIndex

	return event
}

func (a *IndexAggregator) totalWeights() decimal.Decimal {
	total := decimal.Zero
	for _, w := range a.activeWeights {
		total = total.Add(w)
	}
	return total
}

func EMA20(lastEMA, newPrice decimal.Decimal) decimal.Decimal {
	return EMA(lastEMA, newPrice, 20)
}

func EMA(previousEMA, price decimal.Decimal, intervals int32) decimal.Decimal {
	if intervals <= 0 {
		return decimal.Zero
	}

	alpha := alpha(intervals)
	return alpha.Mul(price.Sub(previousEMA)).Add(previousEMA)
}

func alpha(intervals int32) decimal.Decimal {
	smoothingFactor := decimal.NewFromInt(2)
	alpha := smoothingFactor.Div(decimal.NewFromInt32(intervals).Add(decimal.NewFromInt(1)))
	return alpha
}
