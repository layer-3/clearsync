package quotes

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Drivers []DriverType `yaml:"drivers" env:"QUOTES_DRIVERS" env-default:"binance,syncswap"`
	Index   IndexConfig  `yaml:"index" env-prefix:"QUOTES_INDEX_"`

	Binance       BinanceConfig       `yaml:"binance" env-prefix:"QUOTES_BINANCE_"`
	Kraken        KrakenConfig        `yaml:"kraken" env-prefix:"QUOTES_KRAKEN_"`
	Opendax       OpendaxConfig       `yaml:"opendax" env-prefix:"QUOTES_OPENDAX_"`
	Bitfaker      BitfakerConfig      `yaml:"bitfaker" env-prefix:"QUOTES_BITFAKER_"`
	UniswapV3Api  UniswapV3ApiConfig  `yaml:"uniswap_v3_api" env-prefix:"QUOTES_UNISWAP_V3_API_"`
	UniswapV3Geth UniswapV3GethConfig `yaml:"uniswap_v3_geth" env-prefix:"QUOTES_UNISWAP_V3_GETH_"`
	Syncswap      SyncswapConfig      `yaml:"syncswap" env-prefix:"QUOTES_SYNCSWAP_"`
	Quickswap     QuickswapConfig     `yaml:"quickswap" env-prefix:"QUOTES_QUICKSWAP_"`
}

func (config Config) GetByDriverType(driver DriverType) (Config, error) {
	driverFound := false
	for _, d := range config.Drivers {
		if d == driver {
			driverFound = true
			break
		}
	}
	if !driverFound {
		return Config{}, fmt.Errorf("driver is not configured: %s", driver)
	}

	switch driver {
	case DriverBinance:
		return Config{Drivers: []DriverType{DriverBinance}, Binance: config.Binance}, nil
	case DriverKraken:
		return Config{Drivers: []DriverType{DriverKraken}, Kraken: config.Kraken}, nil
	case DriverOpendax:
		return Config{Drivers: []DriverType{DriverOpendax}, Opendax: config.Opendax}, nil
	case DriverBitfaker:
		return Config{Drivers: []DriverType{DriverBitfaker}, Bitfaker: config.Bitfaker}, nil
	case DriverUniswapV3Api:
		return Config{Drivers: []DriverType{DriverUniswapV3Api}, UniswapV3Api: config.UniswapV3Api}, nil
	case DriverUniswapV3Geth:
		return Config{Drivers: []DriverType{DriverUniswapV3Geth}, UniswapV3Geth: config.UniswapV3Geth}, nil
	case DriverSyncswap:
		return Config{Drivers: []DriverType{DriverSyncswap}, Syncswap: config.Syncswap}, nil
	case DriverQuickswap:
		return Config{Drivers: []DriverType{DriverQuickswap}, Quickswap: config.Quickswap}, nil
	default:
		return config, nil
	}
}

func NewConfigFromFile(path string) (Config, error) {
	var config Config
	return config, cleanenv.ReadConfig(path, &config)
}

func NewConfigFromEnv() (Config, error) {
	var config Config
	return config, cleanenv.ReadEnv(&config)
}

type IndexConfig struct {
	TradesCached   int                 `yaml:"trades_cached" env:"TRADES_CACHED" env-default:"20"`
	BufferMinutes  int                 `yaml:"buffer_minutes" env:"BUFFER_MINUTES" env-default:"15"`
	MarketsMapping map[string][]string `yaml:"markets_mapping" env:"MARKETS_MAPPING"`
}

type BinanceConfig struct {
	USDCtoUSDT bool         `yaml:"usdc_to_usdt" env:"USDC_TO_USDT" env-default:"true"`
	Filter     FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

type KrakenConfig struct {
	URL             string        `yaml:"url" env:"URL" env-default:"wss://ws.kraken.com"`
	ReconnectPeriod time.Duration `yaml:"period" env:"RECONNECT_PERIOD" env-default:"5s"`
	Filter          FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type OpendaxConfig struct {
	URL             string        `yaml:"url" env:"URL" env-default:"wss://alpha.yellow.org/api/v1/finex/ws"`
	ReconnectPeriod time.Duration `yaml:"period" env:"RECONNECT_PERIOD" env-default:"5s"`
	Filter          FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type BitfakerConfig struct {
	Period time.Duration `yaml:"period" env:"PERIOD" env-default:"5s"`
	Filter FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type UniswapV3ApiConfig struct {
	URL        string        `yaml:"url" env:"URL" env-default:"https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"`
	WindowSize time.Duration `yaml:"window_size" env:"WINDOW_SIZE" env-default:"2s"`
	Filter     FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type UniswapV3GethConfig struct {
	URL            string       `yaml:"url" env:"URL"`
	AssetsURL      string       `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	FactoryAddress string       `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0x1F98431c8aD98523631AE4a59f267346ea31F984"`
	Filter         FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

type SyncswapConfig struct {
	URL                       string       `yaml:"url" env:"URL"`
	AssetsURL                 string       `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/e76e8afce1b9d5d62465e0f568c66e1301b5b4b4/networks/59144/assets.json"`
	MappingURL                string       `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/e76e8afce1b9d5d62465e0f568c66e1301b5b4b4/networks/59144/mapping.json"`
	ClassicPoolFactoryAddress string       `yaml:"classic_pool_factory_address" env:"CLASSIC_POOL_FACTORY_ADDRESS" env-default:"0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d"`
	StablePoolFactoryAddress  string       `yaml:"stable_pool_factory_address" env:"STABLE_POOL_FACTORY_ADDRESS" env-default:"0xE4CF807E351b56720B17A59094179e7Ed9dD3727"`
	StablePoolMarkets         []string     `yaml:"stable_pool_markets" env:"STABLE_POOL_MARKETS" env-default:"usdt/usdc"` // `env-default` tag value is a comma separated list of markets as in `usdt/usdc,usdc/dai`
	Filter                    FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

type QuickswapConfig struct {
	URL       string `yaml:"url" env:"URL"`
	AssetsURL string `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	// PoolFactoryAddress is the address of the factory contract.
	// See docs at https://docs.quickswap.exchange/technical-reference/smart-contracts/v3/factory.
	// Note that the contract used in this lib is compiled from https://github.com/code-423n4/2022-09-quickswap.
	PoolFactoryAddress string       `yaml:"pool_factory_address" env:"POOL_FACTORY_ADDRESS" env-default:"0x411b0fAcC3489691f28ad58c47006AF5E3Ab3A28"`
	Filter             FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

type SamplerFilterConfig struct {
	Percentage int64 `yaml:"percentage" env:"PERCENTAGE" env-default:"5"`
}

type PriceDiffFilterConfig struct {
	Threshold string `yaml:"threshold" env:"THRESHOLD" env-default:"0.05"`
}

type FilterConfig struct {
	FilterType FilterType `yaml:"filter_type" env:"TYPE" env-default:"disabled"`

	SamplerFilter   SamplerFilterConfig   `yaml:"sampler" env-prefix:"SAMPLER_"`
	PriceDiffFilter PriceDiffFilterConfig `yaml:"price_diff" env-prefix:"PRICE_DIFF_"`
}
