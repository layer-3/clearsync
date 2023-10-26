package quotes

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"

	"github.com/layer-3/neodax/finex/models/market"
	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
	"github.com/layer-3/neodax/finex/pkg/config"
	"github.com/layer-3/neodax/finex/pkg/event"
	"github.com/layer-3/neodax/finex/pkg/websocket/client"
	"github.com/layer-3/neodax/finex/pkg/websocket/protocol"
)

type Opendax struct {
	conn        client.WSTransport
	dialer      client.WSDialer
	url         string
	marketCache cache.Market
	outbox      chan trade.Event
	output      chan<- event.Event
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
	TakerType trade.TakerType
	CreatedAt int64
}

func (o *Opendax) Init(markets cache.Market, outbox chan trade.Event, output chan<- event.Event, config config.QuoteFeed, dialer client.WSDialer) error {
	o.url = config.URL
	o.outbox = outbox
	o.output = output
	o.marketCache = markets
	o.period = time.Duration(config.Period) * time.Second
	o.reqId = 1
	o.dialer = dialer
	return nil
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
			logger.Warnf("Websocket.Dial: can't connect to opendax, reason: %s", err.Error())
		}

		time.Sleep(o.period)
	}
}

func (o *Opendax) IsConnected() bool {
	o.mu.Lock()
	defer o.mu.Unlock()

	return o.isConnected
}

func (o *Opendax) Subscribe(base, quote string) error {
	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", base, quote)

	message := client.SubscribePublic(o.reqId, resource)
	o.reqId++

	byteMsg, err := message.Encode()
	if err != nil {
		logger.Warn(err)
		return err
	}

	err = o.conn.WriteMessage(websocket.TextMessage, byteMsg)
	if err != nil {
		logger.Warn(err)
		return err
	}

	_, _, err = o.conn.ReadMessage()
	if err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}

func (o *Opendax) marketSubscribe(marketList map[cache.MarketKey]market.Market) {
	for _, marketModel := range marketList {

		err := o.Subscribe(marketModel.BaseUnit, marketModel.QuoteUnit)
		if err != nil {
			logger.Infof("market %s doesn't exist in opendax", marketModel.Symbol)
		}

		logger.Infof("quotes service connected to Opendax %s market", marketModel.Symbol)
	}
}
func (o *Opendax) Start() error {
	o.Connect()

	marketList, err := o.marketCache.GetActive()
	if err != nil {
		logger.Warn(err.Error())
		return err
	}

	o.marketSubscribe(marketList)
	o.receiveOpendaxMsg(marketList)
	return nil
}

func (o *Opendax) receiveOpendaxMsg(marketList map[cache.MarketKey]market.Market) {
	for {
		if o.IsConnected() {
			_, message, err := o.conn.ReadMessage()
			if err != nil {
				logger.Warn("Error reading from connection", err)
				o.Connect()
				o.marketSubscribe(marketList)
				continue
			}

			trEvent, err := o.parseOpendaxMsg(message)
			if err != nil {
				logger.Warn(err)
				continue
			}

			// Skip system messages
			if trEvent.Market == "" || trEvent.Price == decimal.Zero {
				continue
			}

			event, err := GetRoutingEvent(*trEvent)
			if err != nil {
				logger.Warn(err)
				continue
			}
			o.outbox <- *trEvent
			o.output <- *event
		}
	}
}

func (o *Opendax) parseOpendaxMsg(message []byte) (*trade.Event, error) {
	if message == nil {
		return &trade.Event{}, nil
	}
	parsedMsg, err := protocol.Parse(message)
	if err != nil {
		return nil, err
	}

	tr, err := client.Parse(parsedMsg)
	if err == errors.New(client.ParseError) || parsedMsg.Method == protocol.MethodSubscribe || parsedMsg.Method == protocol.EventSystem {
		return &trade.Event{}, nil
	}
	if err != nil {
		return nil, err
	}

	tradeEvent, ok := tr.(*trade.Event)
	if !ok {
		return nil, errors.New("error parsing opendax message")
	}

	return tradeEvent, nil
}

func (o *Opendax) Close() error {
	o.mu.Lock()
	if o.conn != nil {
		o.conn = nil
	}

	o.isConnected = false
	o.mu.Unlock()
	return nil
}
