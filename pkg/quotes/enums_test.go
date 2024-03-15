package quotes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestDriverType_Marshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input        DriverType
		expectedJSON string
		expectedYAML string
	}{
		{
			input:        DriverBinance,
			expectedJSON: `"binance"`,
			expectedYAML: "binance\n",
		}, {
			input:        DriverKraken,
			expectedJSON: `"kraken"`,
			expectedYAML: "kraken\n",
		}, {
			input:        DriverOpendax,
			expectedJSON: `"opendax"`,
			expectedYAML: "opendax\n",
		}, {
			input:        DriverBitfaker,
			expectedJSON: `"bitfaker"`,
			expectedYAML: "bitfaker\n",
		}, {
			input:        DriverUniswapV3Api,
			expectedJSON: `"uniswap_v3_api"`,
			expectedYAML: "uniswap_v3_api\n",
		}, {
			input:        DriverUniswapV3Geth,
			expectedJSON: `"uniswap_v3_geth"`,
			expectedYAML: "uniswap_v3_geth\n",
		}, {
			input:        DriverSyncswap,
			expectedJSON: `"syncswap"`,
			expectedYAML: "syncswap\n",
		}, {
			input:        DriverQuickswap,
			expectedJSON: `"quickswap"`,
			expectedYAML: "quickswap\n",
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

func TestDriverType_Unmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		inputJSON string
		inputYAML string
		expected  DriverType
		shouldErr bool
	}{
		{
			inputJSON: `"binance"`,
			inputYAML: "binance",
			expected:  DriverBinance,
			shouldErr: false,
		}, {
			inputJSON: `"kraken"`,
			inputYAML: "kraken",
			expected:  DriverKraken,
			shouldErr: false,
		}, {
			inputJSON: `"opendax"`,
			inputYAML: "opendax",
			expected:  DriverOpendax,
			shouldErr: false,
		}, {
			inputJSON: `"bitfaker"`,
			inputYAML: "bitfaker",
			expected:  DriverBitfaker,
			shouldErr: false,
		}, {
			inputJSON: `"uniswap_v3_api"`,
			inputYAML: "uniswap_v3_api",
			expected:  DriverUniswapV3Api,
			shouldErr: false,
		}, {
			inputJSON: `"uniswap_v3_geth"`,
			inputYAML: "uniswap_v3_geth",
			expected:  DriverUniswapV3Geth,
			shouldErr: false,
		}, {
			inputJSON: `"syncswap"`,
			inputYAML: "syncswap",
			expected:  DriverSyncswap,
			shouldErr: false,
		}, {
			inputJSON: `"quickswap"`,
			inputYAML: "quickswap",
			expected:  DriverQuickswap,
			shouldErr: false,
		}, {
			inputJSON: `"invalid"`,
			inputYAML: "invalid",
			expected:  DriverType{},
			shouldErr: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("JSON %s", test.expected.String()), func(t *testing.T) {
			t.Parallel()

			var fromJSON DriverType
			if err := json.Unmarshal([]byte(test.inputJSON), &fromJSON); (err != nil) != test.shouldErr {
				t.Errorf("UnmarshalJSON %v, expected error: %v, got error: %v", test.expected, test.shouldErr, err)
			}
			require.Equal(t, test.expected, fromJSON)
		})

		t.Run(fmt.Sprintf("YAML %s", test.expected.String()), func(t *testing.T) {
			t.Parallel()

			var fromYAML DriverType
			if err := yaml.Unmarshal([]byte(test.inputYAML), &fromYAML); (err != nil) != test.shouldErr {
				t.Errorf("UnmarshalYAML %v, expected error: %v, got error: %v", test.expected, test.shouldErr, err)
			}
			require.Equal(t, test.expected, fromYAML)
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
