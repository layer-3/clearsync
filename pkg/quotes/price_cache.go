package quotes

import (
	"time"

	"github.com/layer-3/clearsync/pkg/safe"
	"github.com/shopspring/decimal"
)

type trade struct {
	Price     decimal.Decimal
	Volume    decimal.Decimal
	Weight    decimal.Decimal
	Source    DriverType
	Timestamp time.Time
}

type marketKey struct {
	baseUnit  string
	quoteUnit string
}

type PriceCache struct {
	weights    safe.Map[DriverType, decimal.Decimal]
	market     safe.Map[marketKey, []trade]
	lastPrice  safe.Map[marketKey, decimal.Decimal]
	nTrades    int
	bufferTime time.Duration
}

// newPriceCache initializes a new cache to store last n trades for each market.
func newPriceCache(driversWeights map[DriverType]decimal.Decimal, nTrades int, bufferTime time.Duration) *PriceCache {
	cache := new(PriceCache)
	cache.market = safe.NewMap[marketKey, []trade]()
	cache.weights = safe.NewMapWithData(driversWeights)
	cache.lastPrice = safe.NewMap[marketKey, decimal.Decimal]()
	cache.nTrades = nTrades
	cache.bufferTime = bufferTime

	return cache
}

// AddTrade adds a new trade to the cache for a market.
func (p *PriceCache) AddTrade(market Market, price, volume decimal.Decimal, timestamp time.Time, source DriverType) {
	key := marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit}
	p.market.UpdateInTx(func(m map[marketKey][]trade) {
		driversTrades, ok := m[key]
		if !ok {
			driversTrades = []trade{}
		}

		newTradesList := []trade{{Price: price, Volume: volume, Weight: decimal.Zero, Timestamp: timestamp, Source: source}}
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

		tradesList := []trade{}
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

func (p *PriceCache) setLastPrice(market Market, newPrice decimal.Decimal) {
	key := marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit}
	p.lastPrice.UpdateInTx(func(m map[marketKey]decimal.Decimal) {
		m[key] = newPrice
	})
}

func (p *PriceCache) getLastPrice(market Market) decimal.Decimal {
	record, ok := p.lastPrice.Load(marketKey{baseUnit: market.baseUnit, quoteUnit: market.quoteUnit})
	if !ok {
		return decimal.Zero
	}
	return record
}

func (p *PriceCache) GetIndexPrice(event *TradeEvent) (decimal.Decimal, bool) {
	trades, ok := p.market.Load(marketKey{baseUnit: event.Market.baseUnit, quoteUnit: event.Market.quoteUnit})
	if !ok || len(trades) == 0 {
		return event.Price, false
	}

	top := decimal.Zero
	bottom := decimal.Zero

	for _, t := range trades {
		top = top.Add(t.Price.Mul(t.Weight))
		bottom = bottom.Add(t.Weight)
	}

	if bottom.Equal(decimal.Zero) {
		return decimal.Zero, false
	}

	return top.Div(bottom), true
}
