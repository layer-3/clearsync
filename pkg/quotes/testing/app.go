package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gorilla/mux"
	"github.com/ipfs/go-log/v2"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/quotes"
	"github.com/layer-3/clearsync/pkg/safe"
)

var (
	blocks    = flag.Uint64("blocks", 100, "number of blocks to fetch") // TODO: handle this flag
	allTrades = safe.NewMap[string, []quotes.TradeEvent]()
)

// Usage example: `go run . binance syncswap -n 100`
func main() {
	go func() {
		// Start server for klines and pprof
		r := mux.NewRouter()
		r.HandleFunc("/kline/{base}/{quote}", HandleKline)
		http.Handle("/", r)

		const url = "localhost:8080"
		slog.Info("listening at " + url)
		if err := http.ListenAndServe(url, nil); err != nil {
			panic(err)
		}
	}()

	if err := log.SetLogLevel("*", "info"); err != nil {
		panic(err)
	}

	flag.Parse()

	var drivers []quotes.DriverType
	if len(os.Args) >= 2 {
		drivers = make([]quotes.DriverType, 0, len(os.Args[1:]))
		for _, arg := range FilterPositionalArgs(os.Args[1:]) {
			parsedDriver, err := quotes.ToDriverType(arg)
			if err != nil {
				panic(err)
			}
			drivers = append(drivers, parsedDriver)
		}
	}

	config, err := quotes.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}
	if len(drivers) > 0 {
		// Override default values only if drivers are provided
		config.Drivers = drivers
	}

	outbox := make(chan quotes.TradeEvent, 128)
	outboxStop := make(chan struct{}, 1)
	go func() {
		// Process trades
		for trade := range outbox {
			slog.Info("new trade",
				"source", trade.Source,
				"market", trade.Market,
				"side", trade.TakerType.String(),
				"price", trade.Price.String(),
				"amount", trade.Amount.String())
			market := strings.ToLower(trade.Market.String())
			allTrades.UpdateInTx(func(m map[string][]quotes.TradeEvent) {
				m[market] = append(m[market], trade)
			})
		}
		outboxStop <- struct{}{}
		close(outboxStop)
	}()

	driver, err := quotes.NewDriver(config, outbox, nil, nil)
	if err != nil {
		panic(err)
	}

	slog.Info("starting", "config", config)

	if err := driver.Start(); err != nil {
		panic(err)
	}

	markets := []quotes.Market{
		quotes.NewMarket("eth", "usd"),
		quotes.NewMarket("btc", "usd"),
		quotes.NewMarket("lube", "usdc"),
		quotes.NewMarket("linda", "usdc"),
	}

	var atLeastOne atomic.Bool
	for _, market := range markets {
		market := market
		go func() {
			if err = driver.Subscribe(market); err != nil {
				slog.Warn("failed to subscribe", "market", market, "err", err)
				return
			}
			atLeastOne.CompareAndSwap(false, true)
			slog.Info("subscribed", "market", market.String())
		}()
	}
	if !atLeastOne.Load() {
		panic("failed to subscribe to at least one market")
	}

	slog.Info("waiting for trades")
	<-outboxStop
}

// FilterPositionalArgs skips flags and their values from the list of arguments
// and returns only positional arguments.
func FilterPositionalArgs(args []string) (positionalArgs []string) {
	skipNext := false
	for _, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}
		if arg[0] == '-' {
			// Check if the next argument is the value for this flag
			skipNext = len(arg) == 2 || (len(arg) > 2 && arg[1] != '-')
		} else {
			positionalArgs = append(positionalArgs, arg)
		}
	}
	return positionalArgs
}

func HandleKline(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	base, quote := vars["base"], vars["quote"]
	market := strings.ToLower(base + "/" + quote)

	// Check if the market exists
	trades, ok := allTrades.Load(market)
	if !ok {
		http.Error(w, "Market not found", http.StatusNotFound)
		return
	}

	ohlc := BuildOHLC(trades, 5*time.Minute)
	chart := BuildChart(ohlc)

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta http-equiv="refresh" content="1">
		<title>Live Chart</title>
	</head>
	<body>
		<div id="chart-container">` + RenderToString(chart) + `</div>
	</body>
	</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

// RenderToString renders the chart to an HTML string.
func RenderToString(chart *charts.Kline) string {
	writer := &strings.Builder{}
	chart.Render(writer)
	return writer.String()
}

// OHLC represents Open, High, Low, Close prices along with volume and timestamp.
// It is used to represent candlestick data.
type OHLC struct {
	Open      decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Close     decimal.Decimal
	Volume    decimal.Decimal
	Timestamp time.Time
}

// BuildOHLC builds a slice of OHLC from a slice of TradeEvent
func BuildOHLC(trades []quotes.TradeEvent, interval time.Duration) []OHLC {
	if len(trades) == 0 {
		return nil
	}

	var ohlcData []OHLC
	var currentOHLC *OHLC

	for _, trade := range trades {
		tradeTime := trade.CreatedAt.Truncate(interval)
		if currentOHLC == nil || tradeTime.After(currentOHLC.Timestamp) {
			if currentOHLC != nil {
				ohlcData = append(ohlcData, *currentOHLC)
			}
			currentOHLC = &OHLC{
				Open:      trade.Price,
				High:      trade.Price,
				Low:       trade.Price,
				Close:     trade.Price,
				Volume:    trade.Amount,
				Timestamp: tradeTime,
			}
		} else {
			if trade.Price.GreaterThan(currentOHLC.High) {
				currentOHLC.High = trade.Price
			}
			if trade.Price.LessThan(currentOHLC.Low) {
				currentOHLC.Low = trade.Price
			}
			currentOHLC.Close = trade.Price
			currentOHLC.Volume = currentOHLC.Volume.Add(trade.Amount)
		}
	}
	if currentOHLC != nil {
		ohlcData = append(ohlcData, *currentOHLC)
	}

	return ohlcData
}

func BuildChart(ohlc []OHLC) *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for _, o := range ohlc {
		x = append(x, o.Timestamp.Format(time.RFC3339))

		value := []any{
			o.Open.InexactFloat64(),
			o.Close.InexactFloat64(),
			o.High.InexactFloat64(),
			o.Low.InexactFloat64(),
		}
		y = append(y, opts.KlineData{Value: value})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Kline-example",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	kline.SetXAxis(x).AddSeries("kline", y)
	return kline
}
