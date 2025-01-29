package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var loggerQuickswap = log.Logger("quickswap")

type quickswapV3Event = quickswap_v3_pool.IQuickswapV3PoolSwap
type quickswapV3Iterator = *quickswap_v3_pool.IQuickswapV3PoolSwapIterator

type quickswap struct {
	poolFactoryAddress common.Address
	factory            *quickswap_v3_factory.IQuickswapV3Factory

	driver base.DexReader
}

func newQuickswap(rpcUrl string, config QuickswapConfig, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	hooks := &quickswap{
		poolFactoryAddress: common.HexToAddress(config.PoolFactoryAddress),
	}

	params := base.DexConfig[quickswapV3Event, quickswapV3Iterator]{
		Params: base.DexParams{
			Type:       quotes_common.DriverQuickswap,
			RPC:        rpcUrl,
			AssetsURL:  config.AssetsURL,
			MappingURL: config.MappingURL,
			MarketsURL: config.MarketsURL,
			IdlePeriod: config.IdlePeriod,
		},
		Hooks: base.DexHooks[quickswapV3Event, quickswapV3Iterator]{
			PostStart:   hooks.postStart,
			GetPool:     hooks.getPool,
			BuildParser: hooks.buildParser,
			DerefIter:   hooks.derefIter,
		},
		// State
		Outbox:  outbox,
		Logger:  loggerQuickswap,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (s *quickswap) postStart(driver *base.DEX[quickswapV3Event, quickswapV3Iterator]) (err error) {
	s.driver = driver

	// Check addresses here: https://quickswap.gitbook.io/quickswap/smart-contracts/smart-contracts
	s.factory, err = quickswap_v3_factory.NewIQuickswapV3Factory(s.poolFactoryAddress, s.driver.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickwap Factory contract: %w", err)
	}
	return nil
}

func (s *quickswap) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[quickswapV3Event, quickswapV3Iterator], error) {
	baseToken, quoteToken, err := base.GetTokens(s.driver.Assets(), market, loggerQuickswap)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	err = debounce.Debounce(ctx, loggerQuickswap, func(ctx context.Context) error {
		poolAddress, err = s.factory.PoolByPair(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerQuickswap.Infow("found pool", "market", market, "address", poolAddress)

	poolContract, err := quickswap_v3_pool.NewIQuickswapV3Pool(poolAddress, s.driver.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to build Quickswap pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(ctx, loggerQuickswap, func(ctx context.Context) error {
		basePoolToken, err = poolContract.Token0(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Quickswap pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(ctx, loggerQuickswap, func(ctx context.Context) error {
		quotePoolToken, err = poolContract.Token1(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Quickswap pool: %w", err)
	}

	isDirect := baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken
	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken

	pools := []*base.DexPool[quickswapV3Event, quickswapV3Iterator]{{
		Contract:   poolContract,
		Address:    poolAddress,
		BaseToken:  baseToken,
		QuoteToken: quoteToken,
		Market:     market,
		Reversed:   isReversed,
	}}

	// Return pools if the token addresses match direct or reversed configurations
	if isDirect || isReversed {
		return pools, nil
	}
	return nil, fmt.Errorf("failed to build Quickswap pool for market %s: %w", market, err)
}

func (s *quickswap) buildParser(
	swap *quickswapV3Event,
	pool *base.DexPool[quickswapV3Event, quickswapV3Iterator],
) base.SwapParser {
	return &base.SwapV3[quickswapV3Event, quickswapV3Iterator]{
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.Price,
		Pool:            pool,
	}
}

func (s *quickswap) derefIter(iter quickswapV3Iterator) *quickswap_v3_pool.IQuickswapV3PoolSwap {
	return iter.Event
}
