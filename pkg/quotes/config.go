package quotes

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Drivers []DriverType `yaml:"drivers" env:"QUOTES_DRIVERS" env-default:"binance,syncswap"`
	Rpc     RpcConfig    `yaml:"rpc" env-prefix:"QUOTES_RPC_"`
	Index   IndexConfig  `yaml:"index" env-prefix:"QUOTES_INDEX_"`

	Binance   BinanceConfig   `yaml:"binance" env-prefix:"QUOTES_BINANCE_"`
	Kraken    KrakenConfig    `yaml:"kraken" env-prefix:"QUOTES_KRAKEN_"`
	Mexc      MexcConfig      `yaml:"mexc" env-prefix:"QUOTES_MEXC_"`
	Opendax   OpendaxConfig   `yaml:"opendax" env-prefix:"QUOTES_OPENDAX_"`
	Bitfaker  BitfakerConfig  `yaml:"bitfaker" env-prefix:"QUOTES_BITFAKER_"`
	UniswapV3 UniswapV3Config `yaml:"uniswap_v3" env-prefix:"QUOTES_UNISWAP_V3_"`
	Syncswap  SyncswapConfig  `yaml:"syncswap" env-prefix:"QUOTES_SYNCSWAP_"`
	Quickswap QuickswapConfig `yaml:"quickswap" env-prefix:"QUOTES_QUICKSWAP_"`
	SectaV2   SectaV2Config   `yaml:"secta_v2" env-prefix:"QUOTES_SECTA_V2_"`
	SectaV3   SectaV3Config   `yaml:"secta_v3" env-prefix:"QUOTES_SECTA_V3_"`
	LynexV2   LynexV2Config   `yaml:"lynex_v2" env-prefix:"QUOTES_LYNEX_V2_"`
	LynexV3   LynexV3Config   `yaml:"lynex_v3" env-prefix:"QUOTES_LYNEX_V3_"`
}

func NewConfigFromFile(path string) (config Config, err error) {
	return config, cleanenv.ReadConfig(path, &config)
}

func NewConfigFromEnv() (config Config, err error) {
	return config, cleanenv.ReadEnv(&config)
}

type RpcConfig struct {
	Ethereum string `yaml:"ethereum" env:"ETHEREUM"`
	Polygon  string `yaml:"polygon" env:"POLYGON"`
	Linea    string `yaml:"linea" env:"LINEA"`
}

type IndexConfig struct {
	TradesCached   int                 `yaml:"trades_cached" env:"TRADES_CACHED" env-default:"20"`
	BufferSeconds  int                 `yaml:"buffer_minutes" env:"BUFFER_MINUTES" env-default:"5"`
	MarketsMapping map[string][]string `yaml:"markets_mapping" env:"MARKETS_MAPPING"`
	// MaxPriceDiff has default of `0.2` because our default leverage is 5x,
	// and so if the user opens order on his full balance, he'll get liquidated on 20% price change.
	MaxPriceDiff string `yaml:"max_price_diff" env:"MAX_PRICE_DIFF" env-default:"0.2"`
}

type BinanceConfig struct {
	USDCtoUSDT         bool          `yaml:"usdc_to_usdt" env:"USDC_TO_USDT" env-default:"true"`
	BatchPeriod        time.Duration `yaml:"batch_period" env:"BATCH_PERIOD" env-default:"5s"`
	AssetsUpdatePeriod time.Duration `yaml:"assets_update_period" env:"ASSETS_UPDATE_PERIOD" env-default:"5m"`
	IdlePeriod         time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter             FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

// TODO: add idle period
type MexcConfig struct {
	USDCtoUSDT         bool          `yaml:"usdc_to_usdt" env:"USDC_TO_USDT" env-default:"true"`
	BatchPeriod        time.Duration `yaml:"batch_period" env:"BATCH_PERIOD" env-default:"5s"`
	AssetsUpdatePeriod time.Duration `yaml:"assets_update_period" env:"ASSETS_UPDATE_PERIOD" env-default:"5m"`
	URL                string        `yaml:"url" env:"URL" env-default:"wss://ws.mexc.com/ws"`
	ReconnectPeriod    time.Duration `yaml:"period" env:"RECONNECT_PERIOD" env-default:"5s"`
	Filter             FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

// TODO: add idle period
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
	Filter  FilterConfig                `yaml:"filter" env-prefix:"FILTER_"`
	Markets map[string]FakeMarketConfig `yaml:"markets" env:"MARKETS"`
}

type FakeMarketConfig struct {
	StartPrice       float64       `yaml:"start_price" env:"START_PRICE"`
	PriceVolatility  float64       `yaml:"price_volatility" env:"PRICE_VOLATILITY"`
	StartAmount      float64       `yaml:"start_pamount" env:"START_AMOUNT"`
	AmountVolatility float64       `yaml:"amount_volatility" env:"AMOUNT_VOLATILITY"`
	Period           time.Duration `yaml:"period" env:"PERIOD"`
}

type UniswapV3Config struct {
	AssetsURL      string        `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/1/assets.json"`
	MappingURL     string        `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/1/mapping.json"`
	MarketsURL     string        `yaml:"markets_url" env:"MARKETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/1/markets.json"`
	FactoryAddress string        `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0x1F98431c8aD98523631AE4a59f267346ea31F984"`
	IdlePeriod     time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter         FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type SyncswapConfig struct {
	AssetsURL                 string        `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	MappingURL                string        `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/mapping.json"`
	MarketsURL                string        `yaml:"markets_url" env:"MARKETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/markets.json"`
	ClassicPoolFactoryAddress string        `yaml:"classic_pool_factory_address" env:"CLASSIC_POOL_FACTORY_ADDRESS" env-default:"0x37BAc764494c8db4e54BDE72f6965beA9fa0AC2d"`
	StablePoolFactoryAddress  string        `yaml:"stable_pool_factory_address" env:"STABLE_POOL_FACTORY_ADDRESS" env-default:"0xE4CF807E351b56720B17A59094179e7Ed9dD3727"`
	StablePoolMarkets         []string      `yaml:"stable_pool_markets" env:"STABLE_POOL_MARKETS" env-default:"usdt/usdc"` // `env-default` tag value is a comma separated list of markets as in `usdt/usdc, usdc/dai`
	IdlePeriod                time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter                    FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type QuickswapConfig struct {
	AssetsURL  string `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/137/assets.json"`
	MappingURL string `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/137/mapping.json"`
	MarketsURL string `yaml:"markets_url" env:"MARKETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/137/markets.json"`
	// PoolFactoryAddress is the address of the factory contract.
	// See docs at https://docs.quickswap.exchange/technical-reference/smart-contracts/v3/factory.
	// Note that the contract used in this lib is compiled from https://github.com/code-423n4/2022-09-quickswap.
	PoolFactoryAddress string        `yaml:"pool_factory_address" env:"POOL_FACTORY_ADDRESS" env-default:"0x411b0fAcC3489691f28ad58c47006AF5E3Ab3A28"`
	IdlePeriod         time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter             FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type SectaV2Config struct {
	AssetsURL      string        `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	MappingURL     string        `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/mapping.json"`
	MarketsURL     string        `yaml:"markets_url" env:"MARKETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/markets.json"`
	FactoryAddress string        `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0x8Ad39bf99765E24012A28bEb0d444DE612903C43"`
	IdlePeriod     time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter         FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type SectaV3Config struct {
	AssetsURL      string        `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	MappingURL     string        `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/mapping.json"`
	MarketsURL     string        `yaml:"markets_url" env:"MARKETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/markets.json"`
	FactoryAddress string        `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0x9BD425a416A276C72a13c13bBd8145272680Cf07"`
	IdlePeriod     time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter         FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type LynexV2Config struct {
	AssetsURL         string        `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	MappingURL        string        `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/mapping.json"`
	FactoryAddress    string        `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0xBc7695Fd00E3b32D08124b7a4287493aEE99f9ee"`
	StablePoolMarkets []string      `yaml:"stable_pool_markets" env:"STABLE_POOL_MARKETS" env-default:"usdt/usdc"` // `env-default` tag value is a comma separated list of markets as in `usdt/usdc, usdc/dai`
	IdlePeriod        time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter            FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type LynexV3Config struct {
	AssetsURL      string        `yaml:"assets_url" env:"ASSETS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/assets.json"`
	MappingURL     string        `yaml:"mappings_url" env:"MAPPINGS_URL" env-default:"https://raw.githubusercontent.com/layer-3/clearsync/master/networks/59144/mapping.json"`
	FactoryAddress string        `yaml:"factory_address" env:"FACTORY_ADDRESS" env-default:"0x622b2c98123D303ae067DB4925CD6282B3A08D0F"`
	IdlePeriod     time.Duration `yaml:"idle_period" env:"IDLE_PERIOD" env-default:"30m"`
	Filter         FilterConfig  `yaml:"filter" env-prefix:"FILTER_"`
}

type SamplerFilterConfig struct {
	Percentage int64 `yaml:"percentage" env:"PERCENTAGE" env-default:"5"`
}

type PriceDiffFilterConfig struct {
	Threshold string `yaml:"threshold" env:"THRESHOLD" env-default:"5"`
}

type FilterConfig struct {
	FilterType FilterType `yaml:"filter_type" env:"TYPE" env-default:"disabled"`

	SamplerFilter   SamplerFilterConfig   `yaml:"sampler" env-prefix:"SAMPLER_"`
	PriceDiffFilter PriceDiffFilterConfig `yaml:"price_diff" env-prefix:"PRICE_DIFF_"`
}
