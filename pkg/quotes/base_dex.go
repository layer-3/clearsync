package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"

	"github.com/layer-3/clearsync/pkg/safe"
)

// rpcRateLimiter limits the number of requests to the RPC provider for DEX drivers.
// As of spring 2024, Infura enables 10 req/s rate limit for free plan.
// A lower limit of 5 req/s is used here just to be safe.
var rpcRateLimiter = rate.NewLimiter(5, 1)

const errInfuraRateLimit = "project ID request rate exceeded"

type baseDEX[Event any, Contract any] struct {
	// Params
	driverType DriverType
	url        string
	assetsURL  string
	mappingURL string

	// Hooks
	postStart func(*baseDEX[Event, Contract]) error
	getPool   func(Market) ([]*dexPool[Event], error)
	parse     func(*Event, *dexPool[Event]) (TradeEvent, error)

	// State
	once    *once
	client  *ethclient.Client
	outbox  chan<- TradeEvent
	logger  *log.ZapEventLogger
	filter  Filter
	streams safe.Map[Market, event.Subscription]
	assets  safe.Map[string, poolToken]
	mapping safe.Map[string, []string]
}

type baseDexConfig[Event any, Contract any] struct {
	// Params
	DriverType DriverType
	URL        string
	AssetsURL  string
	MappingURL string

	// Hooks
	PostStartHook func(*baseDEX[Event, Contract]) error
	PoolGetter    func(Market) ([]*dexPool[Event], error)
	EventParser   func(*Event, *dexPool[Event]) (TradeEvent, error)

	// State
	Outbox chan<- TradeEvent
	Logger *log.ZapEventLogger
	Filter FilterConfig
}

func newBaseDEX[Event any, Contract any](config baseDexConfig[Event, Contract]) *baseDEX[Event, Contract] {
	return &baseDEX[Event, Contract]{
		// Params
		driverType: config.DriverType,
		url:        config.URL,
		assetsURL:  config.AssetsURL,
		mappingURL: config.MappingURL,

		// Hooks
		postStart: config.PostStartHook,
		getPool:   config.PoolGetter,
		parse:     config.EventParser,

		// State
		once:    newOnce(),
		client:  nil,
		outbox:  config.Outbox,
		logger:  config.Logger,
		filter:  NewFilter(config.Filter),
		streams: safe.NewMap[Market, event.Subscription](),
		assets:  safe.NewMap[string, poolToken](),
		mapping: safe.NewMap[string, []string](),
	}
}

func (b *baseDEX[Event, Contract]) Client() *ethclient.Client {
	return b.client
}

func (b *baseDEX[Event, Contract]) Assets() *safe.Map[string, poolToken] {
	return &b.assets
}

func (b *baseDEX[Event, Contract]) ActiveDrivers() []DriverType {
	return []DriverType{b.driverType}
}

func (b *baseDEX[Event, Contract]) ExchangeType() ExchangeType {
	return ExchangeTypeDEX
}

func (b *baseDEX[Event, Contract]) Start() error {
	var startErr error
	started := b.once.Start(func() {
		// Connect to the RPC provider

		if !(strings.HasPrefix(b.url, "ws://") || strings.HasPrefix(b.url, "wss://")) {
			startErr = fmt.Errorf("%s (got '%s')", ErrInvalidWsUrl, b.url)
			return
		}

		client, err := ethclient.Dial(b.url)
		if err != nil {
			startErr = fmt.Errorf("failed to connect to the Ethereum client: %w", err)
			return
		}
		b.client = client

		// Fetch assets

		assets, err := getAssets(b.assetsURL)
		if err != nil {
			startErr = fmt.Errorf("failed to fetch assets: %w", err)
			return
		}
		for _, asset := range assets {
			b.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}

		// Fetch mappings

		mapping, err := getMapping(b.mappingURL)
		if err != nil {
			startErr = fmt.Errorf("failed to fetch mapping: %w", err)
			return
		}
		for key, mapItem := range mapping {
			b.mapping.Store(key, mapItem)
		}

		// Run post-start hook

		if err := b.postStart(b); err != nil {
			startErr = err
			return
		}
	})

	if !started {
		return ErrAlreadyStarted
	}
	return startErr
}

func (b *baseDEX[Event, Contract]) Stop() error {
	var stopErr error
	stopped := b.once.Stop(func() {
		var g errgroup.Group
		g.SetLimit(10)

		b.streams.Range(func(market Market, _ event.Subscription) bool {
			if err := b.Unsubscribe(market); err != nil {
				stopErr = err
			}
			return true
		})

		stopErr = g.Wait()
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return stopErr
}

func (b *baseDEX[Event, Contract]) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return ErrNotStarted
	}

	// mapping map[BTC:[WBTC] ETH:[WETH] USD:[USDT USDC TUSD]]
	var mappingErr error
	b.mapping.Range(func(token string, mappings []string) bool {
		if token != strings.ToUpper(market.Quote()) {
			return true
		}

		for _, mappedToken := range mappings {
			market := NewMarketWithMainQuote(market.Base(), mappedToken, market.Quote())
			if err := debounce(b.logger, func() error { return b.Subscribe(market) }); err != nil {
				b.logger.Errorf("failed to subscribe to market %s: %s", market, err)
				mappingErr = err
			}
		}

		return true
	})
	if mappingErr != nil {
		return fmt.Errorf("failed to subscribe to helper markets: %w", mappingErr)
	}

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	pools, err := b.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %s: %s", market.StringWithoutMain(), err)
	}

	for _, pool := range pools {
		sink := make(chan *Event, 128)

		var sub event.Subscription
		err := debounce(b.logger, func() error {
			opts := &bind.WatchOpts{Context: context.TODO()}
			sub, err = pool.Contract.WatchSwap(opts, sink, []common.Address{}, []common.Address{})
			return err
		})
		if err != nil {
			close(sink)
			return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
		}

		go b.watchSwap(pool, sink, sub)
		go b.streams.Store(market, sub) // to not block the loop since it's a blocking call with mutex under the hood
		recordSubscribed(b.driverType, market)
	}

	return nil
}

func (b *baseDEX[Event, Contract]) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return ErrNotStarted
	}

	stream, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	stream.Unsubscribe()

	b.streams.Delete(market)
	recordUnsubscribed(b.driverType, market)
	return nil
}

// debounce is a wrapper around the rate limiter
// that retries the request if it fails with rate limit error.
func debounce(logger *log.ZapEventLogger, f func() error) error {
	for {
		if err := rpcRateLimiter.Wait(context.TODO()); err != nil {
			logger.Warnf("failed to aquire rate limiter: %s", err)
		}

		err := f()
		if err == nil {
			return nil
		}
		if strings.Contains(err.Error(), errInfuraRateLimit) {
			logger.Infow("rate limit exceeded, retrying", "error", err)
			continue // retry the request after a while
		}
		return err
	}
}

func (b *baseDEX[Event, Contract]) SetInbox(_ <-chan TradeEvent) {}

func (b *baseDEX[Event, Contract]) watchSwap(
	pool *dexPool[Event],
	sink chan *Event,
	sub event.Subscription,
) {
	defer close(sink)
	for {
		select {
		case err := <-sub.Err():
			b.logger.Warnw("connection failed, resubscribing", "market", pool.Market, "err", err)
			if _, ok := b.streams.Load(pool.Market); !ok {
				break // market was unsubscribed earlier
			}
			if err := b.Unsubscribe(pool.Market); err != nil {
				b.logger.Errorw("failed to resubscribe", "market", pool.Market, "err", err)
			}
			if err := b.Subscribe(pool.Market); err != nil {
				b.logger.Errorw("failed to resubscribe", "market", pool.Market, "err", err)
			}
			return
		case swap := <-sink:
			trade, err := b.parse(swap, pool)
			if err != nil {
				b.logger.Errorw("failed to parse swap event",
					"error", err,
					"market", pool.Market,
					"swap", swap,
					"pool", pool,
				)
				continue
			}

			skip := !b.filter.Allow(trade)
			b.logger.Infow("parsed trade",
				"skip", skip,
				"trade", trade,
				"swap", swap,
				"pool", pool)

			if skip {
				continue
			}
			b.outbox <- trade
		}
	}
}

type dexPool[Event any] struct {
	Contract   dexEventWatcher[Event]
	Address    common.Address // not used in code but is useful for logging
	BaseToken  poolToken
	QuoteToken poolToken
	Market     Market
	Reversed   bool
}

type dexEventWatcher[Event any] interface {
	WatchSwap(opts *bind.WatchOpts, sink chan<- *Event, from []common.Address, to []common.Address) (event.Subscription, error)
}

type poolToken struct {
	Name     string
	Address  common.Address
	Symbol   string
	Decimals decimal.Decimal
	ChainId  uint
	LogoURI  string
}

func getTokens(
	assets *safe.Map[string, poolToken],
	market Market,
	logger *log.ZapEventLogger,
) (baseToken poolToken, quoteToken poolToken, err error) {
	baseToken, ok := assets.Load(strings.ToUpper(market.Base()))
	if !ok {
		return poolToken{}, poolToken{}, fmt.Errorf("base tokens does not exist for market %s", market.StringWithoutMain())
	}
	logger.Infow("found base token", "address", baseToken.Address, "market", market.StringWithoutMain())

	quoteToken, ok = assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		return poolToken{}, poolToken{}, fmt.Errorf("quote tokens does not exist for market %s", market.StringWithoutMain())
	}
	logger.Infow("found quote token", "address", quoteToken.Address, "market", market.StringWithoutMain())

	return baseToken, quoteToken, nil
}

func getAssets(assetsURL string) ([]poolToken, error) {
	resp, err := http.Get(assetsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var assets map[string][]poolToken
	if err := json.Unmarshal(body, &assets); err != nil {
		return nil, err
	}
	return assets["tokens"], nil
}

func getMapping(url string) (map[string][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mappings map[string]map[string][]string
	if err := json.Unmarshal(body, &mappings); err != nil {
		return nil, err
	}
	return mappings["tokens"], nil
}
