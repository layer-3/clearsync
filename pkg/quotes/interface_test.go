package quotes

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_pair"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_pair"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_pool"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_pool"
)

func TestNewDriver(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		driverType DriverType
		expected   interface{}
	}{
		// Centralized exchanges
		{DriverBinance.String(), DriverBinance, (*binance)(nil)},
		{DriverKraken.String(), DriverKraken, (*kraken)(nil)},
		{DriverMexc.String(), DriverMexc, (*mexc)(nil)},
		{DriverOpendax.String(), DriverOpendax, (*opendax)(nil)},
		{DriverBitfaker.String(), DriverBitfaker, (*bitfaker)(nil)},

		// Decentralized exchanges
		{DriverUniswapV3.String(), DriverUniswapV3, (*baseDEX[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator])(nil)},
		{DriverSyncswap.String(), DriverSyncswap, (*baseDEX[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool, *isyncswap_pool.ISyncSwapPoolSwapIterator])(nil)},
		{DriverQuickswap.String(), DriverQuickswap, (*baseDEX[quickswap_v3_pool.IQuickswapV3PoolSwap, quickswap_v3_pool.IQuickswapV3Pool, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator])(nil)},
		{DriverSectaV2.String(), DriverSectaV2, (*baseDEX[isecta_v2_pair.ISectaV2PairSwap, isecta_v2_pair.ISectaV2Pair, *isecta_v2_pair.ISectaV2PairSwapIterator])(nil)},
		{DriverSectaV3.String(), DriverSectaV3, (*baseDEX[isecta_v3_pool.ISectaV3PoolSwap, isecta_v3_pool.ISectaV3Pool, *isecta_v3_pool.ISectaV3PoolSwapIterator])(nil)},
		{DriverLynex.String(), DriverLynex, (*baseDEX[ilynex_pair.ILynexPairSwap, ilynex_pair.ILynexPair])(nil)},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			config, err := NewConfigFromEnv()
			require.NoError(t, err)
			config.Drivers = []DriverType{tc.driverType}

			outbox := make(chan<- TradeEvent, 1)

			priceFeeds, err := NewDriver(config, outbox, nil, nil)
			require.NoError(t, err)

			actualType := reflect.TypeOf(priceFeeds)
			expectedType := reflect.TypeOf(tc.expected)
			assert.Equal(t, expectedType, actualType)
		})
	}
}
