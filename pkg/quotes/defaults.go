package quotes

import "github.com/shopspring/decimal"

var (
	DefaultWeightsMap = map[DriverType]decimal.Decimal{
		DriverKraken:        decimal.NewFromInt(15),
		DriverBinance:       decimal.NewFromInt(20),
		DriverUniswapV3Api:  decimal.NewFromInt(50),
		DriverUniswapV3Geth: decimal.NewFromInt(50),
		DriverSyncswap:      decimal.NewFromInt(50),
		DriverQuickswap:     decimal.NewFromInt(50),
	}
)

var (
	DefaultMarketsMapping = map[string][]string{"usdc": {"eth", "weth", "matic"}}
)
