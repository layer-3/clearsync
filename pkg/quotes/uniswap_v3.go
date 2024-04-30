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

	params := baseDexConfig[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool]{
		DriverType: DriverUniswapV3,
		URL:        config.URL,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		Outbox:     outbox,
		Filter:     config.Filter,
		Logger:     loggerUniswapV3,
		// Hooks
		PostStartHook: hooks.postStart,
		PoolGetter:    hooks.getPool,
		EventParser:   hooks.parseSwap,
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
		poolAddress, err := u.factory.GetPool(nil, baseToken.Address, quoteToken.Address, big.NewInt(int64(feeTier)))
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
			return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
		}

		basePoolToken, err := poolContract.Token0(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
		}

		quotePoolToken, err := poolContract.Token1(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
		}

		isReverted := quoteToken.Address == basePoolToken && baseToken.Address == quotePoolToken
		pool := &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]{
			contract:   poolContract,
			baseToken:  baseToken,
			quoteToken: quoteToken,
			reverted:   isReverted,
		}

		// Append pool if the token addresses match direct or reversed configurations
		if (baseToken.Address == basePoolToken && quoteToken.Address == quotePoolToken) || isReverted {
			pools = append(pools, pool)
		}
	}

	return pools, nil
}

func (u *uniswapV3) parseSwap(
	swap *iuniswap_v3_pool.IUniswapV3PoolSwap,
	pool *dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap],
) (trade TradeEvent, err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := "recovered in from panic during swap parsing"
			loggerUniswapV3.Errorw(msg, "swap", swap)
			err = fmt.Errorf("%s: %s", msg, r)
		}
	}()

	return buildV3Trade(
		DriverUniswapV3,
		swap.Amount0,
		swap.Amount1,
		pool,
	)
}

func buildV3Trade[Event any](
	driver DriverType,
	rawAmount0, rawAmount1 *big.Int,
	pool *dexPool[Event],
) (TradeEvent, error) {
	if !isValidNonZero(rawAmount0) || !isValidNonZero(rawAmount1) {
		return TradeEvent{}, fmt.Errorf("either Amount0 (%s) or Amount1 (%s) is invalid", rawAmount0, rawAmount1)
	}

	if pool.reverted {
		// For USDC/ETH:
		// Amount0 = -52052662345          = USDC removed from pool
		// Amount1 = +16867051239984403529 = ETH added into pool
		rawAmount0, rawAmount1 = rawAmount1, rawAmount0
	}

	// Normalize swap amounts
	ten := decimal.NewFromInt(10)
	baseDecimals := pool.baseToken.Decimals
	quoteDecimals := pool.quoteToken.Decimals
	amount0 := decimal.NewFromBigInt(rawAmount0, 0).Div(ten.Pow(baseDecimals))
	amount1 := decimal.NewFromBigInt(rawAmount1, 0).Div(ten.Pow(quoteDecimals))

	// Calculate price and order side
	price := amount1.Div(amount0) // NOTE: may panic here if `amount0` is zero
	amount := amount0
	takerType := TakerTypeBuy
	if amount0.Sign() < 0 {
		takerType = TakerTypeSell
	}
	amount = amount.Abs()
	price = price.Abs()

	if amount.Equals(decimal.Zero) {
		return TradeEvent{}, fmt.Errorf("amount is zero")
	}

	tr := TradeEvent{
		Source:    driver,
		Market:    pool.market,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return tr, nil
}
