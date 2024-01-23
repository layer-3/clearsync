package quotes

import (
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type bitfaker struct {
	once         *once
	mu           sync.RWMutex
	streams      []Market
	outbox       chan<- TradeEvent
	stopCh       chan struct{}
	period       time.Duration
	tradeSampler tradeSampler
}

func newBitfaker(config Config, outbox chan<- TradeEvent) *bitfaker {
	return &bitfaker{
		once:         newOnce(),
		streams:      make([]Market, 0),
		outbox:       outbox,
		stopCh:       make(chan struct{}, 1),
		period:       5 * time.Second,
		tradeSampler: *newTradeSampler(config.TradeSampler),
	}
}

func (b *bitfaker) Start() error {
	started := b.once.Start(func() {
		go func() {
			for {
				select {
				case <-b.stopCh:
					return
				default:
				}

				b.mu.RLock()
				for _, v := range b.streams {
					b.createTradeEvent(v)
				}
				b.mu.RUnlock()
				<-time.After(b.period)
			}
		}()
	})

	if !started {
		return errAlreadyStarted
	}
	return nil
}

func (b *bitfaker) Stop() error {
	stopped := b.once.Stop(func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		b.stopCh <- struct{}{}
		b.streams = make([]Market, 0)
	})

	if !stopped {
		return errAlreadyStopped
	}
	return nil
}

func (b *bitfaker) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return errNotStarted
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	for _, m := range b.streams {
		if m == market {
			return fmt.Errorf("%s: %w", market, errAlreadySubbed)
		}
	}

	b.streams = append(b.streams, market)
	return nil
}

func (b *bitfaker) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return errNotStarted
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	index := -1
	for i, m := range b.streams {
		if market == m {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	b.streams = append(b.streams[:index], b.streams[index+1:]...)
	return nil
}

func (b *bitfaker) createTradeEvent(market Market) {
	price := decimal.NewFromFloat(2.213)
	amount := decimal.NewFromFloat(2.213)
	takerType := TakerTypeBuy

	tr := TradeEvent{
		Source:    DriverBitfaker,
		Market:    market.BaseUnit + market.QuoteUnit,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}

	b.outbox <- tr
}
