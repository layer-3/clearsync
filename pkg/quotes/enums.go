package quotes

import "fmt"

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

func ToDriverType(raw string) (*DriverType, error) {
	allDrivers := map[string]DriverType{
		DriverBinance.String():   DriverBinance,
		DriverKraken.String():    DriverKraken,
		DriverOpendax.String():   DriverOpendax,
		DriverBitfaker.String():  DriverBitfaker,
		DriverUniswapV3.String(): DriverUniswapV3,
	}

	driver, ok := allDrivers[raw]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", raw)
	}
	return &driver, nil
}

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

func ToTakerType(raw string) (*TakerType, error) {
	allTypes := map[string]TakerType{
		TakerTypeUnknown.String(): TakerTypeUnknown,
		TakerTypeBuy.String():     TakerTypeBuy,
		TakerTypeSell.String():    TakerTypeSell,
	}

	typ, ok := allTypes[raw]
	if !ok {
		return nil, fmt.Errorf("invalid taker type: %v", raw)
	}
	return &typ, nil
}
