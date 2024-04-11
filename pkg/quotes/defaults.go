package quotes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shopspring/decimal"
)

var (
	DefaultWeightsMap = map[DriverType]decimal.Decimal{
		DriverKraken:        decimal.NewFromInt(15),
		DriverBinance:       decimal.NewFromInt(20),
		DriverUniswapV3Api:  decimal.NewFromInt(50),
		DriverUniswapV3Geth: decimal.NewFromInt(50),
		DriverSyncswap:      decimal.NewFromInt(50),
		DriverQuickswap:     decimal.NewFromInt(50),
	}

	DefaultMarketsMapping = map[string][]string{"usd": {"eth", "weth", "matic"}}
)

func getMapping(url string) (map[string][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mappings map[string]map[string][]string
	if err := json.Unmarshal(body, &mappings); err != nil {
		return nil, err
	}
	return mappings["tokens"], nil
}
