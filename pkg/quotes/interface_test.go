package quotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDriver(t *testing.T) {
	t.Parallel()

	t.Run(DriverBinance.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverBinance}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*binance)
		assert.True(t, ok)
	})

	t.Run(DriverKraken.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverKraken}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*kraken)
		assert.True(t, ok)
	})

	t.Run(DriverBitfaker.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverBitfaker}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*bitfaker)
		assert.True(t, ok)
	})

	t.Run(DriverOpendax.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverOpendax}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*opendax)
		assert.True(t, ok)
	})

	t.Run(DriverUniswapV3Api.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverUniswapV3Api}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*uniswapV3Api)
		assert.True(t, ok)
	})

	t.Run(DriverUniswapV3Geth.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverUniswapV3Geth}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*uniswapV3Geth)
		assert.True(t, ok)
	})

	t.Run(DriverSyncswap.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverSyncswap}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*syncswap)
		assert.True(t, ok)
	})

	t.Run(DriverQuickswap.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Drivers: []DriverType{DriverQuickswap}}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*quickswap)
		assert.True(t, ok)
	})
}
