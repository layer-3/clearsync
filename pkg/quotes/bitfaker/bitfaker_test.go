package bitfaker

import (
	"sync"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

  "github.com/layer-3/clearsync/pkg/quotes/common"
)

func TestBitfaker_Subscribe(t *testing.T) {
	t.Parallel()

	t.Run("Single market", func(t *testing.T) {
		t.Parallel()

		ch := make(chan common.TradeEvent, 16)
		client := Bitfaker{outbox: ch}

		m := common.Market{BaseUnit: "btc", QuoteUnit: "usd"}
		err := client.Subscribe(m)
		require.Nil(t, err)

		expectedMarkets := []common.Market{m}
		assert.Equal(t, client.streams, expectedMarkets)
	})

	t.Run("Multiple markets", func(t *testing.T) {
		t.Parallel()

		outbox := make(chan common.TradeEvent, 16)
		client := Bitfaker{outbox: outbox}

		market1 := common.Market{BaseUnit: "btc", QuoteUnit: "usd"}
		err := client.Subscribe(market1)
		require.Nil(t, err)

		market2 := common.Market{BaseUnit: "eth", QuoteUnit: "usd"}
		err = client.Subscribe(market2)
		require.Nil(t, err)

		expectedMarkets := []common.Market{market1, market2}
		assert.Equal(t, client.streams, expectedMarkets)
	})

	t.Run("Subscribe to a market already subscribed to", func(t *testing.T) {
		t.Parallel()

		ch := make(chan common.TradeEvent, 16)
		client := Bitfaker{outbox: ch}

		market := common.Market{BaseUnit: "btc", QuoteUnit: "usd"}
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

		ch := make(chan common.TradeEvent, 16)
		client := Bitfaker{outbox: ch}

		market1 := common.Market{BaseUnit: "btc", QuoteUnit: "usd"}
		market2 := common.Market{BaseUnit: "eth", QuoteUnit: "usd"}
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))
		require.NoError(t, client.Unsubscribe(market2))

		assert.NotContains(t, client.streams, market1)
		assert.NotContains(t, client.streams, market2)
	})

	t.Run("Unsubscribe from a market not subscribed to", func(t *testing.T) {
		t.Parallel()

		ch := make(chan common.TradeEvent, 16)
		client := Bitfaker{outbox: ch}

		market := common.Market{BaseUnit: "xrp", QuoteUnit: "usd"}
		err := client.Unsubscribe(market)

		require.Error(t, err)
	})

	t.Run("No effect on other subscriptions after unsubscribing", func(t *testing.T) {
		t.Parallel()

		ch := make(chan common.TradeEvent, 16)
		client := Bitfaker{outbox: ch}

		market1 := common.Market{BaseUnit: "btc", QuoteUnit: "usd"}
		market2 := common.Market{BaseUnit: "eth", QuoteUnit: "usd"}
		require.NoError(t, client.Subscribe(market1))
		require.NoError(t, client.Subscribe(market2))

		require.NoError(t, client.Unsubscribe(market1))

		assert.NotContains(t, client.streams, market1)
		assert.Contains(t, client.streams, market2)
	})
}

func TestBitfaker_Start(t *testing.T) {
	t.Parallel()

	outbox := make(chan common.TradeEvent, 1)
	tradeSampler := *common.NewTradeSampler(common.TradeSamplerConfig{
		Enabled:           false,
		DefaultPercentage: 0,
	})
	client := Bitfaker{
		once:         common.NewOnce(),
		outbox:       outbox,
		period:       0 * time.Second,
		tradeSampler: tradeSampler,
	}
	market := common.Market{BaseUnit: "btc", QuoteUnit: "usd"}

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
	assert.Equal(t, event.Market, "btcusd")
}

func TestCreateTradeEvent(t *testing.T) {
	t.Parallel()

	outbox := make(chan common.TradeEvent)
	client := Bitfaker{outbox: outbox}

	go func() { client.createTradeEvent(common.Market{BaseUnit: "btc", QuoteUnit: "usd"}) }()

	event := <-outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, event.Market, "btcusd")
	assert.Equal(t, event.Source, common.DriverBitfaker)
	assert.Equal(t, event.Price, decimal.NewFromFloat(2.213))
}
