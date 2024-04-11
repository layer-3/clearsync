package quotes

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/isyncswap_factory"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerSyncswap = log.Logger("syncswap")

type syncswap struct {
	once       *once
	url        string
	assetsURL  string
	mappingURL string
	client     *ethclient.Client

	classicPoolFactoryAddress string
	classicFactory            *isyncswap_factory.ISyncSwapFactory
	stablePoolMarkets         map[Market]struct{}
	stablePoolFactoryAddress  string
	stableFactory             *isyncswap_factory.ISyncSwapFactory

	outbox  chan<- TradeEvent
	filter  Filter
	streams safe.Map[Market, event.Subscription]
	assets  safe.Map[string, poolToken]
	mapping safe.Map[string, []string]
}

func newSyncswap(config SyncswapConfig, outbox chan<- TradeEvent) Driver {
	stablePoolMarkets := make(map[Market]struct{})
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := NewMarketFromString(rawMarket)
		if !ok {
			loggerSyncswap.Errorf("failed to parse stable pool market `%s`", rawMarket)
			continue
		}
		stablePoolMarkets[market] = struct{}{}
	}

	return &syncswap{
		once:       newOnce(),
		url:        config.URL,
		assetsURL:  config.AssetsURL,
		mappingURL: config.MappingURL,

		classicPoolFactoryAddress: config.ClassicPoolFactoryAddress,
		classicFactory:            nil,
		stablePoolMarkets:         stablePoolMarkets,
		stablePoolFactoryAddress:  config.StablePoolFactoryAddress,
		stableFactory:             nil,

		outbox:  outbox,
		filter:  NewFilter(config.Filter),
		streams: safe.NewMap[Market, event.Subscription](),
		assets:  safe.NewMap[string, poolToken](),
		mapping: safe.NewMap[string, []string](),
	}
}

func (s *syncswap) Name() DriverType {
	return DriverSyncswap
}

func (b *syncswap) Type() Type {
	return TypeDEX
}

func (s *syncswap) Start() error {
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

		// Check addresses here: https://syncswap.gitbook.io/syncswap/smart-contracts/smart-contracts
		classicPoolFactoryAddress := common.HexToAddress(s.classicPoolFactoryAddress)
		classicFactory, err := isyncswap_factory.NewISyncSwapFactory(classicPoolFactoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to instantiate a Quickswap Factory contract: %w", err)
			return
		}
		s.classicFactory = classicFactory

		stablePoolFactoryAddress := common.HexToAddress(s.stablePoolFactoryAddress)
		stableFactory, err := isyncswap_factory.NewISyncSwapFactory(stablePoolFactoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to instantiate a Quickswap Factory contract: %w", err)
			return
		}
		s.stableFactory = stableFactory

		assets, err := getAssets(s.assetsURL)
		if err != nil {
			startErr = fmt.Errorf("failed to fetch assets: %w", err)
			return
		}

		for _, asset := range assets {
			s.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}

		mapping, err := getMapping(s.mappingURL)

		if err != nil {
			startErr = fmt.Errorf("failed to fetch mapping: %w", err)
			return
		}

		for key, mapItem := range mapping {
			s.mapping.Store(key, mapItem)
		}
	})

	if !started {
		return errAlreadyStarted
	}
	return startErr
}
func (s *syncswap) Stop() error {
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

func (s *syncswap) Subscribe(market Market) error {
	if !s.once.Subscribe() {
		return errNotStarted
	}

	// mapping map[BTC:[WBTC] ETH:[WETH] USD:[USDT USDC TUSD]]
	s.mapping.Range(func(token string, mappings []string) bool {
		if token == strings.ToUpper(market.Quote()) {
			for _, mappedToken := range mappings {
				err := s.Subscribe(NewMarketWithLegacyQuote(market.Base(), mappedToken, market.Quote()))
				if err != nil {
					loggerSyncswap.Errorf("failed to subscribe to market %s: %s", market, err)
				}
			}
		}
		return true
	})

	if _, ok := s.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	pool, err := s.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %v: %s", market.String(), err)
	}

	sink := make(chan *isyncswap_pool.ISyncSwapPoolSwap, 128)
	sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(sink)
		for {
			select {
			case err := <-sub.Err():
				loggerSyncswap.Warnf("connection failed for market %s, resubscribing: %s", market, err)
				if _, ok := s.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := s.Unsubscribe(market); err != nil {
					loggerSyncswap.Errorf("market %s: failed to resubscribe: %s", market, err)
				}
				if err := s.Subscribe(market); err != nil {
					loggerSyncswap.Errorf("market %s: failed to resubscribe: %s", market, err)
				}
				return
			case swap := <-sink:
				if pool.reverted {
					s.flipSwap(swap)
				}

				tr, err := s.parseSwap(swap, market, pool)
				if err != nil {
					loggerSyncswap.Errorf("market %s: failed to parse swap: %s", market.String(), err)
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

func (*syncswap) flipSwap(swap *isyncswap_pool.ISyncSwapPoolSwap) {
	amount0In, amount0Out := swap.Amount0In, swap.Amount0Out
	swap.Amount0In, swap.Amount0Out = swap.Amount1In, swap.Amount1Out
	swap.Amount1In, swap.Amount1Out = amount0In, amount0Out
}

func (*syncswap) parseSwap(swap *isyncswap_pool.ISyncSwapPoolSwap, market Market, pool *syncswapPoolWrapper) (TradeEvent, error) {
	var takerType TakerType
	var price decimal.Decimal
	var amount decimal.Decimal
	var total decimal.Decimal

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered after parse swap panic. Swap = %+v\n", swap)
		}
	}()

	switch {
	case isValidNonZero(swap.Amount0In) && isValidNonZero(swap.Amount1Out):
		amount1Out := decimal.NewFromBigInt(swap.Amount1Out, 0).Div(decimal.NewFromInt(10).Pow(pool.quoteToken.Decimals))
		amount0In := decimal.NewFromBigInt(swap.Amount0In, 0).Div(decimal.NewFromInt(10).Pow(pool.baseToken.Decimals))

		takerType = TakerTypeSell
		price = amount1Out.Div(amount0In)
		total = amount1Out
		amount = amount0In

	case isValidNonZero(swap.Amount0Out) && isValidNonZero(swap.Amount1In):
		amount0Out := decimal.NewFromBigInt(swap.Amount0Out, 0).Div(decimal.NewFromInt(10).Pow(pool.baseToken.Decimals))
		amount1In := decimal.NewFromBigInt(swap.Amount1In, 0).Div(decimal.NewFromInt(10).Pow(pool.quoteToken.Decimals))

		takerType = TakerTypeBuy
		price = amount1In.Div(amount0Out)
		total = amount1In
		amount = amount0Out
	default:
		loggerSyncswap.Errorf("market %s: unknown swap type", market.String())
		return TradeEvent{}, fmt.Errorf("market %s: unknown swap type", market.String())
	}

	amount = amount.Abs()
	tr := TradeEvent{
		Source:    DriverSyncswap,
		Market:    market.ApplyLegacyQuote(),
		Price:     price,
		Amount:    amount,
		Total:     total,
		TakerType: takerType,
		CreatedAt: time.Now(),
	}

	return tr, nil
}

func (s *syncswap) Unsubscribe(market Market) error {
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

type syncswapPoolWrapper struct {
	contract   *isyncswap_pool.ISyncSwapPool
	baseToken  poolToken
	quoteToken poolToken
	reverted   bool
}

func (s *syncswap) getPool(market Market) (*syncswapPoolWrapper, error) {
	baseToken, quoteToken, err := s.getTokens(market)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	if _, ok := s.stablePoolMarkets[market]; ok {
		loggerSyncswap.Infof("market %s is a stable pool", market)
		poolAddress, err = s.stableFactory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
		)
	} else {
		loggerSyncswap.Infof("market %s is a classic pool", market)
		poolAddress, err = s.classicFactory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get classic pool address: %w", err)
	}
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("classic pool for market %s does not exist", market)
	}
	loggerSyncswap.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := isyncswap_pool.NewISyncSwapPool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}

	basePoolToken, err := poolContract.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}

	quotePoolToken, err := poolContract.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}

	pool := &syncswapPoolWrapper{
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
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}
}

func (s *syncswap) getTokens(market Market) (baseToken poolToken, quoteToken poolToken, err error) {
	baseToken, ok := s.assets.Load(strings.ToUpper(market.Base()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Base())
		return baseToken, quoteToken, err
	}
	loggerSyncswap.Infof("market %s: base token address is %s", market, baseToken.Address)

	quoteToken, ok = s.assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Quote())
		return baseToken, quoteToken, err
	}
	loggerSyncswap.Infof("market %s: quote token address is %s", market, quoteToken.Address)

	return baseToken, quoteToken, nil
}

func isValidNonZero(x *big.Int) bool {
	return x != nil && x.Sign() != 0
}

// Not implemented
func (s *syncswap) SetInbox(inbox <-chan TradeEvent) {
}
