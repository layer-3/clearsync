package quotes

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDriver(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		driverType DriverType
		expected   interface{}
	}{
		{DriverBinance.String(), DriverBinance, (*binance)(nil)},
		{DriverKraken.String(), DriverKraken, (*kraken)(nil)},
		{DriverOpendax.String(), DriverOpendax, (*opendax)(nil)},
		{DriverBitfaker.String(), DriverBitfaker, (*bitfaker)(nil)},
		{DriverUniswapV3Api.String(), DriverUniswapV3Api, (*uniswapV3Api)(nil)},
		{DriverUniswapV3Geth.String(), DriverUniswapV3Geth, (*uniswapV3Geth)(nil)},
		{DriverSyncswap.String(), DriverSyncswap, (*syncswap)(nil)},
		{DriverQuickswap.String(), DriverQuickswap, (*quickswap)(nil)},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			config := Config{Drivers: []DriverType{tc.driverType}}
			outbox := make(chan<- TradeEvent, 1)

			priceFeeds, err := NewDriver(config, outbox)
			require.NoError(t, err)

			actualType := reflect.TypeOf(priceFeeds)
			expectedType := reflect.TypeOf(tc.expected)
			assert.Equal(t, expectedType, actualType)
		})
	}
}
