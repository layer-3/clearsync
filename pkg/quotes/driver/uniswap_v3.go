package driver

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	quotes_common "github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/driver/base"
)

var (
	loggerUniswapV3 = log.Logger("uniswap_v3")
	// Uniswap v3 protocol has the 0.01%, 0.05%, 0.3%, and 1% fee tiers.
	uniswapV3FeeTiers = []uint{100, 500, 3000, 10000}
	wethContract      = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)

type uniswapV3Event = iuniswap_v3_pool.IUniswapV3PoolSwap
type uniswapV3Iterator = *iuniswap_v3_pool.IUniswapV3PoolSwapIterator

type uniswapV3 struct {
	factory *iuniswap_v3_factory.IUniswapV3Factory
	driver  base.DexReader
}

func newUniswapV3(rpcUrl string, config UniswapV3Config, outbox chan<- quotes_common.TradeEvent, history quotes_common.HistoricalDataDriver) (quotes_common.Driver, error) {
	var hooks uniswapV3
	params := base.DexConfig[uniswapV3Event, uniswapV3Iterator]{
		Params: base.DexParams{
			Type:       quotes_common.DriverUniswapV3,
			RPC:        rpcUrl,
			AssetsURL:  config.AssetsURL,
			MappingURL: config.MappingURL,
			MarketsURL: config.MarketsURL,
			IdlePeriod: config.IdlePeriod,
		},
		Hooks: base.DexHooks[uniswapV3Event, uniswapV3Iterator]{
			BuildPoolContracts: hooks.buildPoolContracts,
			BuildParser:        hooks.buildParser,
			DerefIter:          hooks.derefIter,
		},
		// State
		Outbox:  outbox,
		Logger:  loggerUniswapV3,
		Filter:  config.Filter,
		History: history,
	}
	driver, err := base.NewDEX(params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Uniswap v3 driver")
	}
	hooks.driver = driver // wire up the driver

	// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
	factory, err := iuniswap_v3_factory.NewIUniswapV3Factory(common.HexToAddress(config.FactoryAddress), driver.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate a Uniswap v3 pool factory contract")
	}
	hooks.factory = factory

	return driver, nil
}

func (u *uniswapV3) buildPoolContracts(ctx context.Context, market quotes_common.Market) ([]common.Address, []base.DexEventWatcher[uniswapV3Event, uniswapV3Iterator], error) {
	baseToken, quoteToken, err := base.GetTokens(u.driver.Assets(), market, loggerUniswapV3)
	if err != nil {
		return nil, nil, err
	}

	if strings.ToLower(baseToken.Symbol) == "eth" {
		baseToken.Address = wethContract
	}

	poolAddresses := make([]common.Address, 0, len(uniswapV3FeeTiers))
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range uniswapV3FeeTiers {
		var poolAddress common.Address
		err = debounce.Debounce(ctx, loggerUniswapV3, func(ctx context.Context) error {
			poolAddress, err = u.factory.GetPool(&bind.CallOpts{Context: ctx}, baseToken.Address, quoteToken.Address, big.NewInt(int64(feeTier)))
			return err
		})
		if err != nil {
			return nil, nil, err
		}

		if poolAddress != zeroAddress {
			loggerUniswapV3.Infow("found pool",
				"market", market,
				"selected fee tier", fmt.Sprintf("%.2f%%", float64(feeTier)/10000),
				"address", poolAddress)
			poolAddresses = append(poolAddresses, poolAddress)
		}
	}

	poolContracts := make([]base.DexEventWatcher[uniswapV3Event, uniswapV3Iterator], 0, len(poolAddresses))
	for _, poolAddress := range poolAddresses {
		poolContract, err := iuniswap_v3_pool.NewIUniswapV3Pool(poolAddress, u.driver.Client())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to build Uniswap v3 pool contract: %w", err)
		}
		poolContracts = append(poolContracts, poolContract)
	}

	return poolAddresses, poolContracts, nil
}

func (u *uniswapV3) buildParser(
	swap *uniswapV3Event,
	pool *base.DexPool[uniswapV3Event, uniswapV3Iterator],
) base.SwapParser {
	return &base.SwapV3[uniswapV3Event, uniswapV3Iterator]{
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.SqrtPriceX96,
		Pool:            pool,
	}
}

func (u *uniswapV3) derefIter(iter uniswapV3Iterator) *uniswapV3Event {
	return iter.Event
}
