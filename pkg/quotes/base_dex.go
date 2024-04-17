package quotes

import (
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

	"github.com/layer-3/clearsync/pkg/safe"
)

type baseDEX[Event any, Contract any] struct {
	// Params
	once       *once
	driverType DriverType
	url        string
	assetsURL  string
	mappingURL string
	logger     *log.ZapEventLogger

	// Hooks
	postStart func(*baseDEX[Event, Contract]) error
	getPool   func(Market) ([]*dexPool[Event], error)
	parse     func(*Event, *dexPool[Event]) (TradeEvent, error)

	// State
	client  *ethclient.Client
	outbox  chan<- TradeEvent
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
	Logger     *log.ZapEventLogger

	// Hooks
	PostStartHook func(*baseDEX[Event, Contract]) error
	PoolGetter    func(Market) ([]*dexPool[Event], error)
	EventParser   func(*Event, *dexPool[Event]) (TradeEvent, error)

	// State
	Outbox chan<- TradeEvent
	Filter FilterConfig
}

func newBaseDEX[Event any, Contract any](config baseDexConfig[Event, Contract]) *baseDEX[Event, Contract] {
	return &baseDEX[Event, Contract]{
		// Params
		once:       newOnce(),
		driverType: config.DriverType,
		url:        config.URL,
		assetsURL:  config.AssetsURL,
		mappingURL: config.MappingURL,
		logger:     config.Logger,

		// Hooks
		postStart: config.PostStartHook,
		getPool:   config.PoolGetter,
		parse:     config.EventParser,

		// State
		client:  nil,
		outbox:  config.Outbox,
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
	stopped := b.once.Stop(func() {
		b.streams.Range(func(market Market, _ event.Subscription) bool {
			err := b.Unsubscribe(market)
			return err == nil
		})
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return nil
}

func (b *baseDEX[Event, Contract]) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return ErrNotStarted
	}

	// mapping map[BTC:[WBTC] ETH:[WETH] USD:[USDT USDC TUSD]]
	b.mapping.Range(func(token string, mappings []string) bool {
		if token == strings.ToUpper(market.Quote()) {
			for _, mappedToken := range mappings {
				err := b.Subscribe(NewMarketWithMainQuote(market.Base(), mappedToken, market.Quote()))
				if err != nil {
					b.logger.Errorf("failed to subscribe to market %s: %s", market, err)
				}
			}
		}
		return true
	})

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	pools, err := b.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %s: %s", market.StringWithoutMain(), err)
	}

	for _, pool := range pools {
		sink := make(chan *Event, 128)
		sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
		if err != nil {
			return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
		}

		go b.watchSwap(market, pool, sink, sub)
		go b.streams.Store(market, sub) // to not block the loop
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
	return nil
}

func (b *baseDEX[Event, Contract]) SetInbox(_ <-chan TradeEvent) {
	// TODO: implement me
}

func (b *baseDEX[Event, Contract]) watchSwap(
	market Market,
	pool *dexPool[Event],
	sink chan *Event,
	sub event.Subscription,
) {
	defer close(sink)
	for {
		select {
		case err := <-sub.Err():
			b.logger.Warnw("connection failed, resubscribing", "market", market.StringWithoutMain(), "err", err)
			if _, ok := b.streams.Load(market); !ok {
				break // market was unsubscribed earlier
			}
			if err := b.Unsubscribe(market); err != nil {
				b.logger.Errorw("failed to resubscribe", "market", market.StringWithoutMain(), "err", err)
			}
			if err := b.Subscribe(market); err != nil {
				b.logger.Errorw("failed to resubscribe", "market", market.StringWithoutMain(), "err", err)
			}
			return
		case swap := <-sink:
			b.logger.Debugw("raw swap", "swap", swap)

			tr, err := b.parse(swap, pool)
			if err != nil {
				b.logger.Errorw("failed to parse swap event", "market", market.StringWithoutMain(), "err", err)
				continue
			}
			tr.Market = market.ApplyMainQuote()

			if !b.filter.Allow(tr) {
				continue
			}

			b.logger.Debugw("parsed trade", "trade", tr)
			b.outbox <- tr
		}
	}
}

type dexPool[Event any] struct {
	contract   dexEventWatcher[Event]
	baseToken  poolToken
	quoteToken poolToken
	reverted   bool
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

func (pool dexPool[Event]) Market() Market {
	return NewMarket(pool.baseToken.Symbol, pool.quoteToken.Symbol)
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
