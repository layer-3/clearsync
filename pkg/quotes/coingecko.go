package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/shopspring/decimal"
)

type Prices map[string]map[string]float64

type Asset struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type TokenNetwork struct {
	TokenAddress string
	ChainID      uint32
}

// TODO: add support for Pro subscription (API endpoint + Header name)
func setAuthHeader(req *http.Request, apiKey string) {
	if apiKey != "" {
		req.Header.Set("x-cg-demo-api-key", apiKey)
	}
}

// FetchTokens fetches the list of all available tokens from CoinGecko.
func FetchTokens(apiKey string) ([]Asset, error) {
	url := "https://api.coingecko.com/api/v3/coins/list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	setAuthHeader(req, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
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

// FetchPrices fetches the current prices for a map of tokens from CoinGecko (map[address]CoinGeckoID).
func FetchPrices(tokens map[TokenNetwork]string, apiKey string) (map[TokenNetwork]decimal.Decimal, error) {
	// Filter out recurring IDs, to lower the number of credits spent
	uniqueIDs := make(map[string]struct{})
	for _, id := range tokens {
		uniqueIDs[id] = struct{}{}
	}

	// Convert map to slice
	ids := make([]string, len(tokens))
	for _, id := range tokens {
		ids = append(ids, id)
	}
	idsQuery := strings.Join(ids, ",")

	url := fmt.Sprintf("%s?ids=%s&vs_currencies=usd", "https://api.coingecko.com/api/v3/simple/price", idsQuery)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	setAuthHeader(req, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
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

	tokenPrices := make(map[TokenNetwork]decimal.Decimal)
	for id, price := range prices {
		price, ok := price["usd"]
		if ok {
			for addr, apiID := range tokens {
				if apiID == id {
					tokenPrices[addr] = decimal.NewFromFloat(price)
				}
			}
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
