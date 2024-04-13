package quotes

import (
	"sync"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
)

var (
	loggerIndex           = log.Logger("index-aggregator")
	defaultMarketsMapping = map[string][]string{"usdc": {"eth", "weth", "matic"}}
)

type indexAggregator struct {
	drivers        []Driver
	marketsMapping map[string][]string
	inbox          <-chan TradeEvent
	aggregated     chan TradeEvent
}

type priceCalculator interface {
	calculateIndexPrice(trade TradeEvent) (decimal.Decimal, bool)
}

// newIndexAggregator creates a new instance of IndexAggregator.
func newIndexAggregator(config Config, marketsMapping map[string][]string, strategy priceCalculator, outbox chan<- TradeEvent) Driver {
	aggregated := make(chan TradeEvent, 128)

	drivers := make([]Driver, 0, len(config.Drivers))
	for _, d := range config.Drivers {
		driverConfig, err := config.GetByDriverType(d)
		if err != nil {
			panic(err) // impossible case if config structure is not amended
		}

		driver, err := NewDriver(driverConfig, aggregated)
		if err != nil {
			loggerIndex.Warnf("failed to create driver %s: %s", d, err.Error())
			continue
		}
		drivers = append(drivers, driver)
	}

	go func() {
		for event := range aggregated {
			indexPrice, ok := strategy.calculateIndexPrice(event)
			if ok && event.Source != DriverInternal {
				if event.Market.convertTo != "" {
					event.Market.quoteUnit = event.Market.convertTo
				}
				event.Price = indexPrice
				event.Source = DriverType{"index/" + event.Source.String()}
				outbox <- event
			}
		}
	}()

	return &indexAggregator{
		drivers:        drivers,
		marketsMapping: marketsMapping,
		aggregated:     aggregated,
	}
}

// newIndex creates a new instance of IndexAggregator with VWA strategy and default drivers weights.
func newIndex(config Config, outbox chan<- TradeEvent) Driver {
	marketsMapping := config.Index.MarketsMapping
	if marketsMapping == nil {
		marketsMapping = defaultMarketsMapping
	}

	return newIndexAggregator(
		config,
		marketsMapping,
		newStrategyVWA(withCustomPriceCacheVWA(newPriceCacheVWA(defaultWeightsMap, config.Index.TradesCached, time.Duration(config.Index.BufferMinutes)*time.Minute))),
		outbox,
	)
}

func (a *indexAggregator) SetInbox(inbox <-chan TradeEvent) {
	a.inbox = inbox
}

func (a *indexAggregator) ActiveDrivers() []DriverType {
	drivers := make([]DriverType, 0, len(a.drivers))
	for _, d := range a.drivers {
		drivers = append(drivers, d.ActiveDrivers()...)
	}
	return drivers
}

func (b *indexAggregator) ExchangeType() ExchangeType {
	return ExchangeTypeHybrid
}

// Start starts all drivers from the provided config.
func (a *indexAggregator) Start() error {
	var wg sync.WaitGroup

	go func() {
		for t := range a.inbox {
			a.aggregated <- t
		}
	}()

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
	for _, d := range a.drivers {
		loggerIndex.Infof("subscribing to %s", d.ActiveDrivers()[0])
		if err := d.Subscribe(m); err != nil {
			if d.ExchangeType() == ExchangeTypeDEX {
				for _, convertFrom := range a.marketsMapping[m.quoteUnit] {
					if err := d.Subscribe(NewDerivedMerket(m.baseUnit, convertFrom, m.quoteUnit)); err != nil {
						loggerIndex.Infof("%s: skipping %s :", d.ActiveDrivers()[0], convertFrom, err.Error())
						continue
					}
					loggerIndex.Infof("%s helper market found: %s/%s", d.ActiveDrivers()[0], m.baseUnit, convertFrom)
				}
			}
			loggerIndex.Warnf("failed to subscribe for %s %s market: %s: ", d.ActiveDrivers()[0], m, err.Error())
		}
	}
	return nil
}

func (a *indexAggregator) Unsubscribe(m Market) error {
	for _, d := range a.drivers {
		if err := d.Unsubscribe(m); err != nil {
			loggerIndex.Warnf("failed to unsubscribe from market %s of %s: %s", m, d.ActiveDrivers()[0], err.Error())
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
