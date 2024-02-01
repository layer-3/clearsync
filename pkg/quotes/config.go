package quotes

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Driver DriverType `yaml:"driver" env:"QUOTES_DRIVER" env-default:"binance"`

	Binance       BinanceConfig       `yaml:"binance"`
	Kraken        KrakenConfig        `yaml:"kraken"`
	Opendax       OpendaxConfig       `yaml:"opendax"`
	Bitfaker      BitfakerConfig      `yaml:"bitfaker"`
	UniswapV3Api  UniswapV3ApiConfig  `yaml:"uniswap_v3_api"`
	UniswapV3Geth UniswapV3GethConfig `yaml:"uniswap_v3_geth"`
	Syncswap      SyncswapConfig      `yaml:"syncswap"`
}

func NewConfigFromFile(path string) (Config, error) {
	var config Config
	return config, cleanenv.ReadConfig(path, &config)
}

func NewConfigFromEnv() (Config, error) {
	var config Config
	return config, cleanenv.ReadEnv(&config)
}

type BinanceConfig struct {
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

type KrakenConfig struct {
	URL             string             `yaml:"url" env:"QUOTES_KRAKEN_URL" env-default:"wss://ws.kraken.com"`
	ReconnectPeriod time.Duration      `yaml:"period" env:"QUOTES_KRAKEN_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler    TradeSamplerConfig `yaml:"trade_sampler"`
}

type OpendaxConfig struct {
	URL             string             `yaml:"url" env:"QUOTES_OPENDAX_URL" env-default:"wss://alpha.yellow.org/api/v1/finex/ws"`
	ReconnectPeriod time.Duration      `yaml:"period" env:"QUOTES_OPENDAX_RECONNECT_PERIOD" env-default:"5s"`
	TradeSampler    TradeSamplerConfig `yaml:"trade_sampler"`
}

type BitfakerConfig struct {
	Period       time.Duration      `yaml:"period" env:"QUOTES_BITFAKER_PERIOD" env-default:"5s"`
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

type UniswapV3ApiConfig struct {
	URL          string             `yaml:"url" env:"QUOTES_UNISWAP_V3_API_URL" env-default:"https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"`
	WindowSize   time.Duration      `yaml:"window_size" env:"QUOTES_UNISWAP_V3_API_WINDOW_SIZE" env-default:"2s"`
	TradeSampler TradeSamplerConfig `yaml:"trade_sampler"`
}

type UniswapV3GethConfig struct {
	URL            string             `yaml:"url" env:"QUOTES_UNISWAP_V3_GETH_URL" env-default:""`
	AssetsURL      string             `yaml:"assets_url" env:"QUOTES_UNISWAP_V3_GETH_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	FactoryAddress string             `yaml:"factory_address" env:"QUOTES_UNISWAP_V3_GETH_FACTORY_ADDRESS" env-default:"0x1F98431c8aD98523631AE4a59f267346ea31F984"`
	TradeSampler   TradeSamplerConfig `yaml:"trade_sampler"`
}

type SyncswapConfig struct {
	URL                       string             `yaml:"url" env:"QUOTES_SYNCSWAP_URL" env-default:""`
	AssetsURL                 string             `yaml:"assets_url" env:"QUOTES_SYNCSWAP_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	ClassicPoolFactoryAddress string             `yaml:"classic_pool_factory_address" env:"QUOTES_SYNCSWAP_CLASSIC_POOL_FACTORY_ADDRESS" env-default:"0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d"`
	TradeSampler              TradeSamplerConfig `yaml:"trade_sampler"`
}

type TradeSamplerConfig struct {
	Enabled           bool `yaml:"enabled" env:"QUOTES_TRADE_SAMPLER_ENABLED"`
	DefaultPercentage int  `yaml:"default_percentage" env:"QUOTES_TRADE_SAMPLER_DEFAULT_PERCENTAGE"`
}
