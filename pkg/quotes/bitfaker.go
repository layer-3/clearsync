package quotes

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/layer-3/neodax/finex/models/market"
	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
	"github.com/layer-3/neodax/finex/pkg/config"
	"github.com/layer-3/neodax/finex/pkg/event"
	"github.com/layer-3/neodax/finex/pkg/websocket/client"
)

type Bitfaker struct {
	outbox       chan trade.Event
	output       chan<- event.Event
	markets      []string
	marketCache  cache.Market
	period       time.Duration
	tradeSampler *TradeSampler
}

func (b *Bitfaker) Init(markets cache.Market, outbox chan trade.Event, output chan<- event.Event, config config.QuoteFeed, dialer client.WSDialer) error {
	b.outbox = outbox
	b.output = output
	b.markets = make([]string, 0)
	b.marketCache = markets
	b.period = 5 * time.Second
	b.tradeSampler = NewTradeSampler(config.TradeSampler)
	return nil
}

func (b *Bitfaker) Start() error {
	func() {
		for {
			markets, _ := b.marketCache.GetActive()
			for _, v := range markets {
				b.createTradeEvent(v)
			}

			<-time.After(b.period)
		}
	}()
	return nil
}

func (b *Bitfaker) Subscribe(base, quote string) error {
	b.markets = append(b.markets, fmt.Sprintf("%s%s", base, quote))
	return nil
}

func (b *Bitfaker) createTradeEvent(market market.Market) {
	tr := trade.Event{
		Market: market.Symbol,
		Price:  decimal.NewFromFloat(2.213),
		Source: "Bitfaker",
	}

	b.outbox <- tr
	event, err := GetRoutingEvent(tr)
	if err != nil {
		logger.Warn(err)
	}
	b.output <- *event
}

func (b *Bitfaker) Close() error {
	return nil
}
