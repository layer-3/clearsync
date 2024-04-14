package quotes

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_factory"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
)

var (
	loggerUniswapV3Geth = log.Logger("uniswap_v3_geth")
	// Uniswap v3 protocol has the 0.01%, 0.05%, 0.3%, and 1% fee tiers.
	uniswapV3FeeTiers = []uint{100, 500, 3000, 10000}
)

type uniswapV3Geth struct {
	base *baseDEX[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool]

	factoryAddress common.Address
	factory        *iuniswap_v3_factory.IUniswapV3Factory
}

func newUniswapV3Geth(config UniswapV3GethConfig, outbox chan<- TradeEvent) Driver {
	hooks := &uniswapV3Geth{
		factoryAddress: common.HexToAddress(config.FactoryAddress),
	}

	params := baseDexConfig[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool]{
		DriverType: DriverUniswapV3Geth,
		URL:        config.URL,
		AssetsURL:  config.AssetsURL,
		MappingURL: config.MappingURL,
		Outbox:     outbox,
		Filter:     config.Filter,
		Logger:     loggerUniswapV3Geth,
		// Hooks
		StartHook:   hooks.start,
		PoolGetter:  hooks.getPool,
		EventParser: hooks.parseSwap,
	}

	driver := newBaseDEX[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool](params)
	hooks.base = driver

	return driver
}

func (u *uniswapV3Geth) start() (err error) {
	// Check addresses here: https://docs.uniswap.org/contracts/v3/reference/deployments
	u.factory, err = iuniswap_v3_factory.NewIUniswapV3Factory(u.factoryAddress, u.base.Client())
	if err != nil {
		return fmt.Errorf("failed to build Uniswap v3 factory: %w", err)
	}
	return nil
}

func (u *uniswapV3Geth) getPool(market Market) (*dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap], error) {
	baseToken, quoteToken, err := getTokens(u.base.Assets(), market, loggerUniswapV3Geth)
	if err != nil {
		return nil, err
	}

	var poolAddress common.Address
	zeroAddress := common.HexToAddress("0x0")
	for _, feeTier := range uniswapV3FeeTiers {
		poolAddress, err = u.factory.GetPool(
			nil,
			common.HexToAddress(baseToken.Address),
			common.HexToAddress(quoteToken.Address),
			big.NewInt(int64(feeTier)),
		)
		if err != nil {
			return nil, err
		}
		if poolAddress != zeroAddress {
			loggerUniswapV3Geth.Infof("market %s: selected fee tier: %.2f%%", market, float64(feeTier)/10000)
			break
		}
	}
	loggerUniswapV3Geth.Infof("got pool %s for market %s", poolAddress, market)

	poolContract, err := iuniswap_v3_pool.NewIUniswapV3Pool(poolAddress, u.base.Client())
	if err != nil {
		return nil, fmt.Errorf("failed to build Uniswap v3 pool: %w", err)
	}

	return &dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]{
		contract:   poolContract,
		baseToken:  baseToken,
		quoteToken: quoteToken,
	}, nil
}

func (u *uniswapV3Geth) parseSwap(swap *iuniswap_v3_pool.IUniswapV3PoolSwap, pool *dexPool[iuniswap_v3_pool.IUniswapV3PoolSwap]) (TradeEvent, error) {
	amount := decimal.NewFromBigInt(swap.Amount0, 0)
	price := calculatePrice(
		decimal.NewFromBigInt(swap.SqrtPriceX96, 0),
		pool.baseToken.Decimals,
		pool.quoteToken.Decimals)
	takerType := TakerTypeBuy
	if amount.Sign() < 0 {

		takerType = TakerTypeSell
	}

	amount = amount.Abs()
	tr := TradeEvent{
		Source:    DriverUniswapV3Geth,
		Market:    pool.Market(),
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}
	return tr, nil
}
