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
				event.Source = DriverType{"index/" + event.Source.slug}
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

func (a *indexAggregator) Name() DriverType {
	return DriverIndex
}

func (b *indexAggregator) Type() Type {
	return TypeHybrid
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
		loggerIndex.Info("subscribing ", d.Name().slug)
		if err := d.Subscribe(m); err != nil {
			if d.Type() == TypeDEX {
				for _, convertFrom := range a.marketsMapping[m.quoteUnit] {
					if err := d.Subscribe(NewDerivedMerket(m.baseUnit, convertFrom, m.quoteUnit)); err != nil {
						loggerIndex.Infof("%s: skipping %s :", d.Name().slug, convertFrom, err.Error())
						continue
					}
					loggerIndex.Infof("%s helper market found: %s/%s", d.Name().slug, m.baseUnit, convertFrom)
				}
			}
			loggerIndex.Warnf("failed to subscribe for %s %s market: %s: ", d.Name().slug, m, err.Error())
		}
	}
	return nil
}

func (a *indexAggregator) Unsubscribe(m Market) error {
	for _, d := range a.drivers {
		if err := d.Unsubscribe(m); err != nil {
			loggerIndex.Warnf("%s unsubsctiption error: ", d.Name().slug, err.Error())
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
