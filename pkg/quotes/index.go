package quotes

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type IndexAggregator struct {
	weightsMap    map[DriverType]decimal.Decimal
	activeWeights map[DriverType]decimal.Decimal

	drivers []Driver
	prices  PriceCache

	aggregated chan TradeEvent
	outbox     chan<- TradeEvent
}

func NewIndex(driverConfigs []Config, weightsMap map[DriverType]decimal.Decimal, outbox chan<- TradeEvent) (*IndexAggregator, error) {
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
		prices:        NewPricesCache(),
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
	lastEMA, timestamp := a.prices.GetEMA(event.Market)
	if lastEMA == decimal.Zero {
		lastEMA = event.Price
	}

	// TODO: delete
	fmt.Println("eventPrice            :", event.Price)

	// Calculate an EMA based on new price
	newEMA := EMA20(lastEMA, event.Price)

	// Get difference between new and previous values
	delta := lastEMA.Sub(newEMA)

	weight := a.activeWeights[event.Source]
	// Influence ema based on amount change multiplied by weight
	weightedNewEMA := lastEMA.Add(delta.Mul(weight))
	// TODO: find the way to include trade amount in this logic

	// Initial weights logic that didn't work:
	// newEMA := EMA20(lastEMA, event.Price.Mul(event.Amount.Mul(weight).Div(a.totalWeights())))

	if time.Now().Unix() >= timestamp+5 { // 5 seconds for testing
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\nupdate EMA")
		a.prices.UpdateEMA(event.Market, weightedNewEMA)
	}

	event.Price = weightedNewEMA
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

func EMA(lastEMA, newPrice decimal.Decimal, intervals int) decimal.Decimal {
	if intervals <= 0 {
		return decimal.Zero
	}
	smoothingFactor := decimal.NewFromInt(2)
	alpha := smoothingFactor.Div(decimal.NewFromInt32(int32(intervals)).Add(decimal.NewFromInt(1)))
	return alpha.Mul(newPrice.Sub(lastEMA)).Add(lastEMA)
}
