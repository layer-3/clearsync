package quotes

import "github.com/shopspring/decimal"

var (
	DefaultWeightsMap = map[DriverType]decimal.Decimal{
		DriverBinance:   decimal.NewFromInt(3),
		DriverKraken:    decimal.NewFromInt(3),
		DriverOpendax:   decimal.NewFromInt(2),
		DriverBitfaker:  decimal.NewFromInt(1),
		DriverUniswapV3: decimal.NewFromInt(2),
	}

	AllDrivers = []Config{
		{
			Driver: DriverBinance,
		},
		{
			Driver: DriverUniswapV3,
		},
	}
)
