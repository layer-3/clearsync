package quotes

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/neodax/finex/models/trade"
)

func TestDriverFromName(t *testing.T) {
	t.Run("Bitfaker", func(t *testing.T) {
		priceFeeds, err := DriverFromName("bitfaker")
		require.NoError(t, err)
		assert.Equal(t, &Bitfaker{}, priceFeeds)
	})

	t.Run("Opendax", func(t *testing.T) {
		priceFeeds, err := DriverFromName("opendax")
		require.NoError(t, err)
		assert.Equal(t, &Opendax{}, priceFeeds)
	})

	t.Run("Unknown", func(t *testing.T) {
		dialer, err := DriverFromName("test")
		require.EqualError(t, errors.New("unknown driver test"), err.Error())
		assert.Nil(t, dialer)
	})
}

func TestGetRoutingEvent(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		tradeEvent := trade.Event{
			ID:        0,
			Market:    "btcusd",
			Price:     decimal.Decimal{},
			Amount:    decimal.Decimal{},
			Total:     decimal.Decimal{},
			TakerType: "",
			CreatedAt: 0,
			Source:    "",
		}

		routingEvent, err := GetRoutingEvent(tradeEvent)
		require.NoError(t, err)
		require.NotNil(t, routingEvent)
	})
}
