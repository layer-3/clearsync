package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type Prices map[string]map[string]float64

type Asset struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// FetchTokens fetches the list of all available tokens from CoinGecko.
func FetchTokens() ([]Asset, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch coin list: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var asset []Asset
	if err := json.Unmarshal(body, &asset); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return asset, nil
}

// FetchPrices fetches the current prices for a map of tokens from CoinGecko (map[CoinGeckoID]address).
func FetchPrices(tokens map[string]common.Address) (map[common.Address]decimal.Decimal, error) {
	ids := make([]string, len(tokens))
	for id := range tokens {
		ids = append(ids, id)
	}
	idsQuery := strings.Join(ids, ",")

	url := fmt.Sprintf("%s?ids=%s&vs_currencies=usd", "https://api.coingecko.com/api/v3/simple/price", idsQuery)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var prices Prices
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	tokenPrices := make(map[common.Address]decimal.Decimal)
	for id, price := range prices {
		price, ok := price["usd"]
		if ok {
			tokenPrices[tokens[id]] = decimal.NewFromFloat(price)
			continue
		}
	}

	return tokenPrices, nil
}

// TokenExists checks if a token is in the list.
func TokenExists(tokenID string, tokens []Asset) bool {
	for _, t := range tokens {
		if t.ID == tokenID {
			return true
		}
	}
	return false
}
