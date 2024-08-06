package quotes

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestCoingecko_GetPrices(t *testing.T) {
	t.Run("Fetch token prices", func(t *testing.T) {
		// Each request supports up to 250 tokens.
		tokens := map[string]common.Address{
			"bitcoin":       {1},
			"ethereum":      {2},
			"osis":          {3},
			"matic-network": {4},
			"duckies":       {5},
		}

		assets, err := FetchTokens()
		if err != nil {
			fmt.Printf("Error fetching coin list: %v\n", err)
			return
		}

		validTokens := make(map[string]common.Address)
		for id, token := range tokens {
			if TokenExists(id, assets) {
				validTokens[id] = token
			} else {
				fmt.Printf("Token not found on CoinGecko: %s (%s)\n", id, token)
			}
		}

		if len(validTokens) == 0 {
			fmt.Println("No valid tokens found to fetch prices for.")
			return
		}

		prices, err := FetchPrices(validTokens)
		if err != nil {
			fmt.Printf("Error fetching prices: %v\n", err)
			return
		}

		for id, token := range validTokens {
			price, ok := prices[token]
			if !ok {
				fmt.Printf("Price for %s not found\n", id)
			} else {
				fmt.Printf("Price for %s is %s USD\n", id, price)
			}
		}

		assert.Equal(t, len(validTokens), 5)
	})
}
