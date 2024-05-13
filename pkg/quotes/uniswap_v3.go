package quotes

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/debounce"
	"github.com/layer-3/clearsync/pkg/safe"
)

var (
	loggerUniswapV3 = log.Logger("uniswap_v3")
	// Uniswap v3 protocol has the 0.01%, 0.05%, 0.3%, and 1% fee tiers.
	uniswapV3FeeTiers = []uint{100, 500, 3000, 10000}
	wethContract      = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)

type uniswapV3 struct {
	factoryAddress common.Address
	factory        *iuniswap_v3_factory.IUniswapV3Factory

	assets *safe.Map[string, poolToken]
	client *ethclient.Client
}

func newUniswapV3(config UniswapV3Config, outbox chan<- TradeEvent) Driver {
	hooks := &uniswapV3{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := baseDexConfig[
		iuniswap_v3_pool.IUniswapV3PoolSwap,
		iuniswap_v3_pool.IUniswapV3Pool,
	]{
		// Params
		DriverType: DriverUniswapV3,
		URL:        config.URL,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		IdlePeriod: config.IdlePeriod,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
		// State
		Outbox: outbox,
		Logger: loggerUniswapV3,
		Filter: config.Filter,
	}
	return newBaseDEX(params)
}

func (u *uniswapV3) postStart(driver *baseDEX[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool]) (err error) {
	u.client = driver.Client()
	u.assets = driver.Assets()

	// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
	u.factory, err = iuniswap_v3_factory.NewIUniswapV3Factory(u.factoryAddress, u.client)
	if err != nil {
		return fmt.Errorf("failed to build Uniswap v3 factory: %w", err)
	}
	return nil
}

func (u *uniswapV3) getPool(market Market) ([]*dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap], error) {
	baseToken, quoteToken, err := getTokens(u.assets, market, loggerUniswapV3)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(baseToken.Symbol) == "eth" {
		baseToken.Address = wethContract
	}

	poolAddresses := make([]common.Address, 0, len(uniswapV3FeeTiers))
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range uniswapV3FeeTiers {
		var poolAddress common.Address
		err = debounce.Debounce(loggerUniswapV3, func() error {
			poolAddress, err = u.factory.GetPool(nil, baseToken.Address, quoteToken.Address, big.NewInt(int64(feeTier)))
			return err
		})
		if err != nil {
			return nil, err
		}

		if poolAddress != zeroAddress {
			loggerUniswapV3.Infow("found pool",
				"market", market,
				"selected fee tier", fmt.Sprintf("%.2f%%", float64(feeTier)/10000),
				"address", poolAddress)
			poolAddresses = append(poolAddresses, poolAddress)
		}
	}

	pools := make([]*dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap], 0, len(poolAddresses))
	for _, poolAddress := range poolAddresses {
		poolContract, err := iuniswap_v3_pool.NewIUniswapV3Pool(poolAddress, u.client)
		if err != nil {
			return nil, fmt.Errorf("failed to build Uniswap v3 pool contract: %w", err)
		}

		var basePoolToken common.Address
		err = debounce.Debounce(loggerUniswapV3, func() error {
			basePoolToken, err = poolContract.Token0(nil)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get base token address for Uniswap v3 pool: %w", err)
		}

		var quotePoolToken common.Address
		err = debounce.Debounce(loggerUniswapV3, func() error {
			quotePoolToken, err = poolContract.Token1(nil)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get quote token address for Uniswap v3 pool: %w", err)
		}

		isReversed := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
		pool := &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]{
			Contract:   poolContract,
			Address:    poolAddress,
			BaseToken:  baseToken,
			QuoteToken: quoteToken,
			Market:     market,
			Reversed:   isReversed,
		}

		// Append pool if the token addresses match direct or reversed configurations
		if (baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken) || isReversed {
			pools = append(pools, pool)
		}
	}

	return pools, nil
}

func (u *uniswapV3) parseSwap(
	swap *iuniswap_v3_pool.IUniswapV3PoolSwap,
	pool *dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap],
) (trade TradeEvent, err error) {
	opts := v3TradeOpts[iuniswap_v3_pool.IUniswapV3PoolSwap]{
		Driver:          DriverUniswapV3,
		RawAmount0:      swap.Amount0,
		RawAmount1:      swap.Amount1,
		RawSqrtPriceX96: swap.SqrtPriceX96,
		Pool:            pool,
		Swap:            swap,
		Logger:          loggerUniswapV3,
	}
	return buildV3Trade(opts)
}

type v3TradeOpts[Event any] struct {
	Driver          DriverType
	RawAmount0      *big.Int
	RawAmount1      *big.Int
	RawSqrtPriceX96 *big.Int
	Pool            *dexPool[Event]
	Swap            *Event
	Logger          *log.ZapEventLogger
}

func buildV3Trade[Event any](o v3TradeOpts[Event]) (trade TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			o.Logger.Errorw(ErrSwapParsing.Error(), "swap", o.Swap, "pool", o.Pool)
			err = fmt.Errorf("%s: %s", ErrSwapParsing, r)
		}
	}()

	if !isValidNonZero(o.RawAmount0) {
		return TradeEvent{}, fmt.Errorf("raw amount0 (%s) is not a valid non-zero number", o.RawAmount0)
	}
	amount0 := decimal.NewFromBigInt(o.RawAmount0, 0)

	if !isValidNonZero(o.RawAmount1) {
		return TradeEvent{}, fmt.Errorf("raw amount1 (%s) is not a valid non-zero number", o.RawAmount0)
	}
	amount1 := decimal.NewFromBigInt(o.RawAmount1, 0)

	if !isValidNonZero(o.RawSqrtPriceX96) {
		return TradeEvent{}, fmt.Errorf("raw sqrtPriceX96 (%s) is not a valid non-zero number", o.RawSqrtPriceX96)
	}
	sqrtPriceX96 := decimal.NewFromBigInt(o.RawSqrtPriceX96, 0)

	if o.Pool.Reversed {
		amount0, amount1 = amount1, amount0
	}

	// Normalize swap amounts.
	baseDecimals, quoteDecimals := o.Pool.BaseToken.Decimals, o.Pool.QuoteToken.Decimals
	amount0Normalized := amount0.Div(ten.Pow(baseDecimals)).Abs()
	amount1Normalized := amount1.Div(ten.Pow(quoteDecimals)).Abs()

	// Calculate swap price.
	price := calculatePrice(sqrtPriceX96, baseDecimals, quoteDecimals, o.Pool.Reversed)
	// Apply a fallback strategy in case the primary one fails.
	// This should never happen, but just in case.
	if price.IsZero() {
		price = amount1Normalized.Div(amount0Normalized)
	}

	// Calculate trade side, amount and total.
	takerType := TakerTypeBuy
	amount, total := amount0Normalized, amount1Normalized
	if (!o.Pool.Reversed && amount0.Sign() < 0) || (o.Pool.Reversed && amount1.Sign() < 0) {
		takerType = TakerTypeSell
	}

	tr := TradeEvent{
		Source:    o.Driver,
		Market:    o.Pool.Market,
		Price:     price,
		Amount:    amount, // amount of BASE token received
		Total:     total,  // total cost in QUOTE token
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return tr, nil
}

var (
	two      = decimal.NewFromInt(2)
	ten      = decimal.NewFromInt(10)
	priceX96 = two.Pow(decimal.NewFromInt(96))
)

// calculatePrice method calculates the price per token at which the swap was performed
// using the sqrtPriceX96 value supplied with every on-chain swap event.
//
// General formula is as follows:
// price = ((sqrtPriceX96 / 2**96)**2) / (10**decimal1 / 10**decimal0)
//
// See the math explained at https://blog.uniswap.org/uniswap-v3-math-primer
func calculatePrice(sqrtPriceX96, baseDecimals, quoteDecimals decimal.Decimal, reversedPool bool) decimal.Decimal {
	if reversedPool {
		baseDecimals, quoteDecimals = quoteDecimals, baseDecimals
	}

	// Simplification for denominator calculations:
	// 10**decimal1 / 10**decimal0 -> 10**(decimal1 - decimal0)
	decimals := quoteDecimals.Sub(baseDecimals)

	numerator := sqrtPriceX96.Div(priceX96).Pow(two)
	denominator := ten.Pow(decimals)

	if reversedPool {
		return denominator.Div(numerator)
	}
	return numerator.Div(denominator)
}
