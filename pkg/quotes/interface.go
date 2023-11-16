// Package quotes implements multiple price feed adapters.
package quotes

import (
	"fmt"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
)

var logger = log.Logger("quotes")

type Driver interface {
	Subscribe(market Market) error
	Start(markets []Market) error
	Stop() error
}

func NewDriver(config QuotesConfig, outbox chan<- TradeEvent) (Driver, error) {
	allDrivers := map[DriverType]Driver{
		DriverBinance:  NewBinance(config, outbox),
		DriverKraken:   NewKraken(config, outbox),
		DriverOpendax:  NewOpendax(config, outbox),
		DriverBitfaker: NewBitfaker(config, outbox),
	}

	driver, ok := allDrivers[config.Driver]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", config.Driver.String())
	}
	return driver, nil
}

type QuotesConfig struct {
	URL          string             `yaml:"url" env:"FINEX_QUOTES_URL" env-default:"wss://alpha.yellow.org/api/v1/finex/ws"`
	Driver       DriverType         `yaml:"driver" env:"FINEX_QUOTES_DRIVER" env-default:"opendax"`
	APIKey       string             `yaml:"api_key" env:"FINEX_QUOTES_API_KEY"`
	Period       time.Duration      `yaml:"period" env:"FINEX_QUOTES_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

type TradeSamplerConfig struct {
	Enabled           bool
	DefaultPercentage int
}

type Market struct {
	BaseUnit  string // e.g. `usdt`
	QuoteUnit string // e.g. `btc`
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
