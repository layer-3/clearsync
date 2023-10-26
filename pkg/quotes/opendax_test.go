package quotes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/neodax/finex/models/market"
	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
	"github.com/layer-3/neodax/finex/pkg/event"
	"github.com/layer-3/neodax/finex/pkg/websocket/client"
	"github.com/layer-3/neodax/finex/pkg/websocket/protocol"
)

type ODAPIMockMsg struct {
	m []byte
}

type ODAPIMock struct {
	messages []ODAPIMockMsg
}

type ODMarketsMock struct{}

type ODDialerSuccessMock struct {
	connRetryAttempted chan bool
	m                  []ODAPIMockMsg
}

type ODDialerFailMock struct {
	connRetryAttempted chan bool
	m                  []ODAPIMockMsg
}

func (dialer *ODDialerSuccessMock) Dial(_ string, _ http.Header) (client.WSTransport, *http.Response, error) {
	select {
	case dialer.connRetryAttempted <- true:
	default:
	}

	return &ODAPIMock{messages: dialer.m}, &http.Response{}, nil
}

func (dialer *ODDialerFailMock) Dial(_ string, _ http.Header) (client.WSTransport, *http.Response, error) {
	select {
	case dialer.connRetryAttempted <- true: // return error on the first try
		return &ODAPIMock{messages: dialer.m}, &http.Response{}, errors.New("cannot connect")
	default: // return no error on consequent retries
		return &ODAPIMock{messages: dialer.m}, &http.Response{}, nil
	}
}

func (m *ODMarketsMock) Get(id, typ string) (market.Market, error) {
	return market.Market{}, nil
}

func (m *ODMarketsMock) GetActive() (map[cache.MarketKey]market.Market, error) {
	return map[cache.MarketKey]market.Market{
		{ID: "btcusd", Type: "spot"}: {Symbol: "btcusd", BaseUnit: "btc", QuoteUnit: "usd"},
		{ID: "ethusd", Type: "spot"}: {Symbol: "ethusd", BaseUnit: "eth", QuoteUnit: "usd"},
	}, nil
}

func (m *ODMarketsMock) GetActiveWithFormat() ([]interface{}, error) {
	return nil, errors.New("cannot get active markets")
}

func (m *ODMarketsMock) SubscribeUpdates() chan cache.MarketUpdateNotification {
	return make(chan cache.MarketUpdateNotification, 1)
}

func (c *ODAPIMock) ReadMessage() (messageType int, p []byte, err error) {
	if len(c.messages) < 1 {
		return 0, []byte{}, fmt.Errorf("connection closed")
	}

	m := c.messages[0]
	if len(c.messages) == 1 {
		c.messages = []ODAPIMockMsg{}
	} else {
		c.messages = c.messages[1:]
	}
	return 0, m.m, nil
}

func (c *ODAPIMock) WriteMessage(messageType int, data []byte) error {
	return nil
}

func (c *ODAPIMock) Close() error {
	return nil
}

func TestOpendax_parseOpendaxMsg(t *testing.T) {
	opendax := &Opendax{}

	t.Run("Successful test", func(t *testing.T) {
		message := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{"btcusd", 1, 1, 1, 1, 1, "buy", "Opendax"},
		}
		byteMsg, err := message.Encode()
		require.NoError(t, err)

		tradeEvent, err := opendax.parseOpendaxMsg(byteMsg)
		require.NoError(t, err)

		number, _ := decimal.NewFromString("1")
		expVal := &trade.Event{
			ID:        1,
			Market:    "btcusd",
			Price:     number,
			Amount:    number,
			Total:     number,
			TakerType: "OrderBid",
			CreatedAt: 1,
			Source:    "Opendax",
		}
		assert.Equal(t, expVal, tradeEvent)
	})

	t.Run("invalid message", func(t *testing.T) {
		trade := ""
		byteMsg, err := json.Marshal(trade)
		require.NoError(t, err)

		_, err = opendax.parseOpendaxMsg(byteMsg)
		require.EqualError(t, errors.New("could not parse message: json: cannot unmarshal string into Go value of type []interface {}"), err.Error())
	})

	t.Run("invalid message args", func(t *testing.T) {
		trade := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{""},
		}
		byteMsg, err := trade.Encode()
		require.NoError(t, err)

		_, err = opendax.parseOpendaxMsg(byteMsg)
		require.Error(t, err)
	})
}

func TestOpendax_marketSubscribe(t *testing.T) {
	opendax := &Opendax{
		marketCache: &ODMarketsMock{},
		conn:        &ODAPIMock{},
	}
	marketList, err := opendax.marketCache.GetActive()
	require.NoError(t, err)
	opendax.marketSubscribe(marketList)
}

func TestOpendax_Subscribe(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		client := &Opendax{
			conn: &ODAPIMock{messages: []ODAPIMockMsg{{}}},
		}

		err := client.Subscribe("btc", "usdt")
		require.NoError(t, err)
	})

	t.Run("Fail test", func(t *testing.T) {
		client := &Opendax{
			conn: &ODAPIMock{messages: []ODAPIMockMsg{}},
		}

		err := client.Subscribe("btc", "usdt")
		require.Error(t, err)
	})
}

func TestOpendax_Close(t *testing.T) {
	client := Opendax{conn: &ODAPIMock{}}
	err := client.Close()
	require.NoError(t, err)
}

func TestOpendax_IsConnected(t *testing.T) {
	client := &Opendax{
		mu:          sync.RWMutex{},
		isConnected: true,
	}
	require.True(t, client.IsConnected())
}

func TestOpendax_Connect(t *testing.T) {
	t.Run("Successful case", func(t *testing.T) {
		connMock := &ODDialerSuccessMock{connRetryAttempted: make(chan bool, 1)}
		client := Opendax{
			conn:        nil,
			mu:          sync.RWMutex{},
			period:      0,
			isConnected: false,
			dialer:      connMock,
		}

		client.Connect()
		require.True(t, <-connMock.connRetryAttempted)
		require.NotNil(t, client.conn)
		require.True(t, client.isConnected)
	})

	t.Run("Fail case", func(t *testing.T) {
		connMock := &ODDialerFailMock{connRetryAttempted: make(chan bool, 1)}
		client := Opendax{
			conn:        nil,
			mu:          sync.RWMutex{},
			period:      0,
			isConnected: false,
			dialer:      connMock,
		}

		client.Connect()
		require.True(t, <-connMock.connRetryAttempted)
		require.NotNil(t, client.conn)
		require.True(t, client.isConnected)
	})
}

// FIXME: update test
// func TestOpendax_Start(t *testing.T) {
// 	t.Run("No active markets error", func(t *testing.T) {
// 		client := &Opendax{marketCache: &MarketsMockError{}}
// 		err := client.Start()
// 		require.Error(t, err)
// 	})
// }

func TestOpendax_receiveOpendaxMsg(t *testing.T) {
	mock := ODMarketsMock{}
	activeMarkets, _ := mock.GetActive()

	t.Run("Error reading from connection", func(t *testing.T) {
		dialer := ODDialerFailMock{connRetryAttempted: make(chan bool, 1)}
		client := &Opendax{
			conn:        &ODAPIMock{messages: []ODAPIMockMsg{}},
			dialer:      &dialer,
			isConnected: true,
			period:      0,
		}

		go func() {
			client.receiveOpendaxMsg(activeMarkets)
		}()

		// the function will try to reestablish the connection,
		// so the number of retries can be measured
		for {
			select {
			case retryAttempted := <-dialer.connRetryAttempted:
				require.True(t, retryAttempted)
				return
			default:
				continue
			}
		}
	})

	t.Run("Successful test", func(t *testing.T) {
		update := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{"btcusd", 1, 1, 1, 1, 1, "buy", "Opendax"},
		}

		rawMsg, err := update.Encode()
		require.NoError(t, err)

		client := &Opendax{
			conn: &ODAPIMock{
				messages: []ODAPIMockMsg{{m: rawMsg}},
			},
			dialer:      &ODDialerSuccessMock{},
			isConnected: true,
			period:      0,
			outbox:      make(chan trade.Event, 1),
			output:      make(chan<- event.Event, 1),
		}

		go func() {
			client.receiveOpendaxMsg(activeMarkets)
		}()

		for {
			select {
			case tradeEvent := <-client.outbox:
				require.NotNil(t, tradeEvent)
				return
			default:
				continue
			}
		}
	})
}

// FIXME: update test
// func TestOpendax_Init(t *testing.T) {
// 	dialer := &ODDialerSuccessMock{connRetryAttempted: make(chan bool, 1)}
// 	marketCache := &cache.MarketCache{}
// 	outbox := make(chan trade.Event)
// 	output := make(chan<- event.Event)
// 	quoteFeed := config.QuoteFeed{
// 		URL:    "whatever",
// 		Period: 0,
// 	}
//
// 	client := &Opendax{
// 		conn:        nil,
// 		url:         "",
// 		mu:          sync.RWMutex{},
// 		period:      0,
// 		reqId:       0,
// 		isConnected: false,
// 	}
//
// 	err := client.Init(marketCache, outbox, output, quoteFeed, dialer)
// 	require.NoError(t, err)
// 	require.True(t, <-dialer.connRetryAttempted)
//
// 	require.Equal(t, quoteFeed.URL, client.url)
// 	require.Equal(t, outbox, client.outbox)
// 	require.Equal(t, output, client.output)
// 	require.Equal(t, marketCache, client.marketCache)
// 	require.Equal(t, time.Duration(0), client.period)
// 	require.Equal(t, uint64(1), client.reqId)
// 	require.Equal(t, dialer, client.dialer)
// }
