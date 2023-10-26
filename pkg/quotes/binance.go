package quotes

import (
	"strings"
	"sync"

	"github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
	"github.com/layer-3/neodax/finex/pkg/config"
	"github.com/layer-3/neodax/finex/pkg/event"
	"github.com/layer-3/neodax/finex/pkg/websocket/client"
)

type Binance struct {
	streams map[string]chan struct{}
	mu      sync.Mutex

	marketCache  cache.Market
	tradeSampler *TradeSampler
	outbox       chan trade.Event
	output       chan<- event.Event
}

func (b *Binance) Init(
	markets cache.Market,
	outbox chan trade.Event,
	output chan<- event.Event,
	config config.QuoteFeed,
	_ client.WSDialer,
) error {
	binance.WebsocketKeepalive = true
	b.streams = make(map[string]chan struct{})

	b.marketCache = markets
	b.tradeSampler = NewTradeSampler(config.TradeSampler)
	b.outbox = outbox
	b.output = output

	return nil
}

func (b *Binance) Start() error {
	marketList, err := b.marketCache.GetActive()
	if err != nil {
		return err
	}

	for _, m := range marketList {
		m := m
		go func() {
			if err := b.Subscribe(m.BaseUnit, m.QuoteUnit); err != nil {
				logger.Warnf("failed to subscribe to %s market: %v", m.Symbol, err)
			}
		}()
	}

	return nil
}

func (b *Binance) Subscribe(base, quote string) error {
	pair := strings.ToUpper(base) + strings.ToUpper(quote)
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
					if err := b.Subscribe(base, quote); err == nil {
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

func (b *Binance) Close() error {
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
	routingEvent, err := GetRoutingEvent(tradeEvent)
	if err != nil {
		logger.Warn(err)
		return
	}
	b.output <- *routingEvent
}

func buildBinanceEvent(tr *binance.WsTradeEvent) (trade.Event, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		logger.Warn(err)
		return trade.Event{}, err
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		logger.Warn(err)
		return trade.Event{}, err
	}

	// IsBuyerMaker: true => the trade was initiated by the sell-side; the buy-side was the order book already.
	// IsBuyerMaker: false => the trade was initiated by the buy-side; the sell-side was the order book already.
	takerType := trade.Buy
	if tr.IsBuyerMaker {
		takerType = trade.Sell
	}

	return trade.Event{
		Market:    strings.ToLower(tr.Symbol),
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: tr.TradeTime,
		Source:    "Binance",
	}, nil
}
