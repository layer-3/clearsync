package quotes

import "github.com/shopspring/decimal"

var (
	DefaultWeightsMap = map[DriverType]decimal.Decimal{
		DriverKraken:        decimal.NewFromInt(15),
		DriverBinance:       decimal.NewFromInt(20),
		DriverUniswapV3Api:  decimal.NewFromInt(50),
		DriverUniswapV3Geth: decimal.NewFromInt(50),
		DriverSyncswap:      decimal.NewFromInt(50),
	}

	AllDrivers = []DriverConfig{
		KrakenConfig{},
		BinanceConfig{},
		UniswapV3ApiConfig{},
		SyncswapConfig{},
	}
)
