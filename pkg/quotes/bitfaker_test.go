package quotes

import (
	"sync"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBitfaker_Subscribe(t *testing.T) {
	t.Run("Single market", func(t *testing.T) {
		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch}

		m := Market{BaseUnit: "btc", QuoteUnit: "usd"}
		err := client.Subscribe(m)
		require.Nil(t, err)

		expectedMarkets := []Market{m}
		assert.Equal(t, client.markets, expectedMarkets)
	})

	t.Run("Multiple markets", func(t *testing.T) {
		outbox := make(chan TradeEvent, 16)
		client := bitfaker{outbox: outbox}

		market1 := Market{BaseUnit: "btc", QuoteUnit: "usd"}
		err := client.Subscribe(market1)
		require.Nil(t, err)

		market2 := Market{BaseUnit: "eth", QuoteUnit: "usd"}
		err = client.Subscribe(market2)
		require.Nil(t, err)

		expectedMarkets := []Market{market1, market2}
		assert.Equal(t, client.markets, expectedMarkets)
	})

	t.Run("Subscribe to a market already subscribed to", func(t *testing.T) {
		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch}

		market := Market{BaseUnit: "btc", QuoteUnit: "usd"}
		err := client.Subscribe(market)
		require.Nil(t, err)

		err = client.Subscribe(market)
		require.Error(t, err)
	})
}

func TestBitfaker_Unsubscribe(t *testing.T) {
	t.Run("Unsubscribe from multiple markets", func(t *testing.T) {
		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch}

		market1 := Market{BaseUnit: "btc", QuoteUnit: "usd"}
		market2 := Market{BaseUnit: "eth", QuoteUnit: "usd"}
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))
		require.NoError(t, client.Unsubscribe(market2))

		assert.NotContains(t, client.markets, market1)
		assert.NotContains(t, client.markets, market2)
	})

	t.Run("Unsubscribe from a market not subscribed to", func(t *testing.T) {
		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch}

		market := Market{BaseUnit: "xrp", QuoteUnit: "usd"}
		err := client.Unsubscribe(market)

		require.Error(t, err)
	})

	t.Run("No effect on other subscriptions after unsubscribing", func(t *testing.T) {
		ch := make(chan TradeEvent, 16)
		client := bitfaker{outbox: ch}

		market1 := Market{BaseUnit: "btc", QuoteUnit: "usd"}
		market2 := Market{BaseUnit: "eth", QuoteUnit: "usd"}
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))

		assert.NotContains(t, client.markets, market1)
		assert.Contains(t, client.markets, market2)
	})
}

func TestBitfaker_Start(t *testing.T) {
	outbox := make(chan TradeEvent, 1)
	tradeSampler := *newTradeSampler(TradeSamplerConfig{
		Enabled:           false,
		DefaultPercentage: 0,
	})
	client := bitfaker{outbox: outbox, period: 0 * time.Second, tradeSampler: tradeSampler}
	market := Market{BaseUnit: "btc", QuoteUnit: "usd"}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err := client.Start([]Market{market})
		require.NoError(t, err)
		wg.Done()
	}()
	go func() {
		client.createTradeEvent(market)
		wg.Done()
	}()

	event := <-outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, event.Market, "btcusd")
}

func TestCreateTradeEvent(t *testing.T) {
	outbox := make(chan TradeEvent)
	client := bitfaker{outbox: outbox}

	go func() { client.createTradeEvent(Market{BaseUnit: "btc", QuoteUnit: "usd"}) }()

	event := <-outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, event.Market, "btcusd")
	assert.Equal(t, event.Source, DriverBitfaker)
	assert.Equal(t, event.Price, decimal.NewFromFloat(2.213))
}
