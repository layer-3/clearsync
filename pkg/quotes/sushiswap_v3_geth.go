package quotes

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/isushiswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/isushiswap_v3_pool"
)

var (
	loggerSushiswapV3Geth = log.Logger("sushiswap_v3_geth")
	// Sushiwap v3 protocol uses the same fee tiers as Uniswap v3,
	// that is 0.01%, 0.05%, 0.3%, and 1%.
	sushiswapV3FeeTiers = []uint{100, 500, 3000, 10000}
)

type sushiswapV3Geth struct {
	once           *once
	url            string
	assetsURL      string
	factoryAddress string
	client         *ethclient.Client
	factory        *isushiswap_v3_factory.ISushiswapV3Factory

	outbox       chan<- TradeEvent
	streams      sync.Map
	assets       sync.Map
	tradeSampler tradeSampler
}

func newSushiswapV3Geth(config SushiswapV3GethConfig, outbox chan<- TradeEvent) Driver {
	return &sushiswapV3Geth{
		once:           newOnce(),
		url:            config.URL,
		assetsURL:      config.AssetsURL,
		factoryAddress: config.FactoryAddress,

		outbox:       outbox,
		tradeSampler: *newTradeSampler(config.TradeSampler),
	}
}

func (s *sushiswapV3Geth) Start() error {
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

		factoryAddress := common.HexToAddress(s.factoryAddress)
		factory, err := isushiswap_v3_factory.NewISushiswapV3Factory(factoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to instantiate a factory contract: %w", err)
			return
		}
		s.factory = factory

		assets, err := fetchAssets(s.assetsURL)
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

func (s *sushiswapV3Geth) Stop() error {
	stopped := s.once.Stop(func() {
		s.streams.Range(func(market, stream any) bool {
			err := s.Unsubscribe(market.(Market))
			return err == nil
		})

		s.streams = sync.Map{} // delete all stopped streams
	})

	if !stopped {
		return errAlreadyStopped
	}
	return nil
}

func (s *sushiswapV3Geth) Subscribe(market Market) error {
	if !s.once.Subscribe() {
		return errNotStarted
	}
	symbol := market.BaseUnit + market.QuoteUnit

	if _, ok := s.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	pool, err := s.getPool(market)
	if err != nil {
		return fmt.Errorf("failed get pool for market %v: %s", symbol, err)
	}

	sink := make(chan *isushiswap_v3_pool.ISushiswapV3PoolSwap, 128)
	sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(sink)
		for {
			select {
			case err := <-sub.Err():
				loggerSushiswapV3Geth.Errorf("market %s: %s", symbol, err)
				if _, ok := s.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := s.Subscribe(market); err != nil {
					loggerSushiswapV3Geth.Errorf("market %s: failed to resubscribe: %s", symbol, err)
				}
				return
			case swap := <-sink:
				amount := decimal.NewFromBigInt(swap.Amount0, 0)
				price := calculatePrice(
					decimal.NewFromBigInt(swap.SqrtPriceX96, 0),
					pool.baseToken.Decimals,
					pool.quoteToken.Decimals)
				takerType := TakerTypeBuy
				if amount.Sign() < 0 {
					// When amount0 is negative (and amount1 is positive),
					// it means token0 is leaving the pool in exchange for token1.
					// This is equivalent to a "sell" of token0 (or a "buy" of token1).
					takerType = TakerTypeSell
				}

				amount = amount.Abs()
				tr := TradeEvent{
					Source:    DriverSushiswapV3Geth,
					Market:    symbol,
					Price:     price,
					Amount:    amount,
					Total:     price.Mul(amount),
					TakerType: takerType,
					CreatedAt: time.Now(),
				}

				if s.tradeSampler.allow(tr) {
					s.outbox <- tr
				}
				s.outbox <- tr
			}
		}
	}()

	s.streams.Store(market, sub)
	return nil
}

func (s *sushiswapV3Geth) Unsubscribe(market Market) error {
	panic("implement me")
}

type sushiswapV3GethPoolWrapper struct {
	contract   *isushiswap_v3_pool.ISushiswapV3Pool
	baseToken  poolToken
	quoteToken poolToken
}

func (s *sushiswapV3Geth) getPool(market Market) (*sushiswapV3GethPoolWrapper, error) {
	baseToken, quoteToken, err := getAssetsFromCache(market, &s.assets)
	if err != nil {
		return nil, err
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range sushiswapV3FeeTiers {
		poolAddress, err = s.factory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
			big.NewInt(int64(feeTier)),
		)
		if err != nil {
			return nil, err
		}
		if poolAddress != zeroAddress {
			loggerSushiswapV3Geth.Infof("market %s: selected fee tier: %.2f%%", market, float64(feeTier)/10000)
			break
		}
	}
	loggerSushiswapV3Geth.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := isushiswap_v3_pool.NewISushiswapV3Pool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Sushiswap v3 pool: %w", err)
	}
	return &sushiswapV3GethPoolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}
