package bitfaker

import (
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

type Bitfaker struct {
	once         *common.Once
	mu           sync.RWMutex
	streams      []common.Market
	outbox       chan<- common.TradeEvent
	stopCh       chan struct{}
	period       time.Duration
	tradeSampler common.TradeSampler
}

func NewBitfaker(config common.Config, outbox chan<- common.TradeEvent) *Bitfaker {
	return &Bitfaker{
		streams:      make([]common.Market, 0),
		outbox:       outbox,
		stopCh:       make(chan struct{}, 1),
		period:       5 * time.Second,
		tradeSampler: *common.NewTradeSampler(config.TradeSampler),
	}
}

func (b *Bitfaker) Start() error {
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

func (b *Bitfaker) Stop() error {
	b.once.Stop(func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		b.stopCh <- struct{}{}
		b.streams = make([]common.Market, 0)
	})
	return nil
}

func (b *Bitfaker) Subscribe(market common.Market) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, m := range b.streams {
		if m == market {
			return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
		}
	}

	b.streams = append(b.streams, market)
	return nil
}

func (b *Bitfaker) Unsubscribe(market common.Market) error {
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
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
	}

	b.streams = append(b.streams[:index], b.streams[index+1:]...)
	return nil
}

func (b *Bitfaker) createTradeEvent(market common.Market) {
	tr := common.TradeEvent{
		Market: market.BaseUnit + market.QuoteUnit,
		Price:  decimal.NewFromFloat(2.213),
		Source: common.DriverBitfaker,
	}

	b.outbox <- tr
}
