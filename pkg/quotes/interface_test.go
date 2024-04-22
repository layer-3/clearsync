package quotes

import (
	"reflect"
	"testing"

	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_pair"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_pool"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/iquickswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
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
		{DriverUniswapV3.String(), DriverUniswapV3, (*baseDEX[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool])(nil)},
		{DriverSyncswap.String(), DriverSyncswap, (*baseDEX[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool])(nil)},
		{DriverQuickswap.String(), DriverQuickswap, (*baseDEX[iquickswap_v3_pool.IQuickswapV3PoolSwap, iquickswap_v3_pool.IQuickswapV3Pool])(nil)},
		{DriverSectaV2.String(), DriverSectaV2, (*baseDEX[isecta_v2_pair.ISectaV2PairSwap, isecta_v2_pair.ISectaV2Pair])(nil)},
		{DriverSectaV3.String(), DriverSectaV3, (*baseDEX[isecta_v3_pool.ISectaV3PoolSwap, isecta_v3_pool.ISectaV3Pool])(nil)},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			config, err := NewConfigFromEnv()
			require.NoError(t, err)
			config.Drivers = []DriverType{tc.driverType}

			outbox := make(chan<- TradeEvent, 1)

			priceFeeds, err := NewDriver(config, outbox)
			require.NoError(t, err)

			actualType := reflect.TypeOf(priceFeeds)
			expectedType := reflect.TypeOf(tc.expected)
			assert.Equal(t, expectedType, actualType)
		})
	}
}
