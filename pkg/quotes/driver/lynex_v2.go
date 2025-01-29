package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_factory"
	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_pair"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var loggerLynexV2 = log.Logger("lynex_v2")

type lynexV2Event = ilynex_v2_pair.ILynexPairSwap
type lynexV2Iterator = *ilynex_v2_pair.ILynexPairSwapIterator

type lynexV2 struct {
	stablePoolMarkets map[quotes_common.Market]struct{}
	factory           *ilynex_v2_factory.ILynexFactory
	driver            base.DexReader
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

	hooks := lynexV2{
		stablePoolMarkets: stablePoolMarkets,
	}

	params := base.DexConfig[lynexV2Event, lynexV2Iterator]{
		Params: base.DexParams{
			Type:       quotes_common.DriverLynexV2,
			RPC:        rpcUrl,
			AssetsURL:  config.AssetsURL,
			MappingURL: config.MappingURL,
			MarketsURL: config.MarketsURL,
			IdlePeriod: config.IdlePeriod,
		},
		Hooks: base.DexHooks[lynexV2Event, lynexV2Iterator]{
			BuildPoolContracts: hooks.buildPoolContracts,
			BuildParser:        hooks.buildParser,
			DerefIter:          hooks.derefIter,
		},
		// State
		Outbox:  outbox,
		Logger:  loggerLynexV2,
		Filter:  config.Filter,
		History: history,
	}
	driver, err := base.NewDEX(params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Lynex v2 driver")
	}
	hooks.driver = driver // wire up the driver

	// Check addresses here: https://lynex.gitbook.io/lynex-docs/security/contracts
	factory, err := ilynex_v2_factory.NewILynexFactory(common.HexToAddress(config.FactoryAddress), driver.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate a Lynex v2 pool factory contract")
	}
	hooks.factory = factory

	return driver, nil
}

func (l *lynexV2) buildPoolContracts(ctx context.Context, market quotes_common.Market) ([]common.Address, []base.DexEventWatcher[lynexV2Event, lynexV2Iterator], error) {
	baseToken, quoteToken, err := base.GetTokens(l.driver.Assets(), market, loggerLynexV2)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	_, isStablePool := l.stablePoolMarkets[market]

	loggerLynexV2.Infow("searching for pool", "market", market)
	err = debounce.Debounce(ctx, loggerLynexV2, func(ctx context.Context) error {
		poolAddress, err = l.factory.GetPair(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address, isStablePool)
		return err
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerLynexV2.Infow("found pool",
		"market", market,
		"address", poolAddress,
		"is_stable", isStablePool)

	poolContract, err := ilynex_v2_pair.NewILynexPair(poolAddress, l.driver.Client())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build Lynex v2 pool contract: %w", err)
	}

	return []common.Address{poolAddress}, []base.DexEventWatcher[lynexV2Event, lynexV2Iterator]{poolContract}, nil
}

func (l *lynexV2) buildParser(
	swap *lynexV2Event,
	pool *base.DexPool[lynexV2Event, lynexV2Iterator],
) base.SwapParser {
	return &base.SwapV2[lynexV2Event, lynexV2Iterator]{
		RawAmount0In:  swap.Amount0In,
		RawAmount0Out: swap.Amount0Out,
		RawAmount1In:  swap.Amount1In,
		RawAmount1Out: swap.Amount1Out,
		Pool:          pool,
	}
}

func (l *lynexV2) derefIter(iter lynexV2Iterator) *ilynex_v2_pair.ILynexPairSwap {
	return iter.Event
}
