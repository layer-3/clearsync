package driver

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"

	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_factory"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_pair"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var loggerSectaV2 = log.Logger("secta_v2")

type sectaV2Event = isecta_v2_pair.ISectaV2PairSwap
type sectaV2Iterator = *isecta_v2_pair.ISectaV2PairSwapIterator

type sectaV2 struct {
	factory *isecta_v2_factory.ISectaV2Factory
	driver  base.DexReader
}

func newSectaV2(rpcUrl string, config SectaV2Config, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	var hooks sectaV2
	params := base.DexConfig[sectaV2Event, sectaV2Iterator]{
		Params: base.DexParams{
			Type:       quotes_common.DriverSectaV2,
			RPC:        rpcUrl,
			AssetsURL:  config.AssetsURL,
			MappingURL: config.MappingURL,
			MarketsURL: config.MarketsURL,
			IdlePeriod: config.IdlePeriod},
		Hooks: base.DexHooks[sectaV2Event, sectaV2Iterator]{
			BuildPoolContracts: hooks.buildPoolContracts,
			BuildParser:        hooks.buildParser,
			DerefIter:          hooks.derefIter,
		},
		// State
		Outbox:  outbox,
		Logger:  loggerSectaV2,
		Filter:  config.Filter,
		History: history,
	}
	driver, err := base.NewDEX(params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Secta v2 driver")
	}
	hooks.driver = driver // wire up the driver

	factory, err := isecta_v2_factory.NewISectaV2Factory(common.HexToAddress(config.FactoryAddress), driver.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate a Secta v2 pool factory contract")
	}
	hooks.factory = factory

	return driver, nil
}

func (s *sectaV2) buildPoolContracts(ctx context.Context, market quotes_common.Market) ([]common.Address, []base.DexEventWatcher[sectaV2Event, sectaV2Iterator], error) {
	baseToken, quoteToken, err := base.GetTokens(s.driver.Assets(), market, loggerSectaV2)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	err = debounce.Debounce(ctx, loggerSectaV2, func(ctx context.Context) error {
		poolAddress, err = s.factory.GetPair(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address)
		return err
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get pool address: %w", err)
	}

	zeroAddress := common.HexToAddress("0x0")
	if poolAddress == zeroAddress {
		return nil, nil, fmt.Errorf("pool for market %s does not exist", market.StringWithoutMain())
	}
	loggerSectaV2.Infow("pool found", "market", market.StringWithoutMain(), "address", poolAddress)

	poolContract, err := isecta_v2_pair.NewISectaV2Pair(poolAddress, s.driver.Client())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to instantiate a Secta v2 pool contract: %w", err)
	}

	return []common.Address{poolAddress}, []base.DexEventWatcher[sectaV2Event, sectaV2Iterator]{poolContract}, nil
}

func (s *sectaV2) buildParser(
	swap *sectaV2Event,
	pool *base.DexPool[sectaV2Event, sectaV2Iterator],
) base.SwapParser {
	return &base.SwapV2[sectaV2Event, sectaV2Iterator]{
		RawAmount0In:  swap.Amount0In,
		RawAmount0Out: swap.Amount0Out,
		RawAmount1In:  swap.Amount1In,
		RawAmount1Out: swap.Amount1Out,
		Pool:          pool,
	}
}

func (s *sectaV2) derefIter(iter sectaV2Iterator) *sectaV2Event {
	return iter.Event
}
