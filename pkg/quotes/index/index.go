package index

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver"
)

var (
	logger                = log.Logger("index-aggregator")
	defaultMarketsMapping = map[string][]string{"usd": {"weth"}}

	maxAllowedPrice  = decimal.NewFromFloat(1e6)
	minAllowedAmount = decimal.NewFromFloat(1e-18)
)

type Config struct {
	TradesCached   int                 `yaml:"trades_cached" env:"TRADES_CACHED" env-default:"20"`
	BufferWindow   time.Duration       `yaml:"buffer_window" env:"BUFFER_WINDOW" env-default:"5s"`
	MarketsMapping map[string][]string `yaml:"markets_mapping" env:"MARKETS_MAPPING"`
	// MaxPriceDiff has default of `0.2` because our default leverage is 5x,
	// and so if the user opens order on his full balance, he'll get liquidated on 20% price change.
	MaxPriceDiff string `yaml:"max_price_diff" env:"MAX_PRICE_DIFF" env-default:"0.2"`
}

// New creates an instance of IndexAggregator with VWA strategy and default drivers weights.
//
// Params:
//   - drivers: a list of drivers to aggregate trades from
//   - config: index configuration
//   - outbox: a channel where the driver sends aggregated trades to
//   - inbox:  an optional channel where the package user can send trades from his own source.
//     If you don't { have / need } your own source, pass `nil` here.
//   - external: an optional adapter to fetch historical data from instead of querying RPC provider,
//     If you don't { have / need } a historical data adapter, pass `nil` here.
func New(
	drivers []driver.Driver,
	config Config,
	outbox chan<- common.TradeEvent,
	inbox <-chan common.TradeEvent,
	external driver.HistoricalDataDriver,
) (driver.Driver, error) {
	marketsMapping := config.MarketsMapping
	if marketsMapping == nil {
		marketsMapping = defaultMarketsMapping
	}

	priceCache := newPriceCache(
		defaultWeightsMap,
		config.TradesCached,
		config.BufferWindow,
	)
	strategy := newIndexStrategy(withCustomPriceCache(priceCache))

	return newIndexAggregator(
		drivers,
		config,
		marketsMapping,
		strategy,
		outbox,
		inbox,
		external,
	)
}

type indexAggregator struct {
	drivers        []driver.Driver
	marketsMapping map[string][]string
	maxPriceDiff   decimal.Decimal
	strategy       priceCalculator
	aggregator     chan common.TradeEvent
	outbox         chan<- common.TradeEvent
	history        driver.HistoricalDataDriver
}

type priceCalculator interface {
	calculateIndexPrice(trade common.TradeEvent) (decimal.Decimal, bool)
	getLastPrice(market common.Market) decimal.Decimal
	setLastPrice(market common.Market, price decimal.Decimal)
}

// newIndexAggregator creates a new instance of IndexAggregator.
func newIndexAggregator(
	drivers []driver.Driver,
	config Config,
	marketsMapping map[string][]string,
	strategy priceCalculator,
	outbox chan<- common.TradeEvent,
	inbox <-chan common.TradeEvent,
	history driver.HistoricalDataDriver,
) (driver.Driver, error) {
	aggregator := make(chan common.TradeEvent, 128)
	if inbox != nil {
		go func() {
			for tradeEvent := range inbox {
				aggregator <- tradeEvent
			}
		}()
	}

	maxPriceDiff, err := decimal.NewFromString(config.MaxPriceDiff)
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
			logger.Warnw("skipping trade with price out of range",
				"max_price", maxAllowedPrice,
				"event", event)
			continue
		}
		if event.Amount.LessThan(minAllowedAmount) {
			logger.Warnw("skipping trade with amount out of range",
				"min_amount", minAllowedAmount,
				"event", event)
			continue
		}
		if event.Total.IsZero() {
			logger.Warnw("skipping zeroes trade", "event", event)
			continue
		}

		lastPrice := a.strategy.getLastPrice(event.Market)
		if !lastPrice.IsZero() && isPriceOutOfRange(event.Price, lastPrice, a.maxPriceDiff) {
			logger.Warnw("skipping incoming outlier trade",
				"event", event,
				"last_price", lastPrice)
			continue
		}

		if event.Market.mainQuote != "" {
			event.Market.quoteUnit = event.Market.mainQuote
		}

		indexPrice, ok := a.strategy.calculateIndexPrice(event)
		if !ok || event.Source == common.DriverInternal {
			continue
		}

		if event.Market.convertTo != "" {
			event.Market.quoteUnit = event.Market.convertTo
		}
		event.Price = indexPrice
		event.Source = common.DriverType{"index/" + event.Source.String()}

		a.strategy.setLastPrice(event.Market, event.Price)

		baseMarkets, ok := defaultMarketsMapping[event.Market.quoteUnit]
		if !ok || slices.Contains(baseMarkets, event.Market.baseUnit) {
			continue
		}

		// Double check to avoid broken trades
		if event.Amount.IsZero() || event.Price.IsZero() || event.Total.IsZero() {
			logger.Warnw("skipping zeroed trade", "event", event)
			continue
		}
		a.outbox <- event
	}
}

func (a *indexAggregator) ActiveDrivers() []common.DriverType {
	drivers := make([]common.DriverType, 0, len(a.drivers))
	for _, d := range a.drivers {
		drivers = append(drivers, d.ActiveDrivers()...)
	}
	return drivers
}

func (a *indexAggregator) ExchangeType() common.ExchangeType {
	return common.ExchangeTypeHybrid
}

// Start starts all drivers from the provided config.
func (a *indexAggregator) Start() error {
	var g errgroup.Group
	g.SetLimit(-1)

	for _, d := range a.drivers {
		d := d
		g.Go(func() error {
			logger.Infow("starting driver for index", "driver", d.ActiveDrivers()[0])
			if err := d.Start(); err != nil {
				logger.Errorw("failed to start driver", "driver", d.ActiveDrivers()[0], "error", err)
				return err
			}

			go func() {
				for quoteMarket, baseMarkets := range defaultMarketsMapping {
					for _, baseMarket := range baseMarkets {
						m := common.NewMarket(baseMarket, quoteMarket)
						if err := d.Subscribe(m); err != nil {
							logger.Errorw("failed to subscribe to default market",
								"driver", d.ActiveDrivers()[0],
								"market", m,
								"error", err)
							continue
						}
					}
				}
			}()
			return nil
		})
	}

	return g.Wait()
}

func (a *indexAggregator) Subscribe(m common.Market) error {
	for _, d := range a.drivers {
		logger.Infow("subscribing", "driver", d.ActiveDrivers()[0], "market", m)

		if err := d.Subscribe(m); err != nil {
			if d.ExchangeType() == common.ExchangeTypeDEX {
				for _, convertFrom := range a.marketsMapping[m.quoteUnit] {
					// TODO: check if base and quote are same
					m := common.NewMarketDerived(m.baseUnit, convertFrom, m.quoteUnit)
					if err := d.Subscribe(m); err != nil {
						logger.Infow("skipping market",
							"driver", d.ActiveDrivers()[0],
							"market", m,
							"is_disabled", errors.Is(err, common.ErrMarketDisabled),
							"error", err)
						continue
					}
					logger.Infow("subscribed to helper market",
						"driver", d.ActiveDrivers()[0],
						"market", fmt.Sprintf("%s/%s", m.baseUnit, convertFrom))
				}
			}
		}
		logger.Infow("subscribed", "driver", d.ActiveDrivers()[0], "market", m)
	}
	return nil
}

func (a *indexAggregator) Unsubscribe(m common.Market) error {
	var g errgroup.Group

	for _, d := range a.drivers {
		d := d
		m := m
		g.Go(func() error {
			if err := d.Unsubscribe(m); err != nil {
				logger.Warnw("failed to unsubscribe", "driver", d.ActiveDrivers()[0], "market", m, "error", err.Error())
				return err
			}
			return nil
		})
	}

	return g.Wait()
}

func (a *indexAggregator) Stop() error {
	var g errgroup.Group
	g.SetLimit(-1)

	for _, d := range a.drivers {
		d := d
		g.Go(func() error { return d.Stop() })
	}

	return g.Wait()
}

func (a *indexAggregator) HistoricalData(ctx context.Context, market common.Market, window time.Duration, limit uint64) ([]common.TradeEvent, error) {
	base := market.Base()
	quote := market.Quote()
	if quote == "usd" {
		quote += "t" // to be USDT
	}
	m := common.NewMarket(base, quote)

	trades, err := driver.FetchHistoryDataFromExternalSource(ctx, a.history, m, window, limit, logger)
	if err == nil && len(trades) > 0 {
		return trades, nil
	}

	trades, err = a.fetchHistoricalDataByExchangeType(ctx, common.ExchangeTypeCEX, m, window, limit)
	if err != nil {
		return nil, err
	}

	// NOTE: DEXes are not reliable enough in terms of market trend stability
	// to be used as a source of historical data.
	if len(trades) == 0 {
		logger.Infow("no historical data found on CEXes, querying DEXes",
			"market", m,
			"window", window)
		trades, err = a.fetchHistoricalDataByExchangeType(ctx, common.ExchangeTypeDEX, m, window, limit)
		if err != nil {
			return nil, err
		}
	}

	logger.Infow("fetched historical data",
		"market", m,
		"window", window,
		"trades_num", len(trades))

	// Trades need to be sorted
	// since they are fetched from different sources
	// and may be not ordered by time.
	common.SortTradeEventsInPlace(trades)

	return trades, nil
}

func (a *indexAggregator) fetchHistoricalDataByExchangeType(
	ctx context.Context,
	typ common.ExchangeType,
	market common.Market,
	window time.Duration,
	limit uint64,
) ([]common.TradeEvent, error) {
	var trades []common.TradeEvent

	for _, d := range a.drivers {
		if d.ExchangeType() != typ {
			continue
		}

		data, err := d.HistoricalData(ctx, market, window, limit)
		if err != nil {
			continue
		}
		trades = append(trades, data...)
	}

	return trades, nil
}

func isPriceOutOfRange(eventPrice, lastPrice, maxPriceDiff decimal.Decimal) bool {
	diff := eventPrice.Sub(lastPrice).Abs().Div(lastPrice)
	return diff.GreaterThan(maxPriceDiff)
}
