package base

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/filter"
	"github.com/layer-3/clearsync/pkg/safe"
)

// DexReader defines a set of methods that DEX hooks may use to interact with
// the DEX driver. It's important to limit the methods to read-only operations,
// since by design hooks are not allowed to impede the driver's operations and
// modify its state, thus the absence of methods like `Subscribe`, `Stop`, etc.
type DexReader interface {
	Client() *ethclient.Client
	Assets() *safe.Map[string, DexPoolToken]
}

type DEX[Event any, Iterator dexEventIterator] struct {
	// Params
	driverType quotes_common.DriverType
	rpc        string
	assetsURL  string
	mappingURL string
	marketsURL string
	idlePeriod time.Duration

	// Hooks
	postStart   func(*DEX[Event, Iterator]) error
	getPool     func(context.Context, quotes_common.Market) ([]*DexPool[Event, Iterator], error)
	buildParser func(*Event, *DexPool[Event, Iterator]) SwapParser
	derefIter   func(Iterator) *Event

	// State
	once    *quotes_common.Once
	client  *ethclient.Client
	outbox  chan<- quotes_common.TradeEvent
	logger  *log.ZapEventLogger
	filter  filter.Filter
	history quotes_common.HistoricalDataDriver
	// streams maps market to a map of DEX pools.
	// The value of the map is a pointer to disallow copying of the underlying mutex
	streams safe.Map[quotes_common.Market, *safe.Map[common.Address, dexStream[Event]]]
	assets  safe.Map[string, DexPoolToken]
	// disabledMarkets is a set of markets that are enabled for DEXes.
	// The map is assumed to be read-only,
	// so there's no need for extra thread safety.
	disabledMarkets map[string]struct{}
	mapping         safe.Map[string, []string]
}

type dexStream[Event any] struct {
	sub  event.Subscription
	sink chan *Event
}

type DexConfig[Event any, Iterator dexEventIterator] struct {
	// Params
	DriverType quotes_common.DriverType
	RPC        string
	AssetsURL  string
	MappingURL string
	MarketsURL string
	IdlePeriod time.Duration

	// Hooks
	PostStartHook func(*DEX[Event, Iterator]) error
	PoolGetter    func(context.Context, quotes_common.Market) ([]*DexPool[Event, Iterator], error)
	ParserFactory func(*Event, *DexPool[Event, Iterator]) SwapParser
	IterDeref     func(Iterator) *Event

	// State
	Outbox  chan<- quotes_common.TradeEvent
	Logger  *log.ZapEventLogger
	Filter  filter.Config
	History quotes_common.HistoricalDataDriver
}

func NewDEX[Event any, Iterator dexEventIterator](
	config DexConfig[Event, Iterator],
) (*DEX[Event, Iterator], error) {
	if !(strings.HasPrefix(config.RPC, "ws://") || strings.HasPrefix(config.RPC, "wss://")) {
		return nil, fmt.Errorf("%s (got '%s')", quotes_common.ErrInvalidWsUrl, config.RPC)
	}

	tradesFilter, err := filter.New(config.Filter, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create filter: %w", err)
	}

	return &DEX[Event, Iterator]{
		// Params
		driverType: config.DriverType,
		rpc:        config.RPC,
		assetsURL:  config.AssetsURL,
		mappingURL: config.MappingURL,
		marketsURL: config.MarketsURL,
		idlePeriod: config.IdlePeriod,

		// Hooks
		postStart:   config.PostStartHook,
		getPool:     config.PoolGetter,
		buildParser: config.ParserFactory,
		derefIter:   config.IterDeref,

		// State
		once:            quotes_common.NewOnce(),
		client:          nil,
		outbox:          config.Outbox,
		logger:          config.Logger,
		filter:          tradesFilter,
		history:         config.History,
		streams:         safe.NewMap[quotes_common.Market, *safe.Map[common.Address, dexStream[Event]]](),
		assets:          safe.NewMap[string, DexPoolToken](),
		disabledMarkets: make(map[string]struct{}),
		mapping:         safe.NewMap[string, []string](),
	}, nil
}

func (b *DEX[Event, Iterator]) Client() *ethclient.Client {
	return b.client
}

func (b *DEX[Event, Iterator]) Assets() *safe.Map[string, DexPoolToken] {
	return &b.assets
}

func (b *DEX[Event, Iterator]) Type() (quotes_common.DriverType, quotes_common.ExchangeType) {
	return b.driverType, quotes_common.ExchangeTypeDEX
}

func (b *DEX[Event, Iterator]) Start() error {
	return b.once.Start(func() error {
		// Connect to the RPC provider

		client, err := ethclient.Dial(b.rpc)
		if err != nil {
			return fmt.Errorf("failed to connect to the Ethereum client: %w", err)
		}
		b.client = client

		// Fetch assets

		allAssets, err := fetch[map[string][]DexPoolToken](b.assetsURL)
		if err != nil {
			return fmt.Errorf("failed to fetch assets: %w", err)
		}
		tokens, ok := allAssets["tokens"]
		if !ok {
			return fmt.Errorf("failed to fetch assets: `tokens` key not found")
		}
		for _, asset := range tokens {
			b.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}

		// Fetch mappings

		mappings, err := fetch[map[string]map[string][]string](b.mappingURL)
		if err != nil {
			return fmt.Errorf("failed to fetch mapping: %w", err)
		}
		mapping, ok := mappings["tokens"]
		if !ok {
			return fmt.Errorf("failed to fetch mappings: `tokens` key not found")
		}
		for key, mapItem := range mapping {
			b.mapping.Store(key, mapItem)
		}

		// Fetch markets

		markets, err := fetch[[]marketSymbol](b.marketsURL)
		if err != nil {
			return fmt.Errorf("failed to fetch markets: %w", err)
		}
		for _, market := range markets {
			if market.Quotes.Dexs {
				continue
			}

			// Strip prefix to get base and quote tokens.
			// Market is assumed to be in the following format:
			// `<type>://<base>/<quote>`, like `spot://btc/usd`.
			if !strings.HasPrefix(market.Symbol, "spot://") {
				return fmt.Errorf("invalid market symbol in markets config: %s", market.Symbol)
			}
			baseQuote := strings.Split(market.Symbol, "spot://")[1]
			tokens := strings.Split(baseQuote, "/")
			if len(tokens) != 2 {
				return fmt.Errorf("invalid market symbol in markets config: %s", market.Symbol)
			}
			market := quotes_common.NewMarket(tokens[0], tokens[1])
			b.disabledMarkets[market.String()] = struct{}{}
		}

		// Run post-start hook
		return b.postStart(b)
	})
}

func (b *DEX[Event, Iterator]) Stop() error {
	return b.once.Stop(func() (stopErr error) {
		b.streams.Range(func(market quotes_common.Market, _ *safe.Map[common.Address, dexStream[Event]]) bool {
			if err := b.Unsubscribe(market); err != nil {
				stopErr = err
			}
			return true
		})
		return stopErr
	})
}

func (b *DEX[Event, Iterator]) Subscribe(market quotes_common.Market) error {
	ctx := context.TODO()

	if !b.once.IsStarted() {
		return quotes_common.ErrNotStarted
	}

	// Subscribe to associated markets
	// mapping map[BTC:[WBTC] ETH:[WETH] USD:[USDT USDC TUSD]]
	var mappingErr error
	b.mapping.Range(func(token string, mappings []string) bool {
		if token != strings.ToUpper(market.Quote()) {
			return true
		}

		for _, mappedToken := range mappings {
			market := quotes_common.NewMarketWithMainQuote(market.Base(), mappedToken, market.Quote())
			if err := debounce.Debounce(ctx, b.logger, func(_ context.Context) error { return b.Subscribe(market) }); err != nil {
				b.logger.Errorf("failed to subscribe to market %s: %s", market, err)
				mappingErr = err
			}
		}

		return true
	})
	if mappingErr != nil {
		return fmt.Errorf("failed to subscribe to helper markets: %w", mappingErr)
	}

	// Verify the market is available

	if _, ok := b.streams.Load(market); ok {
		fmt.Println("Market already subscribed", market)
		return fmt.Errorf("%s: %w", market, quotes_common.ErrAlreadySubbed)
	}

	// Check if market is enabled for DEXes
	if _, ok := b.disabledMarkets[market.String()]; ok && CexConfigured.Load() {
		return fmt.Errorf("%w: %s", quotes_common.ErrMarketDisabled, market)
	}

	pools, err := b.getPool(ctx, market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %s: %s", market.StringWithoutMain(), err)
	}

	// Publish the last trade for a given pool using historical data
	// since it may take too long to receive the first swap with DEXes.
	if trades, err := b.HistoricalData(context.TODO(), market, 12*time.Hour, 1); err == nil && len(trades) > 0 {
		b.outbox <- trades[0]
	}

	// Subscribe to all pools
	for _, pool := range pools {
		if err := b.subscribePool(pool); err != nil {
			return err
		}
	}

	return nil
}

func (b *DEX[Event, Iterator]) subscribePool(pool *DexPool[Event, Iterator]) error {
	watchCtx, cancel := context.WithCancel(context.TODO())
	sink := make(chan *Event, 128)

	var sub event.Subscription
	var err error
	err = debounce.Debounce(watchCtx, b.logger, func(ctx context.Context) error {
		opts := &bind.WatchOpts{Context: ctx}
		sub, err = pool.Contract.WatchSwap(opts, sink, []common.Address{}, []common.Address{})
		return err
	})
	if err != nil {
		cancel()
		close(sink)
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", pool.Market, err)
	}

	pools := safe.NewMap[common.Address, dexStream[Event]]()
	stream, _ := b.streams.LoadOrStore(pool.Market, &pools)
	stream.Store(pool.Address, dexStream[Event]{sub: sub, sink: sink})

	RecordSubscribed(b.driverType, pool.Market)
	go b.watchSwap(cancel, pool, sink, sub)
	return nil
}

func (b *DEX[Event, Iterator]) Unsubscribe(market quotes_common.Market) error {
	if !b.once.IsStarted() {
		return quotes_common.ErrNotStarted
	}

	stream, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, quotes_common.ErrNotSubbed)
	}
	stream.UpdateInTx(func(stream map[common.Address]dexStream[Event]) {
		for _, s := range stream {
			s.sub.Unsubscribe()
			s.sub = nil
			// do not delete the sink channel
		}
	})

	RecordUnsubscribed(b.driverType, market)
	return nil
}

func (b *DEX[Event, Iterator]) HistoricalData(
	ctx context.Context,
	market quotes_common.Market,
	window time.Duration,
	limit uint64,
) ([]quotes_common.TradeEvent, error) {
	trades, err := FetchHistoryDataFromExternalSource(ctx, b.history, market, window, limit, b.logger)
	if err == nil && len(trades) > 0 {
		return trades, nil
	}

	// Calculate the block range

	m := quotes_common.NewMarket(market.Base(), market.Quote())
	if strings.ToLower(market.Quote()) == "usd" {
		m = quotes_common.NewMarket(market.Base(), market.Quote()+"t") // convert USD quote to USDT
	}
	pools, err := b.getPool(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool for market %s: %w", m, err)
	}

	now := time.Now()
	from := now.Add(-window)
	block, err := b.findBlockByTimestamp(ctx, b.client, from)
	if err != nil {
		return nil, fmt.Errorf("failed to find block by timestamp %d: %w", from.Unix(), err)
	}

	// Fetch all swaps in the block range
	for i, pool := range pools {
		var iter Iterator

		err = debounce.Debounce(ctx, b.logger, func(ctx context.Context) error {
			opts := &bind.FilterOpts{Start: block.Uint64(), Context: ctx}
			iter, err = pool.Contract.FilterSwap(opts, nil, nil)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to fetch historical swaps: %w", err)
		}
		defer iter.Close() // to avoid memory leak

		// Convert swaps into trade events
		for iter.Next() {
			swap := b.derefIter(iter)
			if swap == nil {
				b.logger.Debugw("failed to deref iter", "iter", iter, "market", m)
				continue
			}

			parser := b.buildParser(swap, pools[i])
			logger := b.logger.With("swap", swap)
			trade, err := parser.ParseSwap(logger)
			if err != nil {
				return nil, fmt.Errorf("failed to parse historical swap: %s (`%+v`)", err, swap)
			}

			trades = append(trades, trade)
			if uint64(len(trades)) >= limit {
				break
			}
		}
		if iter.Error() != nil {
			return nil, fmt.Errorf("failed to fetch historical swaps: %w", iter.Error())
		}
	}

	quotes_common.SortTradeEventsInPlace(trades)

	return trades, nil
}

// findBlockByTimestamp performs a binary search over the range of block numbers
// to find the block whose timestamp is closest to but not greater than the given timestamp.
// It returns a block number at or immediately before the given timestamp.
func (b *DEX[Event, Iterator]) findBlockByTimestamp(
	ctx context.Context,
	client *ethclient.Client,
	target time.Time,
) (*big.Int, error) {
	currentTime := time.Now()
	if target.After(currentTime) {
		return nil, fmt.Errorf("provided time %v is in the future", target)
	}

	var header *types.Header
	var err error
	err = debounce.Debounce(ctx, b.logger, func(ctx context.Context) error {
		header, err = client.HeaderByNumber(ctx, nil)
		return err
	})
	if err != nil {
		return nil, err
	}

	high := header.Number
	low := big.NewInt(0)
	targetTimestamp := target.Unix()

	for low.Cmp(high) < 0 {
		mid := new(big.Int).Add(low, high)
		mid.Div(mid, big.NewInt(2))

		err = debounce.Debounce(ctx, b.logger, func(ctx context.Context) error {
			header, err = client.HeaderByNumber(ctx, mid)
			return err
		})
		if err != nil {
			return nil, err
		}

		blockTime := header.Time

		if blockTime < uint64(targetTimestamp) {
			low = mid.Add(mid, big.NewInt(1))
		} else if blockTime > uint64(targetTimestamp) {
			high = mid
		} else {
			return mid, nil // The exact block number with the matching timestamp
		}
	}

	return high, nil // The closest block number to the desired timestamp
}

func (b *DEX[Event, Iterator]) watchSwap(
	cancel context.CancelFunc,
	pool *DexPool[Event, Iterator],
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

			parser := b.buildParser(swap, pool)
			logger := b.logger.With("swap", swap)
			trade, err := parser.ParseSwap(logger)
			if err != nil {
				b.logger.Errorw("failed to parse swap event",
					"error", err,
					"market", pool.Market,
					"swap", swap,
					"pool", pool,
					"error", err,
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
			b.logger.Warnw("market inactivity detected",
				"market", pool.Market,
				"pool_address", pool.Address,
				"base_token", pool.BaseToken.Symbol,
				"quote_token", pool.QuoteToken.Symbol)
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

func (b *DEX[Event, Iterator]) resubscribe(pool *DexPool[Event, Iterator]) error {
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

func fetch[T any](url string) (data T, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

type marketConfig struct {
	Dexs bool `json:"dexs"`
}

type marketSymbol struct {
	Symbol string        `json:"symbol"`
	Quotes *marketConfig `json:"quotes,omitempty"`
}

func (s *marketSymbol) UnmarshalJSON(data []byte) error {
	type Alias marketSymbol
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if s.Quotes == nil {
		// Default value for `Dexs` is `true`.
		// If the field is not present in the JSON,
		// then it is overridden.
		s.Quotes = &marketConfig{Dexs: true}
	}
	return nil
}

type DexPool[Event any, Iterator dexEventIterator] struct {
	Contract   dexEvent[Event, Iterator]
	Address    common.Address // not used in code but is useful for logging
	BaseToken  DexPoolToken
	QuoteToken DexPoolToken
	Market     quotes_common.Market
	Reversed   bool
}

// dexEvent represents an interface for interacting with DEX contract events.
// When a smart contract with events is processed using `abigen`, the generated
// event bindings conform to this interface. It defines a subset of methods
// available in the event binding.
//
// Event: A generic type representing the specific event structure.
// EventIterator: A generic type representing the iterator for filtering events.
type dexEvent[Event any, Iterator dexEventIterator] interface {
	// WatchSwap subscribes to the "Swap" event, streaming events to the provided sink channel.
	// Parameters:
	// - opts: Options for configuring the subscription, such as context and block start/end.
	// - sink: Channel to receive the streamed event data.
	// - from, to: Filter the event by sender and recipient addresses.
	// Returns:
	// - A subscription object for managing the event stream.
	// - An error, if the subscription fails.
	WatchSwap(opts *bind.WatchOpts, sink chan<- *Event, from, to []common.Address) (event.Subscription, error)

	// FilterSwap retrieves past "Swap" events matching the provided filter criteria.
	// Parameters:
	// - opts: Options for configuring the filter, such as block range.
	// - sender, to: Filter the events by sender and recipient addresses.
	// Returns:
	// - An iterator for accessing the matching events.
	// - An error, if the filtering fails.
	FilterSwap(opts *bind.FilterOpts, sender, to []common.Address) (Iterator, error)
}

type dexEventIterator interface {
	Next() bool
	Error() error
	io.Closer
}

type DexPoolToken struct {
	Name     string
	Address  common.Address
	Symbol   string
	Decimals decimal.Decimal
	ChainId  uint
	LogoURI  string
}

func GetTokens(
	assets *safe.Map[string, DexPoolToken],
	market quotes_common.Market,
	logger *log.ZapEventLogger,
) (baseToken DexPoolToken, quoteToken DexPoolToken, err error) {
	baseToken, ok := assets.Load(strings.ToUpper(market.Base()))
	if !ok {
		return baseToken, quoteToken, fmt.Errorf("base token does not exist for market %s", market.StringWithoutMain())
	}
	logger.Infow("found base token", "address", baseToken.Address, "market", market.StringWithoutMain())

	quoteToken, ok = assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		return baseToken, quoteToken, fmt.Errorf("quote tokens does not exist for market %s", market.StringWithoutMain())
	}
	logger.Infow("found quote token", "address", quoteToken.Address, "market", market.StringWithoutMain())

	return baseToken, quoteToken, nil
}

func FetchHistoryDataFromExternalSource(
	ctx context.Context,
	source quotes_common.HistoricalDataDriver,
	market quotes_common.Market,
	window time.Duration,
	limit uint64,
	logger *log.ZapEventLogger,
) ([]quotes_common.TradeEvent, error) {
	if source == nil {
		return nil, nil // no data source is not an error
	}
	logger.Infow("fetching historical data from external source", "market", market, "window", window.String())

	trades, err := source.HistoricalData(ctx, market, window, limit)
	if err != nil {
		logger.Warnw("failed to fetch historical data from external source",
			"market", market,
			"window", window.String(),
			"error", err)
		return nil, fmt.Errorf("failed to fetch historical data from external source: %w", err)
	}

	if len(trades) == 0 {
		logger.Infow("the external source returned no trades",
			"market", market,
			"window", window.String())
		return nil, nil
	}

	quotes_common.SortTradeEventsInPlace(trades)

	stale := time.Now().Add(-10 * time.Minute)
	lastTrade := trades[len(trades)-1]
	if stale.Before(lastTrade.CreatedAt) {
		logger.Infow("successfully fetched historical data from external source",
			"market", market,
			"window", window.String(),
			"trades_num", len(trades))
		return trades, nil
	}

	logger.Infow("the external source returned stale trades",
		"market", market,
		"window", window.String(),
		"stale_before", stale.String(),
		"trades_num", len(trades),
		"last_created_at", lastTrade.CreatedAt)
	return nil, nil
}
