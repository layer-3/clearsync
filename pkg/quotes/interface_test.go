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

		outbox := make(chan<- TradeEvent, 1)
		priceFeeds, err := NewDriver(Config{Driver: DriverBinance}, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*binance)
		assert.True(t, ok)
	})

	t.Run(DriverKraken.String(), func(t *testing.T) {
		t.Parallel()

		outbox := make(chan<- TradeEvent, 1)
		priceFeeds, err := NewDriver(Config{Driver: DriverKraken}, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*kraken)
		assert.True(t, ok)
	})

	t.Run(DriverBitfaker.String(), func(t *testing.T) {
		t.Parallel()

		outbox := make(chan<- TradeEvent, 1)
		priceFeeds, err := NewDriver(Config{Driver: DriverBitfaker}, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*bitfaker)
		assert.True(t, ok)
	})

	t.Run(DriverOpendax.String(), func(t *testing.T) {
		t.Parallel()

		outbox := make(chan<- TradeEvent, 1)
		priceFeeds, err := NewDriver(Config{Driver: DriverOpendax}, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*opendax)
		assert.True(t, ok)
	})

	t.Run(DriverUniswapV3Api.String(), func(t *testing.T) {
		t.Parallel()

		outbox := make(chan<- TradeEvent, 1)
		priceFeeds, err := NewDriver(Config{Driver: DriverUniswapV3Api}, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*uniswapV3Api)
		assert.True(t, ok)
	})

	t.Run(DriverUniswapV3Geth.String(), func(t *testing.T) {
		t.Parallel()

		outbox := make(chan<- TradeEvent, 1)
		priceFeeds, err := NewDriver(Config{Driver: DriverUniswapV3Geth}, outbox)
		require.NoError(t, err)

		_, ok := priceFeeds.(*uniswapV3Geth)
		assert.True(t, ok)
	})
}
