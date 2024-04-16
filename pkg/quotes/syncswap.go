package quotes

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/layer-3/clearsync/pkg/safe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/isyncswap_factory"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
)

var loggerSyncswap = log.Logger("syncswap")

type syncswap struct {
	stablePoolMarkets         map[Market]struct{}
	classicPoolFactoryAddress common.Address
	stablePoolFactoryAddress  common.Address
	classicFactory            *isyncswap_factory.ISyncSwapFactory
	stableFactory             *isyncswap_factory.ISyncSwapFactory
	assets                    *safe.Map[string, poolToken]
	client                    *ethclient.Client
}

func newSyncswap(config SyncswapConfig, outbox chan<- TradeEvent) Driver {
	stablePoolMarkets := make(map[Market]struct{})
	for _, rawMarket := range config.StablePoolMarkets {
		market, ok := NewMarketFromString(rawMarket)
		if !ok {
			loggerSyncswap.Errorw("failed to parse stable pool from market", "market", rawMarket)
			continue
		}
		stablePoolMarkets[market] = struct{}{}
	}
	loggerSyncswap.Debugw("configured stable pool markets", "markets", stablePoolMarkets)

	hooks := &syncswap{
		stablePoolMarkets:         stablePoolMarkets,
		classicPoolFactoryAddress: common.HexToAddress(config.ClassicPoolFactoryAddress),
		stablePoolFactoryAddress:  common.HexToAddress(config.StablePoolFactoryAddress),
	}

	params := baseDexConfig[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool]{
		DriverType: DriverSyncswap,
		URL:        config.URL,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		Outbox:     outbox,
		Filter:     config.Filter,
		Logger:     loggerSyncswap,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
	}

	return newBaseDEX(params)
}

func (s *syncswap) postStart(driver *baseDEX[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool]) (err error) {
	s.client = driver.Client()
	s.assets = driver.Assets()

	// Check addresses here: https://syncswap.gitbook.io/syncswap/smart-contracts/smart-contracts
	s.classicFactory, err = isyncswap_factory.NewISyncSwapFactory(s.classicPoolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap classic pool factory contract: %w", err)
	}

	s.stableFactory, err = isyncswap_factory.NewISyncSwapFactory(s.stablePoolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap stable pool factory contract: %w", err)
	}
	return nil
}

func (s *syncswap) getPool(market Market) ([]*dexPool[isyncswap_pool.ISyncSwapPoolSwap], error) {
	baseToken, quoteToken, err := getTokens(s.assets, market, loggerSyncswap)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	if _, ok := s.stablePoolMarkets[market]; ok {
		loggerSyncswap.Infow("found stable pool", "market", market.StringWithoutMain())
		poolAddress, err = s.stableFactory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
		)
	} else {
		loggerSyncswap.Infow("found classic pool", "market", market.StringWithoutMain())
		poolAddress, err = s.classicFactory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get classic pool address: %w", err)
	}
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("classic pool for market %s does not exist", market.StringWithoutMain())
	}
	loggerSyncswap.Infow("pool found", "address", poolAddress, "market", market.StringWithoutMain())

	poolContract, err := isyncswap_pool.NewISyncSwapPool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}

	basePoolToken, err := poolContract.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}

	quotePoolToken, err := poolContract.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}

	baseAddress := common.HexToAddress(baseToken.Address)
	quoteAddress := common.HexToAddress(quoteToken.Address)
	isReverted := quoteAddress == basePoolToken && baseAddress == quotePoolToken
	pools := []*dexPool[isyncswap_pool.ISyncSwapPoolSwap]{{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
		reverted:   isReverted,
	}}

	// Return pools if the token addresses match direct or reversed configurations
	if (baseAddress == basePoolToken && quoteAddress == quotePoolToken) || isReverted {
		return pools, nil
	}
	return nil, fmt.Errorf("failed to build Syncswap pool for market %s: %w", market, err)
}

func (s *syncswap) parseSwap(
	swap *isyncswap_pool.ISyncSwapPoolSwap,
	pool *dexPool[isyncswap_pool.ISyncSwapPoolSwap],
) (TradeEvent, error) {
	var takerType TakerType
	var price decimal.Decimal
	var amount decimal.Decimal
	var total decimal.Decimal

	defer func() {
		if r := recover(); r != nil {
			loggerSyncswap.Errorw("recovered in from panic during swap parsing", "swap", swap)
		}
	}()

	if pool.reverted {
		s.flipSwap(swap)
	}

	baseDecimals := pool.baseToken.Decimals
	quoteDecimals := pool.quoteToken.Decimals

	switch {
	case isValidNonZero(swap.Amount0In) && isValidNonZero(swap.Amount1Out):
		amount1Out := decimal.NewFromBigInt(swap.Amount1Out, 0).Div(decimal.NewFromInt(10).Pow(quoteDecimals))
		amount0In := decimal.NewFromBigInt(swap.Amount0In, 0).Div(decimal.NewFromInt(10).Pow(baseDecimals))

		takerType = TakerTypeSell
		price = amount1Out.Div(amount0In)
		total = amount1Out
		amount = amount0In

	case isValidNonZero(swap.Amount0Out) && isValidNonZero(swap.Amount1In):
		amount0Out := decimal.NewFromBigInt(swap.Amount0Out, 0).Div(decimal.NewFromInt(10).Pow(baseDecimals))
		amount1In := decimal.NewFromBigInt(swap.Amount1In, 0).Div(decimal.NewFromInt(10).Pow(quoteDecimals))

		takerType = TakerTypeBuy
		price = amount1In.Div(amount0Out)
		total = amount1In
		amount = amount0Out
	default:
		loggerSyncswap.Errorw("unknown swap type", "market", pool.Market())
		return TradeEvent{}, fmt.Errorf("market %s: unknown swap type", pool.Market())
	}

	trade := TradeEvent{
		Source:    DriverSyncswap,
		Market:    pool.Market(),
		Price:     price,
		Amount:    amount.Abs(),
		Total:     total,
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return trade, nil
}

func (*syncswap) flipSwap(swap *isyncswap_pool.ISyncSwapPoolSwap) {
	amount0In, amount0Out := swap.Amount0In, swap.Amount0Out
	swap.Amount0In, swap.Amount0Out = swap.Amount1In, swap.Amount1Out
	swap.Amount1In, swap.Amount1Out = amount0In, amount0Out
}

func isValidNonZero(x *big.Int) bool {
	return x != nil && x.Sign() != 0
}
