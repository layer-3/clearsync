package quotes

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
)

var loggerOpendax = log.Logger("opendax")

type opendax struct {
	once   *once
	conn   wsTransport
	dialer wsDialer
	url    string

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
		once:   newOnce(),
		url:    url,
		outbox: outbox,
		period: config.ReconnectPeriod * time.Second,
		reqID:  atomic.Uint64{},
		dialer: wsDialWrapper{},
	}
}

func (o *opendax) Start() error {
	o.once.Start(func() {
		o.connect()
		go o.listen()
	})
	return nil
}

func (o *opendax) Stop() error {
	var stopErr error
	o.once.Stop(func() {
		conn := o.conn
		o.conn = nil

		if conn == nil {
			return
		}
		stopErr = conn.Close()
	})
	return stopErr
}

func (o *opendax) Subscribe(market Market) error {
	if _, ok := o.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewSubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, errFailedSub, err)
	}

	o.streams.Store(market, struct{}{})
	return nil
}

func (o *opendax) Unsubscribe(market Market) error {
	if _, ok := o.streams.Load(market); !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewUnsubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, errFailedUnsub, err)
	}

	o.streams.Delete(market)
	return nil
}

func (o *opendax) connect() {
	for {
		var err error
		o.conn, _, err = o.dialer.Dial(o.url, nil)
		if err == nil {
			return
		}

		loggerBinance.Warnf("connection attempt failed: %s", err.Error())
		time.Sleep(o.period)
	}
}

func (o *opendax) writeConn(message *protocol.Msg) error {
	byteMsg, err := message.Encode()
	if err != nil {
		loggerBinance.Warn(err)
		return err
	}

	if err := o.conn.WriteMessage(websocket.TextMessage, byteMsg); err != nil {
		loggerBinance.Warn(err)
		return err
	}

	if _, _, err := o.conn.ReadMessage(); err != nil {
		loggerBinance.Warn(err)
		return err
	}
	return nil
}

func (o *opendax) listen() {
	for {
		if o.conn == nil {
			return
		}
		if !o.conn.IsConnected() {
			continue
		}

		_, message, err := o.conn.ReadMessage()
		if err != nil {
			loggerBinance.Warn("error reading from connection", err)

			o.connect()
			o.streams.Range(func(m, value any) bool {
				market := m.(Market)
				if err := o.Subscribe(market); err != nil {
					loggerBinance.Warnf("error subscribing to market %s: %s", market, err)
					return false
				}
				return true
			})

			continue
		}

		trEvent, err := o.parse(message)
		if err != nil {
			loggerBinance.Warn(err)
			continue
		}

		// Skip system messages
		if trEvent.Market == "" || trEvent.Price == decimal.Zero {
			continue
		}
		o.outbox <- *trEvent
	}
}

func (o *opendax) parse(message []byte) (*TradeEvent, error) {
	if message == nil {
		return &TradeEvent{}, nil
	}
	parsedMsg, err := protocol.ParseRaw(message)
	if err != nil {
		return nil, err
	}

	tradeEvent, err := o.convertToTrade(parsedMsg.Args)
	if errors.Is(err, errors.New(protocol.ParseError)) || parsedMsg.Method == protocol.MethodSubscribe || parsedMsg.
		Method == protocol.EventSystem {
		return &TradeEvent{}, nil
	}
	if err != nil {
		return nil, err
	}

	return tradeEvent, nil
}

func (o *opendax) convertToTrade(args []interface{}) (*TradeEvent, error) {
	it := protocol.NewArgIterator(args)

	market := it.NextString()
	_ = it.NextUint64() // trade id
	price := it.NextDecimal()
	amount := it.NextDecimal()
	total := it.NextDecimal()
	ts := it.NextTimestamp()
	tSide := it.NextString()
	takerSide, err := o.recognizeSide(tSide)
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

func (o *opendax) recognizeSide(side string) (TakerType, error) {
	switch side {
	case protocol.OrderSideSell:
		return TakerTypeSell, nil
	case protocol.OrderSideBuy:
		return TakerTypeBuy, nil
	default:
		return TakerTypeUnknown, errors.New("order side invalid: " + side)
	}
}
