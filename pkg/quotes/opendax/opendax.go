package opendax

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/opendax/protocol"
)

var logger = log.Logger("opendax")

type Opendax struct {
	once   *common.Once
	conn   common.WsTransport
	dialer common.WsDialer
	url    string

	outbox  chan<- common.TradeEvent
	period  time.Duration
	reqID   atomic.Uint64
	streams sync.Map
}

func NewOpendax(config common.Config, outbox chan<- common.TradeEvent) *Opendax {
	url := "wss://alpha.yellow.org/api/v1/finex/ws"
	if config.URL != "" {
		url = config.URL
	}

	return &Opendax{
		once:   common.NewOnce(),
		url:    url,
		outbox: outbox,
		period: config.ReconnectPeriod * time.Second,
		reqID:  atomic.Uint64{},
		dialer: common.WsDialWrapper{},
	}
}

func (o *Opendax) Start() error {
	o.once.Start(func() {
		o.connect()
		go o.listen()
	})
	return nil
}

func (o *Opendax) Stop() error {
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

func (o *Opendax) Subscribe(market common.Market) error {
	if _, ok := o.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
	}

	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewSubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, common.ErrFailedSub, err)
	}

	o.streams.Store(market, struct{}{})
	return nil
}

func (o *Opendax) Unsubscribe(market common.Market) error {
	if _, ok := o.streams.Load(market); !ok {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
	}

	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewUnsubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, common.ErrFailedUnsub, err)
	}

	o.streams.Delete(market)
	return nil
}

func (o *Opendax) connect() {
	for {
		var err error
		o.conn, _, err = o.dialer.Dial(o.url, nil)
		if err == nil {
			return
		}

    logger.Warnf("connection attempt failed: %s", err.Error())
		time.Sleep(o.period)
	}
}

func (o *Opendax) writeConn(message *protocol.Msg) error {
	byteMsg, err := message.Encode()
	if err != nil {
		logger.Warn(err)
		return err
	}

	if err := o.conn.WriteMessage(websocket.TextMessage, byteMsg); err != nil {
		logger.Warn(err)
		return err
	}

	if _, _, err := o.conn.ReadMessage(); err != nil {
		logger.Warn(err)
		return err
	}
	return nil
}

func (o *Opendax) listen() {
	for {
		if o.conn == nil {
			return
		}
		if !o.conn.IsConnected() {
			continue
		}

		_, message, err := o.conn.ReadMessage()
		if err != nil {
			logger.Warn("error reading from connection", err)

			o.connect()
			o.streams.Range(func(m, value any) bool {
				market := m.(common.Market)
				if err := o.Subscribe(market); err != nil {
					logger.Warnf("error subscribing to market %s: %s", market, err)
					return false
				}
				return true
			})

			continue
		}

		trEvent, err := parse(message)
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

func parse(message []byte) (*common.TradeEvent, error) {
	if message == nil {
		return &common.TradeEvent{}, nil
	}
	parsedMsg, err := protocol.ParseRaw(message)
	if err != nil {
		return nil, err
	}

	tradeEvent, err := convertToTrade(parsedMsg.Args)
	if errors.Is(err, errors.New(protocol.ParseError)) || parsedMsg.Method == protocol.MethodSubscribe || parsedMsg.
		Method == protocol.EventSystem {
		return &common.TradeEvent{}, nil
	}
	if err != nil {
		return nil, err
	}

	return tradeEvent, nil
}

func convertToTrade(args []interface{}) (*common.TradeEvent, error) {
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

	return &common.TradeEvent{
		Source:    common.DriverOpendax,
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     total,
		CreatedAt: time.Unix(ts, 0),
		TakerType: takerSide,
	}, it.Err()
}

func recognizeSide(side string) (common.TakerType, error) {
	switch side {
	case protocol.OrderSideSell:
		return common.TakerTypeSell, nil
	case protocol.OrderSideBuy:
		return common.TakerTypeBuy, nil
	default:
		return common.TakerTypeUnknown, errors.New("order side invalid: " + side)
	}
}
