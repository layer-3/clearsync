package quotes

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"
)

type Binance struct {
	mu      sync.Mutex
	streams map[string]chan struct{}

	tradeSampler *TradeSampler
	outbox       chan<- TradeEvent
}

func NewBinance(config Config, outbox chan<- TradeEvent) *Binance {
	binance.WebsocketKeepalive = true
	return &Binance{
		streams:      make(map[string]chan struct{}),
		tradeSampler: NewTradeSampler(config.TradeSampler),
		outbox:       outbox,
	}
}

func (b *Binance) Start(markets []Market) error {
	if len(markets) == 0 {
		return errors.New("no markets specified")
	}

	for _, m := range markets {
		m := m
		go func() {
			if err := b.Subscribe(m); err != nil {
				symbol := m.BaseUnit + m.QuoteUnit
				logger.Warnf("failed to subscribe to %s market: %v", symbol, err)
			}
		}()
	}

	return nil
}

func (b *Binance) Subscribe(m Market) error {
	pair := strings.ToUpper(m.BaseUnit) + strings.ToUpper(m.QuoteUnit)
	handleErr := func(err error) {
		logger.Errorf("error for Binance market %s: %v", pair, err)
	}

	doneCh, stopCh, err := binance.WsTradeServe(pair, b.handleTrade, handleErr)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-doneCh:
				for {
					if err := b.Subscribe(m); err == nil {
						break
					}
				}
				return
			}
		}
	}()

	b.mu.Lock()
	defer b.mu.Unlock()
	b.streams[pair] = stopCh

	logger.Infof("subscribed to Binance %s market", strings.ToUpper(pair))
	return nil
}

func (b *Binance) Stop() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, stopCh := range b.streams {
		stopCh <- struct{}{}
		close(stopCh)
	}

	b.streams = make(map[string]chan struct{}) // delete all stopped streams
	return nil
}

func (b *Binance) handleTrade(event *binance.WsTradeEvent) {
	tradeEvent, err := buildBinanceEvent(event)
	if err != nil {
		logger.Error(err)
		return
	}

	if !b.tradeSampler.Allow(tradeEvent) {
		return
	}

	b.outbox <- tradeEvent
}

func buildBinanceEvent(tr *binance.WsTradeEvent) (TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		logger.Warn(err)
		return TradeEvent{}, err
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		logger.Warn(err)
		return TradeEvent{}, err
	}

	// IsBuyerMaker: true => the trade was initiated by the sell-side; the buy-side was the order book already.
	// IsBuyerMaker: false => the trade was initiated by the buy-side; the sell-side was the order book already.
	takerType := TakerTypeBuy
	if tr.IsBuyerMaker {
		takerType = TakerTypeSell
	}

	return TradeEvent{
		Source:    DriverBinance,
		Market:    strings.ToLower(tr.Symbol),
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Unix(tr.TradeTime, 0),
	}, nil
}
