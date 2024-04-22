package quotes

import (
	"fmt"
	"strings"
	"time"

	gobinance "github.com/adshao/go-binance/v2"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/safe"
)

var loggerBinance = log.Logger("binance")

type binance struct {
	once         *once
	streams      safe.Map[Market, chan struct{}]
	filter       Filter
	combineInbox chan<- TradeEvent
	outbox       chan<- TradeEvent
	usdcToUSDT   bool

	symbolToMarket safe.Map[string, Market]
}

func newBinance(config BinanceConfig, outbox chan<- TradeEvent) Driver {
	combineInbox := make(chan TradeEvent, 1024)
	go batcher(config.BatchPeriod, combineInbox, outbox)

	return &binance{
		once:           newOnce(),
		streams:        safe.NewMap[Market, chan struct{}](),
		filter:         NewFilter(config.Filter),
		usdcToUSDT:     config.USDCtoUSDT,
		combineInbox:   combineInbox,
		outbox:         outbox,
		symbolToMarket: safe.NewMap[string, Market](),
	}
}

func batcher(batchPeriod time.Duration, inbox <-chan TradeEvent, outbox chan<- TradeEvent) {
	marketTrades := make(map[Market][]TradeEvent)
	timer := time.NewTimer(batchPeriod)
	defer timer.Stop()

	for {
		select {
		case trade := <-inbox:
			marketTrades[trade.Market] = append(marketTrades[trade.Market], trade)
		case <-timer.C:
			for market, trades := range marketTrades {
				if event := combineTrades(trades); event != nil {
					marketTrades[market] = nil
					outbox <- *event
				}
			}
			timer.Reset(batchPeriod)
		}
	}
}

func combineTrades(trades []TradeEvent) *TradeEvent {
	if len(trades) == 0 {
		return nil
	}

	totalAmount := decimal.Zero
	totalValue := decimal.Zero
	netAmount := decimal.Zero

	for _, trade := range trades {
		totalAmount = totalAmount.Add(trade.Amount)
		totalValue = totalValue.Add(trade.Amount.Mul(trade.Price))

		// Update net amount to determine net side (buy or sell)
		if trade.TakerType == TakerTypeBuy {
			netAmount = netAmount.Add(trade.Amount)
		} else if trade.TakerType == TakerTypeSell {
			netAmount = netAmount.Sub(trade.Amount)
		}
	}

	if totalAmount.Equal(decimal.Zero) {
		return nil
	}

	avgPrice := totalValue.Div(totalAmount)
	// Determine net side (buy or sell)
	var side TakerType
	if netAmount.GreaterThanOrEqual(decimal.Zero) {
		side = TakerTypeSell // "buy" (yes, it looks inverted)
	} else {
		side = TakerTypeBuy // "sell"
		netAmount = netAmount.Abs()
	}

	return &TradeEvent{
		Source:    trades[0].Source,
		Market:    trades[0].Market,
		Price:     avgPrice,
		Amount:    totalAmount,
		Total:     avgPrice.Mul(totalAmount),
		TakerType: side,
		CreatedAt: time.Now(),
	}
}

func (b *binance) ActiveDrivers() []DriverType {
	return []DriverType{DriverBinance}
}

func (b *binance) ExchangeType() ExchangeType {
	return ExchangeTypeCEX
}

func (b *binance) Start() error {
	if started := b.once.Start(func() {}); !started {
		return ErrAlreadyStarted
	}
	return nil
}

func (b *binance) Stop() error {
	stopped := b.once.Stop(func() {
		b.streams.Range(func(market Market, _ chan struct{}) bool {
			err := b.Unsubscribe(market)
			return err == nil
		})

		b.streams = safe.NewMap[Market, chan struct{}]()
	})

	if !stopped {
		return ErrAlreadyStopped
	}
	return nil
}

func (b *binance) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return ErrNotStarted
	}

	if b.usdcToUSDT && market.Quote() == "usd" {
		if err := b.Subscribe(NewMarket(market.Base(), "usdt")); err != nil {
			loggerBinance.Warnw("failed to subscribe to USDT", "market", market, "error", err)
		}

		if err := b.Subscribe(NewMarket(market.Base(), "usdc")); err != nil {
			loggerBinance.Warnw("failed to subscribe to USDC", "market", market, "error", err)
		}
		return nil
	}

	pair := strings.ToUpper(market.Base()) + strings.ToUpper(market.Quote())
	b.symbolToMarket.Store(strings.ToLower(pair), market)

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, ErrAlreadySubbed)
	}

	doneCh, stopCh, err := gobinance.WsTradeServe(pair, b.handleTrade, b.handleErr(market))
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, ErrFailedSub, err)
	}
	b.streams.Store(market, stopCh)

	go func() {
		defer close(doneCh)
		<-doneCh

		loggerBinance.Infow("resubscribing", "market", market)
		if _, ok := b.streams.Load(market); !ok {
			return // market was unsubscribed earlier
		}

		loggerBinance.Warnw("connection failed, resubscribing", "market", market)
		if err := b.Unsubscribe(market); err != nil {
			loggerBinance.Errorw("failed to unsubscribe", "market", pair, "error", err)
		}

		for {
			err := b.Subscribe(market)
			if err == nil {
				break
			}
			loggerBinance.Errorw("failed to resubscribe", "market", pair, "error", err)
			<-time.After(5 * time.Second)
		}
	}()

	loggerBinance.Infow("subscribed", "market", strings.ToUpper(pair))
	return nil
}

func (b *binance) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return ErrNotStarted
	}

	stopCh, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, ErrNotSubbed)
	}

	stopCh <- struct{}{}
	close(stopCh)

	b.streams.Delete(market)
	return nil
}

func (b *binance) handleTrade(event *gobinance.WsTradeEvent) {
	tradeEvent, err := b.buildEvent(event)
	if err != nil {
		loggerBinance.Errorw("failed to build trade event", "event", event, "error", err)
		return
	}

	if !b.filter.Allow(tradeEvent) {
		return
	}
	b.combineInbox <- tradeEvent
}

func (b *binance) handleErr(market Market) func(error) {
	return func(err error) {
		loggerBinance.Errorw("received error", "market", market, "error", err)
	}
}

func (b *binance) buildEvent(tr *gobinance.WsTradeEvent) (TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		return TradeEvent{}, fmt.Errorf("failed to parse price: %+v", tr.Price)
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		return TradeEvent{}, fmt.Errorf("failed to parse quantity: %+v", tr.Quantity)
	}

	market, ok := b.symbolToMarket.Load(strings.ToLower(tr.Symbol))
	if !ok {
		return TradeEvent{}, fmt.Errorf("failed to load market: %+v", tr.Symbol)
	}

	if b.usdcToUSDT && (market.quoteUnit == "usdt" || market.quoteUnit == "usdc") {
		market.quoteUnit = "usd"
	}

	// IsBuyerMaker: true => the trade was initiated by the sell-side; the buy-side was the order book already.
	// IsBuyerMaker: false => the trade was initiated by the buy-side; the sell-side was the order book already.
	takerType := TakerTypeBuy
	if tr.IsBuyerMaker {
		takerType = TakerTypeSell
	}

	return TradeEvent{
		Source:    DriverBinance,
		Market:    market,
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.UnixMilli(tr.TradeTime),
	}, nil
}

func (b *binance) SetInbox(_ <-chan TradeEvent) {}
