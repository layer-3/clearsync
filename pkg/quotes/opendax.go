package quotes

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ipfs/go-log/v2"

	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerOpendax = log.Logger("opendax")

type opendax struct {
	once   *once
	conn   wsTransport
	dialer wsDialer
	url    string

	outbox         chan<- TradeEvent
	filter         Filter
	history        HistoricalData
	period         time.Duration
	reqID          atomic.Uint64
	streams        safe.Map[Market, struct{}]
	symbolToMarket safe.Map[string, Market]
}

func newOpendax(config OpendaxConfig, outbox chan<- TradeEvent, history HistoricalData) (Driver, error) {
	if !(strings.HasPrefix(config.URL, "ws://") || strings.HasPrefix(config.URL, "wss://")) {
		return nil, fmt.Errorf("%s (got '%s')", ErrInvalidWsUrl, config.URL)
	}

	return &opendax{
		once:           newOnce(),
		url:            config.URL,
		outbox:         outbox,
		filter:         NewFilter(config.Filter),
		history:        history,
		period:         config.ReconnectPeriod * time.Second,
		reqID:          atomic.Uint64{},
		dialer:         &wsDialWrapper{},
		streams:        safe.NewMap[Market, struct{}](),
		symbolToMarket: safe.NewMap[string, Market](),
	}, nil
}

func (o *opendax) ActiveDrivers() []DriverType {
	return []DriverType{DriverOpendax}
}

func (b *opendax) ExchangeType() ExchangeType {
	return ExchangeTypeHybrid
}

func (o *opendax) Start() error {
	var startErr error
	started := o.once.Start(func() {
		o.connect()
		go o.listen()

		cexConfigured.CompareAndSwap(false, true)
	})

	if !started {
		return ErrAlreadyStarted
	}
	return startErr
}

func (o *opendax) Stop() error {
	var stopErr error
	stopped := o.once.Stop(func() {
		conn := o.conn
		o.conn = nil

		if conn == nil {
			return
		}

		stopErr = conn.Close()
		cexConfigured.CompareAndSwap(true, false)
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return stopErr
}

func (o *opendax) Subscribe(market Market) error {
	if !o.once.Subscribe() {
		return ErrNotStarted
	}

	if _, ok := o.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	if err := o.subscribeUnchecked(market); err != nil {
		return err
	}

	o.streams.Store(market, struct{}{})
	o.symbolToMarket.Store(market.Base()+market.Quote(), market)
	return nil
}

func (o *opendax) subscribeUnchecked(market Market) error {
	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.Base(), market.Quote())
	message := protocol.NewSubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedSub, err)
	}

	return nil
}

func (o *opendax) Unsubscribe(market Market) error {
	if !o.once.Unsubscribe() {
		return ErrNotStarted
	}

	if _, ok := o.streams.Load(market); !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	if err := o.unsubscribeUnchecked(market); err != nil {
		return err
	}

	o.streams.Delete(market)
	o.symbolToMarket.Delete(market.Base() + market.Quote())
	return nil
}

func (*opendax) HistoricalData(_ context.Context, _ Market, _ time.Duration, _ uint64) ([]TradeEvent, error) {
	return nil, errors.New("not implemented")
}

func (o *opendax) unsubscribeUnchecked(market Market) error {
	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.Base(), market.Quote())
	message := protocol.NewUnsubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedUnsub, err)
	}

	return nil
}

func (o *opendax) connect() {
	for {
		var err error
		o.conn, _, err = o.dialer.Dial(o.url, nil)
		if err == nil {
			return
		}

		loggerOpendax.Warnf("connection attempt failed: %s", err.Error())
		time.Sleep(o.period)
	}
}

func (o *opendax) writeConn(message *protocol.Msg) error {
	byteMsg, err := message.Encode()
	if err != nil {
		loggerOpendax.Warn(err)
		return err
	}

	if err := o.conn.WriteMessage(websocket.TextMessage, byteMsg); err != nil {
		loggerOpendax.Warn(err)
		return err
	}

	if _, _, err := o.conn.ReadMessage(); err != nil {
		loggerOpendax.Warn(err)
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
			loggerOpendax.Warn("error reading from connection", err)

			o.connect()
			o.streams.Range(func(market Market, _ struct{}) bool {
				if err := o.subscribeUnchecked(market); err != nil {
					loggerOpendax.Warnf("error subscribing to market %s: %s", market, err)
					return false
				}
				return true
			})

			continue
		}

		loggerOpendax.Infow("raw event", "event", string(message))

		tr, err := o.parse(message)
		if err != nil {
			loggerOpendax.Warn(err)
			continue
		}

		// Skip system messages
		if tr.Market.IsEmpty() || tr.Price.IsZero() {
			continue
		}

		if !o.filter.Allow(*tr) {
			continue
		}
		o.outbox <- *tr
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

func (o *opendax) convertToTrade(args []any) (*TradeEvent, error) {
	it := protocol.NewArgIterator(args)

	marketSymbol := it.NextString()
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

	market, ok := NewMarketFromString(marketSymbol)
	if !ok {
		market, ok = o.symbolToMarket.Load(marketSymbol)
	}

	// market is unparsable and is missing in the map
	if !ok {
		return nil, fmt.Errorf("failed to get market: %s", marketSymbol)
	}

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
