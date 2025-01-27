package filter

import "github.com/layer-3/clearsync/pkg/quotes/common"

type Filter interface {
	Allow(trade common.TradeEvent) bool
}

type Type string

const (
	TypeSampler   Type = "sampler"
	TypePriceDiff Type = "price_diff"
	TypeDisabled  Type = "disabled"
)

type Config struct {
	FilterType Type `yaml:"filter_type" env:"TYPE" env-default:"disabled"`

	SamplerFilter   SamplerConfig   `yaml:"sampler" env-prefix:"SAMPLER_"`
	PriceDiffFilter PriceDiffConfig `yaml:"price_diff" env-prefix:"PRICE_DIFF_"`
}

func New(conf Config) Filter {
	switch conf.FilterType {
	case TypeSampler:
		return newSamplerFilter(conf.SamplerFilter)
	case TypePriceDiff:
		return newPriceDiffFilter(conf.PriceDiffFilter)
	default:
		return newDisabledFilter()
	}
}

type DisabledFilter struct{}

func newDisabledFilter() DisabledFilter {
	return DisabledFilter{}
}

func (f DisabledFilter) Allow(common.TradeEvent) bool {
	return true
}
