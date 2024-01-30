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
	DriverBinance         = DriverType{"binance"}
	DriverKraken          = DriverType{"kraken"}
	DriverOpendax         = DriverType{"opendax"}
	DriverBitfaker        = DriverType{"bitfaker"}
	DriverUniswapV3Api    = DriverType{"uniswap_v3_api"}
	DriverUniswapV3Geth   = DriverType{"uniswap_v3_geth"}
	DriverSyncswap        = DriverType{"syncswap"}
	DriverSushiswapV2Geth = DriverType{"sushiswap_v2_geth"}
	DriverSushiswapV3Geth = DriverType{"sushiswap_v3_geth"}
)

func ToDriverType(raw string) (DriverType, error) {
	allDrivers := map[string]DriverType{
		DriverBinance.String():         DriverBinance,
		DriverKraken.String():          DriverKraken,
		DriverOpendax.String():         DriverOpendax,
		DriverBitfaker.String():        DriverBitfaker,
		DriverUniswapV3Api.String():    DriverUniswapV3Api,
		DriverUniswapV3Geth.String():   DriverUniswapV3Geth,
		DriverSyncswap.String():        DriverSyncswap,
		DriverSushiswapV2Geth.String(): DriverSushiswapV2Geth,
		DriverSushiswapV3Geth.String(): DriverSushiswapV3Geth,
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

// TakerType is enum that represents
// the side of taker in a trade.
type TakerType struct {
	slug string
}

func (t TakerType) String() string {
	return t.slug
}

var (
	TakerTypeUnknown = TakerType{""}
	TakerTypeBuy     = TakerType{"sell"}
	TakerTypeSell    = TakerType{"buy"}
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
