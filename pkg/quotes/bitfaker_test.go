package quotes

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/neodax/finex/models/market"
	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
)

func TestBitfaker_Subscribe_SingleMarket(t *testing.T) {
	ch := make(chan trade.Event)
	client := Bitfaker{outbox: ch}

	err := client.Subscribe("btc", "usd")
	require.Nil(t, err)

	expectedMarkets := make([]string, 0)
	expectedMarkets = append(expectedMarkets, "btcusd")
	assert.Equal(t, client.markets, expectedMarkets)
}

func TestNewBitfakerPriceSourceClient(t *testing.T) {
	ch := make(chan trade.Event)
	client := Bitfaker{outbox: ch}

	assert.Equal(t, client.outbox, ch)
}

func TestBitfaker_Start(t *testing.T) {
	ch := make(chan trade.Event)
	marketCache := &cache.TestMarketCache{}
	marketCache.
		On("GetActive").
		Return(map[cache.MarketKey]market.Market{}, nil)
	sampl := TradeSampler{enabled: false, defaultPercentage: 0.0}
	client := Bitfaker{outbox: ch, marketCache: marketCache, period: 0 * time.Second, tradeSampler: &sampl}

	err := client.Subscribe("btc", "usd")
	require.Nil(t, err)
	go func() { client.createTradeEvent(market.Market{Symbol: "btcusd"}) }()

	event := <-client.outbox
	go func() {
		err := client.Start()
		require.Nil(t, err)
	}()
	assert.NotEmpty(t, event)
	assert.Equal(t, event.Market, "btcusd")
}

func TestBitfaker_Subscribe_MultipleMarkets(t *testing.T) {
	ch := make(chan trade.Event)
	client := Bitfaker{outbox: ch}

	err := client.Subscribe("btc", "usd")
	require.Nil(t, err)
	err = client.Subscribe("eth", "usd")
	require.Nil(t, err)

	expectedMarkets := make([]string, 0)
	expectedMarkets = append(expectedMarkets, "btcusd")
	expectedMarkets = append(expectedMarkets, "ethusd")
	assert.Equal(t, client.markets, expectedMarkets)
}

func TestCreateTradeEvent(t *testing.T) {
	ch := make(chan trade.Event)
	client := Bitfaker{outbox: ch}

	go func() { client.createTradeEvent(market.Market{Symbol: "btcusd"}) }()

	event := <-client.outbox
	assert.NotEmpty(t, event)
	assert.Equal(t, event.Market, "btcusd")
	assert.Equal(t, event.Source, "Bitfaker")
	assert.Equal(t, event.Price, decimal.NewFromFloat(2.213))
}
