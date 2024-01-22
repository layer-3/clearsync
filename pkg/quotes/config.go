package quotes

import "time"

type Config struct {
	URL             string             `yaml:"url" env:"QUOTES_URL" env-default:""`
	AssetsURL       string             `yaml:"assets_url" env:"QUOTES_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	Driver          DriverType         `yaml:"driver" env:"QUOTES_DRIVER" env-default:"binance"`
	ReconnectPeriod time.Duration      `yaml:"period" env:"QUOTES_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler    TradeSamplerConfig `yaml:"trade_sampler"`
}

type TradeSamplerConfig struct {
	Enabled           bool `yaml:"enabled" env:"QUOTES_TRADE_SAMPLER_ENABLED"`
	DefaultPercentage int  `yaml:"default_percentage" env:"QUOTES_TRADE_SAMPLER_DEFAULT_PERCENTAGE"`
}
