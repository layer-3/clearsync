package quotes

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerQuickswap = log.Logger("quickswap")

type quickswap struct {
	poolFactoryAddress common.Address
	factory            *quickswap_v3_factory.IQuickswapV3Factory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newQuickswap(config QuickswapConfig, outbox chan<- TradeEvent) Driver {
	hooks := &quickswap{
		poolFactoryAddress: common.HexToAddress(config.PoolFactoryAddress),
	}

	params := baseDexConfig[quickswap_v3_pool.IQuickswapV3PoolSwap, quickswap_v3_pool.IQuickswapV3Pool]{
		DriverType: DriverQuickswap,
		URL:        config.URL,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		Outbox:     outbox,
		Filter:     config.Filter,
		Logger:     loggerQuickswap,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
	}
	return newBaseDEX(params)
}

func (s *quickswap) postStart(driver *baseDEX[quickswap_v3_pool.IQuickswapV3PoolSwap, quickswap_v3_pool.IQuickswapV3Pool]) (err error) {
	s.client = driver.Client()
	s.assets = driver.Assets()

	// Check addresses here: https://quickswap.gitbook.io/quickswap/smart-contracts/smart-contracts
	s.factory, err = quickswap_v3_factory.NewIQuickswapV3Factory(s.poolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickwap Factory contract: %w", err)
	}
	return nil
}

func (s *quickswap) getPool(market Market) ([]*dexPool[quickswap_v3_pool.IQuickswapV3PoolSwap], error) {
	baseToken, quoteToken, err := getTokens(s.assets, market, loggerQuickswap)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	poolAddress, err = s.factory.PoolByPair(nil, baseToken.Address, quoteToken.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerQuickswap.Infow("found pool", "market", market, "address", poolAddress)

	poolContract, err := quickswap_v3_pool.NewIQuickswapV3Pool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Quickswap pool: %w", err)
	}

	basePoolToken, err := poolContract.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build Quickswap pool: %w", err)
	}

	quotePoolToken, err := poolContract.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build Quickswap pool: %w", err)
	}

	isReverted := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
	pools := []*dexPool[quickswap_v3_pool.IQuickswapV3PoolSwap]{{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
		reverted:   isReverted,
	}}

	// Return pools if the token addresses match direct or reversed configurations
	if (baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken) || isReverted {
		return pools, nil
	}
	return nil, fmt.Errorf("failed to build Quickswap pool for market %s: %w", market, err)
}

func (s *quickswap) parseSwap(
	swap *quickswap_v3_pool.IQuickswapV3PoolSwap,
	pool *dexPool[quickswap_v3_pool.IQuickswapV3PoolSwap],
) (trade TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := "recovered in from panic during swap parsing"
			loggerQuickswap.Errorw(msg, "swap", swap)
			err = fmt.Errorf("%s: %s", msg, r)
		}
	}()

	return buildV3Trade(
		DriverQuickswap,
		swap.Amount0,
		swap.Amount1,
		pool,
	)
}
