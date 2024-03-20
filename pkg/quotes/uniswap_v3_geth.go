package quotes

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/safe"
)

var (
	loggerUniswapV3Geth = log.Logger("uniswap_v3_geth")
	// Uniswap v3 protocol has the 0.01%, 0.05%, 0.3%, and 1% fee tiers.
	uniswapV3FeeTiers = []uint{100, 500, 3000, 10000}
)

type uniswapV3Geth struct {
	once           *once
	url            string
	assetsURL      string
	factoryAddress string
	client         *ethclient.Client
	factory        *iuniswap_v3_factory.IUniswapV3Factory

	outbox  chan<- TradeEvent
	filter  Filter
	streams safe.Map[Market, event.Subscription]
	assets  safe.Map[string, poolToken]
}

func newUniswapV3Geth(config UniswapV3GethConfig, outbox chan<- TradeEvent) Driver {
	return &uniswapV3Geth{
		once:           newOnce(),
		url:            config.URL,
		assetsURL:      config.AssetsURL,
		factoryAddress: config.FactoryAddress,

		outbox:  outbox,
		filter:  NewFilter(config.Filter),
		streams: safe.NewMap[Market, event.Subscription](),
		assets:  safe.NewMap[string, poolToken](),
	}
}

func (u *uniswapV3Geth) Name() DriverType {
	return DriverUniswapV3Geth
}

func (b *uniswapV3Geth) Type() Type {
	return TypeDEX
}

func (u *uniswapV3Geth) Start() error {
	var startErr error
	started := u.once.Start(func() {
		if !strings.HasPrefix(u.url, "ws") {
			startErr = fmt.Errorf("websocket URL must start with ws:// or wss:// (got %s)", u.url)
			return
		}

		client, err := ethclient.Dial(u.url)
		if err != nil {
			startErr = fmt.Errorf("failed to connect to the Ethereum client: %w", err)
			return
		}
		u.client = client

		// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
		factoryAddress := common.HexToAddress(u.factoryAddress)
		uniswapFactory, err := iuniswap_v3_factory.NewIUniswapV3Factory(factoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to build Uniswap v3 factory: %w", err)
			return
		}
		u.factory = uniswapFactory

		assets, err := getAssets(u.assetsURL)
		if err != nil {
			startErr = fmt.Errorf("failed to fetch assets: %w", err)
			return
		}
		for _, asset := range assets {
			u.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}
	})

	if !started {
		return errAlreadyStarted
	}
	return startErr
}

func (u *uniswapV3Geth) Stop() error {
	stopped := u.once.Stop(func() {
		u.streams.Range(func(market Market, _ event.Subscription) bool {
			err := u.Unsubscribe(market)
			return err == nil
		})

		u.streams = safe.NewMap[Market, event.Subscription]() // delete all stopped streams
	})

	if !stopped {
		return errAlreadyStopped
	}
	return nil
}

func (u *uniswapV3Geth) Subscribe(market Market) error {
	if !u.once.Subscribe() {
		return errNotStarted
	}

	if _, ok := u.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	pool, err := u.getPool(market)
	if err != nil {
		return fmt.Errorf("failed get pool for market %v: %s", market.String(), err)
	}

	go func() {
		block, err := u.client.BlockNumber(context.Background())
		if err != nil {
			loggerSyncswap.Errorf("failed to get block number: %s", err)
		}

		iter, err := pool.contract.FilterSwap(
			&bind.FilterOpts{Start: block - 1000},
			[]common.Address{},
			[]common.Address{})
		if err != nil {
			loggerSyncswap.Errorf("failed to filter swap events for market %s: %s", market, err)
		}
		if ok := iter.Next(); ok {
			tr, err := u.parseSwap(iter.Event, pool)
			if err != nil {
				loggerSyncswap.Errorf("failed to parse swap event for market %s: %s", market, err)
			}
			u.outbox <- tr
		}
	}()

	sink := make(chan *iuniswap_v3_pool.IUniswapV3PoolSwap, 128)
	sub, err := pool.contract.WatchSwap(nil, sink, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(sink)
		for {
			select {
			case err := <-sub.Err():
				loggerUniswapV3Geth.Errorf("market %s: %s", market.String(), err)
				if _, ok := u.streams.Load(market); !ok {
					break // market was unsubscribed earlier
				}
				if err := u.Subscribe(market); err != nil {
					loggerUniswapV3Geth.Errorf("market %s: failed to resubscribe: %s", market.String(), err)
				}
				return
			case swap := <-sink:
				tr, err := u.parseSwap(swap, pool)
				if err != nil {
					loggerUniswapV3Geth.Errorf("failed to parse swap event for market %s: %w", market, err)
				}

				if !u.filter.Allow(tr) {
					continue
				}
				u.outbox <- tr
			}
		}
	}()

	u.streams.Store(market, sub)
	return nil
}

func (u *uniswapV3Geth) Unsubscribe(market Market) error {
	if !u.once.Unsubscribe() {
		return errNotStarted
	}

	stream, ok := u.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	stream.Unsubscribe()
	u.streams.Delete(market)

	return nil
}

func (*uniswapV3Geth) parseSwap(swap *iuniswap_v3_pool.IUniswapV3PoolSwap, pool *uniswapV3GethPoolWrapper) (TradeEvent, error) {
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
	return TradeEvent{
		Source: DriverUniswapV3Geth,
		Market: Market{
			baseUnit:  pool.baseToken.Symbol,
			quoteUnit: pool.quoteToken.Symbol,
		},
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}, nil
}

type poolToken struct {
	Name     string
	Address  string
	Symbol   string
	Decimals decimal.Decimal
	ChainId  uint
	LogoURI  string
}

type uniswapV3GethPoolWrapper struct {
	contract   *iuniswap_v3_pool.IUniswapV3Pool
	baseToken  poolToken
	quoteToken poolToken
}

func (u *uniswapV3Geth) getPool(market Market) (*uniswapV3GethPoolWrapper, error) {
	baseToken, quoteToken, err := u.getTokens(market)
	if err != nil {
		return nil, err
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range uniswapV3FeeTiers {
		poolAddress, err = u.factory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
			big.NewInt(int64(feeTier)),
		)
		if err != nil {
			return nil, err
		}
		if poolAddress != zeroAddress {
			loggerUniswapV3Geth.Infof("market %s: selected fee tier: %.2f%%", market, float64(feeTier)/10000)
			break
		}
	}
	loggerUniswapV3Geth.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := iuniswap_v3_pool.NewIUniswapV3Pool(poolAddress, u.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
	}
	return &uniswapV3GethPoolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}

func (u *uniswapV3Geth) getTokens(market Market) (baseToken poolToken, quoteToken poolToken, err error) {
	baseToken, ok := u.assets.Load(strings.ToUpper(market.Base()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Base())
		return
	}
	loggerUniswapV3Geth.Infof("market %s: base token address is %s", market, baseToken.Address)

	quoteToken, ok = u.assets.Load(strings.ToUpper(market.Quote()))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.Quote())
		return
	}
	loggerUniswapV3Geth.Infof("market %s: quote token address is %s", market, quoteToken.Address)

	return baseToken, quoteToken, nil
}

func getAssets(assetsURL string) ([]poolToken, error) {
	return []poolToken{{
		Address:  "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619",
		Symbol:   "weth",
		Decimals: decimal.NewFromInt(18),
	}, {
		Address:  "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
		Symbol:   "matic",
		Decimals: decimal.NewFromInt(18),
	}, {
		Address:  "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6",
		Symbol:   "wbtc",
		Decimals: decimal.NewFromInt(8),
	}}, nil
}
