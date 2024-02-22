package quotes

import "github.com/shopspring/decimal"

var (
	defaultWeightsMap = map[DriverType]decimal.Decimal{
		DriverKraken:        decimal.NewFromInt(15),
		DriverBinance:       decimal.NewFromInt(20),
		DriverUniswapV3Api:  decimal.NewFromInt(50),
		DriverUniswapV3Geth: decimal.NewFromInt(50),
		DriverSyncswap:      decimal.NewFromInt(50),
	}
)

func defaultIndexDrivers(config Config) []DriverConfig {
	return []DriverConfig{config.Kraken, config.Binance, config.UniswapV3Api, config.Syncswap}
}
