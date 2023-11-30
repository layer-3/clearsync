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

func NewDriver(config Config, outbox chan<- TradeEvent) (Driver, error) {
	allDrivers := map[DriverType]Driver{
		DriverBinance:   NewBinance(config, outbox),
		DriverKraken:    NewKraken(config, outbox),
		DriverOpendax:   NewOpendax(config, outbox),
		DriverBitfaker:  NewBitfaker(config, outbox),
		DriverUniswapV3: NewUniswapV3(config, outbox),
	}

	driver, ok := allDrivers[config.Driver]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", config.Driver.String())
	}
	return driver, nil
}

type Config struct {
	URL             string             `yaml:"url" env:"QUOTES_URL" env-default:""`
	Driver          DriverType         `yaml:"driver" env:"QUOTES_DRIVER" env-default:"binance"`
	ReconnectPeriod time.Duration      `yaml:"period" env:"QUOTES_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler    TradeSamplerConfig `yaml:"trade_sampler"`
}

type TradeSamplerConfig struct {
	Enabled           bool `yaml:"enabled" env:"QUOTES_TRADE_SAMPLER_ENABLED"`
	DefaultPercentage int  `yaml:"default_percentage" env:"QUOTES_TRADE_SAMPLER_DEFAULT_PERCENTAGE"`
}

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
