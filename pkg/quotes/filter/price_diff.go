package filter

import (
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

var loggerPriceDiffFilter = log.Logger("price_diff_filter")

type PriceDiffConfig struct {
	Threshold string `yaml:"threshold" env:"THRESHOLD" env-default:"5"`
}

type PriceDiffFilter struct {
	threshold decimal.Decimal

	previousQuotes map[common.Market]decimal.Decimal
}

var (
	defaultThreshold = decimal.NewFromFloat(0.05)
)

func newPriceDiffFilter(conf PriceDiffConfig) *PriceDiffFilter {
	threshold, err := decimal.NewFromString(conf.Threshold)
	if err != nil {
		threshold = defaultThreshold
		loggerPriceDiffFilter.Warnf("failed to parse threshold: `%s`, using default: %s", conf.Threshold, defaultThreshold)
	}

	return &PriceDiffFilter{
		threshold:      threshold,
		previousQuotes: make(map[common.Market]decimal.Decimal),
	}
}

func (f *PriceDiffFilter) Allow(trade common.TradeEvent) bool {
	previousQuote, ok := f.previousQuotes[trade.Market]
	if !ok {
		f.previousQuotes[trade.Market] = trade.Price
		return true
	}

	diff := trade.Price.Sub(previousQuote).Abs()
	marketTreshold := previousQuote.Mul(f.threshold)
	if diff.GreaterThanOrEqual(marketTreshold) {
		f.previousQuotes[trade.Market] = trade.Price
		return true
	}

	loggerPriceDiffFilter.Infow("skipping trade", "trade", trade, "diff", diff, "threshold", f.threshold)
	return false
}
