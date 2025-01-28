package driver

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
	"github.com/layer-3/clearsync/pkg/safe"
)

var (
	loggerUniswapV3 = log.Logger("uniswap_v3")
	// Uniswap v3 protocol has the 0.01%, 0.05%, 0.3%, and 1% fee tiers.
	uniswapV3FeeTiers = []uint{100, 500, 3000, 10000}
	wethContract      = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)

type uniswapV3 struct {
	factoryAddress common.Address
	factory        *iuniswap_v3_factory.IUniswapV3Factory

	assets *safe.Map[string, base.DexPoolToken]
	client *ethclient.Client
}

func newUniswapV3(rpcUrl string, config UniswapV3Config, outbox chan<- quotes_common.TradeEvent, history base.HistoricalDataDriver) (base.Driver, error) {
	hooks := &uniswapV3{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := base.DexConfig[
		iuniswap_v3_pool.IUniswapV3PoolSwap,
		iuniswap_v3_pool.IUniswapV3Pool,
		*iuniswap_v3_pool.IUniswapV3PoolSwapIterator,
	]{
		// Params
		DriverType: quotes_common.DriverUniswapV3,
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
		Logger:  loggerUniswapV3,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (u *uniswapV3) postStart(driver *base.DEX[
	iuniswap_v3_pool.IUniswapV3PoolSwap,
	iuniswap_v3_pool.IUniswapV3Pool,
	*iuniswap_v3_pool.IUniswapV3PoolSwapIterator,
]) (err error) {
	u.client = driver.Client()
	u.assets = driver.Assets()

	// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
	u.factory, err = iuniswap_v3_factory.NewIUniswapV3Factory(u.factoryAddress, u.client)
	if err != nil {
		return fmt.Errorf("failed to build Uniswap v3 factory: %w", err)
	}
	return nil
}

func (u *uniswapV3) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator], error) {
	baseToken, quoteToken, err := base.GetTokens(u.assets, market, loggerUniswapV3)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(baseToken.Symbol) == "eth" {
		baseToken.Address = wethContract
	}

	poolAddresses := make([]common.Address, 0, len(uniswapV3FeeTiers))
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range uniswapV3FeeTiers {
		var poolAddress common.Address
		err = debounce.Debounce(ctx, loggerUniswapV3, func(ctx context.Context) error {
			poolAddress, err = u.factory.GetPool(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address, big.NewInt(int64(feeTier)))
			return err
		})
		if err != nil {
			return nil, err
		}

		if poolAddress != zeroAddress {
			loggerUniswapV3.Infow("found pool",
				"market", market,
				"selected fee tier", fmt.Sprintf("%.2f%%", float64(feeTier)/10000),
				"address", poolAddress)
			poolAddresses = append(poolAddresses, poolAddress)
		}
	}

	pools := make([]*base.DexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator], 0, len(poolAddresses))
	for _, poolAddress := range poolAddresses {
		poolContract, err := iuniswap_v3_pool.NewIUniswapV3Pool(poolAddress, u.client)
		if err != nil {
			return nil, fmt.Errorf("failed to build Uniswap v3 pool contract: %w", err)
		}

		var basePoolToken common.Address
		err = debounce.Debounce(ctx, loggerUniswapV3, func(ctx context.Context) error {
			basePoolToken, err = poolContract.Token0(&bind.CallOpts{Context: ctx})
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get base token address for Uniswap v3 pool: %w", err)
		}

		var quotePoolToken common.Address
		err = debounce.Debounce(ctx, loggerUniswapV3, func(ctx context.Context) error {
			quotePoolToken, err = poolContract.Token1(&bind.CallOpts{Context: ctx})
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get quote token address for Uniswap v3 pool: %w", err)
		}

		isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
		pool := &base.DexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]{
			Contract:   poolContract,
			Address:    poolAddress,
			BaseToken:  baseToken,
			QuoteToken: quoteToken,
			Market:     market,
			Reversed:   isReversed,
		}

		// Append pool if the token addresses match direct or reversed configurations
		if (baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken) || isReversed {
			pools = append(pools, pool)
		}
	}

	return pools, nil
}

func (u *uniswapV3) parseSwap(
	swap *iuniswap_v3_pool.IUniswapV3PoolSwap,
	pool *base.DexPool[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator],
) (trade quotes_common.TradeEvent, err error) {
	opts := base.V3TradeOpts[iuniswap_v3_pool.IUniswapV3PoolSwap, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator]{
		Driver:          quotes_common.DriverUniswapV3,
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.SqrtPriceX96,
		Pool:            pool,
		Swap:            swap,
		Logger:          loggerUniswapV3,
	}
	return base.BuildV3Trade(opts)
}

func (u *uniswapV3) derefIter(
	iter *iuniswap_v3_pool.IUniswapV3PoolSwapIterator,
) *iuniswap_v3_pool.IUniswapV3PoolSwap {
	return iter.Event
}
