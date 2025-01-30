package filter

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"time"

	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes/common"
)

var (
	loggerMADFilter = log.Logger("mad_filter")
	emptyTime       = time.Time{}

	consistencyConstant = decimal.RequireFromString("1.4826") // as for normal distribution
)

type MADConfig struct {
	Factor     int64         `yaml:"factor" env:"FACTOR" env-default:"4"`
	TimeWindow time.Duration `yaml:"time_window" env:"TIME_WINDOW" env-default:"30m"`
}

type madFilter struct {
	factor  decimal.Decimal
	window  time.Duration
	trades  map[common.Market][]madTrade
	history common.HistoricalDataDriver
}

func newMADFilter(config MADConfig, history common.HistoricalDataDriver) (madFilter, error) {
	if config.Factor <= 0 {
		return madFilter{}, fmt.Errorf("factor must be greater than 0")
	}
	if config.TimeWindow <= 0 {
		return madFilter{}, fmt.Errorf("time window must be greater than 0")
	}
	if history == nil {
		return madFilter{}, fmt.Errorf("historical data driver is required")
	}

	return madFilter{
		factor:  decimal.NewFromInt(config.Factor),
		window:  config.TimeWindow,
		trades:  make(map[common.Market][]madTrade),
		history: history,
	}, nil
}

func (f madFilter) Allow(trade common.TradeEvent) bool {
	now := time.Now()
	if trade.CreatedAt == emptyTime {
		trade.CreatedAt = now
	}

	if time.Since(trade.CreatedAt) > f.window {
		return false
	}

	trades := f.filterTrades(trade.Market)
	prices := make([]decimal.Decimal, len(trades))
	for i, t := range trades {
		prices[i] = t.price
	}

	// If no historical trades were fetched,
	// there's no way to know if the incoming trade is outlier or not,
	// that's why the trade is allowed.
	market := trade.Market
	cachedTrade := madTrade{
		price:     trade.Price,
		createdAt: trade.CreatedAt,
	}
	if len(prices) == 0 {
		loggerMADFilter.Infow("trades cache is empty, storing incoming price", "market", market)
		f.trades[market] = append(f.trades[market], cachedTrade)
		return true
	}

	median := calculateMedian(sortableDecimal{prices})
	mad := calculateMAD(prices)
	mmad := calculateMMAD(mad)
	lowerBound, upperBound := calculateOutlierBounds(median, mmad, f.factor)

	if trade.Price.LessThan(lowerBound) || upperBound.LessThan(trade.Price) {
		loggerMADFilter.Infow("outlier detected",
			"lower_bound", lowerBound,
			"upper_bound", upperBound,
			"trade", trade)
		return false
	}

	// Market is guaranteed to be in the map
	// as it was inserted in `filterTrades` method.
	f.trades[market] = append(f.trades[market], cachedTrade)

	// The trade may be within the time window
	// but not the last trade in the list by timestamp.
	// In this (pretty rare) case, we need to sort the trades.
	sortMADtrades(f.trades[market])

	return true
}

func (f madFilter) filterTrades(market common.Market) []madTrade {
	// +1 to account for trade that will be appended if the trade passes filtering.
	// This is done to optimize memory usage by avoiding frequent allocations.
	filteredCapacity := len(f.trades[market]) + 1
	filteredTrades := make([]madTrade, 0, filteredCapacity)

	if _, ok := f.trades[market]; !ok {
		const limit = 10000 // TODO: this is arbitrary value, make it configurable
		data, err := f.history.HistoricalData(context.TODO(), market, f.window, limit)
		if err != nil {
			loggerMADFilter.Errorw("failed to get historical data", "market", market, "err", err)
		}

		filteredTrades = make([]madTrade, len(data))
		for i, t := range data {
			filteredTrades[i] = madTrade{
				price:     t.Price,
				createdAt: t.CreatedAt,
			}
		}
	}

	trades := f.trades[market]

	// Filter trades to include only those within the time window
	windowStart := time.Now().Add(f.window)
	for _, t := range trades {
		if windowStart.After(t.createdAt) {
			filteredTrades = append(filteredTrades, t)
		}
	}

	// Sorting is required to accurately represent the market trend.
	sortMADtrades(filteredTrades)

	f.trades[market] = filteredTrades // store only the filtered trades
	return filteredTrades
}

type madTrade struct {
	price     decimal.Decimal
	createdAt time.Time
}

func sortMADtrades(trades []madTrade) {
	slices.SortFunc(trades, func(a, b madTrade) int {
		if a.createdAt.Before(b.createdAt) {
			return -1
		}
		if a.createdAt.After(b.createdAt) {
			return 1
		}
		return 0
	})
}

type sortableDecimal struct {
	slice []decimal.Decimal
}

func (sortable sortableDecimal) Len() int {
	return len(sortable.slice)
}

func (sortable sortableDecimal) Less(i, j int) bool {
	return sortable.slice[i].LessThan(sortable.slice[j])
}

func (sortable sortableDecimal) Swap(i, j int) {
	sortable.slice[i], sortable.slice[j] = sortable.slice[j], sortable.slice[i]
}

// calculateMedian calculates the median of a slice of float64.
func calculateMedian(data sortableDecimal) decimal.Decimal {
	if data.Len() == 0 {
		return decimal.Zero
	}

	sort.Sort(data)
	n := data.Len()
	if n%2 == 0 {
		return data.slice[n/2-1].Add(data.slice[n/2]).Div(decimal.RequireFromString("2"))
	}
	return data.slice[n/2]
}

// calculateAbsoluteDeviations calculates the absolute deviations from the median for a slice of float64.
func calculateAbsoluteDeviations(data []decimal.Decimal, median decimal.Decimal) []decimal.Decimal {
	deviations := make([]decimal.Decimal, len(data))
	for i, v := range data {
		deviations[i] = v.Sub(median).Abs()
	}
	return deviations
}

// calculateMAD calculates the median absolute deviation of a slice of float64.
func calculateMAD(data []decimal.Decimal) decimal.Decimal {
	median := calculateMedian(sortableDecimal{data})
	deviations := calculateAbsoluteDeviations(data, median)
	return calculateMedian(sortableDecimal{deviations})
}

// calculateMMAD calculates the modified median absolute deviation.
func calculateMMAD(mad decimal.Decimal) decimal.Decimal {
	return mad.Mul(consistencyConstant)
}

// calculateOutlierBounds calculates lower and upper bounds for outlier detection.
func calculateOutlierBounds(median, mmad, factor decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	delta := factor.Mul(mmad)
	lowerBound := median.Sub(delta)
	upperBound := median.Add(delta)
	return lowerBound, upperBound
}
