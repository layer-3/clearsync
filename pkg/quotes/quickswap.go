package quotes

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/layer-3/clearsync/pkg/abi/iquickswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iquickswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

var loggerQuickswap = log.Logger("quickswap")

type quickswap struct {
	once               *once
	url                string
	assetsURL          string
	poolFactoryAddress string
	client             *ethclient.Client
	factory            *iquickswap_v3_factory.IQuickswapV3Factory

	outbox  chan<- TradeEvent
	filter  Filter
	streams safe.Map[Market, event.Subscription]
	assets  safe.Map[string, poolToken]
}

func newQuickswap(config QuickswapConfig, outbox chan<- TradeEvent) Driver {
	return &quickswap{
		once:               newOnce(),
		url:                config.URL,
		assetsURL:          config.AssetsURL,
		poolFactoryAddress: config.PoolFactoryAddress,

		outbox:  outbox,
		filter:  FilterFactory(config.Filter),
		streams: safe.NewMap[Market, event.Subscription](),
		assets:  safe.NewMap[string, poolToken](),
	}
}

func (s *quickswap) Name() DriverType {
	return DriverQuickswap
}

func (s *quickswap) Start() error {
	var startErr error
	started := s.once.Start(func() {
		if !(strings.HasPrefix(s.url, "ws://") || strings.HasPrefix(s.url, "wss://")) {
			startErr = fmt.Errorf("%s (got '%s')", errInvalidWsURL, s.url)
			return
		}

		client, err := ethclient.Dial(s.url)
		if err != nil {
			startErr = fmt.Errorf("failed to connect to the Ethereum client: %w", err)
			return
		}
		s.client = client

		// Check addresses here: https://quickswap.gitbook.io/quickswap/smart-contracts/smart-contracts
		poolFactoryAddress := common.HexToAddress(s.poolFactoryAddress)
		factory, err := iquickswap_v3_factory.NewIQuickswapV3Factory(poolFactoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to instantiate a Quickwap Factory contract: %w", err)
			return
		}
		s.factory = factory

		assets, err := getAssets(s.assetsURL)
		if err != nil {
			startErr = fmt.Errorf("failed to fetch assets: %w", err)
			return
		}
		for _, asset := range assets {
			s.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}
	})

	if !started {
		return errAlreadyStarted
	}
	return startErr
}

func (s *quickswap) Stop() error {
	stopped := s.once.Stop(func() {
		s.streams.Range(func(market Market, _ event.Subscription) bool {
			err := s.Unsubscribe(market)
			return err == nil
		})

		s.streams = safe.NewMap[Market, event.Subscription]() // delete all stopped streams
	})

	if !stopped {
		return errAlreadyStopped
	}
	return nil
}

func (s *quickswap) Subscribe(market Market) error {
	if !s.once.Subscribe() {
		return errNotStarted
	}

	if _, ok := s.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	pool, err := s.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %v: %s", market.String(), err)
	}

	sink := make(chan *iquickswap_v3_pool.IQuickswapV3PoolSwap, 128)
	sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(sink)
		for {
			select {
			case err := <-sub.Err():
				loggerQuickswap.Errorf("market %s: %s", market.String(), err)
				if _, ok := s.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := s.Subscribe(market); err != nil {
					loggerQuickswap.Errorf("market %s: failed to resubscribe: %s", market.String(), err)
				}
				return
			case swap := <-sink:
				tr, err := s.parseSwap(swap, pool)
				if err != nil {
					loggerQuickswap.Errorf("market %s: failed to parse swap: %s", market.String(), err)
					continue
				}

				if !s.filter.Allow(tr) {
					continue
				}
				s.outbox <- tr
			}
		}
	}()

	s.streams.Store(market, sub)
	return nil
}

func (s *quickswap) Unsubscribe(market Market) error {
	if !s.once.Unsubscribe() {
		return errNotStarted
	}

	stream, ok := s.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	stream.Unsubscribe()

	s.streams.Delete(market)
	return nil
}

func (s *quickswap) parseSwap(swap *iquickswap_v3_pool.IQuickswapV3PoolSwap, pool *quickswapPoolWrapper) (TradeEvent, error) {
	if !isValidNonZero(swap.Amount0) || !isValidNonZero(swap.Amount1) {
		return TradeEvent{}, fmt.Errorf("either Amount0 (%s) or Amount1 (%s) is invalid", swap.Amount0, swap.Amount1)
	}

	fmt.Printf("swap = %+v\n", swap)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in from panic during swap parsing in Quickswap (swap = %+v)\n", swap)
		}
	}()

	// Normalize swap amounts
	amount0 := decimal.NewFromBigInt(swap.Amount0, 0).Div(decimal.NewFromInt(10).Pow(pool.baseToken.Decimals))
	amount1 := decimal.NewFromBigInt(swap.Amount1, 0).Div(decimal.NewFromInt(10).Pow(pool.quoteToken.Decimals))

	takerType := TakerTypeBuy
	price := calculatePrice(
		decimal.NewFromBigInt(swap.Price, 0),
		pool.baseToken.Decimals,
		pool.quoteToken.Decimals)
	amount := amount0
	total := amount1

	if price.IsZero() { // then it's a sell trade
		takerType = TakerTypeSell
		price = calculatePrice(
			decimal.NewFromBigInt(swap.Price, 0),
			pool.quoteToken.Decimals,
			pool.baseToken.Decimals)
		amount = amount1
		total = amount0
	}

	tr := TradeEvent{
		Source: DriverQuickswap,
		Market: Market{
			baseUnit:  pool.baseToken.Symbol,
			quoteUnit: pool.quoteToken.Symbol,
		},
		Price:     price,
		Amount:    amount.Abs(),
		Total:     total.Abs(),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}

	fmt.Printf("trade = %+v\n", tr)
	return tr, nil
}

type quickswapPoolWrapper struct {
	contract   *iquickswap_v3_pool.IQuickswapV3Pool
	baseToken  poolToken
	quoteToken poolToken
	reverted   bool
}

func (s *quickswap) getPool(market Market) (*quickswapPoolWrapper, error) {
	baseToken, quoteToken, err := s.getTokens(market)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	poolAddress, err = s.factory.PoolByPair(
		nil,
		common.HexToAddress(baseToken.Address),
		common.HexToAddress(quoteToken.Address),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerQuickswap.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := iquickswap_v3_pool.NewIQuickswapV3Pool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}

	basePoolToken, err := poolContract.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}

	quotePoolToken, err := poolContract.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}

	pool := &quickswapPoolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
		reverted:   false,
	}

	if common.HexToAddress(baseToken.Address) == basePoolToken && common.HexToAddress(quoteToken.Address) == quotePoolToken {
		return pool, nil
	} else if common.HexToAddress(quoteToken.Address) == basePoolToken && common.HexToAddress(baseToken.Address) == quotePoolToken {
		pool.reverted = true
		return pool, nil
	} else {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}
}

func (s *quickswap) getTokens(market Market) (baseToken poolToken, quoteToken poolToken, err error) {
	baseToken, ok := s.assets.Load(strings.ToUpper(market.Base()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Base())
		return
	}
	loggerQuickswap.Infof("market %s: base token address is %s", market, baseToken.Address)

	quoteToken, ok = s.assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Quote())
		return
	}
	loggerQuickswap.Infof("market %s: quote token address is %s", market, quoteToken.Address)

	return baseToken, quoteToken, nil
}
