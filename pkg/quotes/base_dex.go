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
	logger     *log.ZapEventLogger

	// Hooks
	start   func() error
	getPool func(Market) (*dexPool[Event], error)
	parse   func(*Event, *dexPool[Event]) (TradeEvent, error)

	// State
	client  *ethclient.Client
	outbox  chan<- TradeEvent
	filter  Filter
	streams safe.Map[Market, event.Subscription]
	assets  safe.Map[string, poolToken]
}

func newBaseDEX[Event any, Contract any](
	driverType DriverType,
	url string,
	assetsURL string,
	outbox chan<- TradeEvent,
	config FilterConfig,
	logger *log.ZapEventLogger,

	startHook func() error,
	poolGetter func(Market) (*dexPool[Event], error),
	eventParser func(*Event, *dexPool[Event]) (TradeEvent, error),
) *baseDEX[Event, Contract] {
	return &baseDEX[Event, Contract]{
		// Params
		once:       newOnce(),
		driverType: driverType,
		url:        url,
		assetsURL:  assetsURL,

		// Hooks
		start:   startHook,
		getPool: poolGetter,
		parse:   eventParser,

		// State
		client:  nil,
		outbox:  outbox,
		filter:  NewFilter(config),
		streams: safe.NewMap[Market, event.Subscription](),
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
		if !(strings.HasPrefix(b.url, "ws://") || strings.HasPrefix(b.url, "wss://")) {
			startErr = fmt.Errorf("%s (got '%s')", errInvalidWsURL, b.url)
			return
		}

		client, err := ethclient.Dial(b.url)
		if err != nil {
			startErr = fmt.Errorf("failed to connect to the Ethereum client: %w", err)
			return
		}
		b.client = client

		assets, err := getAssets(b.assetsURL)
		if err != nil {
			startErr = fmt.Errorf("failed to fetch assets: %w", err)
			return
		}
		for _, asset := range assets {
			b.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}

		if err := b.start(); err != nil {
			startErr = err
			return
		}
	})

	if !started {
		return errAlreadyStarted
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
		return errAlreadyStopped
	}
	return nil
}

func (b *baseDEX[Event, Contract]) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return errNotStarted
	}

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	pool, err := b.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %v: %s", market.String(), err)
	}

	sink := make(chan *Event, 128)
	sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(sink)
		for {
			select {
			case err := <-sub.Err():
				b.logger.Warnw("connection failed, resubscribing", "market", market, "err", err)
				if _, ok := b.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := b.Unsubscribe(market); err != nil {
					b.logger.Errorw("failed to resubscribe", "market", market, "err", err)
				}
				if err := b.Subscribe(market); err != nil {
					b.logger.Errorw("failed to resubscribe", "market", market, "err", err)
				}
				return
			case swap := <-sink:
				tr, err := b.parse(swap, pool)
				if err != nil {
					b.logger.Errorw("failed to parse swap event", "market", market, "err", err)
					continue
				}

				if !b.filter.Allow(tr) {
					continue
				}
				b.outbox <- tr
			}
		}
	}()

	b.streams.Store(market, sub)
	return nil
}

func (b *baseDEX[Event, Contract]) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return errNotStarted
	}

	stream, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	stream.Unsubscribe()

	b.streams.Delete(market)
	return nil
}

func (b *baseDEX[Event, Contract]) SetInbox(inbox <-chan TradeEvent) {
	// TODO: implement me
	panic("not implemented")
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
	Address  string
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
		err = fmt.Errorf("tokens '%s' does not exist", market.Base())
		return
	}
	logger.Infow("found base token", "address", baseToken.Address, "market", market)

	quoteToken, ok = assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Quote())
		return
	}
	logger.Infow("found quote token", "address", quoteToken.Address, "market", market)

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
