package quotes

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"

	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
)

type opendax struct {
	connMu      sync.RWMutex
	conn        wsTransport
	dialer      wsDialer
	isConnected atomic.Bool
	url         string

	outbox  chan<- TradeEvent
	period  time.Duration
	reqID   atomic.Uint64
	streams sync.Map
}

func newOpendax(config Config, outbox chan<- TradeEvent) *opendax {
	url := "wss://alpha.yellow.org/api/v1/finex/ws"
	if config.URL != "" {
		url = config.URL
	}

	return &opendax{
		url:    url,
		outbox: outbox,
		period: config.ReconnectPeriod * time.Second,
		reqID:  atomic.Uint64{},
		dialer: wsDialWrapper{},
	}
}

func (o *opendax) Subscribe(market Market) error {
	if _, ok := o.streams.Load(market); ok {
		return fmt.Errorf("market %s already subscribed", market)
	}

	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewSubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeOpendaxMsg(message)
	if err != nil {
		return err
	}

	o.streams.Store(market, struct{}{})
	return nil
}

func (o *opendax) Unsubscribe(market Market) error {
	if _, ok := o.streams.Load(market); !ok {
		return fmt.Errorf("market %s not subscribed", market)
	}

	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewUnsubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeOpendaxMsg(message)
	if err != nil {
		return err
	}

	o.streams.Delete(market)
	return nil
}

func (o *opendax) Start() error {
	o.connect()
	go o.readOpendaxMsg()
	return nil
}

func (o *opendax) Stop() error {
	o.connMu.Lock()
	defer o.connMu.Unlock()

	if o.conn != nil {
		o.conn = nil
	}

	o.isConnected.Store(false)
	return nil
}

func (o *opendax) connect() {
	for {
		wsConn, _, err := o.dialer.Dial(o.url, nil)
		o.connMu.Lock()
		o.conn = wsConn
		o.connMu.Unlock()
		o.isConnected.Store(err == nil)

		if err == nil {
			return
		} else {
			logger.Warnf("Websocket.Dial: can't connect to Opendax, reason: %s", err.Error())
		}

		time.Sleep(o.period)
	}
}

func (o *opendax) writeOpendaxMsg(message *protocol.Msg) error {
	byteMsg, err := message.Encode()
	if err != nil {
		logger.Warn(err)
		return err
	}

	o.connMu.RLock()
	conn := o.conn
	o.connMu.RUnlock()

	if err := conn.WriteMessage(websocket.TextMessage, byteMsg); err != nil {
		logger.Warn(err)
		return err
	}

	if _, _, err := conn.ReadMessage(); err != nil {
		logger.Warn(err)
		return err
	}
	return nil
}

func (o *opendax) readOpendaxMsg() {
	for {
		if !o.isConnected.Load() {
			continue
		}

		o.connMu.RLock()
		conn := o.conn
		o.connMu.RUnlock()

		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Warn("Error reading from connection", err)

			o.connect()
			o.streams.Range(func(m, value any) bool {
				market := m.(Market)
				if err := o.Subscribe(market); err != nil {
					logger.Warnf("Error subscribing to market %s: %s", market, err)
					return false
				}
				return true
			})

			continue
		}

		trEvent, err := parseOpendaxMsg(message)
		if err != nil {
			logger.Warn(err)
			continue
		}

		// Skip system messages
		if trEvent.Market == "" || trEvent.Price == decimal.Zero {
			continue
		}
		o.outbox <- *trEvent
	}
}

func parseOpendaxMsg(message []byte) (*TradeEvent, error) {
	if message == nil {
		return &TradeEvent{}, nil
	}
	parsedMsg, err := protocol.ParseRaw(message)
	if err != nil {
		return nil, err
	}

	tradeEvent, err := convertToTrade(parsedMsg.Args)
	if errors.Is(err, errors.New(protocol.ParseError)) || parsedMsg.Method == protocol.MethodSubscribe || parsedMsg.
		Method == protocol.EventSystem {
		return &TradeEvent{}, nil
	}
	if err != nil {
		return nil, err
	}

	return tradeEvent, nil
}

func convertToTrade(args []interface{}) (*TradeEvent, error) {
	it := protocol.NewArgIterator(args)

	market := it.NextString()
	_ = it.NextUint64() // trade id
	price := it.NextDecimal()
	amount := it.NextDecimal()
	total := it.NextDecimal()
	ts := it.NextTimestamp()
	tSide := it.NextString()
	takerSide, err := recognizeSide(tSide)
	if err != nil {
		return nil, err
	}
	_ = it.NextString() // trade source

	return &TradeEvent{
		Source:    DriverOpendax,
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     total,
		CreatedAt: time.Unix(ts, 0),
		TakerType: takerSide,
	}, it.Err()
}

func recognizeSide(side string) (TakerType, error) {
	switch side {
	case protocol.OrderSideSell:
		return TakerTypeSell, nil
	case protocol.OrderSideBuy:
		return TakerTypeBuy, nil
	default:
		return TakerTypeUnknown, errors.New("order side invalid: " + side)
	}
}
