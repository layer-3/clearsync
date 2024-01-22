package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	factory "github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	pool "github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
)

var (
	loggerUniswapV3Geth = log.Logger("uniswap_v3_geth")
	// Uniswap v3 protocol has the 1%, 0.3%, 0.05%, and 0.01% fee tiers.
	uniswapV3FeeTiers = []uint{10000, 3000, 500, 100}
)

type uniswapV3Geth struct {
	once      *once
	url       string
	assetsURL string
	client    *ethclient.Client
	factory   *factory.IUniswapV3Factory

	outbox  chan<- TradeEvent
	streams sync.Map
	assets  sync.Map
}

func newUniswapV3Geth(config Config, outbox chan<- TradeEvent) *uniswapV3Geth {
	return &uniswapV3Geth{
		once:      newOnce(),
		url:       config.URL,
		assetsURL: config.AssetsURL,
		outbox:    outbox,
	}
}

func (u *uniswapV3Geth) Start() error {
	var startErr error
	u.once.Start(func() {
		client, err := ethclient.Dial(u.url)
		if err != nil {
			startErr = fmt.Errorf("failed to connect to the Ethereum client: %w", err)
			return
		}
		u.client = client

		// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
		factoryAddress := common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
		uniswapFactory, err := factory.NewIUniswapV3Factory(factoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to build Uniswap v3 factory: %w", err)
			return
		}
		u.factory = uniswapFactory

		// Fetch assets.json

		resp, err := http.Get(u.assetsURL)
		if err != nil {
			startErr = err
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			startErr = err
			return
		}

		var assets []poolToken
		if err := json.Unmarshal(body, &assets); err != nil {
			startErr = err
			return
		}

		for _, asset := range assets {
			u.assets.Store(strings.ToUpper(asset.Symbol), asset)
		}
	})
	return startErr
}

func (u *uniswapV3Geth) Stop() error {
	u.once.Stop(func() {
		u.streams.Range(func(market, stream any) bool {
			err := u.Unsubscribe(market.(Market))
			return err == nil
		})

		u.streams = sync.Map{} // delete all stopped streams
	})
	return nil
}

func (u *uniswapV3Geth) Subscribe(market Market) error {
	symbol := market.BaseUnit + market.QuoteUnit

	if _, ok := u.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	uniswapPool, err := u.getPool(market)
	if err != nil {
		return fmt.Errorf("failed get pool for market %v: %s", symbol, err)
	}

	swapsSink := make(chan *pool.IUniswapV3PoolSwap, 128)
	sub, err := uniswapPool.contract.WatchSwap(
		nil,
		swapsSink,
		[]common.Address{},
		[]common.Address{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to swaps for market %s: %w", market, err)
	}

	go func() {
		defer close(swapsSink)
		for swap := range swapsSink {
			amount := decimal.NewFromBigInt(swap.Amount0, 0)
			price := calculatePrice(
				decimal.NewFromBigInt(swap.SqrtPriceX96, 0),
				uniswapPool.baseToken.Decimals,
				uniswapPool.quoteToken.Decimals)

			u.outbox <- TradeEvent{
				Source:    DriverUniswapV3Geth,
				Market:    symbol,
				Price:     price,
				Amount:    amount,
				Total:     price.Mul(amount),
				TakerType: TakerTypeSell,
				CreatedAt: time.Now(),
			}
		}
	}()

	u.streams.Store(market, sub)
	return nil
}

func (u *uniswapV3Geth) Unsubscribe(market Market) error {
	stream, ok := u.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	sub := stream.(event.Subscription)
	sub.Unsubscribe()

	u.streams.Delete(market)
	return nil
}

type poolToken struct {
	Name     string
	Address  string
	Symbol   string
	Decimals decimal.Decimal
	ChainId  uint
	LogoURI  string
}

type poolWrapper struct {
	contract   *pool.IUniswapV3Pool
	baseToken  poolToken
	quoteToken poolToken
}

func (u *uniswapV3Geth) getPool(market Market) (*poolWrapper, error) {
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
			loggerUniswapV3Geth.Infof("selected %s fee tier for market %s", feeTier, market)
			break
		}
	}
	loggerUniswapV3Geth.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := pool.NewIUniswapV3Pool(poolAddress, u.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
	}
	return &poolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}

func (u *uniswapV3Geth) getTokens(market Market) (baseToken poolToken, quoteToken poolToken, err error) {
	baseAsset, ok := u.assets.Load(strings.ToUpper(market.BaseUnit))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.BaseUnit)
		return
	}
	baseToken = baseAsset.(poolToken)
	loggerUniswapV3Geth.Infof("market %s: base token address is %s", market, baseToken.Address)

	quoteAsset, ok := u.assets.Load(strings.ToUpper(market.QuoteUnit))
	if !ok {
		err = fmt.Errorf("tokens '%s' does not exist", market.QuoteUnit)
		return
	}
	quoteToken = quoteAsset.(poolToken)
	loggerUniswapV3Geth.Infof("market %s: quote token address is %s", market, quoteToken.Address)

	return baseToken, quoteToken, nil
}
