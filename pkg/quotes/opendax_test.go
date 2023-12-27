package quotes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
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

func (dialer *ODDialerSuccessMock) Dial(_ string, _ http.Header) (wsTransport, *http.Response, error) {
	select {
	case dialer.connRetryAttempted <- true:
	default:
	}

	return &ODAPIMock{messages: dialer.m}, &http.Response{}, nil
}

func (dialer *ODDialerFailMock) Dial(_ string, _ http.Header) (wsTransport, *http.Response, error) {
	select {
	case dialer.connRetryAttempted <- true: // return error on the first try
		return &ODAPIMock{messages: dialer.m}, &http.Response{}, errors.New("cannot connect")
	default: // return no error on consequent retries
		return &ODAPIMock{messages: dialer.m}, &http.Response{}, nil
	}
}

func (m *ODMarketsMock) Get(id, typ string) (Market, error) {
	return Market{}, nil
}

func (m *ODMarketsMock) GetActive() ([]Market, error) {
	return []Market{
		{BaseUnit: "btc", QuoteUnit: "usd"},
		{BaseUnit: "eth", QuoteUnit: "usd"},
	}, nil
}

func (m *ODMarketsMock) GetActiveWithFormat() ([]interface{}, error) {
	return nil, errors.New("cannot get active markets")
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
	t.Run("Successful test", func(t *testing.T) {
		message := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{"btcusd", 1, 1, 1, 1, 1, "buy", "Opendax"},
		}
		byteMsg, err := message.Encode()
		require.NoError(t, err)

		tradeEvent, err := parseOpendaxMsg(byteMsg)
		require.NoError(t, err)

		number, _ := decimal.NewFromString("1")
		expVal := &TradeEvent{
			Market:    "btcusd",
			Price:     number,
			Amount:    number,
			Total:     number,
			TakerType: TakerTypeBuy,
			CreatedAt: time.Unix(1, 0),
			Source:    DriverOpendax,
		}
		assert.Equal(t, expVal, tradeEvent)
	})

	t.Run("invalid message", func(t *testing.T) {
		trade := ""
		byteMsg, err := json.Marshal(trade)
		require.NoError(t, err)

		_, err = parseOpendaxMsg(byteMsg)
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

		_, err = parseOpendaxMsg(byteMsg)
		require.Error(t, err)
	})
}

func TestOpendax_marketSubscribe(t *testing.T) {
	opendax := &opendax{conn: &ODAPIMock{}}
	markets := []Market{{BaseUnit: "btc", QuoteUnit: "usdt"}}
	opendax.marketSubscribe(markets)
}

func TestOpendax_Subscribe(t *testing.T) {
	t.Run("Successful test", func(t *testing.T) {
		client := &opendax{
			conn: &ODAPIMock{messages: []ODAPIMockMsg{{}}},
		}

		err := client.Subscribe(Market{BaseUnit: "btc", QuoteUnit: "usdt"})
		require.NoError(t, err)
	})

	t.Run("Fail test", func(t *testing.T) {
		client := &opendax{
			conn: &ODAPIMock{messages: []ODAPIMockMsg{}},
		}

		err := client.Subscribe(Market{BaseUnit: "btc", QuoteUnit: "usdt"})
		require.Error(t, err)
	})
}

func TestOpendax_Stop(t *testing.T) {
	client := opendax{conn: &ODAPIMock{}}
	err := client.Stop()
	require.NoError(t, err)
}

func TestOpendax_Connect(t *testing.T) {
	t.Run("Successful case", func(t *testing.T) {
		connMock := &ODDialerSuccessMock{connRetryAttempted: make(chan bool, 1)}
		client := opendax{
			conn:   nil,
			period: 0,
			dialer: connMock,
		}
		client.isConnected.Store(false)

		client.connect()
		require.True(t, <-connMock.connRetryAttempted)
		require.NotNil(t, client.conn)
		require.True(t, client.isConnected.Load())
	})

	t.Run("Fail case", func(t *testing.T) {
		connMock := &ODDialerFailMock{connRetryAttempted: make(chan bool, 1)}
		client := opendax{
			conn:   nil,
			period: 0,
			dialer: connMock,
		}
		client.isConnected.Store(false)

		client.connect()
		require.True(t, <-connMock.connRetryAttempted)
		require.NotNil(t, client.conn)
		require.True(t, client.isConnected.Load())
	})
}

func TestOpendax_Start(t *testing.T) {
	t.Run("No active markets error", func(t *testing.T) {
		client := &opendax{}
		err := client.Start([]Market{})
		require.Error(t, err)
	})
}

func TestOpendax_receiveOpendaxMsg(t *testing.T) {
	mock := ODMarketsMock{}
	activeMarkets, _ := mock.GetActive()

	t.Run("Error reading from connection", func(t *testing.T) {
		dialer := ODDialerFailMock{connRetryAttempted: make(chan bool, 1)}
		client := &opendax{
			conn:   &ODAPIMock{messages: []ODAPIMockMsg{}},
			dialer: &dialer,
			period: 0,
		}
		client.isConnected.Store(true)

		go func() {
			client.readOpendaxMsg(activeMarkets)
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

		outbox := make(chan TradeEvent, 1)
		client := &opendax{
			conn: &ODAPIMock{
				messages: []ODAPIMockMsg{{m: rawMsg}},
			},
			dialer: &ODDialerSuccessMock{},
			period: 0,
			outbox: outbox,
		}
		client.isConnected.Store(true)

		go func() {
			client.readOpendaxMsg(activeMarkets)
		}()

		for {
			select {
			case tradeEvent := <-outbox:
				require.NotNil(t, tradeEvent)
				return
			default:
				continue
			}
		}
	})
}
