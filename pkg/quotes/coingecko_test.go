package quotes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoingecko_GetPrices(t *testing.T) {
	t.Run("Fetch token prices", func(t *testing.T) {
		// Each request supports up to 250 tokens.
		tokens := []Token{
			{ID: "bitcoin", Symbol: "BTC"},
			{ID: "ethereum", Symbol: "ETH"},
			{ID: "osis", Symbol: "OSIS"},
			{ID: "matic-network", Symbol: "MATIC"},
			{ID: "duckies", Symbol: "DUCKIES"},
		}

		assets, err := FetchTokens()
		if err != nil {
			fmt.Printf("Error fetching coin list: %v\n", err)
			return
		}

		var validTokens []Token
		for _, token := range tokens {
			if TokenExists(token, assets) {
				validTokens = append(validTokens, token)
			} else {
				fmt.Printf("Token not found on CoinGecko: %s (%s)\n", token.Symbol, token.ID)
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

		for _, token := range validTokens {
			price, ok := prices[token.ID]["usd"]
			if !ok {
				fmt.Printf("Price for %s not found\n", token.Symbol)
			} else {
				fmt.Printf("Price for %s is $%.5f USD\n", token.Symbol, price)
			}
		}

		assert.Equal(t, len(validTokens), 5)
	})
}
