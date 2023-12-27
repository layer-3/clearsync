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
		streams:      make([]Market, 0),
		outbox:       outbox,
		stopCh:       make(chan struct{}, 1),
		period:       5 * time.Second,
		tradeSampler: *newTradeSampler(config.TradeSampler),
	}
}

func (b *bitfaker) Start() error {
	b.once.Start(func() {
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
	return nil
}

func (b *bitfaker) Stop() error {
	b.once.Stop(func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		b.stopCh <- struct{}{}
		b.streams = make([]Market, 0)
	})
	return nil
}

func (b *bitfaker) Subscribe(market Market) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, m := range b.streams {
		if m == market {
			return fmt.Errorf("market %s already subscribed", market)
		}
	}

	b.streams = append(b.streams, market)
	return nil
}

func (b *bitfaker) Unsubscribe(market Market) error {
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
		return fmt.Errorf("market %s not found", market)
	}

	b.streams = append(b.streams[:index], b.streams[index+1:]...)
	return nil
}

func (b *bitfaker) createTradeEvent(market Market) {
	tr := TradeEvent{
		Market: market.BaseUnit + market.QuoteUnit,
		Price:  decimal.NewFromFloat(2.213),
		Source: DriverBitfaker,
	}

	b.outbox <- tr
}
