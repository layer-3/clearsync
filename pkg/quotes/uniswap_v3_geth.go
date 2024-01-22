package quotes

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/shopspring/decimal"

	factory "github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	pool "github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
)

type uniswapV3Geth struct {
	once    *once
	url     string
	client  *ethclient.Client
	factory *factory.IUniswapV3Factory

	outbox  chan<- TradeEvent
	streams sync.Map
}

func newUniswapV3Geth(config Config, outbox chan<- TradeEvent) *uniswapV3Geth {
	return &uniswapV3Geth{
		once:   newOnce(),
		url:    config.URL,
		outbox: outbox,
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
				uniswapPool.baseToken.decimals,
				uniswapPool.quoteToken.decimals)

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
	address  string
	decimals decimal.Decimal
}

type poolWrapper struct {
	contract   *pool.IUniswapV3Pool
	baseToken  poolToken
	quoteToken poolToken
}

func (u *uniswapV3Geth) getPool(market Market) (*poolWrapper, error) {
	// TODO: search for token on Etherscan
	allTokens := map[string]poolToken{
		"WETH": {
			address:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			decimals: decimal.RequireFromString("18"),
		},
		"USDT": {
			address:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			decimals: decimal.RequireFromString("6"),
		},
	}

	baseToken, ok := allTokens[market.BaseUnit]
	if !ok {
		return nil, fmt.Errorf("tokens '%s' does not exist", market.BaseUnit)
	}
	quoteToken, ok := allTokens[market.QuoteUnit]
	if !ok {
		return nil, fmt.Errorf("tokens '%s' does not exist", market.QuoteUnit)
	}

	poolAddress, err := u.factory.GetPool(
		nil,
		common.HexToAddress(baseToken.address),
		common.HexToAddress(quoteToken.address),
		big.NewInt(0),
	)
	if err != nil {
		return nil, err
	}

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
