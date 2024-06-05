package quotes

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_factory"
	"github.com/layer-3/clearsync/pkg/abi/ilynex_pair"
	"github.com/layer-3/clearsync/pkg/debounce"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerLynex = log.Logger("lynex")

type lynex struct {
	stablePoolMarkets map[Market]struct{}
	factoryAddress    common.Address
	factory           *ilynex_factory.ILynexFactory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newLynex(config LynexConfig, outbox chan<- TradeEvent) Driver {
	stablePoolMarkets := make(map[Market]struct{})
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := NewMarketFromString(rawMarket)
		if !ok {
			loggerLynex.Errorw("failed to parse stable pool from market", "market", rawMarket)
			continue
		}
		stablePoolMarkets[market] = struct{}{}
	}
	loggerLynex.Debugw("configured stable pool markets", "markets", stablePoolMarkets)

	hooks := &lynex{
		stablePoolMarkets: stablePoolMarkets,
		factoryAddress:    common.HexToAddress(config.FactoryAddress),
	}

	params := baseDexConfig[
		ilynex_pair.ILynexPairSwap,
		ilynex_pair.ILynexPair,
	]{
		// Params
		DriverType: DriverLynex,
		URL:        config.URL,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		IdlePeriod: config.IdlePeriod,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
		// State
		Outbox: outbox,
		Logger: loggerLynex,
		Filter: config.Filter,
	}
	return newBaseDEX(params)
}

func (l *lynex) postStart(driver *baseDEX[
	ilynex_pair.ILynexPairSwap,
	ilynex_pair.ILynexPair,
]) (err error) {
	l.client = driver.Client()
	l.assets = driver.Assets()

	// Check addresses here: https://lynex.gitbook.io/lynex-docs/security/contracts
	l.factory, err = ilynex_factory.NewILynexFactory(l.factoryAddress, l.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Lynex pool factory contract: %w", err)
	}

	return nil
}

func (l *lynex) getPool(market Market) ([]*dexPool[ilynex_pair.ILynexPairSwap], error) {
	baseToken, quoteToken, err := getTokens(l.assets, market, loggerLynex)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	_, isStablePool := l.stablePoolMarkets[market]

	loggerLynex.Infow("searching for classic pool", "market", market)
	err = debounce.Debounce(loggerLynex, func() error {
		poolAddress, err = l.factory.GetPair(nil, baseToken.Address, quoteToken.Address, isStablePool)
		return err
	})
	loggerLynex.Infow("found pool",
		"market", market,
		"address", poolAddress,
		"is_stable", isStablePool)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("classic pool for market %s does not exist", market)
	}
	loggerLynex.Infow("pool found", "market", market, "address", poolAddress)

	poolContract, err := ilynex_pair.NewILynexPair(poolAddress, l.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Lynex pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(loggerLynex, func() error {
		basePoolToken, err = poolContract.Token0(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Lynex pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(loggerLynex, func() error {
		quotePoolToken, err = poolContract.Token1(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Lynex pool: %w", err)
	}

	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*dexPool[ilynex_pair.ILynexPairSwap]{{
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
	return nil, fmt.Errorf("failed to build Lynex pool for market %s: %w", market, err)
}

func (l *lynex) parseSwap(
	swap *ilynex_pair.ILynexPairSwap,
	pool *dexPool[ilynex_pair.ILynexPairSwap],
) (trade TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			loggerLynex.Errorw(ErrSwapParsing.Error(), "swap", swap, "pool", pool)
			err = fmt.Errorf("%s: %s", ErrSwapParsing, r)
		}
	}()

	return buildV2Trade(
		DriverLynex,
		swap.Amount0In,
		swap.Amount0Out,
		swap.Amount1In,
		swap.Amount1Out,
		pool,
	)
}
