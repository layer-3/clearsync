package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/isyncswap_factory"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var loggerSyncswap = log.Logger("syncswap")

type syncswapEvent = isyncswap_pool.ISyncSwapPoolSwap
type syncswapIterator = *isyncswap_pool.ISyncSwapPoolSwapIterator

type syncswap struct {
	stablePoolMarkets         map[quotes_common.Market]struct{}
	classicPoolFactoryAddress common.Address
	stablePoolFactoryAddress  common.Address
	classicFactory            *isyncswap_factory.ISyncSwapFactory
	stableFactory             *isyncswap_factory.ISyncSwapFactory

	driver base.DexReader
}

func newSyncswap(rpcUrl string, config SyncswapConfig, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	stablePoolMarkets := make(map[quotes_common.Market]struct{})
	logStablePoolMarkets := make([]quotes_common.Market, 0, len(config.StablePoolMarkets))
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := quotes_common.NewMarketFromString(rawMarket)
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

	params := base.DexConfig[syncswapEvent, syncswapIterator]{
		Params: base.DexParams{
			Type:       quotes_common.DriverSyncswap,
			RPC:        rpcUrl,
			AssetsURL:  config.AssetsURL,
			MappingURL: config.MappingURL,
			MarketsURL: config.MarketsURL,
			IdlePeriod: config.IdlePeriod,
		},
		Hooks: base.DexHooks[syncswapEvent, syncswapIterator]{
			PostStart:   hooks.postStart,
			GetPool:     hooks.getPool,
			BuildParser: hooks.buildParser,
			DerefIter:   hooks.derefIter,
		},
		// State
		Outbox:  outbox,
		Logger:  loggerSyncswap,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (s *syncswap) postStart(driver *base.DEX[syncswapEvent, syncswapIterator]) (err error) {
	s.driver = driver

	// Check addresses here: https://syncswap.gitbook.io/syncswap/smart-contracts/smart-contracts
	s.classicFactory, err = isyncswap_factory.NewISyncSwapFactory(s.classicPoolFactoryAddress, s.driver.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap classic pool factory contract: %w", err)
	}

	s.stableFactory, err = isyncswap_factory.NewISyncSwapFactory(s.stablePoolFactoryAddress, s.driver.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap stable pool factory contract: %w", err)
	}
	return nil
}

func (s *syncswap) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[syncswapEvent, syncswapIterator], error) {
	baseToken, quoteToken, err := base.GetTokens(s.driver.Assets(), market, loggerSyncswap)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	if _, ok := s.stablePoolMarkets[market]; ok {
		loggerSyncswap.Infow("searching for stable pool", "market", market, "address", poolAddress)
		err = debounce.Debounce(ctx, loggerSyncswap, func(ctx context.Context) error {
			poolAddress, err = s.stableFactory.GetPool(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address)
			return err
		})
		loggerSyncswap.Infow("found stable pool", "market", market)
	} else {
		loggerSyncswap.Infow("searching for classic pool", "market", market)
		err = debounce.Debounce(ctx, loggerSyncswap, func(ctx context.Context) error {
			poolAddress, err = s.classicFactory.GetPool(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address)
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

	poolContract, err := isyncswap_pool.NewISyncSwapPool(poolAddress, s.driver.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(ctx, loggerSyncswap, func(ctx context.Context) error {
		basePoolToken, err = poolContract.Token0(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Syncswap pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(ctx, loggerSyncswap, func(ctx context.Context) error {
		quotePoolToken, err = poolContract.Token1(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Syncswap pool: %w", err)
	}

	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*base.DexPool[syncswapEvent, syncswapIterator]{{
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

func (s *syncswap) buildParser(
	swap *syncswapEvent,
	pool *base.DexPool[syncswapEvent, syncswapIterator],
) base.SwapParser {
	return &base.SwapV2[syncswapEvent, syncswapIterator]{
		RawAmount0In:  swap.Amount0In,
		RawAmount0Out: swap.Amount0Out,
		RawAmount1In:  swap.Amount1In,
		RawAmount1Out: swap.Amount1Out,
		Pool:          pool,
	}
}

func (s *syncswap) derefIter(iter syncswapIterator) *syncswapEvent {
	return iter.Event
}
