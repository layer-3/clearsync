package quotes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type PriceCache interface {
	GetEMA(market string) (decimal.Decimal, int64)
	UpdateEMA(market string, newValue decimal.Decimal)
}

type IndexAggregator struct {
	weightsURL string
	weightsMap map[DriverType]decimal.Decimal
	prices     PriceCache
	drivers    []Driver
	aggregated chan TradeEvent
	outbox     chan<- TradeEvent
}

func NewIndex(driverConfigs []Config, remoteWeights string, outbox chan<- TradeEvent) *IndexAggregator {
	aggregated := make(chan TradeEvent, 128)
	drivers := []Driver{}
	for _, d := range driverConfigs {
		driver, err := NewDriver(d, aggregated)
		if err != nil {
			logger.Info("cannot start a driver: ", err)
			return nil
		}

		drivers = append(drivers, driver)
	}

	return &IndexAggregator{
		prices:     NewPricesCache(),
		drivers:    drivers,
		outbox:     outbox,
		aggregated: aggregated,
	}
}

func (a *IndexAggregator) Start(markets []Market) error {
	go func() {
		for {
			err := a.fetchWeights(context.Background())
			if err != nil {
				logger.Error(err)
			}
			<-time.After(5 * time.Minute)
		}
	}()

	for _, d := range a.drivers {
		go func(d Driver) {
			logger.Info("starting index quotes service")

			markets := make([]Market, 0, len(markets))
			for _, m := range markets {
				market := Market{BaseUnit: m.BaseUnit, QuoteUnit: m.QuoteUnit}
				markets = append(markets, market)
			}

			if err := d.Start(markets); err != nil {
				logger.Warn(err.Error())
			}
		}(d)
	}
	go func() {
		for event := range a.aggregated {
			indexPrice := a.indexPrice(event)
			a.outbox <- indexPrice
		}
	}()
	return nil
}

func (a *IndexAggregator) Subscribe(m Market) error {
	for _, d := range a.drivers {
		err := d.Subscribe(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *IndexAggregator) Stop() error {
	for _, d := range a.drivers {
		err := d.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *IndexAggregator) indexPrice(event TradeEvent) TradeEvent {
	lastEMA, timestamp := a.prices.GetEMA(event.Market)
	weight := a.weightsMap[event.Source]

	newEMA := EMA20(lastEMA, event.Price.Mul(event.Amount.Mul(weight).Div(a.activeWeights())))
	if time.Now().Unix() >= timestamp+60 {
		a.prices.UpdateEMA(event.Market, newEMA)
	}

	event.Price = newEMA
	event.Source = DriverIndex

	return event
}

func (a *IndexAggregator) activeWeights() decimal.Decimal {
	total := decimal.Zero
	for _, w := range a.weightsMap {
		total = total.Add(w)
	}
	return total
}

func (a *IndexAggregator) fetchWeights(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, a.weightsURL, nil)
	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("Error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Cannot fetch markets. HTTP request failed with status code: %d", resp.StatusCode)
	}

	var weights []string
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&weights); err != nil {
		return fmt.Errorf("Error decoding JSON: %w", err)
	}

	weightsMap := make(map[DriverType]decimal.Decimal, 0)
	for _, w := range weights {
		driver, weight, err := parseWeight(w)
		if err != nil {
			return err
		}
		weightsMap[driver] = weight
	}

	a.weightsMap = weightsMap
	return nil
}

func parseWeight(raw string) (DriverType, decimal.Decimal, error) {
	// TODO: define remote config structure and add parsing based on that
	return DriverIndex, decimal.Zero, nil
}

func EMA20(lastEMA, newPrice decimal.Decimal) decimal.Decimal {
	return EMA(lastEMA, newPrice, 20)
}

func EMA(lastEMA, newPrice decimal.Decimal, intervals int) decimal.Decimal {
	if intervals <= 0 {
		return decimal.Zero
	}
	smoothingFactor := decimal.NewFromInt(2)
	alpha := smoothingFactor.Div(decimal.NewFromInt32(int32(intervals)).Add(decimal.NewFromInt(1)))
	return alpha.Mul(newPrice.Sub(lastEMA)).Add(lastEMA)
}
