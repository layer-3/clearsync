package quotes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDriver(t *testing.T) {
	t.Parallel()

	t.Run(DriverIndex.String(), func(t *testing.T) {
		t.Parallel()

		config := IndexConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*indexAggregator)
		assert.True(t, ok)
	})

	t.Run(DriverBinance.String(), func(t *testing.T) {
		t.Parallel()

		config := BinanceConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*binance)
		assert.True(t, ok)
	})

	t.Run(DriverKraken.String(), func(t *testing.T) {
		t.Parallel()

		config := KrakenConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*kraken)
		assert.True(t, ok)
	})

	t.Run(DriverBitfaker.String(), func(t *testing.T) {
		t.Parallel()

		config := BitfakerConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*bitfaker)
		assert.True(t, ok)
	})

	t.Run(DriverOpendax.String(), func(t *testing.T) {
		t.Parallel()

		config := OpendaxConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*opendax)
		assert.True(t, ok)
	})

	t.Run(DriverUniswapV3Api.String(), func(t *testing.T) {
		t.Parallel()

		config := UniswapV3ApiConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*uniswapV3Api)
		assert.True(t, ok)
	})

	t.Run(DriverUniswapV3Geth.String(), func(t *testing.T) {
		t.Parallel()

		config := UniswapV3GethConfig{}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(ToConfig(config), outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*uniswapV3Geth)
		assert.True(t, ok)
	})
}
