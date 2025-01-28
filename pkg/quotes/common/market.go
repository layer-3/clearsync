package common

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	BaseUnit  string // e.g. `lube` // Base currency
	QuoteUnit string // e.g. `usdt` // Quote currency
	MainQuote string // e.g. `usd` // Main quote currency
	// If ConvertTo specified, the index driver will convert quote currency to the specified one.
	ConvertTo string // e.g. `usdc`
}

func NewMarket(base, quote string) Market {
	return Market{
		BaseUnit:  strings.ToLower(base),
		QuoteUnit: strings.ToLower(quote),
	}
}

func NewMarketWithMainQuote(base, quote, mainQuote string) Market {
	return Market{
		BaseUnit:  strings.ToLower(base),
		QuoteUnit: strings.ToLower(quote),
		MainQuote: strings.ToLower(mainQuote),
	}
}

func NewMarketDerived(base, quote, convertQuoteTo string) Market {
	return Market{
		BaseUnit:  strings.ToLower(base),
		QuoteUnit: strings.ToLower(quote),
		ConvertTo: strings.ToLower(convertQuoteTo),
	}
}

// NewMarketFromString returns a new Market from a string
// "btc/usdt" -> Market{btc, usdt}
// NOTE: string must contain "/" delimiter
func NewMarketFromString(s string) (Market, bool) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return Market{}, false
	}
	return NewMarket(parts[0], parts[1]), true
}

// String returns a string representation of the market.
// Example: `Market{btc, usdt}` -> "btc/usdt"
func (m Market) String() string {
	if m.MainQuote != "" {
		return fmt.Sprintf("%s/%s", m.BaseUnit, m.MainQuote)
	}
	return fmt.Sprintf("%s/%s", m.BaseUnit, m.QuoteUnit)
}

func (m Market) StringWithoutMain() string {
	return fmt.Sprintf("%s/%s", m.BaseUnit, m.QuoteUnit)
}

func (m Market) ApplyMainQuote() Market {
	if m.MainQuote == "" {
		return m
	}

	m.QuoteUnit = m.MainQuote
	return m
}

func (m Market) Base() string {
	return m.BaseUnit
}

func (m Market) Quote() string {
	return m.QuoteUnit
}

func (m Market) LegacyQuote() string {
	return m.MainQuote
}

func (m Market) IsEmpty() bool {
	return m.BaseUnit == "" || m.QuoteUnit == ""
}

func (m Market) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *Market) UnmarshalJSON(raw []byte) error {
	var rawParsed string
	if err := json.Unmarshal(raw, &rawParsed); err != nil {
		return err
	}

	parts := strings.Split(rawParsed, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("invalid market format: got '%s' instead of e.g. 'btc/usdt'", rawParsed)
	}

	*m = Market{BaseUnit: parts[0], QuoteUnit: parts[1]}
	return nil
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

// SortTradeEventsInPlace sorts the given trade events
// by their Unix timestamps in descending order (first comes is the newest trade).
// The input slice is modified in place.
// It is safe to call this function with an empty slice.
func SortTradeEventsInPlace(trades []TradeEvent) {
	slices.SortFunc(trades, func(a, b TradeEvent) int {
		if a.CreatedAt.Before(b.CreatedAt) {
			return 1
		}
		if a.CreatedAt.After(b.CreatedAt) {
			return -1
		}
		return 0
	})
}
