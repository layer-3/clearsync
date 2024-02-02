package quotes

import (
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	BaseUnit  string // e.g. `btc` in `btc/usdt`
	QuoteUnit string // e.g. `usdt` in `btc/usdt`
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
