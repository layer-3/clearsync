package quotes

// DriverType is enum that represents
// all available quotes providers.
type DriverType struct {
	slug string
}

func (d *DriverType) String() string {
	return d.slug
}

var (
	DriverBinance   = DriverType{"binance"}
	DriverKraken    = DriverType{"kraken"}
	DriverOpendax   = DriverType{"opendax"}
	DriverBitfaker  = DriverType{"bitfaker"}
	DriverUniswapV3 = DriverType{"uniswap_v3"}
)

// TakerType is enum that represents
// the side of taker in a trade.
type TakerType struct {
	slug string
}

func (t *TakerType) String() string {
	return t.slug
}

var (
	TakerTypeUnknown = TakerType{""}
	TakerTypeBuy     = TakerType{"sell"}
	TakerTypeSell    = TakerType{"buy"}
)
