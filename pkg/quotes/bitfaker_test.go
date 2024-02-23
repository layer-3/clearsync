package quotes

import (
	"sync"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var disabledFilter = FilterFactory(FilterConfig{FilterType: DisabledFilterType})

func TestBitfaker_Subscribe(t *testing.T) {
	t.Parallel()

	t.Run("Single market", func(t *testing.T) {
		t.Parallel()

		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch, once: newOnce(), filter: disabledFilter}
		require.NoError(t, client.Start())

		m := NewMarket("btc", "usd")

		err := client.Subscribe(m)
		require.Nil(t, err)

		expectedMarkets := []Market{m}
		assert.Equal(t, client.streams, expectedMarkets)
	})

	t.Run("Multiple markets", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan TradeEvent, 16)
		client := bitfaker{outbox: outbox, once: newOnce(), filter: disabledFilter}
		require.NoError(t, client.Start())

		market1 := NewMarket("btc", "usd")
		err := client.Subscribe(market1)
		require.Nil(t, err)

		market2 := NewMarket("eth", "usd")
		err = client.Subscribe(market2)
		require.Nil(t, err)

		expectedMarkets := []Market{market1, market2}
		assert.Equal(t, client.streams, expectedMarkets)
	})

	t.Run("Subscribe to a market already subscribed to", func(t *testing.T) {
		t.Parallel()

		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch, once: newOnce(), filter: disabledFilter}
		require.NoError(t, client.Start())

		market := NewMarket("btc", "usd")
		err := client.Subscribe(market)
		require.Nil(t, err)

		err = client.Subscribe(market)
		require.Error(t, err)
	})
}

func TestBitfaker_Unsubscribe(t *testing.T) {
	t.Parallel()

	t.Run("Unsubscribe from multiple markets", func(t *testing.T) {
		t.Parallel()

		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch, once: newOnce(), filter: disabledFilter}
		require.NoError(t, client.Start())

		market1 := NewMarket("btc", "usd")
		market2 := NewMarket("eth", "usd")
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))
		require.NoError(t, client.Unsubscribe(market2))

		assert.NotContains(t, client.streams, market1)
		assert.NotContains(t, client.streams, market2)
	})

	t.Run("Unsubscribe from a market not subscribed to", func(t *testing.T) {
		t.Parallel()

		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch, once: newOnce(), filter: disabledFilter}
		require.NoError(t, client.Start())

		market := NewMarket("xrp", "usd")
		err := client.Unsubscribe(market)

		require.Error(t, err)
	})

	t.Run("No effect on other subscriptions after unsubscribing", func(t *testing.T) {
		t.Parallel()

		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch, once: newOnce(), filter: disabledFilter}
		require.NoError(t, client.Start())

		market1 := NewMarket("btc", "usd")
		market2 := NewMarket("eth", "usd")
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))

		assert.NotContains(t, client.streams, market1)
		assert.Contains(t, client.streams, market2)
	})
}

func TestBitfaker_Start(t *testing.T) {
	t.Parallel()

	outbox := make(chan TradeEvent, 1)
	client := bitfaker{
		once:   newOnce(),
		outbox: outbox,
		period: 0 * time.Second,
		filter: disabledFilter,
	}
	market := NewMarket("btc", "usd")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		require.NoError(t, client.Start())
		wg.Done()
	}()
	go func() {
		client.createTradeEvent(market)
		wg.Done()
	}()

	event := <-outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, "btc/usd", event.Market.String())
}

func TestCreateTradeEvent(t *testing.T) {
	t.Parallel()

	outbox := make(chan TradeEvent)
	client := bitfaker{outbox: outbox, once: newOnce(), filter: disabledFilter}
	require.NoError(t, client.Start())

	go func() { client.createTradeEvent(NewMarket("btc", "usd")) }()

	event := <-outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, "btc/usd", event.Market.String())
	assert.Equal(t, DriverBitfaker, event.Source)
	assert.Equal(t, decimal.NewFromFloat(2.213), event.Price)
}
