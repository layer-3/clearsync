package driver

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ipfs/go-log/v2"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/filter"
	protocol "github.com/layer-3/clearsync/pkg/quotes/opendax_protocol"
	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerOpendax = log.Logger("opendax")

type opendax struct {
	once   *common.Once
	conn   common.WsTransport
	dialer common.WsDialer
	url    string

	outbox         chan<- common.TradeEvent
	filter         filter.Filter
	history        HistoricalDataDriver
	period         time.Duration
	reqID          atomic.Uint64
	streams        safe.Map[common.Market, struct{}]
	symbolToMarket safe.Map[string, common.Market]
}

func newOpendax(config OpendaxConfig, outbox chan<- common.TradeEvent, history HistoricalDataDriver) (Driver, error) {
	if !(strings.HasPrefix(config.URL, "ws://") || strings.HasPrefix(config.URL, "wss://")) {
		return nil, fmt.Errorf("%s (got '%s')", common.ErrInvalidWsUrl, config.URL)
	}

	return &opendax{
		once:           common.NewOnce(),
		url:            config.URL,
		outbox:         outbox,
		filter:         filter.New(config.Filter),
		history:        history,
		period:         config.ReconnectPeriod * time.Second,
		reqID:          atomic.Uint64{},
		dialer:         &common.WsDialWrapper{},
		streams:        safe.NewMap[common.Market, struct{}](),
		symbolToMarket: safe.NewMap[string, common.Market](),
	}, nil
}

func (o *opendax) ActiveDrivers() []common.DriverType {
	return []common.DriverType{common.DriverOpendax}
}

func (b *opendax) ExchangeType() common.ExchangeType {
	return common.ExchangeTypeHybrid
}

func (o *opendax) Start() error {
	var startErr error
	started := o.once.Start(func() {
		o.connect()
		go o.listen()

		cexConfigured.CompareAndSwap(false, true)
	})

	if !started {
		return common.ErrAlreadyStarted
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
		return common.ErrAlreadyStopped
	}
	return stopErr
}

func (o *opendax) Subscribe(market common.Market) error {
	if !o.once.Subscribe() {
		return common.ErrNotStarted
	}

	if _, ok := o.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
	}

	if err := o.subscribeUnchecked(market); err != nil {
		return err
	}

	o.streams.Store(market, struct{}{})
	o.symbolToMarket.Store(market.Base()+market.Quote(), market)
	return nil
}

func (o *opendax) subscribeUnchecked(market common.Market) error {
	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.Base(), market.Quote())
	message := protocol.NewSubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, common.ErrFailedSub, err)
	}

	return nil
}

func (o *opendax) Unsubscribe(market common.Market) error {
	if !o.once.Unsubscribe() {
		return common.ErrNotStarted
	}

	if _, ok := o.streams.Load(market); !ok {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
	}

	if err := o.unsubscribeUnchecked(market); err != nil {
		return err
	}

	o.streams.Delete(market)
	o.symbolToMarket.Delete(market.Base() + market.Quote())
	return nil
}

func (*opendax) HistoricalData(_ context.Context, _ common.Market, _ time.Duration, _ uint64) ([]common.TradeEvent, error) {
	return nil, errors.New("not implemented")
}

func (o *opendax) unsubscribeUnchecked(market common.Market) error {
	// Opendax resource [market].[trades]
	resource := fmt.Sprintf("%s%s.trades", market.Base(), market.Quote())
	message := protocol.NewUnsubscribeMessage(o.reqID.Load(), resource)
	o.reqID.Add(1)

	err := o.writeConn(message)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, common.ErrFailedUnsub, err)
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
			o.streams.Range(func(market common.Market, _ struct{}) bool {
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

func (o *opendax) parse(message []byte) (*common.TradeEvent, error) {
	if message == nil {
		return &common.TradeEvent{}, nil
	}
	parsedMsg, err := protocol.ParseRaw(message)
	if err != nil {
		return nil, err
	}

	tradeEvent, err := o.convertToTrade(parsedMsg.Args)
	if errors.Is(err, errors.New(protocol.ParseError)) || parsedMsg.Method == protocol.MethodSubscribe || parsedMsg.
		Method == protocol.EventSystem {
		return &common.TradeEvent{}, nil
	}
	if err != nil {
		return nil, err
	}

	return tradeEvent, nil
}

func (o *opendax) convertToTrade(args []any) (*common.TradeEvent, error) {
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

	market, ok := common.NewMarketFromString(marketSymbol)
	if !ok {
		market, ok = o.symbolToMarket.Load(marketSymbol)
	}

	// market is unparsable and is missing in the map
	if !ok {
		return nil, fmt.Errorf("failed to get market: %s", marketSymbol)
	}

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

func (o *opendax) recognizeSide(side string) (common.TakerType, error) {
	switch side {
	case protocol.OrderSideSell:
		return common.TakerTypeSell, nil
	case protocol.OrderSideBuy:
		return common.TakerTypeBuy, nil
	default:
		return common.TakerTypeUnknown, errors.New("order side invalid: " + side)
	}
}
