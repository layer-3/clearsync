package driver

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/filter"
)

func TestBitfaker_Subscribe(t *testing.T) {
	t.Parallel()

	disabledFilter := filter.New(filter.Config{Type: filter.TypeDisabled})

	t.Run("Single market", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := bitfaker{
			once:          common.NewOnce(),
			outbox:        outbox,
			streamPeriods: make(map[common.Market]time.Duration),
			streams:       make(map[common.Market]chan struct{}),
			filter:        disabledFilter,
		}
		require.NoError(t, client.Start())

		m := common.NewMarket("btc", "usd")

		err := client.Subscribe(m)
		require.Nil(t, err)

		_, ok := client.streams[m]
		require.True(t, ok)

		assert.Equal(t, len(client.streams), 1)
	})

	t.Run("Multiple markets", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := bitfaker{
			once:          common.NewOnce(),
			outbox:        outbox,
			streamPeriods: make(map[common.Market]time.Duration),
			streams:       make(map[common.Market]chan struct{}),
			filter:        disabledFilter,
		}
		require.NoError(t, client.Start())

		market1 := common.NewMarket("btc", "usd")
		err := client.Subscribe(market1)
		require.Nil(t, err)

		market2 := common.NewMarket("eth", "usd")
		err = client.Subscribe(market2)
		require.Nil(t, err)

		_, ok := client.streams[market1]
		require.True(t, ok)

		_, ok = client.streams[market2]
		require.True(t, ok)

		assert.Equal(t, len(client.streams), 2)
	})

	t.Run("Subscribe to a market already subscribed to", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := bitfaker{
			once:          common.NewOnce(),
			outbox:        outbox,
			streamPeriods: make(map[common.Market]time.Duration),
			streams:       make(map[common.Market]chan struct{}),
			filter:        disabledFilter,
		}
		require.NoError(t, client.Start())

		market := common.NewMarket("btc", "usd")
		err := client.Subscribe(market)
		require.Nil(t, err)

		err = client.Subscribe(market)
		require.Error(t, err)
	})
}

func TestBitfaker_Unsubscribe(t *testing.T) {
	t.Parallel()

	disabledFilter := filter.New(filter.Config{Type: filter.TypeDisabled})

	t.Run("Unsubscribe from multiple markets", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := bitfaker{
			once:          common.NewOnce(),
			outbox:        outbox,
			streamPeriods: make(map[common.Market]time.Duration),
			streams:       make(map[common.Market]chan struct{}),
			filter:        disabledFilter,
		}
		require.NoError(t, client.Start())

		market1 := common.NewMarket("btc", "usd")
		market2 := common.NewMarket("eth", "usd")
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))
		require.NoError(t, client.Unsubscribe(market2))

		assert.NotContains(t, client.streams, market1)
		assert.NotContains(t, client.streams, market2)
	})

	t.Run("Unsubscribe from a market not subscribed to", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := bitfaker{
			once:          common.NewOnce(),
			outbox:        outbox,
			streamPeriods: make(map[common.Market]time.Duration),
			streams:       make(map[common.Market]chan struct{}),
			filter:        disabledFilter,
		}
		require.NoError(t, client.Start())

		market := common.NewMarket("xrp", "usd")
		err := client.Unsubscribe(market)

		require.Error(t, err)
	})

	t.Run("No effect on other subscriptions after unsubscribing", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := bitfaker{
			once:          common.NewOnce(),
			outbox:        outbox,
			streamPeriods: make(map[common.Market]time.Duration),
			streams:       make(map[common.Market]chan struct{}),
			filter:        disabledFilter,
		}
		require.NoError(t, client.Start())

		market1 := common.NewMarket("btc", "usd")
		market2 := common.NewMarket("eth", "usd")
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))

		assert.NotContains(t, client.streams, market1)
		assert.Contains(t, client.streams, market2)
	})
}

func TestCreateTradeEvent(t *testing.T) {
	t.Parallel()

	disabledFilter := filter.New(filter.Config{Type: filter.TypeDisabled})

	outbox := make(chan common.TradeEvent)
	client := bitfaker{
		outbox:        outbox,
		once:          common.NewOnce(),
		filter:        disabledFilter,
		streamPeriods: make(map[common.Market]time.Duration),
		streams:       make(map[common.Market]chan struct{}),
	}
	require.NoError(t, client.Start())

	startPrice := 100.0
	startAmount := 10.0
	priceVolatility := 2.0
	amountVolatility := 1.0

	newPrice := initializeMarket(startPrice, priceVolatility)
	newAmount := initializeMarket(startAmount, amountVolatility)

	go func() { client.createTradeEvent(common.NewMarket("btc", "usd"), newPrice, newAmount) }()

	event := <-outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, "btc/usd", event.Market.String())
	assert.Equal(t, common.DriverBitfaker, event.Source)
}
