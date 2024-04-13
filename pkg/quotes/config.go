package quotes

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Driver DriverType `yaml:"driver" env:"QUOTES_DRIVER" env-default:"index"`

	Binance       BinanceConfig       `yaml:"binance"`
	Kraken        KrakenConfig        `yaml:"kraken"`
	Opendax       OpendaxConfig       `yaml:"opendax"`
	Bitfaker      BitfakerConfig      `yaml:"bitfaker"`
	UniswapV3Api  UniswapV3ApiConfig  `yaml:"uniswap_v3_api"`
	UniswapV3Geth UniswapV3GethConfig `yaml:"uniswap_v3_geth"`
	Syncswap      SyncswapConfig      `yaml:"syncswap"`
	Quickswap     QuickswapConfig     `yaml:"quickswap"`
	Index         IndexConfig         `yaml:"index"`
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
	TradesCached   int                 `yaml:"trades_cached" env:"QUOTES_INDEX_TRADES_CACHED" env-default:"20"`
	BufferMinutes  int                 `yaml:"buffer_minutes" env:"QUOTES_INDEX_BUFFER_MINUTES" env-default:"15"`
	DriverConfigs  []Config            `yaml:"drivers"`
	MarketsMapping map[string][]string `yaml:"markets_mapping" env:"QUOTES_INDEX_MARKETS_MAPPING"`
}

func (IndexConfig) DriverType() DriverType {
	return DriverIndex
}

type BinanceConfig struct {
	Filter     FilterConfig `yaml:"filter" env-prefix:"QUOTES_BINANCE_FILTER_"`
	USDCtoUSDT bool         `yaml:"usdc_to_usdt" env:"QUOTES_BINANCE_USDC_TO_USDT" env-default:"true"`
}

func (BinanceConfig) DriverType() DriverType {
	return DriverBinance
}

type KrakenConfig struct {
	URL             string        `yaml:"url" env:"QUOTES_KRAKEN_URL" env-default:"wss://ws.kraken.com"`
	ReconnectPeriod time.Duration `yaml:"period" env:"QUOTES_KRAKEN_RECONNECT_PERIOD" env-default:"5s"`
	Filter          FilterConfig  `yaml:"filter" env-prefix:"QUOTES_KRAKEN_FILTER_"`
}

func (KrakenConfig) DriverType() DriverType {
	return DriverKraken
}

type OpendaxConfig struct {
	URL             string        `yaml:"url" env:"QUOTES_OPENDAX_URL" env-default:"wss://alpha.yellow.org/api/v1/finex/ws"`
	ReconnectPeriod time.Duration `yaml:"period" env:"QUOTES_OPENDAX_RECONNECT_PERIOD" env-default:"5s"`
	Filter          FilterConfig  `yaml:"filter" env-prefix:"QUOTES_OPENDAX_FILTER_"`
}

func (OpendaxConfig) DriverType() DriverType {
	return DriverOpendax
}

type BitfakerConfig struct {
	Period time.Duration `yaml:"period" env:"QUOTES_BITFAKER_PERIOD" env-default:"5s"`
	Filter FilterConfig  `yaml:"filter" env-prefix:"QUOTES_BITFAKER_FILTER_"`
}

func (BitfakerConfig) DriverType() DriverType {
	return DriverBitfaker
}

type UniswapV3ApiConfig struct {
	URL        string        `yaml:"url" env:"QUOTES_UNISWAP_V3_API_URL" env-default:"https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"`
	WindowSize time.Duration `yaml:"window_size" env:"QUOTES_UNISWAP_V3_API_WINDOW_SIZE" env-default:"2s"`
	Filter     FilterConfig  `yaml:"filter" env-prefix:"QUOTES_UNISWAP_V3_API_FILTER_"`
}

func (UniswapV3ApiConfig) DriverType() DriverType {
	return DriverUniswapV3Api
}

type UniswapV3GethConfig struct {
	URL            string       `yaml:"url" env:"QUOTES_UNISWAP_V3_GETH_URL" env-default:""`
	AssetsURL      string       `yaml:"assets_url" env:"QUOTES_UNISWAP_V3_GETH_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	FactoryAddress string       `yaml:"factory_address" env:"QUOTES_UNISWAP_V3_GETH_FACTORY_ADDRESS" env-default:"0x1F98431c8aD98523631AE4a59f267346ea31F984"`
	Filter         FilterConfig `yaml:"filter" env-prefix:"QUOTES_UNISWAP_V3_GETH_FILTER_"`
}

func (UniswapV3GethConfig) DriverType() DriverType {
	return DriverUniswapV3Geth
}

type SyncswapConfig struct {
	URL                       string       `yaml:"url" env:"QUOTES_SYNCSWAP_URL" env-default:""`
	AssetsURL                 string       `yaml:"assets_url" env:"QUOTES_SYNCSWAP_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	ClassicPoolFactoryAddress string       `yaml:"classic_pool_factory_address" env:"QUOTES_SYNCSWAP_CLASSIC_POOL_FACTORY_ADDRESS" env-default:"0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d"`
	StablePoolFactoryAddress  string       `yaml:"stable_pool_factory_address" env:"QUOTES_SYNCSWAP_STABLE_POOL_FACTORY_ADDRESS" env-default:"0xE4CF807E351b56720B17A59094179e7Ed9dD3727"`
	StablePoolMarkets         []string     `yaml:"stable_pool_markets" env:"QUOTES_SYNCSWAP_STABLE_POOL_MARKETS" env-default:"usdt/usdc"` // `env-default` tag value is a comma separated list of markets as in `usdt/usdc,usdc/dai`
	Filter                    FilterConfig `yaml:"filter" env-prefix:"QUOTES_SYNCSWAP_FILTER_"`
}

func (SyncswapConfig) DriverType() DriverType {
	return DriverSyncswap
}

type QuickswapConfig struct {
	URL       string `yaml:"url" env:"QUOTES_QUICKSWAP_URL" env-default:""`
	AssetsURL string `yaml:"assets_url" env:"QUOTES_QUICKSWAP_ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/mainnet/assets.json"`
	// PoolFactoryAddress is the address of the factory contract.
	// See docs at https://docs.quickswap.exchange/technical-reference/smart-contracts/v3/factory.
	// Note that the contract used in this lib is compiled from https://github.com/code-423n4/2022-09-quickswap.
	PoolFactoryAddress string       `yaml:"pool_factory_address" env:"QUOTES_QUICKSWAP_POOL_FACTORY_ADDRESS" env-default:"0x411b0fAcC3489691f28ad58c47006AF5E3Ab3A28"`
	Filter             FilterConfig `yaml:"filter" env-prefix:"QUOTES_QUICKSWAP_FILTER_"`
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
