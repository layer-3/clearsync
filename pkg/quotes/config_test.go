package quotes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfigFromEnv(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input     DriverType
		expect    Config
		expectErr bool
	}{
		{
			input:     DriverBinance,
			expect:    &BinanceConfig{},
			expectErr: false,
		},
		{
			input:     DriverKraken,
			expect:    &KrakenConfig{},
			expectErr: false,
		},
		{
			input:     DriverOpendax,
			expect:    &OpendaxConfig{},
			expectErr: false,
		},
		{
			input:     DriverBitfaker,
			expect:    &BitfakerConfig{},
			expectErr: false,
		},
		{
			input:     DriverUniswapV3Api,
			expect:    &UniswapV3ApiConfig{},
			expectErr: false,
		},
		{
			input:     DriverUniswapV3Geth,
			expect:    &UniswapV3GethConfig{},
			expectErr: false,
		},
		{
			input:     DriverSyncswap,
			expect:    &SyncswapConfig{},
			expectErr: false,
		},
		{
			input:     DriverSushiswapV2Geth,
			expect:    &SushiswapV2GethConfig{},
			expectErr: false,
		},
		{
			input:     DriverSushiswapV3Geth,
			expect:    &SushiswapV3GethConfig{},
			expectErr: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.input.String(), func(t *testing.T) {
			t.Parallel()

			config, err := NewConfigFromEnv(tt.input)
			require.True(t, (err != nil) == tt.expectErr)
			require.Equal(t, tt.expect.DriverType(), config.DriverType())
		})
	}
}
