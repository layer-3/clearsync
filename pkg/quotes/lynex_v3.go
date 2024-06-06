package quotes

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/ilynex_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerLynexV3 = log.Logger("lynex_v3")

type lynexV3 struct {
	factoryAddress common.Address
	factory        *ilynex_v3_factory.ILynexV3Factory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newLynexV3(rpcUrl string, config LynexV3Config, outbox chan<- TradeEvent, history HistoricalData) (Driver, error) {
	hooks := &lynexV3{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := baseDexConfig[
		ilynex_v3_pool.ILynexV3PoolSwap,
		ilynex_v3_pool.ILynexV3Pool,
		*ilynex_v3_pool.ILynexV3PoolSwapIterator,
	]{
		// Params
		DriverType: DriverLynexV3,
		RPC:        rpcUrl,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		IdlePeriod: config.IdlePeriod,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
		IterDeref:     hooks.derefIter,
		// State
		Outbox:  outbox,
		Logger:  loggerLynexV3,
		Filter:  config.Filter,
		History: history,
	}
	return newBaseDEX(params)
}

func (l *lynexV3) postStart(driver *baseDEX[
	ilynex_v3_pool.ILynexV3PoolSwap,
	ilynex_v3_pool.ILynexV3Pool,
	*ilynex_v3_pool.ILynexV3PoolSwapIterator,
]) (err error) {
	l.client = driver.Client()
	l.assets = driver.Assets()

	// Check addresses here: https://lynex.gitbook.io/lynex-docs/security/contracts
	l.factory, err = ilynex_v3_factory.NewILynexV3Factory(l.factoryAddress, l.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Lynex v3 pool factory contract: %w", err)
	}

	return nil
}

func (l *lynexV3) getPool(market Market) ([]*dexPool[ilynex_v3_pool.ILynexV3PoolSwap, *ilynex_v3_pool.ILynexV3PoolSwapIterator], error) {
	baseToken, quoteToken, err := getTokens(l.assets, market, loggerLynexV3)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	loggerLynexV3.Infow("searching for pool", "market", market)
	err = debounce.Debounce(loggerLynexV3, func() error {
		poolAddress, err = l.factory.PoolByPair(nil, baseToken.Address, quoteToken.Address)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerLynexV3.Infow("found pool",
		"market", market,
		"address", poolAddress)

	poolContract, err := ilynex_v3_pool.NewILynexV3Pool(poolAddress, l.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Lynex v3 pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(loggerLynexV3, func() error {
		basePoolToken, err = poolContract.Token0(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Lynex v3 pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(loggerLynexV3, func() error {
		quotePoolToken, err = poolContract.Token1(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Lynex v3 pool: %w", err)
	}

	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*dexPool[ilynex_v3_pool.ILynexV3PoolSwap, *ilynex_v3_pool.ILynexV3PoolSwapIterator]{{
		Contract:   poolContract,
		Address:    poolAddress,
		BaseToken:  baseToken,
		QuoteToken: quoteToken,
		Market:     market,
		Reversed:   isReversed,
	}}

	// Return pools if the token addresses match direct or reversed configurations
	if (baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken) || isReversed {
		return pools, nil
	}
	return nil, fmt.Errorf("failed to build Lynex v3 pool for market %s: %w", market, err)
}

func (l *lynexV3) parseSwap(
	swap *ilynex_v3_pool.ILynexV3PoolSwap,
	pool *dexPool[ilynex_v3_pool.ILynexV3PoolSwap, *ilynex_v3_pool.ILynexV3PoolSwapIterator],
) (trade TradeEvent, err error) {
	opts := v3TradeOpts[ilynex_v3_pool.ILynexV3PoolSwap, *ilynex_v3_pool.ILynexV3PoolSwapIterator]{
		Driver:          DriverLynexV3,
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.Price,
		Pool:            pool,
		Swap:            swap,
		Logger:          loggerLynexV3,
	}
	return buildV3Trade(opts)
}

func (l *lynexV3) derefIter(
	iter *ilynex_v3_pool.ILynexV3PoolSwapIterator,
) *ilynex_v3_pool.ILynexV3PoolSwap {
	return iter.Event
}
