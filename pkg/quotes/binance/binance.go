package quotes

import (
	"fmt"
	"strings"
	"sync"
	"time"

	gobinance "github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"
	"github.com/ipfs/go-log/v2"

  "github.com/layer-3/clearsync/pkg/quotes/common"
)

var logger = log.Logger("trade_sampler")

type Binance struct {
	once         *common.Once
	streams      sync.Map
	tradeSampler common.TradeSampler
	outbox       chan<- common.TradeEvent
}

func NewBinance(config common.Config, outbox chan<- common.TradeEvent) *Binance {
	gobinance.WebsocketKeepalive = true
	return &Binance{
		once:         common.NewOnce(),
		tradeSampler: *common.NewTradeSampler(config.TradeSampler),
		outbox:       outbox,
	}
}

func (b *Binance) Start() error {
	b.once.Start(func() {})
	return nil
}

func (b *Binance) Stop() error {
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

func (b *Binance) Subscribe(market common.Market) error {
	pair := strings.ToUpper(market.BaseUnit) + strings.ToUpper(market.QuoteUnit)
	if _, ok := b.streams.Load(pair); ok {
		return fmt.Errorf("%s: %w", market, common.ErrAlreadySubbed)
	}

	handleErr := func(err error) {
		logger.Errorf("error for Binance market %s: %v", pair, err)
	}

	doneCh, stopCh, err := gobinance.WsTradeServe(pair, b.handleTrade, handleErr)
	if err != nil {
		return fmt.Errorf("%s: %w: %w", market, common.ErrFailedSub, err)
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

func (b *Binance) Unsubscribe(market common.Market) error {
	pair := strings.ToUpper(market.BaseUnit) + strings.ToUpper(market.QuoteUnit)
	stream, ok := b.streams.Load(pair)
	if !ok {
		return fmt.Errorf("%s: %w", market, common.ErrNotSubbed)
	}

	stopCh := stream.(chan struct{})
	stopCh <- struct{}{}
	close(stopCh)

	b.streams.Delete(pair)
	return nil
}

func (b *Binance) handleTrade(event *gobinance.WsTradeEvent) {
	tradeEvent, err := buildEvent(event)
	if err != nil {
		logger.Error(err)
		return
	}

	if !b.tradeSampler.Allow(tradeEvent) {
		return
	}

	b.outbox <- tradeEvent
}

func buildEvent(tr *gobinance.WsTradeEvent) (common.TradeEvent, error) {
	price, err := decimal.NewFromString(tr.Price)
	if err != nil {
		logger.Warn(err)
		return common.TradeEvent{}, err
	}

	amount, err := decimal.NewFromString(tr.Quantity)
	if err != nil {
		logger.Warn(err)
		return common.TradeEvent{}, err
	}

	// IsBuyerMaker: true => the trade was initiated by the sell-side; the buy-side was the order book already.
	// IsBuyerMaker: false => the trade was initiated by the buy-side; the sell-side was the order book already.
	takerType := common.TakerTypeBuy
	if tr.IsBuyerMaker {
		takerType = common.TakerTypeSell
	}

	return common.TradeEvent{
		Source:    common.DriverBinance,
		Market:    strings.ToLower(tr.Symbol),
		Price:     price,
		Amount:    amount,
		Total:     price.Mul(amount),
		TakerType: takerType,
		CreatedAt: time.Unix(tr.TradeTime, 0),
	}, nil
}
