package quotes

import (
	"fmt"
	"strings"
	"sync"
	"time"

	gobinance "github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"
)

type binance struct {
	once         *once
	streams      sync.Map
	tradeSampler tradeSampler
	outbox       chan<- TradeEvent
}

func newBinance(config Config, outbox chan<- TradeEvent) *binance {
	gobinance.WebsocketKeepalive = true
	return &binance{
		once:         newOnce(),
		tradeSampler: *newTradeSampler(config.TradeSampler),
		outbox:       outbox,
	}
}

func (b *binance) Start() error {
	b.once.Start(func() {})
	return nil
}

func (b *binance) Stop() error {
	b.once.Stop(func() {
		b.streams.Range(func(key, value any) bool {
			stopCh := value.(chan struct{})
			stopCh <- struct{}{}
			close(stopCh)
			return true
		})

		b.streams = sync.Map{}
	})
	return nil
}

func (b *binance) Subscribe(market Market) error {
	pair := strings.ToUpper(market.BaseUnit) + strings.ToUpper(market.QuoteUnit)
	if _, ok := b.streams.Load(pair); ok {
		return fmt.Errorf("market %s already subscribed", pair)
	}

	handleErr := func(err error) {
		logger.Errorf("error for Binance market %s: %v", pair, err)
	}

	doneCh, stopCh, err := gobinance.WsTradeServe(pair, b.handleTrade, handleErr)
	if err != nil {
		return err
	}
	b.streams.Store(pair, stopCh)

	go func() {
		select {
		case <-doneCh:
			for {
				if err := b.Subscribe(market); err == nil {
					return
				}
			}
		}
	}()

	logger.Infof("subscribed to Binance %s market", strings.ToUpper(pair))
	return nil
}

func (b *binance) Unsubscribe(market Market) error {
	pair := strings.ToUpper(market.BaseUnit) + strings.ToUpper(market.QuoteUnit)
	stream, ok := b.streams.Load(pair)
	if !ok {
		return fmt.Errorf("market %s not found", pair)
	}

	stopCh := stream.(chan struct{})
	stopCh <- struct{}{}
	close(stopCh)

	b.streams.Delete(pair)
	return nil
}

func (b *binance) handleTrade(event *gobinance.WsTradeEvent) {
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

func buildBinanceEvent(tr *gobinance.WsTradeEvent) (TradeEvent, error) {
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
