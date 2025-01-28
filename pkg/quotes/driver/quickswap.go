package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerQuickswap = log.Logger("quickswap")

type quickswap struct {
	poolFactoryAddress common.Address
	factory            *quickswap_v3_factory.IQuickswapV3Factory

	assets *safe.Map[string, base.DexPoolToken]
	client *ethclient.Client
}

func newQuickswap(rpcUrl string, config QuickswapConfig, outbox chan<- quotes_common.TradeEvent, history base.HistoricalDataDriver) (base.Driver, error) {
	hooks := &quickswap{
		poolFactoryAddress: common.HexToAddress(config.PoolFactoryAddress),
	}

	params := base.DexConfig[
		quickswap_v3_pool.IQuickswapV3PoolSwap,
		quickswap_v3_pool.IQuickswapV3Pool,
		*quickswap_v3_pool.IQuickswapV3PoolSwapIterator,
	]{
		// Params
		DriverType: quotes_common.DriverQuickswap,
		RPC:        rpcUrl,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		MarketsURL: config.MarketsURL,
		IdlePeriod: config.IdlePeriod,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
		IterDeref:     hooks.derefIter,
		// State
		Outbox:  outbox,
		Logger:  loggerQuickswap,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (s *quickswap) postStart(driver *base.DEX[
	quickswap_v3_pool.IQuickswapV3PoolSwap,
	quickswap_v3_pool.IQuickswapV3Pool,
	*quickswap_v3_pool.IQuickswapV3PoolSwapIterator,
]) (err error) {
	s.client = driver.Client()
	s.assets = driver.Assets()

	// Check addresses here: https://quickswap.gitbook.io/quickswap/smart-contracts/smart-contracts
	s.factory, err = quickswap_v3_factory.NewIQuickswapV3Factory(s.poolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickwap Factory contract: %w", err)
	}
	return nil
}

func (s *quickswap) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[quickswap_v3_pool.IQuickswapV3PoolSwap, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator], error) {
	baseToken, quoteToken, err := base.GetTokens(s.assets, market, loggerQuickswap)
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

	poolContract, err := quickswap_v3_pool.NewIQuickswapV3Pool(poolAddress, s.client)
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

	pools := []*base.DexPool[quickswap_v3_pool.IQuickswapV3PoolSwap, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator]{{
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

func (s *quickswap) parseSwap(
	swap *quickswap_v3_pool.IQuickswapV3PoolSwap,
	pool *base.DexPool[quickswap_v3_pool.IQuickswapV3PoolSwap, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator],
) (trade quotes_common.TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			loggerQuickswap.Errorw(quotes_common.ErrSwapParsing.Error(), "swap", swap, "pool", pool)
			err = fmt.Errorf("%s: %s", quotes_common.ErrSwapParsing, r)
		}
	}()

	opts := base.V3TradeOpts[
		quickswap_v3_pool.IQuickswapV3PoolSwap,
		*quickswap_v3_pool.IQuickswapV3PoolSwapIterator,
	]{
		Driver:          quotes_common.DriverQuickswap,
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.Price,
		Pool:            pool,
		Swap:            swap,
		Logger:          loggerQuickswap,
	}
	return base.BuildV3Trade(opts)
}

func (s *quickswap) derefIter(
	iter *quickswap_v3_pool.IQuickswapV3PoolSwapIterator,
) *quickswap_v3_pool.IQuickswapV3PoolSwap {
	return iter.Event
}
