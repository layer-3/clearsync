package quotes

import (
	"fmt"
	"slices"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

var (
	loggerIndex           = log.Logger("index-aggregator")
	defaultMarketsMapping = map[string][]string{"usd": {"weth", "matic"}}

	maxAllowedPrice  = decimal.NewFromInt(1e6)
	minAllowedAmount = decimal.NewFromInt(1e-18)
)

type indexAggregator struct {
	drivers        []Driver
	marketsMapping map[string][]string
	aggregated     chan TradeEvent
}

type priceCalculator interface {
	calculateIndexPrice(trade TradeEvent) (decimal.Decimal, bool)
	getLastPrice(market Market) decimal.Decimal
	setLastPrice(market Market, price decimal.Decimal)
}

// newIndexAggregator creates a new instance of IndexAggregator.
func newIndexAggregator(
	config Config,
	marketsMapping map[string][]string,
	strategy priceCalculator,
	outbox chan<- TradeEvent,
) Driver {
	inbox := make(chan TradeEvent, 128)

	drivers := make([]Driver, 0, len(config.Drivers))
	for _, d := range config.Drivers {
		loggerIndex.Infow("creating new driver", "driver", d)
		driverConfig, err := config.GetByDriverType(d)
		if err != nil {
			panic(err) // impossible case if config structure is not amended
		}

		driver, err := NewDriver(driverConfig, inbox)
		if err != nil {
			loggerIndex.Warnw("failed to create driver", "driver", d, "error", err)
			continue
		}
		drivers = append(drivers, driver)
	}

	maxPriceDiff, err := decimal.NewFromString(config.Index.MaxPriceDiff)
	if err != nil {
		loggerIndex.Fatalf("invalid max price diff config value", "driver", "error", err)
	}

	index := &indexAggregator{
		drivers:        drivers,
		marketsMapping: marketsMapping,
		aggregated:     inbox,
	}
	go index.computeAggregatePrice(inbox, maxPriceDiff, strategy, outbox)

	return index
}

func (a *indexAggregator) computeAggregatePrice(
	aggregated <-chan TradeEvent,
	maxPriceDiff decimal.Decimal,
	strategy priceCalculator,
	outbox chan<- TradeEvent,
) {
	for event := range aggregated {
		if event.Price.GreaterThanOrEqual(maxAllowedPrice) || event.Amount.LessThan(minAllowedAmount) {
			loggerIndex.Warnw("skipping trades with too big or too small amount", "event", event)
			continue
		}
		if event.Amount.IsZero() || event.Price.IsZero() || event.Total.IsZero() {
			loggerIndex.Warnw("skipping zeroes trades", "event", event)
			continue
		}

		lastPrice := strategy.getLastPrice(event.Market)
		if lastPrice.IsZero() && isPriceOutOfRange(event.Price, lastPrice, maxPriceDiff) {
			loggerIndex.Warnw("skipping incoming outlier trade",
				"event", event,
				"last_price", lastPrice)
			continue
		}

		if event.Market.mainQuote != "" {
			event.Market.quoteUnit = event.Market.mainQuote
		}

		indexPrice, ok := strategy.calculateIndexPrice(event)
		if !ok || event.Source == DriverInternal {
			continue
		}

		if event.Market.convertTo != "" {
			event.Market.quoteUnit = event.Market.convertTo
		}
		event.Price = indexPrice
		event.Source = DriverType{"index/" + event.Source.String()}

		strategy.setLastPrice(event.Market, event.Price)

		baseMarkets, ok := defaultMarketsMapping[event.Market.quoteUnit]
		if !ok || slices.Contains(baseMarkets, event.Market.baseUnit) {
			continue
		}

		// Double check to avoid broken trades
		if event.Amount.IsZero() || event.Price.IsZero() || event.Total.IsZero() {
			loggerIndex.Warnw("skipping zeroed trades", "event", event)
			continue
		}
		outbox <- event
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
		newIndexStrategy(withCustomPriceCache(newPriceCache(defaultWeightsMap, config.Index.TradesCached, time.Duration(config.Index.BufferSeconds)*time.Second))),
		outbox,
	)
}

func (a *indexAggregator) SetInbox(inbox <-chan TradeEvent) {
	go func() {
		for tradeEvent := range inbox {
			a.aggregated <- tradeEvent
		}
	}()
}

func (a *indexAggregator) ActiveDrivers() []DriverType {
	drivers := make([]DriverType, 0, len(a.drivers))
	for _, d := range a.drivers {
		drivers = append(drivers, d.ActiveDrivers()...)
	}
	return drivers
}

func (a *indexAggregator) ExchangeType() ExchangeType {
	return ExchangeTypeHybrid
}

// Start starts all drivers from the provided config.
func (a *indexAggregator) Start() error {
	var g errgroup.Group
	g.SetLimit(10)

	for _, d := range a.drivers {
		d := d
		g.Go(func() error {
			loggerIndex.Infow("starting driver for index", "driver", d.ActiveDrivers()[0])
			if err := d.Start(); err != nil {
				loggerIndex.Errorw("failed to start driver", "driver", d.ActiveDrivers()[0], "error", err)
				return err
			}

			for quoteMarket, baseMarkets := range defaultMarketsMapping {
				for _, baseMarket := range baseMarkets {
					m := NewMarket(baseMarket, quoteMarket)
					if err := d.Subscribe(m); err != nil {
						loggerIndex.Errorw("failed to subscribe to default market",
							"driver", d.ActiveDrivers()[0],
							"market", m,
							"error", err)
						continue
					}
				}
			}
			return nil
		})
	}

	return g.Wait()
}

func (a *indexAggregator) Subscribe(m Market) error {
	for _, d := range a.drivers {
		loggerIndex.Infow("subscribing", "driver", d.ActiveDrivers()[0], "market", m)
		if err := d.Subscribe(m); err != nil {
			if d.ExchangeType() == ExchangeTypeDEX {
				for _, convertFrom := range a.marketsMapping[m.quoteUnit] {
					// TODO: check if base and quote are same
					m := NewMarketDerived(m.baseUnit, convertFrom, m.quoteUnit)
					if err := d.Subscribe(m); err != nil {
						loggerIndex.Infow("skipping market", "driver", d.ActiveDrivers()[0], "market", convertFrom, "error", err)
						continue
					}
					loggerIndex.Infow("subscribed to helper market",
						"driver", d.ActiveDrivers()[0],
						"market", fmt.Sprintf("%s/%s", m.baseUnit, convertFrom))
				}
			}
		}
		loggerIndex.Infow("subscribed", "driver", d.ActiveDrivers()[0], "market", m)
	}
	return nil
}

func (a *indexAggregator) Unsubscribe(m Market) error {
	var g errgroup.Group

	for _, d := range a.drivers {
		d := d
		m := m
		g.Go(func() error {
			if err := d.Unsubscribe(m); err != nil {
				loggerIndex.Warnw("failed to unsubscribe", "driver", d.ActiveDrivers()[0], "market", m, "error", err.Error())
				return err
			}
			return nil
		})
	}

	return g.Wait()
}

func (a *indexAggregator) Stop() error {
	var g errgroup.Group
	g.SetLimit(10)

	for _, d := range a.drivers {
		d := d
		g.Go(func() error { return d.Stop() })
	}

	return g.Wait()
}

func isPriceOutOfRange(eventPrice, lastPrice, maxPriceDiff decimal.Decimal) bool {
	diff := eventPrice.Sub(lastPrice).Abs().Div(lastPrice)
	return diff.GreaterThan(maxPriceDiff)
}
