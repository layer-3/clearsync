package quotes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoingecko_GetPrices(t *testing.T) {
	t.Run("Fetch token prices", func(t *testing.T) {
		// Each request supports up to 250 tokens.
		tokens := map[TokenNetwork]string{
			{"bitcoin_token_address", 1}:  "bitcoin",
			{"ethereum_token_address", 2}: "ethereum",
			{"osis_token_address", 3}:     "osis",
			{"matic_token_address", 4}:    "matic-network",
			{"duckies_token_address", 5}:  "duckies",
		}

		assets, err := FetchTokens("")
		if err != nil {
			fmt.Printf("Error fetching coin list: %v\n", err)
			return
		}

		validTokens := make(map[TokenNetwork]string)
		for token, id := range tokens {
			if TokenExists(id, assets) {
				validTokens[token] = id
			} else {
				fmt.Printf("Token not found on CoinGecko: %s (%s)\n", id, token.TokenAddress)
			}
		}

		if len(validTokens) == 0 {
			fmt.Println("No valid tokens found to fetch prices for.")
			return
		}

		prices, err := FetchPrices(validTokens, "")
		if err != nil {
			fmt.Printf("Error fetching prices: %v\n", err)
			return
		}

		for address, id := range validTokens {
			price, ok := prices[address]
			if !ok {
				fmt.Printf("Price for %s not found\n", id)
			} else {
				fmt.Printf("Price for %s is %s USD\n", id, price)
			}
		}

		assert.Equal(t, len(validTokens), 5)
	})
}
