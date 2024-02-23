package quotes

type Filter interface {
	Allow(trade TradeEvent) bool
}

type FilterType string

const (
	SamplerFilterType   FilterType = "sampler"
	PriceDiffFilterType FilterType = "price_diff"
	DisabledFilterType  FilterType = "disabled"
)

func FilterFactory(conf FilterConfig) Filter {
	switch FilterType(conf.FilterType) {
	case SamplerFilterType:
		return newSamplerFilter(conf.SamplerFilter)
	case PriceDiffFilterType:
		return newPriceDiffFilter(conf.PriceDiffFilter)
	default:
		return newDisabledFilter()
	}
}

type DisabledFilter struct{}

func newDisabledFilter() DisabledFilter {
	return DisabledFilter{}
}

func (f DisabledFilter) Allow(trade TradeEvent) bool {
	return true
}
