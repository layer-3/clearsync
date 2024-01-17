//go:generate sh -c "command -v jq >/dev/null 2>&1 || { echo 'Error: jq is not installed.' >&2; exit 1; } && command -v abigen >/dev/null 2>&1 || { echo 'Error: abigen is not installed.' >&2; exit 1; } && mkdir -p abi/factory && cat ../../../node_modules/@uniswap/v3-core/artifacts/contracts/UniswapV3Factory.sol/UniswapV3Factory.json | jq '.abi' | abigen --pkg=factory --out=./abi/factory/factory.go --abi=-"
//go:generate sh -c "command -v jq >/dev/null 2>&1 || { echo 'Error: jq is not installed.' >&2; exit 1; } && command -v abigen >/dev/null 2>&1 || { echo 'Error: abigen is not installed.' >&2; exit 1; } && mkdir -p abi/factory && cat ../../../node_modules/@uniswap/v3-core/artifacts/contracts/UniswapV3Pool.sol/UniswapV3Pool.json | jq '.abi' | abigen --pkg=pool --out=./abi/pool/pool.go --abi=-"
package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"sync"
	"time"

	geth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/uniswap/abi/factory"
	"github.com/layer-3/clearsync/pkg/quotes/uniswap/abi/pool"
)

type V3Geth struct {
	once    *common.Once
	url     string
	client  *ethclient.Client
	factory *factory.Factory

	outbox  chan<- common.TradeEvent
	streams sync.Map
}

func NewV3Geth(config common.Config, outbox chan<- common.TradeEvent) *V3Geth {
	return &V3Geth{
		once:   common.NewOnce(),
		url:    config.URL,
		outbox: outbox,
	}
}

func (u *V3Geth) Start() error {
	var startErr error
	u.once.Start(func() {
		client, err := ethclient.Dial(u.url)
		if err != nil {
			startErr = fmt.Errorf("failed to connect to the Ethereum client: %w", err)
			return
		}
		u.client = client

		// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
		factoryAddress := geth_common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
		uniswapFactory, err := factory.NewFactory(factoryAddress, client)
		if err != nil {
			startErr = fmt.Errorf("failed to build Uniswap v3 factory: %w", err)
			return
		}
		u.factory = uniswapFactory
	})
	return startErr
}

func (u *V3Geth) Stop() error {
	u.once.Stop(func() {
		u.streams.Range(func(market, stream any) bool {
			err := u.Unsubscribe(market.(common.Market))
			return err == nil
		})

		u.streams = sync.Map{} // delete all stopped streams
	})
	return nil
}

func (u *V3Geth) Subscribe(market common.Market) error {
	symbol := market.BaseUnit + market.QuoteUnit

	if _, ok := u.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
	}

	uniswapPool, err := u.getPool(market)
	if err != nil {
		return fmt.Errorf("failed get pool for market %v: %s", symbol, err)
	}

	swapsSink := make(chan *pool.PoolSwap, 128)
	sub, err := uniswapPool.contract.WatchSwap(
		nil,
		swapsSink,
		[]geth_common.Address{},
		[]geth_common.Address{})
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

			u.outbox <- common.TradeEvent{
				Source:    common.DriverUniswapV3Geth,
				Market:    symbol,
				Price:     price,
				Amount:    amount,
				Total:     price.Mul(amount),
				TakerType: common.TakerTypeSell,
				CreatedAt: time.Now(),
			}
		}
	}()

	u.streams.Store(market, sub)
	return nil
}

func (u *V3Geth) Unsubscribe(market common.Market) error {
	stream, ok := u.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
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
	contract   *pool.Pool
	baseToken  poolToken
	quoteToken poolToken
}

func (u *V3Geth) getPool(market common.Market) (*poolWrapper, error) {
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
		geth_common.HexToAddress(baseToken.address),
		geth_common.HexToAddress(quoteToken.address),
		big.NewInt(0),
	)
	if err != nil {
		return nil, err
	}

	poolContract, err := pool.NewPool(poolAddress, u.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
	}
	return &poolWrapper{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}
