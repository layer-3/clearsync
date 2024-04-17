package quotes

import (
	"fmt"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
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
	getLastPrice(market Market) decimal.Decimal
	setLastPrice(market Market, price decimal.Decimal)
}

// newIndexAggregator creates a new instance of IndexAggregator.
func newIndexAggregator(config Config, marketsMapping map[string][]string, strategy priceCalculator, outbox chan<- TradeEvent) Driver {
	aggregated := make(chan TradeEvent, 128)

	drivers := make([]Driver, 0, len(config.Drivers))
	for _, d := range config.Drivers {
		loggerIndex.Infow("creating driver as an index driver", "driver", d)
		driverConfig, err := config.GetByDriverType(d)
		if err != nil {
			panic(err) // impossible case if config structure is not amended
		}

		driver, err := NewDriver(driverConfig, aggregated)
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

	toCombine := make(chan TradeEvent, 128)
	go func() {
		for event := range aggregated {
			lastPrice := strategy.getLastPrice(event.Market)
			if lastPrice != decimal.Zero {
				if isPriceOutOfRange(event.Price, lastPrice, maxPriceDiff) {
					loggerIndex.Warnf("skipping incoming outlier trade. Source: %s, Market: %s, Price: %s, Amount:%s", event.Source, event.Market, event.Price, event.Amount)
					continue
				}
			}

			indexPrice, ok := strategy.calculateIndexPrice(event)
			if ok && event.Source != DriverInternal {
				if event.Market.convertTo != "" {
					event.Market.quoteUnit = event.Market.convertTo
				}
				event.Price = indexPrice
				if event.Source != DriverBinance {
					loggerIndex.Infow("agregated event", "event", event)
				}
				event.Source = DriverType{"index/" + event.Source.String()}

				toCombine <- event
			}
		}
	}()

	go func() {
		marketTrades := make(map[Market][]TradeEvent)
		timer := time.NewTimer(time.Duration(config.Index.BatchBufferSeconds) * time.Second)
		for {
			select {
			case trade := <-toCombine:
				marketTrades[trade.Market] = append(marketTrades[trade.Market], trade)
			case <-timer.C:
				for market, trades := range marketTrades {
					event := combineTrades(trades)
					if event != nil {
						marketTrades[market] = nil
						strategy.setLastPrice(event.Market, event.Price)
						outbox <- *event
					}
				}
				timer.Reset(time.Duration(config.Index.BatchBufferSeconds) * time.Second)
			}
		}
	}()

	return &indexAggregator{
		drivers:        drivers,
		marketsMapping: marketsMapping,
		aggregated:     aggregated,
	}
}

func combineTrades(trades []TradeEvent) *TradeEvent {
	if len(trades) == 0 {
		return nil
	}

	totalAmount := decimal.Zero
	totalValue := decimal.Zero
	netAmount := decimal.Zero

	for _, trade := range trades {
		totalAmount = totalAmount.Add(trade.Amount)
		totalValue = totalValue.Add(trade.Amount.Mul(trade.Price))

		// Update net amount to determine net side (buy or sell)
		if trade.TakerType == TakerTypeBuy {
			netAmount = netAmount.Add(trade.Amount)
		} else if trade.TakerType == TakerTypeSell {
			netAmount = netAmount.Sub(trade.Amount)
		}
	}

	if totalAmount.Equal(decimal.Zero) {
		return nil
	}

	avgPrice := totalValue.Div(totalAmount)
	// Determine net side (buy or sell)
	var side TakerType
	if netAmount.GreaterThanOrEqual(decimal.Zero) {
		side = TakerTypeSell // "buy" (yes, it looks inverted)
	} else {
		side = TakerTypeBuy // "sell"
		netAmount = netAmount.Abs()
	}

	return &TradeEvent{
		Source:    DriverType{"index"},
		Market:    trades[0].Market,
		Price:     avgPrice,
		Amount:    totalAmount,
		Total:     avgPrice.Mul(totalAmount),
		TakerType: side,
		CreatedAt: time.Now(),
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
	var g errgroup.Group
	g.SetLimit(10)

	go func() {
		for t := range a.inbox {
			a.aggregated <- t
		}
	}()

	for _, d := range a.drivers {
		d := d
		g.Go(func() error {
			loggerIndex.Infow("starting driver for index", "driver", d.ActiveDrivers()[0])
			if err := d.Start(); err != nil {
				loggerIndex.Errorw("failed to start driver", "driver", d.ActiveDrivers()[0], "error", err)
				return err
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
					if err := d.Subscribe(NewDerivedMerket(m.baseUnit, convertFrom, m.quoteUnit)); err != nil {
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
	for _, d := range a.drivers {
		if err := d.Unsubscribe(m); err != nil {
			loggerIndex.Warnw("failed to unsubscribe", "driver", d.ActiveDrivers()[0], "market", m, "error", err.Error())
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

func isPriceOutOfRange(eventPrice, lastPrice, maxPriceDiff decimal.Decimal) bool {
	diff := eventPrice.Sub(lastPrice).Abs().Div(lastPrice)
	return diff.GreaterThan(maxPriceDiff)
}
