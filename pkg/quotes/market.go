package quotes

import (
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	BaseUnit  string // e.g. `btc` in `btcusdt`
	QuoteUnit string // e.g. `usdt` in `btcusdt`
}

// TradeEvent is a generic container
// for trades received from providers.
type TradeEvent struct {
	Source    DriverType
	Market    string // e.g. `btcusdt`
	Price     decimal.Decimal
	Amount    decimal.Decimal
	Total     decimal.Decimal
	TakerType TakerType
	CreatedAt time.Time
}
