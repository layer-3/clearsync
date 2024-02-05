package quotes

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	BaseUnit  string // e.g. `btc` in `btc/usdt`
	QuoteUnit string // e.g. `usdt` in `btc/usdt`
}

// String returns a string representation of the market
// Market{BTC, Usdt} -> "btc/usdt"
func (m Market) String() string {
	return strings.ToLower(fmt.Sprintf("%s/%s", m.BaseUnit, m.QuoteUnit))
}

func NewMarket(base, quote string) Market {
	return Market{
		BaseUnit:  base,
		QuoteUnit: quote,
	}
}

// NewMarketFromString returns a new Market from a string
// "btc/usdt" -> Market{BTC, Usdt}
// NOTE: string should contain "/" delimiter
func NewMarketFromString(s string) (Market, bool) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return Market{}, false
	}
	return NewMarket(parts[0], parts[1]), true
}

// TradeEvent is a generic container
// for trades received from providers.
type TradeEvent struct {
	Source    DriverType
	Market    Market // e.g. `btc/usdt`
	Price     decimal.Decimal
	Amount    decimal.Decimal
	Total     decimal.Decimal
	TakerType TakerType
	CreatedAt time.Time
}
