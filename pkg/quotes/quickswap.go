package quotes

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/layer-3/clearsync/pkg/safe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/iquickswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iquickswap_v3_pool"
)

var loggerQuickswap = log.Logger("quickswap")

type quickswap struct {
	poolFactoryAddress common.Address
	factory            *iquickswap_v3_factory.IQuickswapV3Factory
	assets             *safe.Map[string, poolToken]
	client             *ethclient.Client
}

func newQuickswap(config QuickswapConfig, outbox chan<- TradeEvent) Driver {
	hooks := &quickswap{
		poolFactoryAddress: common.HexToAddress(config.PoolFactoryAddress),
	}

	params := baseDexConfig[iquickswap_v3_pool.IQuickswapV3PoolSwap, iquickswap_v3_pool.IQuickswapV3Pool]{
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

func (s *quickswap) postStart(driver *baseDEX[iquickswap_v3_pool.IQuickswapV3PoolSwap, iquickswap_v3_pool.IQuickswapV3Pool]) (err error) {
	s.client = driver.Client()
	s.assets = driver.Assets()

	// Check addresses here: https://quickswap.gitbook.io/quickswap/smart-contracts/smart-contracts
	s.factory, err = iquickswap_v3_factory.NewIQuickswapV3Factory(s.poolFactoryAddress, s.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Quickwap Factory contract: %w", err)
	}
	return nil
}

func (s *quickswap) getPool(market Market) ([]*dexPool[iquickswap_v3_pool.IQuickswapV3PoolSwap], error) {
	baseToken, quoteToken, err := getTokens(s.assets, market, loggerQuickswap)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens: %w", err)
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	poolAddress, err = s.factory.PoolByPair(
		nil,
		common.HexToAddress(baseToken.Address),
		common.HexToAddress(quoteToken.Address),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool address: %w", err)
	}
	if poolAddress == zeroAddress {
		return nil, fmt.Errorf("pool for market %s does not exist", market)
	}
	loggerQuickswap.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := iquickswap_v3_pool.NewIQuickswapV3Pool(poolAddress, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}

	basePoolToken, err := poolContract.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}

	quotePoolToken, err := poolContract.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build quickswap pool: %w", err)
	}

	baseAddress := common.HexToAddress(baseToken.Address)
	quoteAddress := common.HexToAddress(quoteToken.Address)
	isReverted := quoteAddress == basePoolToken && baseAddress == quotePoolToken
	pools := []*dexPool[iquickswap_v3_pool.IQuickswapV3PoolSwap]{{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
		reverted:   isReverted,
	}}

	// Return pools if the token addresses match direct or reversed configurations
	if (baseAddress == basePoolToken && quoteAddress == quotePoolToken) || isReverted {
		return pools, nil
	}
	return nil, fmt.Errorf("failed to build Quickswap pool for market %s: %w", market, err)
}

func (s *quickswap) parseSwap(
	swap *iquickswap_v3_pool.IQuickswapV3PoolSwap,
	pool *dexPool[iquickswap_v3_pool.IQuickswapV3PoolSwap],
) (TradeEvent, error) {
	if !isValidNonZero(swap.Amount0) || !isValidNonZero(swap.Amount1) {
		return TradeEvent{}, fmt.Errorf("either Amount0 (%s) or Amount1 (%s) is invalid", swap.Amount0, swap.Amount1)
	}

	defer func() {
		if r := recover(); r != nil {
			loggerQuickswap.Errorw("recovered in from panic during swap parsing", "swap", swap)
		}
	}()

	baseDecimals := pool.baseToken.Decimals
	quoteDecimals := pool.quoteToken.Decimals

	// Normalize swap amounts
	amount0 := decimal.NewFromBigInt(swap.Amount0, 0).Div(decimal.NewFromInt(10).Pow(baseDecimals))
	amount1 := decimal.NewFromBigInt(swap.Amount1, 0).Div(decimal.NewFromInt(10).Pow(quoteDecimals))

	// Assume it's a buy trade.
	// If it's not, price will equal to zero
	// and should be recalculated later.
	takerType := TakerTypeSell
	price := calculatePrice(
		decimal.NewFromBigInt(swap.Price, 0),
		baseDecimals,
		quoteDecimals)
	amount := amount0
	total := amount1

	if price.IsZero() { // then it's a buy trade
		takerType = TakerTypeBuy
		price = calculatePrice(
			decimal.NewFromBigInt(swap.Price, 0),
			quoteDecimals,
			baseDecimals)
		amount = amount1
		total = amount0
	}

	tr := TradeEvent{
		Source:    DriverQuickswap,
		Market:    pool.Market(),
		Price:     price,
		Amount:    amount.Abs(),
		Total:     total.Abs(),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return tr, nil
}

var (
	priceX96 = decimal.NewFromInt(2).Pow(decimal.NewFromInt(96))
	ten      = decimal.NewFromInt(10)
)

// calculatePrice method calculates the price per token at which the swap was performed
// using the sqrtPriceX96 value supplied with every on-chain swap event.
//
// General formula is as follows:
// price = ((sqrtPriceX96 / 2**96)**2) / (10**decimal1 / 10**decimal0)
//
// See the math explained at https://blog.uniswap.org/uniswap-v3-math-primer
func calculatePrice(
	sqrtPriceX96 decimal.Decimal,
	baseTokenDecimals decimal.Decimal,
	quoteTokenDecimals decimal.Decimal,
) decimal.Decimal {
	decimals := quoteTokenDecimals.Sub(baseTokenDecimals)

	numerator := sqrtPriceX96.Div(priceX96).Pow(decimal.NewFromInt(2))
	denominator := ten.Pow(decimals)
	return numerator.Div(denominator)
}
