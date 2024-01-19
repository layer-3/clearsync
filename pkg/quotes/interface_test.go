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

		config := Config{Driver: DriverBinance}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*binance)
		assert.True(t, ok)
	})

	t.Run(DriverKraken.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Driver: DriverKraken}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*kraken)
		assert.True(t, ok)
	})

	t.Run(DriverBitfaker.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Driver: DriverBitfaker}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*bitfaker)
		assert.True(t, ok)
	})

	t.Run(DriverOpendax.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Driver: DriverOpendax}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		_, ok := priceFeeds.(*opendax)
		assert.True(t, ok)
	})
}
