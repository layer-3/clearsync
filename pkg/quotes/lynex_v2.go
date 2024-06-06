package quotes

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_factory"
	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_pair"
	"github.com/layer-3/clearsync/pkg/debounce"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerLynexV2 = log.Logger("lynex_v2")

type lynexV2 struct {
	stablePoolMarkets map[Market]struct{}
	factoryAddress    common.Address
	factory           *ilynex_v2_factory.ILynexFactory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newLynexV2(config LynexV2Config, outbox chan<- TradeEvent, history HistoricalData) (Driver, error) {
	stablePoolMarkets := make(map[Market]struct{})
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := NewMarketFromString(rawMarket)
		if !ok {
			loggerLynexV2.Errorw("failed to parse stable pool from market", "market", rawMarket)
			continue
		}
		stablePoolMarkets[market] = struct{}{}
	}
	loggerLynexV2.Debugw("configured stable pool markets", "markets", stablePoolMarkets)

	hooks := &lynexV2{
		stablePoolMarkets: stablePoolMarkets,
		factoryAddress:    common.HexToAddress(config.FactoryAddress),
	}

	params := baseDexConfig[
		ilynex_v2_pair.ILynexPairSwap,
		ilynex_v2_pair.ILynexPair,
		*ilynex_v2_pair.ILynexPairSwapIterator,
	]{
		// Params
		DriverType: DriverLynexV2,
		URL:        config.URL,
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
		Logger:  loggerLynexV2,
		Filter:  config.Filter,
		History: history,
	}
	return newBaseDEX(params)
}

func (l *lynexV2) postStart(driver *baseDEX[
	ilynex_v2_pair.ILynexPairSwap,
	ilynex_v2_pair.ILynexPair,
	*ilynex_v2_pair.ILynexPairSwapIterator,
]) (err error) {
	l.client = driver.Client()
	l.assets = driver.Assets()

	// Check addresses here: https://lynex.gitbook.io/lynex-docs/security/contracts
	l.factory, err = ilynex_v2_factory.NewILynexFactory(l.factoryAddress, l.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Lynex v2 pool factory contract: %w", err)
	}

	return nil
}

func (l *lynexV2) getPool(market Market) ([]*dexPool[ilynex_v2_pair.ILynexPairSwap, *ilynex_v2_pair.ILynexPairSwapIterator], error) {
	baseToken, quoteToken, err := getTokens(l.assets, market, loggerLynexV2)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	_, isStablePool := l.stablePoolMarkets[market]

	loggerLynexV2.Infow("searching for pool", "market", market)
	err = debounce.Debounce(loggerLynexV2, func() error {
		poolAddress, err = l.factory.GetPair(nil, baseToken.Address, quoteToken.Address, isStablePool)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerLynexV2.Infow("found pool",
		"market", market,
		"address", poolAddress,
		"is_stable", isStablePool)

	poolContract, err := ilynex_v2_pair.NewILynexPair(poolAddress, l.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Lynex v2 pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(loggerLynexV2, func() error {
		basePoolToken, err = poolContract.Token0(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Lynex v2 pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(loggerLynexV2, func() error {
		quotePoolToken, err = poolContract.Token1(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Lynex v2 pool: %w", err)
	}

	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*dexPool[ilynex_v2_pair.ILynexPairSwap, *ilynex_v2_pair.ILynexPairSwapIterator]{{
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
	return nil, fmt.Errorf("failed to build Lynex v2 pool for market %s: %w", market, err)
}

func (l *lynexV2) parseSwap(
	swap *ilynex_v2_pair.ILynexPairSwap,
	pool *dexPool[ilynex_v2_pair.ILynexPairSwap, *ilynex_v2_pair.ILynexPairSwapIterator],
) (trade TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			loggerLynexV2.Errorw(ErrSwapParsing.Error(), "swap", swap, "pool", pool)
			err = fmt.Errorf("%s: %s", ErrSwapParsing, r)
		}
	}()

	return buildV2Trade(
		DriverLynexV2,
		swap.Amount0In,
		swap.Amount0Out,
		swap.Amount1In,
		swap.Amount1Out,
		pool,
		swap,
		loggerLynexV2,
	)
}

func (l *lynexV2) derefIter(
	iter *ilynex_v2_pair.ILynexPairSwapIterator,
) *ilynex_v2_pair.ILynexPairSwap {
	return iter.Event
}
