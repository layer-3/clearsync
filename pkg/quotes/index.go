package quotes

import (
	"sync"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
)

var loggerIndex = log.Logger("index-aggregator")

type indexAggregator struct {
	drivers         []Driver
	priceCalculator priceCalculator
	aggregated      chan TradeEvent
	outbox          chan<- TradeEvent
}

type priceCalculator interface {
	calculateIndex(trade TradeEvent) (decimal.Decimal, bool)
}

// NewIndexAggregator creates a new instance of IndexAggregator.
func NewIndexAggregator(driversConfigs []Config, strategy priceCalculator, outbox chan<- TradeEvent) Driver {
	aggregated := make(chan TradeEvent, 128)

	var drivers []Driver
	for _, d := range driversConfigs {
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
		drivers:         drivers,
		priceCalculator: strategy,
		outbox:          outbox,
		aggregated:      aggregated,
	}
}

// newIndex creates a new instance of IndexAggregator with default configs.
func newIndex(config Config, outbox chan<- TradeEvent) Driver {
	return NewIndexAggregator(AllDrivers, NewStrategyVWA(), outbox)
}

// ChangeStrategy allows index price calculation strategy to be changed.
func (a *indexAggregator) ChangeStrategy(newStrategy priceCalculator) {
	a.priceCalculator = newStrategy
}

func (a *indexAggregator) Name() DriverType {
	return DriverIndex
}

// Start starts all drivers from the provided config.
func (a *indexAggregator) Start() error {
	var wg sync.WaitGroup

	for _, d := range a.drivers {
		wg.Add(1)

		go func(d Driver) {
			defer wg.Done()

			if err := d.Start(); err != nil {
				loggerIndex.Warn(err.Error())
			}
		}(d)
	}

	wg.Wait()
	return nil
}

func (a *indexAggregator) Subscribe(m Market) error {
	go func() {
		for event := range a.aggregated {
			indexPrice, ok := a.priceCalculator.calculateIndex(event)
			if ok {
				event.Price = indexPrice
				event.Source = DriverIndex
				a.outbox <- event
			}
		}
	}()

	for _, d := range a.drivers {
		loggerIndex.Info("subscribing ", d.Name().slug)
		if err := d.Subscribe(m); err != nil {
			loggerIndex.Warnf("%s subsctiption error: ", d.Name().slug, err.Error())
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
