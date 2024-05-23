package quotes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestDriverType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input      DriverType
		expectJSON string
		expectYAML string
		expectErr  bool
	}{
		{
			input:      DriverBinance,
			expectJSON: `"binance"`,
			expectYAML: "binance\n",
			expectErr:  false,
		}, {
			input:      DriverKraken,
			expectJSON: `"kraken"`,
			expectYAML: "kraken\n",
			expectErr:  false,
		}, {
			input:      DriverOpendax,
			expectJSON: `"opendax"`,
			expectYAML: "opendax\n",
			expectErr:  false,
		}, {
			input:      DriverBitfaker,
			expectJSON: `"bitfaker"`,
			expectYAML: "bitfaker\n",
			expectErr:  false,
		}, {
			input:      DriverUniswapV3,
			expectJSON: `"uniswap_v3"`,
			expectYAML: "uniswap_v3\n",
			expectErr:  false,
		}, {
			input:      DriverSyncswap,
			expectJSON: `"syncswap"`,
			expectYAML: "syncswap\n",
			expectErr:  false,
		}, {
			input:      DriverQuickswap,
			expectJSON: `"quickswap"`,
			expectYAML: "quickswap\n",
			expectErr:  false,
		}, {
			input:      DriverSectaV2,
			expectJSON: `"secta_v2"`,
			expectYAML: "secta_v2\n",
			expectErr:  false,
		}, {
			input:      DriverSectaV3,
			expectJSON: `"secta_v3"`,
			expectYAML: "secta_v3\n",
			expectErr:  false,
		}, {
			input:      DriverMexc,
			expectJSON: `"mexc"`,
			expectYAML: "mexc\n",
			expectErr:  false,
		}, {
			input:      DriverInternal,
			expectJSON: `"internal"`,
			expectYAML: "internal\n",
			expectErr:  false,
		}, {
			input:      DriverType{},
			expectJSON: `""`,
			expectYAML: "\"\"\n",
			expectErr:  true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Marhsal JSON %s", test.input.String()), func(t *testing.T) {
			t.Parallel()

			dataJSON, err := json.Marshal(&test.input)
			require.NoError(t, err)
			require.Equal(t, test.expectJSON, string(dataJSON))
		})

		t.Run(fmt.Sprintf("Unmarhsal JSON %s", test.input.String()), func(t *testing.T) {
			t.Parallel()

			var fromJSON DriverType
			if err := json.Unmarshal([]byte(test.expectJSON), &fromJSON); (err != nil) != test.expectErr {
				t.Errorf("UnmarshalJSON %v, expected error: %v, got error: %v", test.input, test.expectErr, err)
			}
			require.Equal(t, test.input, fromJSON)
		})

		t.Run(fmt.Sprintf("Marhsal YAML %s", test.input.String()), func(t *testing.T) {
			t.Parallel()

			dataYAML, err := yaml.Marshal(test.input)
			require.NoError(t, err)
			require.Equal(t, test.expectYAML, string(dataYAML))
		})

		t.Run(fmt.Sprintf("Unmarshal YAML %s", test.input.String()), func(t *testing.T) {
			t.Parallel()

			var fromYAML DriverType
			if err := yaml.Unmarshal([]byte(test.expectYAML), &fromYAML); (err != nil) != test.expectErr {
				t.Errorf("UnmarshalYAML %v, expected error: %v, got error: %v", test.input, test.expectErr, err)
			}
			require.Equal(t, test.input, fromYAML)
		})
	}
}

func TestTakerType_Marshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input        TakerType
		expectedJSON string
		expectedYAML string
	}{
		{
			input:        TakerTypeUnknown,
			expectedJSON: `""`,
			expectedYAML: "\"\"\n",
		}, {
			input:        TakerTypeBuy,
			expectedJSON: `"sell"`,
			expectedYAML: "sell\n",
		}, {
			input:        TakerTypeSell,
			expectedJSON: `"buy"`,
			expectedYAML: "buy\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("JSON %s", test.input.String()), func(t *testing.T) {
			t.Parallel()

			dataJSON, err := json.Marshal(&test.input)
			require.NoError(t, err)
			require.Equal(t, test.expectedJSON, string(dataJSON))
		})

		t.Run(fmt.Sprintf("YAML %s", test.input.String()), func(t *testing.T) {
			t.Parallel()

			dataYAML, err := yaml.Marshal(test.input)
			require.NoError(t, err)
			require.Equal(t, test.expectedYAML, string(dataYAML))
		})
	}
}

func TestTakerType_Unmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputJSON string
		inputYAML string
		expected  TakerType
		shouldErr bool
	}{
		{
			name:      "taker type unknown",
			inputJSON: `""`,
			inputYAML: `""`,
			expected:  TakerTypeUnknown,
			shouldErr: false,
		}, {
			name:      "taker type buy",
			inputJSON: `"sell"`,
			inputYAML: "sell",
			expected:  TakerTypeBuy,
			shouldErr: false,
		}, {
			name:      "taker type sell",
			inputJSON: `"buy"`,
			inputYAML: "buy",
			expected:  TakerTypeSell,
			shouldErr: false,
		}, {
			name:      "invalid taker type",
			inputJSON: `"invalid"`,
			inputYAML: "invalid",
			expected:  TakerType{},
			shouldErr: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("JSON %s", test.expected.String()), func(t *testing.T) {
			t.Parallel()

			var fromJSON TakerType
			if err := json.Unmarshal([]byte(test.inputJSON), &fromJSON); (err != nil) != test.shouldErr {
				t.Errorf("UnmarshalJSON %v, expected error: %v, got error: %v", test.expected, test.shouldErr, err)
			}
			require.Equal(t, test.expected, fromJSON)
		})

		t.Run(fmt.Sprintf("YAML %s", test.expected.String()), func(t *testing.T) {
			t.Parallel()

			var fromYAML TakerType
			if err := yaml.Unmarshal([]byte(test.inputYAML), &fromYAML); (err != nil) != test.shouldErr {
				t.Errorf("UnmarshalYAML %v, expected error: %v, got error: %v", test.expected, test.shouldErr, err)
			}
			require.Equal(t, test.expected, fromYAML)
		})
	}
}
