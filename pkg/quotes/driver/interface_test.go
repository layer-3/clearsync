package driver

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/abi/ilynex_v2_pair"
	"github.com/layer-3/clearsync/pkg/abi/ilynex_v3_pool"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v2_pair"
	"github.com/layer-3/clearsync/pkg/abi/isecta_v3_pool"
	"github.com/layer-3/clearsync/pkg/abi/isyncswap_pool"
	"github.com/layer-3/clearsync/pkg/abi/iuniswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/artifacts/quickswap_v3_pool"
	"github.com/layer-3/clearsync/pkg/quotes/common"
)

func TestNew(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		driverType common.DriverType
		expected   interface{}
	}{
		// Centralized exchanges
		{common.DriverBinance.String(), common.DriverBinance, (*binance)(nil)},
		{common.DriverKraken.String(), common.DriverKraken, (*kraken)(nil)},
		{common.DriverMexc.String(), common.DriverMexc, (*mexc)(nil)},
		{common.DriverOpendax.String(), common.DriverOpendax, (*opendax)(nil)},
		{common.DriverBitfaker.String(), common.DriverBitfaker, (*bitfaker)(nil)},

		// Decentralized exchanges
		{common.DriverUniswapV3.String(), common.DriverUniswapV3, (*baseDEX[iuniswap_v3_pool.IUniswapV3PoolSwap, iuniswap_v3_pool.IUniswapV3Pool, *iuniswap_v3_pool.IUniswapV3PoolSwapIterator])(nil)},
		{common.DriverSyncswap.String(), common.DriverSyncswap, (*baseDEX[isyncswap_pool.ISyncSwapPoolSwap, isyncswap_pool.ISyncSwapPool, *isyncswap_pool.ISyncSwapPoolSwapIterator])(nil)},
		{common.DriverQuickswap.String(), common.DriverQuickswap, (*baseDEX[quickswap_v3_pool.IQuickswapV3PoolSwap, quickswap_v3_pool.IQuickswapV3Pool, *quickswap_v3_pool.IQuickswapV3PoolSwapIterator])(nil)},
		{common.DriverSectaV2.String(), common.DriverSectaV2, (*baseDEX[isecta_v2_pair.ISectaV2PairSwap, isecta_v2_pair.ISectaV2Pair, *isecta_v2_pair.ISectaV2PairSwapIterator])(nil)},
		{common.DriverSectaV3.String(), common.DriverSectaV3, (*baseDEX[isecta_v3_pool.ISectaV3PoolSwap, isecta_v3_pool.ISectaV3Pool, *isecta_v3_pool.ISectaV3PoolSwapIterator])(nil)},
		{common.DriverLynexV2.String(), common.DriverLynexV2, (*baseDEX[ilynex_v2_pair.ILynexPairSwap, ilynex_v2_pair.ILynexPair, *ilynex_v2_pair.ILynexPairSwapIterator])(nil)},
		{common.DriverLynexV3.String(), common.DriverLynexV3, (*baseDEX[ilynex_v3_pool.ILynexV3PoolSwap, ilynex_v3_pool.ILynexV3Pool, *ilynex_v3_pool.ILynexV3PoolSwapIterator])(nil)},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			config, err := NewConfigFromEnv()
			require.NoError(t, err)
			config.Rpc = RpcConfig{
				Ethereum: "wss://mainnet.infura.io/ws/v3/changeme",
				Polygon:  "wss://polygon-mainnet.infura.io/ws/v3/changeme",
				Linea:    "wss://linea-mainnet.infura.io/ws/v3/changeme",
			}
			config.Drivers = []common.DriverType{tc.driverType}

			outbox := make(chan<- TradeEvent, 1)

			priceFeeds, err := New(config, outbox, nil)
			require.NoError(t, err)

			actualType := reflect.TypeOf(priceFeeds)
			expectedType := reflect.TypeOf(tc.expected)
			assert.Equal(t, expectedType, actualType)
		})
	}
}
