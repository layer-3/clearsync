package index

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
	"github.com/layer-3/clearsync/pkg/safe"
)

type priceCacheTrade struct {
	Price     decimal.Decimal
	Volume    decimal.Decimal
	Weight    decimal.Decimal
	Source    common.DriverType
	Timestamp time.Time
}

type marketKey struct {
	baseUnit  string
	quoteUnit string
}

type PriceCache struct {
	weights    safe.Map[common.DriverType, decimal.Decimal]
	market     safe.Map[marketKey, []priceCacheTrade]
	lastPrice  safe.Map[marketKey, decimal.Decimal]
	nTrades    int
	bufferTime time.Duration
}

// newPriceCache initializes a new cache to store last n trades for each market.
func newPriceCache(driversWeights map[common.DriverType]decimal.Decimal, nTrades int, bufferTime time.Duration) *PriceCache {
	cache := new(PriceCache)
	cache.market = safe.NewMap[marketKey, []priceCacheTrade]()
	cache.weights = safe.NewMapWithData(driversWeights)
	cache.lastPrice = safe.NewMap[marketKey, decimal.Decimal]()
	cache.nTrades = nTrades
	cache.bufferTime = bufferTime

	return cache
}

// AddTrade adds a new trade to the cache for a market.
func (p *PriceCache) AddTrade(market common.Market, price, volume decimal.Decimal, timestamp time.Time, source common.DriverType) {
	key := marketKey{baseUnit: market.BaseUnit, quoteUnit: market.QuoteUnit}
	p.market.UpdateInTx(func(m map[marketKey][]priceCacheTrade) {
		driversTrades, ok := m[key]
		if !ok {
			driversTrades = []priceCacheTrade{}
		}

		newTradesList := []priceCacheTrade{{Price: price, Volume: volume, Weight: decimal.Zero, Timestamp: timestamp, Source: source}}
		// transfer all existing trades to a new array
		for _, t := range driversTrades {
			if t.Source != source && time.Now().Sub(t.Timestamp) <= p.bufferTime {
				newTradesList = append(newTradesList, t)
			}
		}

		totalWeights := decimal.Zero
		for _, t := range newTradesList {
			w, ok := p.weights.Load(t.Source)
			if !ok {
				continue
			}
			totalWeights = totalWeights.Add(w)
		}

		var tradesList []priceCacheTrade
		for _, t := range newTradesList {
			w, ok := p.weights.Load(t.Source)
			if !ok {
				continue
			}

			if totalWeights != decimal.Zero {
				t.Weight = w.Div(totalWeights)
				tradesList = append(tradesList, t)
			}
		}

		m[key] = tradesList
	})
}

func (p *PriceCache) setLastPrice(market common.Market, newPrice decimal.Decimal) {
	key := marketKey{baseUnit: market.BaseUnit, quoteUnit: market.QuoteUnit}
	p.lastPrice.UpdateInTx(func(m map[marketKey]decimal.Decimal) {
		m[key] = newPrice
	})
}

func (p *PriceCache) getLastPrice(market common.Market) decimal.Decimal {
	record, ok := p.lastPrice.Load(marketKey{baseUnit: market.BaseUnit, quoteUnit: market.QuoteUnit})
	if !ok {
		return decimal.Zero
	}
	return record
}

func (p *PriceCache) GetIndexPrice(event *TradeEvent) (decimal.Decimal, bool) {
	trades, ok := p.market.Load(marketKey{baseUnit: event.Market.BaseUnit, quoteUnit: event.Market.QuoteUnit})
	if !ok || len(trades) == 0 {
		return event.Price, false
	}

	top := decimal.Zero
	bottom := decimal.Zero

	for _, t := range trades {
		top = top.Add(t.Price.Mul(t.Weight))
		bottom = bottom.Add(t.Weight)
	}

	if bottom.IsZero() {
		return decimal.Zero, false
	}

	quotePrice := decimal.NewFromInt(1)
	if event.Market.ConvertTo != "" {
		event.Market = common.Market{BaseUnit: event.Market.QuoteUnit, QuoteUnit: event.Market.ConvertTo}
		quotePrice, ok = p.GetIndexPrice(event)
		if !ok {
			return decimal.Zero, false
		}
	}

	return top.Div(bottom).Mul(quotePrice), true
}
