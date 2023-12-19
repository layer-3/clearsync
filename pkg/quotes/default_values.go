package quotes

import "github.com/shopspring/decimal"

var (
	DefaultWeightsMap = map[DriverType]decimal.Decimal{
		DriverBinance:       decimal.NewFromInt(4),
		DriverKraken:        decimal.NewFromInt(4),
		DriverOpendax:       decimal.NewFromInt(1),
		DriverBitfaker:      decimal.NewFromInt(1),
		DriverUniswapV3Api:  decimal.NewFromInt(2),
		DriverUniswapV3Geth: decimal.NewFromInt(2),
	}

	AllDrivers = []Config{
		{
			Driver: DriverKraken,
		},
		{
			Driver: DriverBinance,
		},
		{
			Driver: DriverUniswapV3Api,
		},
	}
)
