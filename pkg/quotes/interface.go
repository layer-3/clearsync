// Package quotes implements multiple price feed adapters.
package quotes

import (
	"errors"
	"fmt"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"
)

var logger = log.Logger("quotes")

type Driver interface {
	Name() DriverType
	Start() error
	Stop() error
	Subscribe(market Market) error
	Unsubscribe(market Market) error
}

func NewDriver(config Config, outbox chan<- TradeEvent) (Driver, error) {
	allDrivers := map[DriverType]func(Config, chan<- TradeEvent) Driver{
		DriverBinance:   newBinance,
		DriverKraken:    newKraken,
		DriverOpendax:   newOpendax,
		DriverBitfaker:  newBitfaker,
		DriverUniswapV3: newUniswapV3,
		DriverIndex:     newIndex,
	}

	driver, ok := allDrivers[config.Driver]
	if !ok {
		return nil, fmt.Errorf("invalid driver type: %v", config.Driver.String())
	}
	return driver(config, outbox), nil
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

var (
	ErrNotSubbed     = errors.New("market not subscribed")
	ErrAlreadySubbed = errors.New("market already subscribed")
	ErrFailedSub     = errors.New("failed to subscribe to market")
	ErrFailedUnsub   = errors.New("failed to unsubscribe from market")
)
