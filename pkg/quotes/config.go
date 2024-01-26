package quotes

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config interface {
	DriverType() DriverType
}

func NewConfigFromEnv(driver DriverType) (Config, error) {
	switch driver.String() {
	case DriverBinance.String():
		config := BinanceConfig{}
		return config, cleanenv.ReadEnv(&config)
	case DriverKraken.String():
		config := KrakenConfig{}
		return config, cleanenv.ReadEnv(&config)
	case DriverOpendax.String():
		config := OpendaxConfig{}
		return config, cleanenv.ReadEnv(&config)
	case DriverBitfaker.String():
		config := BitfakerConfig{}
		return config, cleanenv.ReadEnv(&config)
	case DriverUniswapV3Api.String():
		config := UniswapV3ApiConfig{}
		return config, cleanenv.ReadEnv(&config)
	case DriverUniswapV3Geth.String():
		config := UniswapV3GethConfig{}
		return config, cleanenv.ReadEnv(&config)
	case DriverSyncswap.String():
		config := SyncswapConfig{}
		return config, cleanenv.ReadEnv(&config)
	default:
		return nil, fmt.Errorf("driver config is not supported: %s", driver)
	}
}

type BinanceConfig struct {
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

func (BinanceConfig) DriverType() DriverType {
	return DriverBinance
}

type KrakenConfig struct {
	URL             string             `yaml:"url" env:"QUOTES_KRAKEN_URL" env-default:"wss://ws.kraken.com"`
	ReconnectPeriod time.Duration      `yaml:"period" env:"QUOTES_KRAKEN_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler    TradeSamplerConfig `yaml:"trade_sampler"`
}

func (KrakenConfig) DriverType() DriverType {
	return DriverKraken
}

type OpendaxConfig struct {
	URL             string             `yaml:"url" env:"QUOTES_OPENDAX_URL" env-default:"wss://alpha.yellow.org/api/v1/finex/ws"`
	ReconnectPeriod time.Duration      `yaml:"period" env:"QUOTES_OPENDAX_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler    TradeSamplerConfig `yaml:"trade_sampler"`
}

func (OpendaxConfig) DriverType() DriverType {
	return DriverOpendax
}

type BitfakerConfig struct {
	Period       time.Duration      `yaml:"period" env:"QUOTES_BITFAKER_PERIOD" env-default:"5s"`
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

func (BitfakerConfig) DriverType() DriverType {
	return DriverBitfaker
}

type UniswapV3ApiConfig struct {
	URL          string             `yaml:"url" env:"QUOTES_UNISWAP_V3_API_URL" env-default:"https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"`
	WindowSize   time.Duration      `yaml:"window_size" env:"QUOTES_UNISWAP_V3_API_WINDOW_SIZE" env-default:"2s"`
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

func (UniswapV3ApiConfig) DriverType() DriverType {
	return DriverUniswapV3Api
}

type UniswapV3GethConfig struct {
	URL            string             `yaml:"url" env:"QUOTES_UNISWAP_V3_GETH_URL" env-default:""`
	AssetsURL      string             `yaml:"assets_url" env:"QUOTES_UNISWAP_V3_GETH_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	FactoryAddress string             `yaml:"factory_address" env:"QUOTES_UNISWAP_V3_GETH_FACTORY_ADDRESS" env-default:"0x1F98431c8aD98523631AE4a59f267346ea31F984"`
	TradeSampler   TradeSamplerConfig `yaml:"trade_sampler"`
}

func (UniswapV3GethConfig) DriverType() DriverType {
	return DriverUniswapV3Geth
}

type SyncswapConfig struct {
	URL                       string             `yaml:"url" env:"QUOTES_SYNCSWAP_URL" env-default:""`
	AssetsURL                 string             `yaml:"assets_url" env:"QUOTES_SYNCSWAP_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	ClassicPoolFactoryAddress string             `yaml:"classic_pool_factory_address" env:"QUOTES_SYNCSWAP_CLASSIC_POOL_FACTORY_ADDRESS" env-default:"0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d"`
	TradeSampler              TradeSamplerConfig `yaml:"trade_sampler"`
}

func (SyncswapConfig) DriverType() DriverType {
	return DriverSyncswap
}

type TradeSamplerConfig struct {
	Enabled           bool `yaml:"enabled" env:"QUOTES_TRADE_SAMPLER_ENABLED"`
	DefaultPercentage int  `yaml:"default_percentage" env:"QUOTES_TRADE_SAMPLER_DEFAULT_PERCENTAGE"`
}
