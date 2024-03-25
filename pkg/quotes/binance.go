package quotes

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	gobinance "github.com/adshao/go-binance/v2"
	"github.com/ipfs/go-log/v2"
	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

var (
	loggerBinance             = log.Logger("binance")
	binanceWebsocketKeepalive = atomic.Bool{}
)

type binance struct {
	once       *once
	streams    safe.Map[Market, chan struct{}]
	filter     Filter
	outbox     chan<- TradeEvent
	usdcToUSDT bool

	symbolToMarket safe.Map[string, Market]
}

func newBinance(config BinanceConfig, outbox chan<- TradeEvent) Driver {
	if binanceWebsocketKeepalive.CompareAndSwap(false, true) {
		gobinance.WebsocketKeepalive = true
	}

	return &binance{
		once:           newOnce(),
		streams:        safe.NewMap[Market, chan struct{}](),
		filter:         NewFilter(config.Filter),
		usdcToUSDT:     config.USDCtoUSDT,
		outbox:         outbox,
		symbolToMarket: safe.NewMap[string, Market](),
	}
}

func (b *binance) Name() DriverType {
	return DriverBinance
}

func (b *binance) Type() Type {
	return TypeCEX
}

func (b *binance) Start() error {
	if started := b.once.Start(func() {}); !started {
		return errAlreadyStarted
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
		return errAlreadyStopped
	}
	return nil
}

func (b *binance) Subscribe(market Market) error {
	if !b.once.Subscribe() {
		return errNotStarted
	}

	if b.usdcToUSDT && market.Quote() == "usdc" {
		_ = b.Subscribe(NewMarket(market.Base(), "usdt"))
	}

	pair := strings.ToUpper(market.Base()) + strings.ToUpper(market.Quote())
	b.symbolToMarket.Store(strings.ToLower(pair), market)

	if _, ok := b.streams.Load(market); ok {
		return fmt.Errorf("%s: %w", market, errAlreadySubbed)
	}

	handleErr := func(err error) {
		loggerBinance.Errorf("error for Binance market %s: %v", pair, err)
	}

	doneCh, stopCh, err := gobinance.WsTradeServe(pair, b.handleTrade, handleErr)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, errFailedSub, err)
	}
	b.streams.Store(market, stopCh)

	go func() {
		<-doneCh
		for {
			if err := b.Subscribe(market); err == nil {
				return
			}
		}
	}()

	loggerBinance.Infof("subscribed to Binance %s market", strings.ToUpper(pair))
	return nil
}

func (b *binance) Unsubscribe(market Market) error {
	if !b.once.Unsubscribe() {
		return errNotStarted
	}

	stopCh, ok := b.streams.Load(market)
	if !ok {
		return fmt.Errorf("%s: %w", market, errNotSubbed)
	}

	stopCh <- struct{}{}
	close(stopCh)

	b.streams.Delete(market)
	return nil
}

func (b *binance) handleTrade(event *gobinance.WsTradeEvent) {
	tradeEvent, err := b.buildEvent(event)
	if err != nil {
		loggerBinance.Error(err)
		return
	}

	if !b.filter.Allow(tradeEvent) {
		return
	}
	b.outbox <- tradeEvent
}

func (b *binance) buildEvent(tr *gobinance.WsTradeEvent) (TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		loggerBinance.Warn(err)
		return TradeEvent{}, err
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		loggerBinance.Warn(err)
		return TradeEvent{}, err
	}

	market, ok := b.symbolToMarket.Load(strings.ToLower(tr.Symbol))
	if !ok {
		return TradeEvent{}, fmt.Errorf("failed to load market: %+v", tr.Symbol)
	}

	if b.usdcToUSDT && market.quoteUnit == "usdt" {
		market.quoteUnit = "usdc"
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
