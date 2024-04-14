package quotes

import (
	"fmt"
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

	go func() {
		marketTrades := make(map[Market][]TradeEvent)
		timer := time.NewTimer(time.Duration(config.Index.TradesBufferSeconds) * time.Second)
		for {
			select {
			case trade := <-aggregated:
				marketTrades[trade.Market] = append(marketTrades[trade.Market], trade)
			case <-timer.C:
				for market, trades := range marketTrades {
					event := combineTrades(trades)
					if event != nil {
						indexPrice, ok := strategy.calculateIndexPrice(*event)
						if ok && event.Source != DriverInternal {
							if event.Market.convertTo != "" {
								event.Market.quoteUnit = event.Market.convertTo
							}
							event.Price = indexPrice
							outbox <- *event
						}
						marketTrades[market] = nil
					}
				}
				timer.Reset(time.Duration(config.Index.TradesBufferSeconds) * time.Second)
			}
		}
	}()

	return &indexAggregator{
		drivers:        drivers,
		marketsMapping: marketsMapping,
		aggregated:     aggregated,
	}
}

// TODO: add driver weights adjustment
func combineTrades(trades []TradeEvent) *TradeEvent {
	if len(trades) == 0 {
		return nil
	}

	totalBuyAmount := decimal.Zero
	totalSellAmount := decimal.Zero
	totalBuyValue := decimal.Zero
	totalSellValue := decimal.Zero

	for _, trade := range trades {
		if trade.TakerType == TakerTypeBuy {
			totalBuyAmount = totalBuyAmount.Add(trade.Amount)
			totalBuyValue = totalBuyValue.Add(trade.Amount.Mul(trade.Price))
		} else if trade.TakerType == TakerTypeSell {
			totalSellAmount = totalSellAmount.Add(trade.Amount)
			totalSellValue = totalSellValue.Add(trade.Amount.Mul(trade.Price))
		}
	}

	if totalBuyAmount.Equal(decimal.Zero) && totalSellAmount.Equal(decimal.Zero) {
		return nil
	}

	var avgBuyPrice, avgSellPrice decimal.Decimal
	if !totalBuyAmount.IsZero() {
		avgBuyPrice = totalBuyValue.Div(totalBuyAmount)
	}
	if !totalSellAmount.IsZero() {
		avgSellPrice = totalSellValue.Div(totalSellAmount)
	}

	var netSide TakerType
	var netPrice decimal.Decimal

	netAmount := totalBuyAmount.Sub(totalSellAmount)
	if netAmount.GreaterThan(decimal.Zero) {
		netSide = TakerTypeBuy
		netPrice = avgBuyPrice
	} else {
		netSide = TakerTypeSell
		netPrice = avgSellPrice
		netAmount = netAmount.Abs()
	}

	return &TradeEvent{
		Source:    DriverType{"index"},
		Market:    trades[0].Market,
		Price:     netPrice,
		Amount:    netAmount,
		Total:     netPrice.Mul(netAmount),
		TakerType: netSide,
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
		newStrategyVWA(withCustomPriceCacheVWA(newPriceCacheVWA(defaultWeightsMap, config.Index.BatchedCached, time.Duration(config.Index.BatchesBufferMinutes)*time.Minute))),
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
			loggerIndex.Infow("starting driver for index", "driver", d.ActiveDrivers()[0])
			if err := d.Start(); err != nil {
				loggerIndex.Warnw("failed to start driver", "driver", d.ActiveDrivers()[0], "error", err)
			}
		}(d)
	}

	wg.Wait()
	return nil
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
