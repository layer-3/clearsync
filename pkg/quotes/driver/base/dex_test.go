package base

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_marketSymbol_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		jsonInput string
		expected  []marketSymbol
	}{
		{
			name: "Dexs is specified as false",
			jsonInput: `[
				{
					"symbol": "spot://ETH/USD",
					"quotes": {
						"dexs": false
					}
				}
			]`,
			expected: []marketSymbol{
				{
					Symbol: "spot://ETH/USD",
					Quotes: &marketConfig{Dexs: false},
				},
			},
		},
		{
			name: "Dexs is not specified",
			jsonInput: `[
				{
					"symbol": "spot://PEPE/USD"
				}
			]`,
			expected: []marketSymbol{
				{
					Symbol: "spot://PEPE/USD",
					Quotes: &marketConfig{Dexs: true},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var symbols []marketSymbol
			err := json.Unmarshal([]byte(tt.jsonInput), &symbols)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, symbols)
		})
	}
}
