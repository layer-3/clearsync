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
		assert.Equal(t, newBinance(config, outbox), priceFeeds)
	})

	t.Run(DriverKraken.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Driver: DriverKraken}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		assert.Equal(t, newKraken(config, outbox), priceFeeds)
	})

	t.Run(DriverBitfaker.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Driver: DriverBitfaker}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		assert.Equal(t, newBitfaker(config, outbox), priceFeeds)
	})

	t.Run(DriverOpendax.String(), func(t *testing.T) {
		t.Parallel()

		config := Config{Driver: DriverOpendax}
		outbox := make(chan<- TradeEvent, 1)

		priceFeeds, err := NewDriver(config, outbox)
		require.NoError(t, err)
		assert.Equal(t, newOpendax(config, outbox), priceFeeds)
	})

	t.Run("Unknown driver", func(t *testing.T) {
		t.Parallel()

		priceFeeds, err := NewDriver(
			Config{Driver: DriverType{"wtf"}},
			make(chan<- TradeEvent, 1),
		)
		require.Error(t, err)
		assert.Nil(t, priceFeeds)
	})
}
