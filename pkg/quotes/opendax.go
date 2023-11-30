package quotes

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"

	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
)

type Opendax struct {
	conn        WSTransport
	dialer      WSDialer
	url         string
	outbox      chan<- TradeEvent
	mu          sync.RWMutex
	period      time.Duration
	reqId       uint64
	isConnected bool
}

type TradeResponse struct {
	ID        uint64
	Market    string
	Price     decimal.Decimal
	Amount    decimal.Decimal
	TakerType TakerType
	CreatedAt int64
}

func NewOpendax(config Config, outbox chan<- TradeEvent) *Opendax {
  url := "wss://alpha.yellow.org/api/v1/finex/ws"
  if config.URL != "" {
    url = config.URL
  }

	return &Opendax{
		url:    url,
		outbox: outbox,
		period: time.Duration(config.ReconnectPeriod) * time.Second,
		reqId:  1,
		dialer: WSDialWrapper{},
	}
}

func (o *Opendax) Connect() {
	for {
		wsConn, _, err := o.dialer.Dial(o.url, nil)
		o.mu.Lock()
		o.conn = wsConn
		o.isConnected = err == nil
		o.mu.Unlock()

		if err == nil {
			return
		} else {
			logger.Warnf("Websocket.Dial: can't connect to Opendax, reason: %s", err.Error())
		}

		time.Sleep(o.period)
	}
}

func (o *Opendax) Subscribe(market Market) error {
	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.BaseUnit, market.QuoteUnit)
	message := protocol.NewSubscribeMessage(o.reqId, resource)
	o.reqId++

	byteMsg, err := message.Encode()
	if err != nil {
		logger.Warn(err)
		return err
	}

	o.reqId++

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

func (o *Opendax) Start(markets []Market) error {
	if len(markets) == 0 {
		return errors.New("no markets specified")
	}

	o.Connect()
	o.marketSubscribe(markets)
	o.receiveOpendaxMsg(markets)
	return nil
}

func (o *Opendax) Stop() error {
	o.mu.Lock()
	if o.conn != nil {
		o.conn = nil
	}

	o.isConnected = false
	o.mu.Unlock()
	return nil
}

func (o *Opendax) marketSubscribe(markets []Market) {
	for _, m := range markets {
		symbol := m.BaseUnit + m.QuoteUnit
		if err := o.Subscribe(m); err != nil {
			logger.Infof("market %s doesn't exist in Opendax", symbol)
		}
		logger.Infof("quotes service connected to Opendax %s market", symbol)
	}
}

func (o *Opendax) receiveOpendaxMsg(markets []Market) {
	for {
		if o.isConnected {
			_, message, err := o.conn.ReadMessage()
			if err != nil {
				logger.Warn("Error reading from connection", err)
				o.Connect()
				o.marketSubscribe(markets)
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
