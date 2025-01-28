package driver

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

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/filter"
	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
	"github.com/layer-3/clearsync/pkg/safe"
)

type ODAPIMockMsg struct {
	m []byte
}

type ODDialerSuccessMock struct {
	reconnectAttempted chan bool
	messages           []ODAPIMockMsg
	isConnected        bool
}

func (dialer *ODDialerSuccessMock) Dial(_ string, _ http.Header) (common.WsTransport, *http.Response, error) {
	select {
	case dialer.reconnectAttempted <- true:
	default:
	}

	// successfully connected
	return &ODAPIMock{
		messages:    dialer.messages,
		isConnected: dialer.isConnected,
	}, &http.Response{}, nil
}

type ODAPIMock struct {
	messages    []ODAPIMockMsg
	isConnected bool
}

func (c *ODAPIMock) IsConnected() bool {
	return c.isConnected
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

func TestOpendax_parse(t *testing.T) {
	t.Parallel()

	t.Run("Successful test", func(t *testing.T) {
		t.Parallel()

		message := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{"btc/usd", 1, 1, 1, 1, 1, "buy", "Opendax"},
		}
		byteMsg, err := message.Encode()
		require.NoError(t, err)

		o := &opendax{}
		tradeEvent, err := o.parse(byteMsg)
		require.NoError(t, err)

		number, _ := decimal.NewFromString("1")
		expVal := &common.TradeEvent{
			Market:    common.NewMarket("btc", "usd"),
			Price:     number,
			Amount:    number,
			Total:     number,
			TakerType: common.TakerTypeBuy,
			CreatedAt: time.Unix(1, 0),
			Source:    common.DriverOpendax,
		}
		assert.Equal(t, expVal, tradeEvent)
	})

	t.Run("invalid message", func(t *testing.T) {
		t.Parallel()

		trade := ""
		byteMsg, err := json.Marshal(trade)
		require.NoError(t, err)

		o := &opendax{}
		_, err = o.parse(byteMsg)
		require.EqualError(t, errors.New("could not parse message: json: cannot unmarshal string into Go value of type []interface {}"), err.Error())
	})

	t.Run("invalid message args", func(t *testing.T) {
		t.Parallel()

		trade := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{""},
		}
		byteMsg, err := trade.Encode()
		require.NoError(t, err)

		o := &opendax{}
		_, err = o.parse(byteMsg)
		require.Error(t, err)
	})
}

func TestOpendax_Subscribe(t *testing.T) {
	t.Parallel()

	t.Run("Successful test", func(t *testing.T) {
		t.Parallel()

		client := &opendax{
			once: common.NewOnce(),
			conn: &ODAPIMock{
				messages:    []ODAPIMockMsg{{}},
				isConnected: true,
			},
			streams:        safe.NewMap[common.Market, struct{}](),
			symbolToMarket: safe.NewMap[string, common.Market](),
		}

		require.NoError(t, client.once.Start(func() error { return nil }))
		err := client.Subscribe(common.NewMarket("btc", "usdt"))
		require.NoError(t, err)
	})

	t.Run("Fail test", func(t *testing.T) {
		t.Parallel()

		client := &opendax{
			once: common.NewOnce(),
			conn: &ODAPIMock{isConnected: false},
		}

		require.NoError(t, client.once.Start(func() error { return nil }))
		err := client.Subscribe(common.NewMarket("btc", "usdt"))
		require.Error(t, err)
	})
}

func TestOpendax_Stop(t *testing.T) {
	t.Parallel()

	client := &opendax{once: common.NewOnce(), conn: &ODAPIMock{}}

	require.NoError(t, client.once.Start(func() error { return nil })) // unblock STOP action
	require.NoError(t, client.Stop())
}

func TestOpendax_connect(t *testing.T) {
	t.Parallel()

	t.Run("Successful case", func(t *testing.T) {
		t.Parallel()

		connMock := &ODDialerSuccessMock{
			reconnectAttempted: make(chan bool, 1),
			messages:           []ODAPIMockMsg{{}},
			isConnected:        true,
		}
		client := &opendax{
			once:   common.NewOnce(),
			conn:   nil,
			period: 0,
			dialer: connMock,
		}

		require.NoError(t, client.once.Start(func() error { return nil }))
		client.connect()
		require.True(t, <-connMock.reconnectAttempted)
		require.NotNil(t, client.conn)
		require.True(t, client.conn.IsConnected())
	})
}

func TestOpendax_listen(t *testing.T) {
	t.Parallel()

	disabledFilter := filter.New(filter.Config{Type: filter.TypeDisabled})

	t.Run("Error reading from connection", func(t *testing.T) {
		t.Parallel()

		// Setup an error message
		outbox := make(chan common.TradeEvent, 1)
		client := &opendax{
			once: common.NewOnce(),
			conn: &ODAPIMock{
				messages:    []ODAPIMockMsg{},
				isConnected: true,
			},
			dialer: &ODDialerSuccessMock{},
			outbox: outbox,
			period: 0,
		}

		require.NoError(t, client.once.Start(func() error { return nil }))
		go client.listen()

		// Allow some time for the goroutine to run
		time.Sleep(1 * time.Second)

		select {
		case tradeEvent := <-outbox:
			// The channel should not receive any message as ReadMessage has failed
			require.Nil(t, tradeEvent)
		default:
			// Test passes if no trade event message is received
		}
	})

	t.Run("Successful test", func(t *testing.T) {
		t.Parallel()

		update := protocol.Msg{
			ReqID:  1,
			Type:   3,
			Method: "trade",
			Args:   []interface{}{"btc/usd", 1, 1, 1, 1, 1, "buy", "Opendax"},
		}

		rawMsg, err := update.Encode()
		require.NoError(t, err)

		outbox := make(chan common.TradeEvent, 1)
		client := &opendax{
			once: common.NewOnce(),
			conn: &ODAPIMock{
				messages:    []ODAPIMockMsg{{m: rawMsg}},
				isConnected: true,
			},
			dialer: &ODDialerSuccessMock{},
			period: 0,
			outbox: outbox,
			filter: disabledFilter,
		}

		require.NoError(t, client.once.Start(func() error { return nil }))
		go client.listen()

		if tradeEvent := <-outbox; true {
			require.NotNil(t, tradeEvent)
		}
	})

	t.Run("Should return if connection is nil", func(t *testing.T) {
		t.Parallel()

		client := &opendax{once: common.NewOnce(), conn: nil}
		require.NoError(t, client.once.Start(func() error { return nil }))
		require.NotPanics(t, func() { client.listen() })
	})
}
