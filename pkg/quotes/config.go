package quotes

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Driver DriverType `yaml:"driver" env:"QUOTES_DRIVER" env-default:"index"`

	Binance       BinanceConfig       `yaml:"binance" env-prefix:"QUOTES_BINANCE_"`
	Kraken        KrakenConfig        `yaml:"kraken" env-prefix:"QUOTES_KRAKEN_"`
	Opendax       OpendaxConfig       `yaml:"opendax" env-prefix:"QUOTES_OPENDAX_"`
	Bitfaker      BitfakerConfig      `yaml:"bitfaker" env-prefix:"QUOTES_BITFAKER_"`
	UniswapV3Api  UniswapV3ApiConfig  `yaml:"uniswap_v3_api" env-prefix:"QUOTES_UNISWAP_V3_API_"`
	UniswapV3Geth UniswapV3GethConfig `yaml:"uniswap_v3_geth" env-prefix:"QUOTES_UNISWAP_V3_GETH_"`
	Syncswap      SyncswapConfig      `yaml:"syncswap" env-prefix:"QUOTES_SYNCSWAP_"`
	Quickswap     QuickswapConfig     `yaml:"quickswap" env-prefix:"QUOTES_QUICKSWAP_"`
	Index         IndexConfig         `yaml:"index" env-prefix:"QUOTES_INDEX_"`
}

type DriverConfig interface {
	DriverType() DriverType
}

func NewConfigFromFile(path string) (Config, error) {
	var config Config
	return config, cleanenv.ReadConfig(path, &config)
}

func NewConfigFromEnv() (Config, error) {
	var config Config
	return config, cleanenv.ReadEnv(&config)
}

func ToConfig(driver DriverConfig) Config {
	config := Config{Driver: driver.DriverType()}

	switch driver.DriverType() {
	case DriverIndex:
		config.Index = driver.(IndexConfig)
	case DriverBinance:
		config.Binance = driver.(BinanceConfig)
	case DriverKraken:
		config.Kraken = driver.(KrakenConfig)
	case DriverOpendax:
		config.Opendax = driver.(OpendaxConfig)
	case DriverBitfaker:
		config.Bitfaker = driver.(BitfakerConfig)
	case DriverUniswapV3Api:
		config.UniswapV3Api = driver.(UniswapV3ApiConfig)
	case DriverUniswapV3Geth:
		config.UniswapV3Geth = driver.(UniswapV3GethConfig)
	case DriverSyncswap:
		config.Syncswap = driver.(SyncswapConfig)
	}
	return config
}

type IndexConfig struct {
	TradesCached   int                 `yaml:"trades_cached" env:"TRADES_CACHED" env-default:"20"`
	BufferMinutes  int                 `yaml:"buffer_minutes" env:"BUFFER_MINUTES" env-default:"1"`
	DriverConfigs  []Config            `yaml:"drivers" env:"DRIVERS"`
	MarketsMapping map[string][]string `yaml:"markets_mapping" env:"MARKETS_MAPPING"`
}

func (IndexConfig) DriverType() DriverType {
	return DriverIndex
}

type BinanceConfig struct {
	Filter     FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
	USDCtoUSDT bool         `yaml:"usdc_to_usdt" env-prefix:"USDC_TO_USDT_" env-default:"true"`
}

func (BinanceConfig) DriverType() DriverType {
	return DriverBinance
}

type KrakenConfig struct {
	URL             string        `yaml:"url" env:"URL" env-default:"wss://ws.kraken.com"`
	ReconnectPeriod time.Duration `yaml:"period" env:"RECONNECT_PERIOD" env-default:"5s"`
	Filter          FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

func (KrakenConfig) DriverType() DriverType {
	return DriverKraken
}

type OpendaxConfig struct {
	URL             string        `yaml:"url" env:"URL" env-default:"wss://alpha.yellow.org/api/v1/finex/ws"`
	ReconnectPeriod time.Duration `yaml:"period" env:"RECONNECT_PERIOD" env-default:"5s"`
	Filter          FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

func (OpendaxConfig) DriverType() DriverType {
	return DriverOpendax
}

type BitfakerConfig struct {
	Period time.Duration `yaml:"period" env:"PERIOD" env-default:"5s"`
	Filter FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

func (BitfakerConfig) DriverType() DriverType {
	return DriverBitfaker
}

type UniswapV3ApiConfig struct {
	URL        string        `yaml:"url" env:"URL" env-default:"https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"`
	WindowSize time.Duration `yaml:"window_size" env:"WINDOW_SIZE" env-default:"2s"`
	Filter     FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

func (UniswapV3ApiConfig) DriverType() DriverType {
	return DriverUniswapV3Api
}

type UniswapV3GethConfig struct {
	URL            string       `yaml:"url" env:"GETH_URL" env-default:""`
	AssetsURL      string       `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	FactoryAddress string       `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0x1F98431c8aD98523631AE4a59f267346ea31F984"`
	Filter         FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

func (UniswapV3GethConfig) DriverType() DriverType {
	return DriverUniswapV3Geth
}

type SyncswapConfig struct {
	URL                       string       `yaml:"url" env:"URL" env-default:""`
	AssetsURL                 string       `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	ClassicPoolFactoryAddress string       `yaml:"classic_pool_factory_address" env:"CLASSIC_POOL_FACTORY_ADDRESS" env-default:"0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d"`
	Filter                    FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

func (SyncswapConfig) DriverType() DriverType {
	return DriverSyncswap
}

type QuickswapConfig struct {
	URL       string `yaml:"url" env:"URL" env-default:""`
	AssetsURL string `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	// PoolFactoryAddress is the address of the factory contract.
	// See docs at https://docs.quickswap.exchange/technical-reference/smart-contracts/v3/factory.
	// Note that the contract used in this lib is compiled from https://github.com/code-423n4/2022-09-quickswap.
	PoolFactoryAddress string       `yaml:"pool_factory_address" env:"POOL_FACTORY_ADDRESS" env-default:"0x411b0fAcC3489691f28ad58c47006AF5E3Ab3A28"`
	Filter             FilterConfig `yaml:"filter" env-prefix:"FILTER_"`
}

func (QuickswapConfig) DriverType() DriverType {
	return DriverQuickswap
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
