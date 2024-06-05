package quotes

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

var (
	loggerIndex           = log.Logger("index-aggregator")
	defaultMarketsMapping = map[string][]string{"usd": {"weth"}}

	maxAllowedPrice  = decimal.NewFromFloat(1e6)
	minAllowedAmount = decimal.NewFromFloat(1e-18)
)

// newIndex creates a new instance of IndexAggregator with VWA strategy and default drivers weights.
func newIndex(config Config, outbox chan<- TradeEvent, inbox <-chan TradeEvent, history HistoricalData) (Driver, error) {
	marketsMapping := config.Index.MarketsMapping
	if marketsMapping == nil {
		marketsMapping = defaultMarketsMapping
	}

	return newIndexAggregator(
		config,
		marketsMapping,
		newIndexStrategy(withCustomPriceCache(newPriceCache(defaultWeightsMap, config.Index.TradesCached, time.Duration(config.Index.BufferSeconds)*time.Second))),
		outbox,
		inbox,
		history,
	)
}

type indexAggregator struct {
	drivers        []Driver
	marketsMapping map[string][]string
	maxPriceDiff   decimal.Decimal
	strategy       priceCalculator
	aggregator     chan TradeEvent
	outbox         chan<- TradeEvent
	history        HistoricalData
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
	inbox <-chan TradeEvent,
	history HistoricalData,
) (Driver, error) {
	aggregator := make(chan TradeEvent, 128)
	if inbox != nil {
		go func() {
			for tradeEvent := range inbox {
				aggregator <- tradeEvent
			}
		}()
	}

	drivers := make([]Driver, 0, len(config.Drivers))
	for _, d := range config.Drivers {
		loggerIndex.Infow("creating new driver", "driver", d)
		driverConfig, err := config.GetByDriverType(d)
		if err != nil {
			return nil, err
		}

		driver, err := NewDriver(driverConfig, aggregator, nil, nil)
		if err != nil {
			loggerIndex.Warnw("failed to create driver", "driver", d, "error", err)
			continue
		}
		drivers = append(drivers, driver)
	}

	maxPriceDiff, err := decimal.NewFromString(config.Index.MaxPriceDiff)
	if err != nil {
		return nil, err
	}

	index := &indexAggregator{
		drivers:        drivers,
		marketsMapping: marketsMapping,
		maxPriceDiff:   maxPriceDiff,
		strategy:       strategy,
		aggregator:     aggregator,
		outbox:         outbox,
		history:        history,
	}

	go index.computeAggregatePrice()

	return index, nil
}

func (a *indexAggregator) computeAggregatePrice() {
	for event := range a.aggregator {
		if event.Price.GreaterThanOrEqual(maxAllowedPrice) {
			loggerIndex.Warnw("skipping trade with price out of range",
				"max_price", maxAllowedPrice,
				"event", event)
			continue
		}
		if event.Amount.LessThan(minAllowedAmount) {
			loggerIndex.Warnw("skipping trade with amount out of range",
				"min_amount", minAllowedAmount,
				"event", event)
			continue
		}
		if event.Total.IsZero() {
			loggerIndex.Warnw("skipping zeroes trade", "event", event)
			continue
		}

		lastPrice := a.strategy.getLastPrice(event.Market)
		if !lastPrice.IsZero() && isPriceOutOfRange(event.Price, lastPrice, a.maxPriceDiff) {
			loggerIndex.Warnw("skipping incoming outlier trade",
				"event", event,
				"last_price", lastPrice)
			continue
		}

		if event.Market.mainQuote != "" {
			event.Market.quoteUnit = event.Market.mainQuote
		}

		indexPrice, ok := a.strategy.calculateIndexPrice(event)
		if !ok || event.Source == DriverInternal {
			continue
		}

		if event.Market.convertTo != "" {
			event.Market.quoteUnit = event.Market.convertTo
		}
		event.Price = indexPrice
		event.Source = DriverType{"index/" + event.Source.String()}

		a.strategy.setLastPrice(event.Market, event.Price)

		baseMarkets, ok := defaultMarketsMapping[event.Market.quoteUnit]
		if !ok || slices.Contains(baseMarkets, event.Market.baseUnit) {
			continue
		}

		// Double check to avoid broken trades
		if event.Amount.IsZero() || event.Price.IsZero() || event.Total.IsZero() {
			loggerIndex.Warnw("skipping zeroed trade", "event", event)
			continue
		}
		a.outbox <- event
	}
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
						loggerIndex.Infow("skipping market", "driver", d.ActiveDrivers()[0], "market", m, "error", err)
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

func (a *indexAggregator) HistoricalData(ctx context.Context, market Market, window time.Duration, limit uint64) ([]TradeEvent, error) {
	base := market.Base()
	quote := market.Quote()
	if quote == "usd" {
		quote += "t" // to be USDT
	}
	m := NewMarket(base, quote)

	trades, err := fetchHistoryDataFromExternalSource(ctx, a.history, m, window, limit, loggerIndex)
	if err == nil && len(trades) > 0 {
		return trades, nil
	}

	trades, err = a.fetchHistoricalDataByExchangeType(ctx, ExchangeTypeCEX, m, window, limit)
	if err != nil {
		return nil, err
	}

	// NOTE: DEXes are not reliable enough in terms of market trend stability
	// to be used as a source of historical data.
	if len(trades) == 0 {
		loggerIndex.Infow("no historical data found on CEXes, querying DEXes",
			"market", m,
			"window", window)
		trades, err = a.fetchHistoricalDataByExchangeType(ctx, ExchangeTypeDEX, m, window, limit)
		if err != nil {
			return nil, err
		}
	}

	loggerIndex.Infow("fetched historical data",
		"market", m,
		"window", window,
		"trades_num", len(trades))

	// Trades need to be sorted
	// since they are fetched from different sources
	// and may be not ordered by time.
	sortTradeEventsInPlace(trades)

	return trades, nil
}

func (a *indexAggregator) fetchHistoricalDataByExchangeType(
	ctx context.Context,
	typ ExchangeType,
	market Market,
	window time.Duration,
	limit uint64,
) ([]TradeEvent, error) {
	var trades []TradeEvent

	for _, driver := range a.drivers {
		if driver.ExchangeType() != typ {
			continue
		}

		data, err := driver.HistoricalData(ctx, market, window, limit)
		if err != nil {
			continue
		}
		trades = append(trades, data...)
	}

	return trades, nil
}

func isValidNonZero(x *big.Int) bool {
	// Note that negative values are allowed
	// as they represent a reduction in the balance of the pool.
	return x != nil && x.Sign() != 0
}

func isPriceOutOfRange(eventPrice, lastPrice, maxPriceDiff decimal.Decimal) bool {
	diff := eventPrice.Sub(lastPrice).Abs().Div(lastPrice)
	return diff.GreaterThan(maxPriceDiff)
}
