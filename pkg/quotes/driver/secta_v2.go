package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_factory"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_pair"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var loggerSectaV2 = log.Logger("secta_v2")

type sectaV2 struct {
	factoryAddress common.Address
	factory        *isecta_v2_factory.ISectaV2Factory

	driver *base.DEX[
		isecta_v2_pair.ISectaV2PairSwap,
		isecta_v2_pair.ISectaV2Pair,
		*isecta_v2_pair.ISectaV2PairSwapIterator,
	]
}

func newSectaV2(rpcUrl string, config SectaV2Config, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	hooks := &sectaV2{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := base.DexConfig[
		isecta_v2_pair.ISectaV2PairSwap,
		isecta_v2_pair.ISectaV2Pair,
		*isecta_v2_pair.ISectaV2PairSwapIterator,
	]{
		// Params
		DriverType: quotes_common.DriverSectaV2,
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
		Logger:  loggerSectaV2,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (s *sectaV2) postStart(driver *base.DEX[
	isecta_v2_pair.ISectaV2PairSwap,
	isecta_v2_pair.ISectaV2Pair,
	*isecta_v2_pair.ISectaV2PairSwapIterator,
]) (err error) {
	s.driver = driver

	s.factory, err = isecta_v2_factory.NewISectaV2Factory(s.factoryAddress, s.driver.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Secta v2 pool factory contract: %w", err)
	}
	return nil
}

func (s *sectaV2) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[isecta_v2_pair.ISectaV2PairSwap, *isecta_v2_pair.ISectaV2PairSwapIterator], error) {
	baseToken, quoteToken, err := base.GetTokens(s.driver.Assets(), market, loggerSectaV2)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	err = debounce.Debounce(ctx, loggerSectaV2, func(ctx context.Context) error {
		poolAddress, err = s.factory.GetPair(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market.StringWithoutMain())
	}
	loggerSectaV2.Infow("pool found", "market", market.StringWithoutMain(), "address", poolAddress)

	poolContract, err := isecta_v2_pair.NewISectaV2Pair(poolAddress, s.driver.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a Secta v2 pool contract: %w", err)
	}

	var basePoolToken common.Address
	err = debounce.Debounce(ctx, loggerSectaV2, func(ctx context.Context) error {
		basePoolToken, err = poolContract.Token0(&bind.CallOpts{Context: ctx})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get base token address for Secta v2 pool: %w", err)
	}

	var quotePoolToken common.Address
	err = debounce.Debounce(ctx, loggerSectaV2, func(ctx context.Context) error {
		quotePoolToken, err = poolContract.Token1(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote token address for Secta v2 pool: %w", err)
	}

	isDirect := baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken
	isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*base.DexPool[isecta_v2_pair.ISectaV2PairSwap, *isecta_v2_pair.ISectaV2PairSwapIterator]{{
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
	return nil, fmt.Errorf("failed to build Secta v2 pool for market %s: %w", market, err)
}

func (s *sectaV2) buildParser(
	swap *isecta_v2_pair.ISectaV2PairSwap,
	pool *base.DexPool[isecta_v2_pair.ISectaV2PairSwap, *isecta_v2_pair.ISectaV2PairSwapIterator],
) base.SwapParser {
	return &base.SwapV2[
		isecta_v2_pair.ISectaV2PairSwap,
		*isecta_v2_pair.ISectaV2PairSwapIterator,
	]{
		RawAmount0In:  swap.Amount0In,
		RawAmount0Out: swap.Amount0Out,
		RawAmount1In:  swap.Amount1In,
		RawAmount1Out: swap.Amount1Out,
		Pool:          pool,
	}
}

func (s *sectaV2) derefIter(
	iter *isecta_v2_pair.ISectaV2PairSwapIterator,
) *isecta_v2_pair.ISectaV2PairSwap {
	return iter.Event
}
