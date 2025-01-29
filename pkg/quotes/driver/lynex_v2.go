package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_factory"
	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_pair"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var loggerLynexV2 = log.Logger("lynex_v2")

type lynexV2 struct {
	stablePoolMarkets map[quotes_common.Market]struct{}
	factoryAddress    common.Address
	factory           *ilynex_v2_factory.ILynexFactory

	driver base.DexReader
}

func newLynexV2(rpcUrl string, config LynexV2Config, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	stablePoolMarkets := make(map[quotes_common.Market]struct{})
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := quotes_common.NewMarketFromString(rawMarket)
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

	params := base.DexConfig[
		ilynex_v2_pair.ILynexPairSwap,
		*ilynex_v2_pair.ILynexPairSwapIterator,
	]{
		// Params
		DriverType: quotes_common.DriverLynexV2,
		RPC:        rpcUrl,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		MarketsURL: config.MarketsURL,
		IdlePeriod: config.IdlePeriod,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		ParserFactory: hooks.buildParser,
		IterDeref:     hooks.derefIter,

		// State
		Outbox:  outbox,
		Logger:  loggerLynexV2,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (l *lynexV2) postStart(driver *base.DEX[
	ilynex_v2_pair.ILynexPairSwap,
	*ilynex_v2_pair.ILynexPairSwapIterator,
]) (err error) {
	l.driver = driver

	// Check addresses here: https://lynex.gitbook.io/lynex-docs/security/contracts
	l.factory, err = ilynex_v2_factory.NewILynexFactory(l.factoryAddress, l.driver.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Lynex v2 pool factory contract: %w", err)
	}

	return nil
}

func (l *lynexV2) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[ilynex_v2_pair.ILynexPairSwap, *ilynex_v2_pair.ILynexPairSwapIterator], error) {
	baseToken, quoteToken, err := base.GetTokens(l.driver.Assets(), market, loggerLynexV2)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	_, isStablePool := l.stablePoolMarkets[market]

	loggerLynexV2.Infow("searching for pool", "market", market)
	err = debounce.Debounce(context.TODO(), loggerLynexV2, func(ctx context.Context) error {
		poolAddress, err = l.factory.GetPair(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address, isStablePool)
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

	poolContract, err := ilynex_v2_pair.NewILynexPair(poolAddress, l.driver.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to build Lynex v2 pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(ctx, loggerLynexV2, func(ctx context.Context) error {
		basePoolToken, err = poolContract.Token0(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Lynex v2 pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(ctx, loggerLynexV2, func(ctx context.Context) error {
		quotePoolToken, err = poolContract.Token1(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Lynex v2 pool: %w", err)
	}

	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*base.DexPool[ilynex_v2_pair.ILynexPairSwap, *ilynex_v2_pair.ILynexPairSwapIterator]{{
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

func (l *lynexV2) buildParser(
	swap *ilynex_v2_pair.ILynexPairSwap,
	pool *base.DexPool[ilynex_v2_pair.ILynexPairSwap, *ilynex_v2_pair.ILynexPairSwapIterator],
) base.SwapParser {
	return &base.SwapV2[
		ilynex_v2_pair.ILynexPairSwap,
		*ilynex_v2_pair.ILynexPairSwapIterator,
	]{
		RawAmount0In:  swap.Amount0In,
		RawAmount0Out: swap.Amount0Out,
		RawAmount1In:  swap.Amount1In,
		RawAmount1Out: swap.Amount1Out,
		Pool:          pool,
	}
}

func (l *lynexV2) derefIter(
	iter *ilynex_v2_pair.ILynexPairSwapIterator,
) *ilynex_v2_pair.ILynexPairSwap {
	return iter.Event
}
