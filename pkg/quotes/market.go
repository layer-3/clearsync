package quotes

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Market struct {
	baseUnit  string // e.g. `lube` // Base currency
	quoteUnit string // e.g. `usdt` // Quote currency
	mainQuote string // e.g. `usd` // Main quote currency
	// If convertTo specified, the index driver will convert quote currency to the specified one.
	convertTo string // e.g. `usdc`
}

func NewMarket(base, quote string) Market {
	return Market{
		baseUnit:  strings.ToLower(base),
		quoteUnit: strings.ToLower(quote),
	}
}

func NewMarketWithMainQuote(base, quote, mainQuote string) Market {
	return Market{
		baseUnit:  strings.ToLower(base),
		quoteUnit: strings.ToLower(quote),
		mainQuote: strings.ToLower(mainQuote),
	}
}

func NewMarketDerived(base, quote, convertQuoteTo string) Market {
	return Market{
		baseUnit:  strings.ToLower(base),
		quoteUnit: strings.ToLower(quote),
		convertTo: strings.ToLower(convertQuoteTo),
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
	if m.mainQuote != "" {
		return fmt.Sprintf("%s/%s", m.baseUnit, m.mainQuote)
	}
	return fmt.Sprintf("%s/%s", m.baseUnit, m.quoteUnit)
}

func (m Market) StringWithoutMain() string {
	return fmt.Sprintf("%s/%s", m.baseUnit, m.quoteUnit)
}

func (m Market) ApplyMainQuote() Market {
	if m.mainQuote == "" {
		return m
	}

	m.quoteUnit = m.mainQuote
	return m
}

func (m Market) Base() string {
	return m.baseUnit
}

func (m Market) Quote() string {
	return m.quoteUnit
}

func (m Market) LegacyQuote() string {
	return m.mainQuote
}

func (m Market) IsEmpty() bool {
	return m.baseUnit == "" || m.quoteUnit == ""
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

	*m = Market{baseUnit: parts[0], quoteUnit: parts[1]}
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
