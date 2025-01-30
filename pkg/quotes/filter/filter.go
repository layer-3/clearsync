package filter

import "github.com/layer-3/clearsync/pkg/quotes/common"

type Filter interface {
	Allow(trade common.TradeEvent) bool
}

type Type string

const (
	TypeMAD       Type = "mad"
	TypeSampler   Type = "sampler"
	TypePriceDiff Type = "price_diff"
	TypeDisabled  Type = "disabled"
)

type Config struct {
	Type Type `yaml:"filter_type" env:"TYPE" env-default:"disabled"`

	Sampler   SamplerConfig   `yaml:"sampler" env-prefix:"SAMPLER_"`
	PriceDiff PriceDiffConfig `yaml:"price_diff" env-prefix:"PRICE_DIFF_"`
	MAD       MADConfig       `yaml:"mad" env-prefix:"MAD_"`
}

// New creates a new filter based on the provided configuration.
//
// Params:
// - conf: the configuration for the filter
// - history: the historical data driver. Required for MAD filter, not used for other filters.
func New(conf Config, history common.HistoricalDataDriver) (Filter, error) {
	switch conf.Type {
	case TypeMAD:
		return newMADFilter(conf.MAD, history)
	case TypeSampler:
		return newSamplerFilter(conf.Sampler), nil
	case TypePriceDiff:
		return newPriceDiffFilter(conf.PriceDiff), nil
	default:
		return newDisabledFilter(), nil
	}
}

type DisabledFilter struct{}

func newDisabledFilter() DisabledFilter {
	return DisabledFilter{}
}

func (f DisabledFilter) Allow(common.TradeEvent) bool {
	return true
}
