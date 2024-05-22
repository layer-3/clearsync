package quotes

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type bitfaker struct {
	once          *once
	mu            sync.RWMutex
	streamPeriods map[Market]time.Duration
	streams       map[Market]chan struct{}
	outbox        chan<- TradeEvent
	stopCh        chan struct{}
	filter        Filter
	marketConfigs map[string]FakeMarketConfig
}

// Config Example:
// config.Bitfaker.Markets = map[string]quotes.FakeMarketConfig{
// 	"btc/usdc":     {StartPrice: 70000.0, PriceVolatility: 2000.0, Period: 5 * time.Second},
// 	"eth/usdc":     {StartPrice: 4000.0, PriceVolatility: 100.0, Period: 2 * time.Second},
// 	"yellow/usdc":  {StartPrice: 0.05, PriceVolatility: 0.01, Period: time.Second},
// 	"duckies/usdc": {StartPrice: 0.000004, PriceVolatility: 0.000003, Period: 500 * time.Millisecond},
// }

func newBitfaker(config BitfakerConfig, outbox chan<- TradeEvent) Driver {
	return &bitfaker{
		once:          newOnce(),
		streamPeriods: make(map[Market]time.Duration),
		streams:       make(map[Market]chan struct{}),
		stopCh:        make(chan struct{}),
		outbox:        outbox,
		filter:        NewFilter(config.Filter),
		marketConfigs: config.Markets,
	}
}

func (b *bitfaker) ActiveDrivers() []DriverType {
	return []DriverType{DriverBitfaker}
}

func (b *bitfaker) ExchangeType() ExchangeType {
	return ExchangeTypeUnspecified
}

func (b *bitfaker) Start() error {
	started := b.once.Start(func() {})
	if !started {
		return ErrAlreadyStarted
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
		return ErrAlreadyStopped
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

func (b *bitfaker) startStream(market Market, period time.Duration, stopCh chan struct{}) {
	startPrice := 100.0
	startAmount := 10.0
	priceVolatility := 2.0
	amountVolatility := 1.0

	cnf, ok := b.marketConfigs[market.String()] // Defauld values are not supported by nested configs.
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

func (b *bitfaker) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return ErrNotStarted
	}

	period := 2 * time.Second
	cnf, ok := b.marketConfigs[market.String()]
	if ok && cnf.Period.Seconds() != 0 {
		period = cnf.Period
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.streams[market]; exists {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	stopStreamCh := make(chan struct{})
	b.streams[market] = stopStreamCh
	b.streamPeriods[market] = period

	go b.startStream(market, period, stopStreamCh)

	return nil
}

func (b *bitfaker) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return ErrNotStarted
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	stopStreamCh, exists := b.streams[market]
	if !exists {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	close(stopStreamCh)
	delete(b.streams, market)
	delete(b.streamPeriods, market)

	return nil
}

func (b *bitfaker) createTradeEvent(market Market, nextPriceFunc, nextAmountFunc func() float64) {
	price := decimal.NewFromFloat(nextPriceFunc())
	amount := decimal.NewFromFloat(nextAmountFunc())
	takerType := TakerTypeBuy

	tr := TradeEvent{
		Source:    DriverBitfaker,
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

func (b *bitfaker) SetInbox(_ <-chan TradeEvent) {}
