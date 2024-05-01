package quotes

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_pool"
	"github.com/layer-3/clearsync/pkg/safe"
)

var (
	loggerSectaV3 = log.Logger("secta_v3")
	// Secta v3 uses 0.01%, 0.05%, 0.25%, and 1% fee tiers.
	sectaV3FeeTiers = []uint{100, 500, 2500, 10000}
)

type sectaV3 struct {
	factoryAddress common.Address
	factory        *isecta_v3_factory.ISectaV3Factory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newSectaV3(config SectaV3Config, outbox chan<- TradeEvent) Driver {
	hooks := &sectaV3{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := baseDexConfig[isecta_v3_pool.ISectaV3PoolSwap, isecta_v3_pool.ISectaV3Pool]{
		DriverType:    DriverSectaV3,
		URL:           config.URL,
		AssetsURL:     config.AssetsURL,
		MappingURL:    config.MappingURL,
		Logger:        loggerSectaV3,
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
		Outbox:        outbox,
		Filter:        config.Filter,
	}
	return newBaseDEX(params)
}

func (s *sectaV3) postStart(driver *baseDEX[isecta_v3_pool.ISectaV3PoolSwap, isecta_v3_pool.ISectaV3Pool]) (err error) {
	s.client = driver.Client()
	s.assets = driver.Assets()

	s.factory, err = isecta_v3_factory.NewISectaV3Factory(s.factoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Secta v3 pool factory contract: %w", err)
	}
	return nil
}

func (s *sectaV3) getPool(market Market) ([]*dexPool[isecta_v3_pool.ISectaV3PoolSwap], error) {
	baseToken, quoteToken, err := getTokens(s.assets, market, loggerSectaV3)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	if strings.ToLower(baseToken.Symbol) == "eth" {
		baseToken.Address = wethContract
	}

	poolAddresses := make([]common.Address, 0, len(sectaV3FeeTiers))
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range sectaV3FeeTiers {
		poolAddress, err := s.factory.GetPool(nil, baseToken.Address, quoteToken.Address, big.NewInt(int64(feeTier)))
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

	pools := make([]*dexPool[isecta_v3_pool.ISectaV3PoolSwap], 0, len(poolAddresses))
	for _, poolAddress := range poolAddresses {
		poolContract, err := isecta_v3_pool.NewISectaV3Pool(poolAddress, s.client)
		if err != nil {
			return nil, fmt.Errorf("failed to build Secta v3 pool: %w", err)
		}

		basePoolToken, err := poolContract.Token0(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to build Secta v3 pool: %w", err)
		}

		quotePoolToken, err := poolContract.Token1(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to build Secta v3 pool: %w", err)
		}

		isReverted := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
		pool := &dexPool[isecta_v3_pool.ISectaV3PoolSwap]{
			contract:   poolContract,
			baseToken:  baseToken,
			quoteToken: quoteToken,
			reverted:   isReverted,
			market:     market,
		}

		// Append pool if the token addresses match direct or reversed configurations
		if (baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken) || isReverted {
			pools = append(pools, pool)
		}
	}

	return pools, nil
}

func (s *sectaV3) parseSwap(
	swap *isecta_v3_pool.ISectaV3PoolSwap,
	pool *dexPool[isecta_v3_pool.ISectaV3PoolSwap],
) (trade TradeEvent, err error) {
	opts := v3TradeOpts[isecta_v3_pool.ISectaV3PoolSwap]{
		driver:          DriverSectaV3,
		rawAmount0:      swap.Amount0,
		rawAmount1:      swap.Amount1,
		rawSqrtPriceX96: swap.SqrtPriceX96,
		pool:            pool,
		swap:            swap,
		logger:          loggerSectaV3,
	}
	return buildV3Trade(opts)
}
