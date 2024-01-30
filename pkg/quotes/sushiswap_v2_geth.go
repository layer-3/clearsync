package quotes

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/isushiswap_v2_factory"
	"github.com/layer-3/clearsync/pkg/abi/isushiswap_v2_pair"
)

var loggerSushiswapV2Geth = log.Logger("sushiswap_v2_geth")

type sushiswapV2Geth struct {
	once           *once
	url            string
	assetsURL      string
	factoryAddress string
	client         *ethclient.Client
	factory        *isushiswap_v2_factory.ISushiswapV2Factory

	outbox  chan<- TradeEvent
	streams sync.Map
	assets  sync.Map
}

func newSushiswapV2Geth(config SushiswapV2GethConfig, outbox chan<- TradeEvent) Driver {
	return &sushiswapV2Geth{
		once:           newOnce(),
		url:            config.URL,
		assetsURL:      config.AssetsURL,
		factoryAddress: config.FactoryAddress,

		outbox: outbox,
	}
}

func (s *sushiswapV2Geth) Start() error {
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
		factory, err := isushiswap_v2_factory.NewISushiswapV2Factory(factoryAddress, client)
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

func (s *sushiswapV2Geth) Stop() error {
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

func (s *sushiswapV2Geth) Subscribe(market Market) error {
	if !s.once.Subscribe() {
		return errNotStarted
	}
	symbol := market.BaseUnit + market.QuoteUnit

	if _, ok := s.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	pool, err := s.getPool(market)
	if err != nil {
		return fmt.Errorf("failed to get pool for market %v: %s", symbol, err)
	}

	sink := make(chan *isushiswap_v2_pair.ISushiswapV2PairSwap, 128)
	sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(sink)
		for {
			select {
			case err := <-sub.Err():
				loggerSushiswapV2Geth.Errorf("market %s: %s", symbol, err)
				if _, ok := s.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := s.Subscribe(market); err != nil {
					loggerSushiswapV2Geth.Errorf("market %s: failed to resubscribe: %s", symbol, err)
				}
				return
			case swap := <-sink:
				var takerType TakerType
				var price decimal.Decimal
				var amount decimal.Decimal

				if swap.Amount0In != nil && swap.Amount1Out != nil {
					amount1Out := decimal.NewFromBigInt(swap.Amount1Out, 0)
					amount0In := decimal.NewFromBigInt(swap.Amount0In, 0)

					takerType = TakerTypeSell
					price = amount1Out.Div(amount0In)
					amount = amount0In
				} else if swap.Amount0Out != nil && swap.Amount1In != nil {
					amount0Out := decimal.NewFromBigInt(swap.Amount0Out, 0)
					amount1In := decimal.NewFromBigInt(swap.Amount1In, 0)

					takerType = TakerTypeBuy
					price = amount0Out.Div(amount1In)
					amount = amount0Out
				} else {
					loggerSushiswapV2Geth.Errorf("market %s: unknown swap type", symbol)
					continue
				}

				amount = amount.Abs()
				s.outbox <- TradeEvent{
					Source:    DriverSushiswapV2Geth,
					Market:    symbol,
					Price:     price,
					Amount:    amount,
					Total:     price.Mul(amount),
					TakerType: takerType,
					CreatedAt: time.Now(),
				}
			}
		}
	}()

	s.streams.Store(market, sub)
	return nil
}

func (s *sushiswapV2Geth) Unsubscribe(market Market) error {
	panic("implement me")
}

type sushiswapV2GethPoolWrapper struct {
	contract   *isushiswap_v2_pair.ISushiswapV2Pair
	baseToken  poolToken
	quoteToken poolToken
}

func (s *sushiswapV2Geth) getPool(market Market) (*sushiswapV2GethPoolWrapper, error) {
	baseToken, quoteToken, err := getAssetsFromCache(market, &s.assets)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	poolAddress, err = s.factory.GetPair(
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
	loggerSushiswapV2Geth.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := isushiswap_v2_pair.NewISushiswapV2Pair(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Sushiswap v2 pool: %w", err)
	}
	return &sushiswapV2GethPoolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}
