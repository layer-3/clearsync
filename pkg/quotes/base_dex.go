package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/layer-3/clearsync/pkg/debounce"
	"github.com/layer-3/clearsync/pkg/safe"
)

type baseDEX[Event any, Contract any] struct {
	// Params
	driverType DriverType
	url        string
	assetsURL  string
	mappingURL string
	idlePeriod time.Duration

	// Hooks
	postStart func(*baseDEX[Event, Contract]) error
	getPool   func(Market) ([]*dexPool[Event], error)
	parse     func(*Event, *dexPool[Event]) (TradeEvent, error)

	// State
	once   *once
	client *ethclient.Client
	outbox chan<- TradeEvent
	logger *log.ZapEventLogger
	filter Filter
	// streams maps market to a map of DEX pools.
	// The value of the map is a pointer to disallow copying of the underlying mutex
	streams safe.Map[Market, *safe.Map[common.Address, dexStream[Event]]]
	assets  safe.Map[string, poolToken]
	mapping safe.Map[string, []string]
}

type dexStream[Event any] struct {
	sub  event.Subscription
	sink chan *Event
}

type baseDexConfig[Event any, Contract any] struct {
	// Params
	DriverType DriverType
	URL        string
	AssetsURL  string
	MappingURL string
	IdlePeriod time.Duration

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
		idlePeriod: config.IdlePeriod,

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
		streams: safe.NewMap[Market, *safe.Map[common.Address, dexStream[Event]]](),
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

		b.streams.Range(func(market Market, _ *safe.Map[common.Address, dexStream[Event]]) bool {
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
			if err := debounce.Debounce(b.logger, func() error { return b.Subscribe(market) }); err != nil {
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
		fmt.Println("Market already subscribed", market)
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	pools, err := b.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %s: %s", market.StringWithoutMain(), err)
	}

	for _, pool := range pools {
		if err := b.subscribePool(pool); err != nil {
			return err
		}
	}

	return nil
}

func (b *baseDEX[Event, Contract]) subscribePool(pool *dexPool[Event]) error {
	watchCtx, cancel := context.WithCancel(context.TODO())
	sink := make(chan *Event, 128)

	var sub event.Subscription
	var err error
	err = debounce.Debounce(b.logger, func() error {
		opts := &bind.WatchOpts{Context: watchCtx}
		sub, err = pool.Contract.WatchSwap(opts, sink, []common.Address{}, []common.Address{})
		return err
	})
	if err != nil {
		close(sink)
		cancel()
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", pool.Market, err)
	}

	pools := safe.NewMap[common.Address, dexStream[Event]]()
	stream, _ := b.streams.LoadOrStore(pool.Market, &pools)
	stream.Store(pool.Address, dexStream[Event]{sub: sub, sink: sink})

	recordSubscribed(b.driverType, pool.Market)
	go b.watchSwap(cancel, pool, sink, sub)
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
	stream.UpdateInTx(func(stream map[common.Address]dexStream[Event]) {
		for _, s := range stream {
			s.sub.Unsubscribe()
			s.sub = nil
			// do not delete the sink channel
		}
	})

	recordUnsubscribed(b.driverType, market)
	return nil
}

func (b *baseDEX[Event, Contract]) SetInbox(_ <-chan TradeEvent) {}

func (b *baseDEX[Event, Contract]) watchSwap(
	cancel context.CancelFunc,
	pool *dexPool[Event],
	sink chan *Event,
	sub event.Subscription,
) {
	timer := time.NewTimer(b.idlePeriod)
	defer timer.Stop()

	for {
		select {
		case err := <-sub.Err():
			if err == nil {
				// A nil error indicates intentional unsubscribe
				b.logger.Infow("intentional unsubscribe received, stopping watch", "market", pool.Market)
				return
			}

			b.logger.Warnw("connection failed, resubscribing", "market", pool.Market, "err", err)
			if _, ok := b.streams.Load(pool.Market); !ok {
				return // market was unsubscribed earlier
			}

			for {
				if err := b.resubscribe(pool); err == nil {
					return
				}
				<-time.After(5 * time.Second)
			}

		case swap := <-sink:
			timer.Reset(b.idlePeriod)
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

		case <-timer.C:
			b.logger.Warnw("market inactivity detected", "market", pool.Market)
			cancel()
			for {
				if err := b.resubscribe(pool); err == nil {
					return
				}
				<-time.After(5 * time.Second)
			}
		}
	}
}

func (b *baseDEX[Event, Contract]) resubscribe(pool *dexPool[Event]) error {
	if _, ok := b.streams.Load(pool.Market); !ok {
		return nil // market was unsubscribed earlier
	}

	if err := b.subscribePool(pool); err != nil {
		b.logger.Errorw("failed to resubscribe", "market", pool.Market, "error", err)
		return err
	}

	b.logger.Infow("resubscribed", "market", pool.Market)
	return nil
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
