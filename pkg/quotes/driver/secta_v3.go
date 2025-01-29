package driver

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var (
	loggerSectaV3 = log.Logger("secta_v3")
	// Secta v3 uses 0.01%, 0.05%, 0.25%, and 1% fee tiers.
	sectaV3FeeTiers = []uint{100, 500, 2500, 10000}
)

type sectaV3Event = isecta_v3_pool.ISectaV3PoolSwap
type sectaV3Iterator = *isecta_v3_pool.ISectaV3PoolSwapIterator

type sectaV3 struct {
	factoryAddress common.Address
	factory        *isecta_v3_factory.ISectaV3Factory

	driver base.DexReader
}

func newSectaV3(rpcUrl string, config SectaV3Config, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	hooks := &sectaV3{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := base.DexConfig[sectaV3Event, sectaV3Iterator]{
		Params: base.DexParams{
			Type:       quotes_common.DriverSectaV3,
			RPC:        rpcUrl,
			AssetsURL:  config.AssetsURL,
			MappingURL: config.MappingURL,
			MarketsURL: config.MarketsURL,
			IdlePeriod: config.IdlePeriod},
		Hooks: base.DexHooks[sectaV3Event, sectaV3Iterator]{
			PostStart:   hooks.postStart,
			GetPool:     hooks.getPool,
			BuildParser: hooks.buildParser,
			DerefIter:   hooks.derefIter,
		},
		// State
		Outbox:  outbox,
		Logger:  loggerSectaV3,
		Filter:  config.Filter,
		History: history,
	}
	return base.NewDEX(params)
}

func (s *sectaV3) postStart(driver *base.DEX[sectaV3Event, sectaV3Iterator]) (err error) {
	s.driver = driver

	s.factory, err = isecta_v3_factory.NewISectaV3Factory(s.factoryAddress, s.driver.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Secta v3 pool factory contract: %w", err)
	}
	return nil
}

func (s *sectaV3) getPool(ctx context.Context, market quotes_common.Market) ([]*base.DexPool[sectaV3Event, sectaV3Iterator], error) {
	baseToken, quoteToken, err := base.GetTokens(s.driver.Assets(), market, loggerSectaV3)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	if strings.ToLower(baseToken.Symbol) == "eth" {
		baseToken.Address = wethContract
	}

	poolAddresses := make([]common.Address, 0, len(sectaV3FeeTiers))
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range sectaV3FeeTiers {
		var poolAddress common.Address
		err = debounce.Debounce(ctx, loggerSectaV3, func(ctx context.Context) error {
			poolAddress, err = s.factory.GetPool(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address, big.NewInt(int64(feeTier)))
			return err
		})
		if err != nil {
			return nil, err
		}

		if poolAddress != zeroAddress {
			loggerSectaV3.Infow("found pool",
				"market", market,
				"selected fee tier", fmt.Sprintf("%.2f%%", float64(feeTier)/10000),
				"address", poolAddress)
			poolAddresses = append(poolAddresses, poolAddress)
		}
	}

	pools := make([]*base.DexPool[sectaV3Event, sectaV3Iterator], 0, len(poolAddresses))
	for _, poolAddress := range poolAddresses {
		poolContract, err := isecta_v3_pool.NewISectaV3Pool(poolAddress, s.driver.Client())
		if err != nil {
			return nil, fmt.Errorf("failed to build Secta v3 pool contract: %w", err)
		}

		var basePoolToken common.Address
		err = debounce.Debounce(ctx, loggerSectaV3, func(ctx context.Context) error {
			basePoolToken, err = poolContract.Token0(&bind.CallOpts{Context: ctx})
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get base token address for Secta v3 pool: %w", err)
		}

		var quotePoolToken common.Address
		err = debounce.Debounce(ctx, loggerSectaV3, func(ctx context.Context) error {
			quotePoolToken, err = poolContract.Token1(&bind.CallOpts{Context: ctx})
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get quote token address for Secta v3 pool: %w", err)
		}

		isDirect := baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken
		isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
		pool := &base.DexPool[sectaV3Event, sectaV3Iterator]{
			Contract:   poolContract,
			Address:    poolAddress,
			BaseToken:  baseToken,
			QuoteToken: quoteToken,
			Market:     market,
			Reversed:   isReversed,
		}

		// Append pool only if the token addresses
		// match direct or reversed configurations
		if isDirect || isReversed {
			pools = append(pools, pool)
		}
	}

	return pools, nil
}

func (s *sectaV3) buildParser(
	swap *sectaV3Event,
	pool *base.DexPool[sectaV3Event, sectaV3Iterator],
) base.SwapParser {
	return &base.SwapV3[sectaV3Event, sectaV3Iterator]{
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.SqrtPriceX96,
		Pool:            pool,
	}
}

func (s *sectaV3) derefIter(iter sectaV3Iterator) *sectaV3Event {
	return iter.Event
}
