package quotes

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// DriverType is enum that represents
// all available quotes providers.
type DriverType struct {
	slug string
}

func (d DriverType) String() string {
	return d.slug
}

var (
	DriverBinance   = DriverType{"binance"}
	DriverKraken    = DriverType{"kraken"}
	DriverOpendax   = DriverType{"opendax"}
	DriverBitfaker  = DriverType{"bitfaker"}
	DriverUniswapV3 = DriverType{"uniswap_v3"}
	DriverSyncswap  = DriverType{"syncswap"}
	DriverQuickswap = DriverType{"quickswap"}
	DriverSectaV3   = DriverType{"secta_v3"}
	DriverInternal  = DriverType{"internal"} // Internal trades
)

func ToDriverType(raw string) (DriverType, error) {
	allDrivers := map[string]DriverType{
		DriverBinance.String():   DriverBinance,
		DriverKraken.String():    DriverKraken,
		DriverOpendax.String():   DriverOpendax,
		DriverBitfaker.String():  DriverBitfaker,
		DriverUniswapV3.String(): DriverUniswapV3,
		DriverSyncswap.String():  DriverSyncswap,
		DriverQuickswap.String(): DriverQuickswap,
		DriverSectaV3.String():   DriverSectaV3,
		DriverInternal.String():  DriverInternal,
	}

	driver, ok := allDrivers[raw]
	if !ok {
		return DriverType{}, fmt.Errorf("invalid driver type: %v", raw)
	}
	return driver, nil
}

func (t DriverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.slug)
}

func (t *DriverType) UnmarshalJSON(raw []byte) error {
	var rawParsed string
	if err := json.Unmarshal(raw, &rawParsed); err != nil {
		return err
	}

	typ, err := ToDriverType(rawParsed)
	if err != nil {
		return err
	}

	*t = typ
	return nil
}

func (t DriverType) MarshalYAML() (any, error) {
	return t.slug, nil
}

func (t *DriverType) UnmarshalYAML(value *yaml.Node) error {
	var raw string
	if err := value.Decode(&raw); err != nil {
		return err
	}

	typ, err := ToDriverType(raw)
	if err != nil {
		return err
	}

	*t = typ
	return nil
}

func (t *DriverType) SetValue(s string) error {
	typ, err := ToDriverType(s)
	if err != nil {
		return err
	}

	*t = typ
	return nil
}

// TakerType is enum that represents
// the side of taker in a trade.
type TakerType struct {
	slug string
}

func (t TakerType) String() string {
	return t.slug
}

var (
	// TakerTypeUnknown represents the trade,
	// for which you can't determine its type,
	// and therefore the taker side cannot be deduced.
	TakerTypeUnknown = TakerType{""}
	// TakerTypeBuy represents a "buy" trade.
	// It's value is set to "sell" because the sell order it
	// was matched with was present in the order book before,
	// therefore the taker is the "sell" side.
	TakerTypeBuy = TakerType{"sell"}
	// TakerTypeSell represents a "sell" trade.
	// It's value is set to "buy" because the buy order it
	// was matched with was present in the order book before,
	// therefore the taker is the "buy" side.
	TakerTypeSell = TakerType{"buy"}
)

func ToTakerType(raw string) (TakerType, error) {
	allTypes := map[string]TakerType{
		TakerTypeUnknown.String(): TakerTypeUnknown,
		TakerTypeBuy.String():     TakerTypeBuy,
		TakerTypeSell.String():    TakerTypeSell,
	}

	typ, ok := allTypes[raw]
	if !ok {
		return TakerType{}, fmt.Errorf("invalid taker type: %v", raw)
	}
	return typ, nil
}

func (t TakerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.slug)
}

func (t *TakerType) UnmarshalJSON(raw []byte) error {
	var rawParsed string
	if err := json.Unmarshal(raw, &rawParsed); err != nil {
		return err
	}

	typ, err := ToTakerType(rawParsed)
	if err != nil {
		return err
	}

	*t = typ
	return nil
}

func (t TakerType) MarshalYAML() (any, error) {
	return t.slug, nil
}

func (t *TakerType) UnmarshalYAML(value *yaml.Node) error {
	var raw string
	if err := value.Decode(&raw); err != nil {
		return err
	}

	typ, err := ToTakerType(raw)
	if err != nil {
		return err
	}

	*t = typ
	return nil
}

type ExchangeType string

var (
	ExchangeTypeUnspecified ExchangeType = ""
	ExchangeTypeCEX         ExchangeType = "cex"
	ExchangeTypeDEX         ExchangeType = "dex"
	ExchangeTypeHybrid      ExchangeType = "hybrid"
)
