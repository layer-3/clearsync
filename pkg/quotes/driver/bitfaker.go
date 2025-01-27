package driver

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/quotes/filter"
)

type bitfaker struct {
	once          *common.Once
	mu            sync.RWMutex
	streamPeriods map[common.Market]time.Duration
	streams       map[common.Market]chan struct{}
	outbox        chan<- common.TradeEvent
	stopCh        chan struct{}
	filter        filter.Filter
	marketConfigs map[string]FakeMarketConfig
}

// Config Example:
// config.Bitfaker.Markets = map[string]quotes.FakeMarketConfig{
// 	"btc/usdc":     {StartPrice: 70000.0, PriceVolatility: 2000.0, Period: 5 * time.Second},
// 	"eth/usdc":     {StartPrice: 4000.0, PriceVolatility: 100.0, Period: 2 * time.Second},
// 	"yellow/usdc":  {StartPrice: 0.05, PriceVolatility: 0.01, Period: time.Second},
// 	"duckies/usdc": {StartPrice: 0.000004, PriceVolatility: 0.000003, Period: 500 * time.Millisecond},
// }

func newBitfaker(config BitfakerConfig, outbox chan<- common.TradeEvent) (Driver, error) {
	return &bitfaker{
		once:          common.NewOnce(),
		streamPeriods: make(map[common.Market]time.Duration),
		streams:       make(map[common.Market]chan struct{}),
		stopCh:        make(chan struct{}),
		outbox:        outbox,
		filter:        filter.New(config.Filter),
		marketConfigs: config.Markets,
	}, nil
}

func (b *bitfaker) Type() (common.DriverType, common.ExchangeType) {
	return common.DriverBitfaker, common.ExchangeTypeHybrid
}

func (b *bitfaker) Start() error {
	started := b.once.Start(func() {})
	if !started {
		return common.ErrAlreadyStarted
	}
	return nil
}

func (b *bitfaker) Stop() error {
	stopped := b.once.Stop(func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		for market, stopStreamCh := range b.streams {
			close(stopStreamCh)
			delete(b.streams, market)
		}
		close(b.stopCh)
	})

	if !stopped {
		return common.ErrAlreadyStopped
	}
	return nil
}

func initializeMarket(startValue, volatility float64) func() float64 {
	currentValue := startValue
	return func() float64 {
		change := rand.NormFloat64() * volatility
		currentValue += change
		if currentValue < 0 {
			currentValue -= change * 2
		}
		return currentValue
	}
}

func (b *bitfaker) startStream(market common.Market, period time.Duration, stopCh chan struct{}) {
	startPrice := 100.0
	startAmount := 10.0
	priceVolatility := 2.0
	amountVolatility := 1.0

	cnf, ok := b.marketConfigs[market.String()] // Default values are not supported by nested configs.
	if ok {
		if cnf.StartPrice != 0 {
			startPrice = cnf.StartPrice
		}
		if cnf.PriceVolatility != 0 {
			priceVolatility = cnf.PriceVolatility
		}
		if cnf.StartAmount != 0 {
			startAmount = cnf.StartAmount
		}
		if cnf.AmountVolatility != 0 {
			amountVolatility = cnf.AmountVolatility
		}
	}

	newPrice := initializeMarket(startPrice, priceVolatility)
	newAmount := initializeMarket(startAmount, amountVolatility)

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-stopCh:
			return
		case <-ticker.C:
			b.mu.RLock()
			b.createTradeEvent(market, newPrice, newAmount)
			b.mu.RUnlock()
		}
	}
}

func (b *bitfaker) Subscribe(market common.Market) error {
	if !b.once.Subscribe() {
		return common.ErrNotStarted
	}

	period := 2 * time.Second
	cnf, ok := b.marketConfigs[market.String()]
	if ok && cnf.Period.Seconds() != 0 {
		period = cnf.Period
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.streams[market]; exists {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
	}

	stopStreamCh := make(chan struct{})
	b.streams[market] = stopStreamCh
	b.streamPeriods[market] = period

	go b.startStream(market, period, stopStreamCh)

	return nil
}

func (b *bitfaker) Unsubscribe(market common.Market) error {
	if !b.once.Unsubscribe() {
		return common.ErrNotStarted
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	stopStreamCh, exists := b.streams[market]
	if !exists {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
	}

	close(stopStreamCh)
	delete(b.streams, market)
	delete(b.streamPeriods, market)

	return nil
}

func (*bitfaker) HistoricalData(_ context.Context, _ common.Market, _ time.Duration, _ uint64) ([]common.TradeEvent, error) {
	return nil, errors.New("by design Bitfaker does not support querying historical data, use a different provider")
}

func (b *bitfaker) createTradeEvent(market common.Market, nextPriceFunc, nextAmountFunc func() float64) {
	price := decimal.NewFromFloat(nextPriceFunc())
	amount := decimal.NewFromFloat(nextAmountFunc())
	takerType := common.TakerTypeBuy

	tr := common.TradeEvent{
		Source:    common.DriverBitfaker,
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Now(),
	}

	if !b.filter.Allow(tr) {
		return
	}
	b.outbox <- tr
}
