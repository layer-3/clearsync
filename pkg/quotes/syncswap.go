package quotes

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/isyncswap_factory"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerSyncswap = log.Logger("syncswap")

type syncswap struct {
	stablePoolMarkets         map[Market]struct{}
	classicPoolFactoryAddress common.Address
	stablePoolFactoryAddress  common.Address
	classicFactory            *isyncswap_factory.ISyncSwapFactory
	stableFactory             *isyncswap_factory.ISyncSwapFactory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newSyncswap(config SyncswapConfig, outbox chan<- TradeEvent, history HistoricalData) (Driver, error) {
	stablePoolMarkets := make(map[Market]struct{})
	logStablePoolMarkets := make([]Market, 0, len(config.StablePoolMarkets))
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := NewMarketFromString(rawMarket)
		if !ok {
			loggerSyncswap.Errorw("failed to parse stable pool from market", "market", rawMarket)
			continue
		}
		stablePoolMarkets[market] = struct{}{}
		logStablePoolMarkets = append(logStablePoolMarkets, market)
	}
	loggerSyncswap.Debugw("configured stable pool markets", "markets", logStablePoolMarkets)

	hooks := &syncswap{
		stablePoolMarkets:         stablePoolMarkets,
		classicPoolFactoryAddress: common.HexToAddress(config.ClassicPoolFactoryAddress),
		stablePoolFactoryAddress:  common.HexToAddress(config.StablePoolFactoryAddress),
	}

	params := baseDexConfig[
		isyncswap_pool.ISyncSwapPoolSwap,
		isyncswap_pool.ISyncSwapPool,
		*isyncswap_pool.ISyncSwapPoolSwapIterator,
	]{
		// Params
		DriverType: DriverSyncswap,
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
		Logger:  loggerSyncswap,
		Filter:  config.Filter,
		History: history,
	}
	return newBaseDEX(params)
}

func (s *syncswap) postStart(driver *baseDEX[
	isyncswap_pool.ISyncSwapPoolSwap,
	isyncswap_pool.ISyncSwapPool,
	*isyncswap_pool.ISyncSwapPoolSwapIterator,
]) (err error) {
	s.client = driver.Client()
	s.assets = driver.Assets()

	// Check addresses here: https://syncswap.gitbook.io/syncswap/smart-contracts/smart-contracts
	s.classicFactory, err = isyncswap_factory.NewISyncSwapFactory(s.classicPoolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap classic pool factory contract: %w", err)
	}

	s.stableFactory, err = isyncswap_factory.NewISyncSwapFactory(s.stablePoolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap stable pool factory contract: %w", err)
	}
	return nil
}

func (s *syncswap) getPool(market Market) ([]*dexPool[isyncswap_pool.ISyncSwapPoolSwap, *isyncswap_pool.ISyncSwapPoolSwapIterator], error) {
	baseToken, quoteToken, err := getTokens(s.assets, market, loggerSyncswap)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	if _, ok := s.stablePoolMarkets[market]; ok {
		loggerSyncswap.Infow("searching for stable pool", "market", market, "address", poolAddress)
		err = debounce.Debounce(loggerSyncswap, func() error {
			poolAddress, err = s.stableFactory.GetPool(nil, baseToken.Address, quoteToken.Address)
			return err
		})
		loggerSyncswap.Infow("found stable pool", "market", market)
	} else {
		loggerSyncswap.Infow("searching for classic pool", "market", market)
		err = debounce.Debounce(loggerSyncswap, func() error {
			poolAddress, err = s.classicFactory.GetPool(nil, baseToken.Address, quoteToken.Address)
			return err
		})
		loggerSyncswap.Infow("found classic pool", "market", market, "address", poolAddress)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get classic pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("classic pool for market %s does not exist", market)
	}
	loggerSyncswap.Infow("pool found", "market", market, "address", poolAddress)

	poolContract, err := isyncswap_pool.NewISyncSwapPool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(loggerSyncswap, func() error {
		basePoolToken, err = poolContract.Token0(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Syncswap pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(loggerSyncswap, func() error {
		quotePoolToken, err = poolContract.Token1(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Syncswap pool: %w", err)
	}

	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*dexPool[isyncswap_pool.ISyncSwapPoolSwap, *isyncswap_pool.ISyncSwapPoolSwapIterator]{{
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
	return nil, fmt.Errorf("failed to build Syncswap pool for market %s: %w", market, err)
}

func (s *syncswap) parseSwap(
	swap *isyncswap_pool.ISyncSwapPoolSwap,
	pool *dexPool[isyncswap_pool.ISyncSwapPoolSwap, *isyncswap_pool.ISyncSwapPoolSwapIterator],
) (trade TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			loggerSyncswap.Errorw(ErrSwapParsing.Error(), "swap", swap, "pool", pool)
			err = fmt.Errorf("%s: %s", ErrSwapParsing, r)
		}
	}()

	return buildV2Trade(
		DriverSyncswap,
		swap.Amount0In,
		swap.Amount0Out,
		swap.Amount1In,
		swap.Amount1Out,
		pool,
		swap,
		loggerSyncswap,
	)
}

func (s *syncswap) derefIter(
	iter *isyncswap_pool.ISyncSwapPoolSwapIterator,
) *isyncswap_pool.ISyncSwapPoolSwap {
	return iter.Event
}
