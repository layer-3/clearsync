package quotes

import (
	"fmt"
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
	once                      *once
	url                       string
	assetsURL                 string
	classicPoolFactoryAddress string
	client                    *ethclient.Client
	factory                   *isyncswap_factory.ISyncSwapFactory

	outbox  chan<- TradeEvent
	streams safe.Map[Market, event.Subscription]
	assets  safe.Map[string, poolToken]
}

func newSyncswap(config SyncswapConfig, outbox chan<- TradeEvent) Driver {
	return &syncswap{
		once:                      newOnce(),
		url:                       config.URL,
		assetsURL:                 config.AssetsURL,
		classicPoolFactoryAddress: config.ClassicPoolFactoryAddress,

		outbox:  outbox,
		streams: safe.NewMap[Market, event.Subscription](),
		assets:  safe.NewMap[string, poolToken](),
	}
}

func (s *syncswap) Name() DriverType {
	return DriverSyncswap
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
		factory, err := isyncswap_factory.NewISyncSwapFactory(classicPoolFactoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to instantiate a TokenFactory contract: %w", err)
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
				loggerSyncswap.Errorf("market %s: %s", market.String(), err)
				if _, ok := s.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := s.Subscribe(market); err != nil {
					loggerSyncswap.Errorf("market %s: failed to resubscribe: %s", market.String(), err)
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
					loggerSyncswap.Errorf("market %s: unknown swap type", market.String())
					continue
				}

				amount = amount.Abs()
				s.outbox <- TradeEvent{
					Source:    DriverSyncswap,
					Market:    market,
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

func (s *syncswap) Unsubscribe(market Market) error {
	if !s.once.Unsubscribe() {
		return errNotStarted
	}

	stream, ok := s.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	sub := stream.(event.Subscription)
	sub.Unsubscribe()

	s.streams.Delete(market)
	return nil
}

type syncswapPoolWrapper struct {
	contract   *isyncswap_pool.ISyncSwapPool
	baseToken  poolToken
	quoteToken poolToken
}

func (s *syncswap) getPool(market Market) (*syncswapPoolWrapper, error) {
	baseToken, quoteToken, err := s.getTokens(market)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	poolAddress, err = s.factory.GetPool(
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
	loggerSyncswap.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := isyncswap_pool.NewISyncSwapPool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}
	return &syncswapPoolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}

func (s *syncswap) getTokens(market Market) (baseToken poolToken, quoteToken poolToken, err error) {
	baseToken, ok := s.assets.Load(strings.ToUpper(market.Base()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Base())
		return
	}
	loggerSyncswap.Infof("market %s: base token address is %s", market, baseToken.Address)

	quoteToken, ok = s.assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Quote())
		return
	}
	loggerSyncswap.Infof("market %s: quote token address is %s", market, quoteToken.Address)

	return baseToken, quoteToken, nil
}
