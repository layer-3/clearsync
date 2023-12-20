package quotes

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type IndexAggregator struct {
	drivers []Driver
	weights map[DriverType]decimal.Decimal
	prices  PriceCache

	aggregated chan TradeEvent
	outbox     chan<- TradeEvent
}

func NewIndex(driverConfigs []Config, weights map[DriverType]decimal.Decimal, outbox chan<- TradeEvent) (*IndexAggregator, error) {
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
		prices:     NewPricesCache(),
		weights:    weights,
		drivers:    drivers,
		outbox:     outbox,
		aggregated: aggregated,
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
	lastEMA, timestamp := a.prices.GetEMA(event.Market)
	if lastEMA == decimal.Zero {
		lastEMA = event.Price
	}
	// weight := a.weights[event.Source]

	fmt.Println("currentEma            :", lastEMA)
	fmt.Println("eventPrice            :", event.Price)
	newEMA := EMA20(lastEMA, event.Price)

	// TODO: fix weights logic
	// newEMA := EMA20(lastEMA, event.Price.Mul(event.Amount.Mul(weight).Div(a.activeWeights())))

	if time.Now().Unix() >= timestamp+5 {
		a.prices.UpdateEMA(event.Market, newEMA)
	}

	event.Price = newEMA
	event.Source = DriverIndex

	return event
}

func (a *IndexAggregator) activeWeights() decimal.Decimal {
	total := decimal.Zero
	for _, w := range a.weights {
		total = total.Add(w)
	}
	return total
}

func EMA20(lastEMA, newPrice decimal.Decimal) decimal.Decimal {
	return EMA(lastEMA, newPrice, 20)
}

func EMA(lastEMA, newPrice decimal.Decimal, intervals int) decimal.Decimal {
	if intervals <= 0 {
		return decimal.Zero
	}
	smoothingFactor := decimal.NewFromInt(2)
	alpha := smoothingFactor.Div(decimal.NewFromInt32(int32(intervals)).Add(decimal.NewFromInt(1)))
	return alpha.Mul(newPrice.Sub(lastEMA)).Add(lastEMA)
}
