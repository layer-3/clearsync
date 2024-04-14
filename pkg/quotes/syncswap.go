package quotes

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/isyncswap_factory"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
)

var loggerSyncswap = log.Logger("syncswap")

type syncswap struct {
	base                      *baseDEX[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool]
	stablePoolMarkets         map[Market]struct{}
	classicPoolFactoryAddress common.Address
	stablePoolFactoryAddress  common.Address
	classicFactory            *isyncswap_factory.ISyncSwapFactory
	stableFactory             *isyncswap_factory.ISyncSwapFactory
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
		StartHook:   hooks.start,
		PoolGetter:  hooks.getPool,
		EventParser: hooks.parseSwap,
	}

	driver := newBaseDEX[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool](params)
	hooks.base = driver

	return driver
}

func (s *syncswap) start() (err error) {
	// Check addresses here: https://syncswap.gitbook.io/syncswap/smart-contracts/smart-contracts
	s.classicFactory, err = isyncswap_factory.NewISyncSwapFactory(s.classicPoolFactoryAddress, s.base.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap classic pool factory contract: %w", err)
	}

	s.stableFactory, err = isyncswap_factory.NewISyncSwapFactory(s.stablePoolFactoryAddress, s.base.Client())
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickswap stable pool factory contract: %w", err)
	}
	return nil
}

func (s *syncswap) getPool(market Market) (*dexPool[isyncswap_pool.ISyncSwapPoolSwap], error) {
	baseToken, quoteToken, err := getTokens(s.base.Assets(), market, loggerSyncswap)
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

	poolContract, err := isyncswap_pool.NewISyncSwapPool(poolAddress, s.base.Client())
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

	pool := &dexPool[isyncswap_pool.ISyncSwapPoolSwap]{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
		reverted:   false,
	}

	baseAddress := common.HexToAddress(baseToken.Address)
	quoteAddress := common.HexToAddress(quoteToken.Address)
	if baseAddress == basePoolToken && quoteAddress == quotePoolToken {
		return pool, nil
	} else if quoteAddress == basePoolToken && baseAddress == quotePoolToken {
		pool.reverted = true
		return pool, nil
	} else {
		return nil, fmt.Errorf("failed to build Syncswap pool: %w", err)
	}
}

func (*syncswap) parseSwap(
	swap *isyncswap_pool.ISyncSwapPoolSwap,
	pool *dexPool[isyncswap_pool.ISyncSwapPoolSwap],
) (TradeEvent, error) {
	var takerType TakerType
	var price decimal.Decimal
	var amount decimal.Decimal
	var total decimal.Decimal

	defer func() {
		if r := recover(); r != nil {
			loggerSyncswap.Errorw("recovered after parse swap panic", "swap", swap)
		}
	}()

	switch {
	case isValidNonZero(swap.Amount0In) && isValidNonZero(swap.Amount1Out):
		amount1Out := decimal.NewFromBigInt(swap.Amount1Out, 0).Div(decimal.NewFromInt(10).Pow(pool.quoteToken.Decimals))
		amount0In := decimal.NewFromBigInt(swap.Amount0In, 0).Div(decimal.NewFromInt(10).Pow(pool.baseToken.Decimals))

		takerType = TakerTypeSell
		price = amount1Out.Div(amount0In)
		total = amount1Out
		amount = amount0In

	case isValidNonZero(swap.Amount0Out) && isValidNonZero(swap.Amount1In):
		amount0Out := decimal.NewFromBigInt(swap.Amount0Out, 0).Div(decimal.NewFromInt(10).Pow(pool.baseToken.Decimals))
		amount1In := decimal.NewFromBigInt(swap.Amount1In, 0).Div(decimal.NewFromInt(10).Pow(pool.quoteToken.Decimals))

		takerType = TakerTypeBuy
		price = amount1In.Div(amount0Out)
		total = amount1In
		amount = amount0Out
	default:
		loggerSyncswap.Errorw("unknown swap type", "market", pool.Market())
		return TradeEvent{}, fmt.Errorf("market %s: unknown swap type", pool.Market())
	}

	return TradeEvent{
		Source:    DriverSyncswap,
		Market:    pool.Market(),
		Price:     price,
		Amount:    amount.Abs(),
		Total:     total,
		TakerType: takerType,
		CreatedAt: time.Now(),
	}, nil
}

func isValidNonZero(x *big.Int) bool {
	return x != nil && x.Sign() != 0
}
