package quotes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"

	"github.com/layer-3/neodax/finex/models/trade"
	"github.com/layer-3/neodax/finex/pkg/cache"
	"github.com/layer-3/neodax/finex/pkg/config"
	"github.com/layer-3/neodax/finex/pkg/event"
	"github.com/layer-3/neodax/finex/pkg/websocket/client"
)

type Kraken struct {
	conn        client.WSTransport
	dialer      client.WSDialer
	retryPeriod time.Duration
	isConnected bool

	marketCache  cache.Market
	tradeSampler *TradeSampler
	outbox       chan trade.Event
	output       chan<- event.Event
	mu           sync.RWMutex
}

func (k *Kraken) Init(markets cache.Market, outbox chan trade.Event, output chan<- event.Event, config config.QuoteFeed, dialer client.WSDialer) error {
	k.dialer = dialer
	k.retryPeriod = time.Duration(config.Period) * time.Second
	k.marketCache = markets
	k.tradeSampler = NewTradeSampler(config.TradeSampler)
	k.outbox = outbox
	k.output = output

	return nil
}

func (k *Kraken) Start() error {
	if err := k.connect(); err != nil {
		return err
	}

	if err := k.subscribe(); err != nil {
		return err
	}

	k.listen()
	return nil
}

type subscribeMessage struct {
	Method string             `json:"method"`
	Params subscriptionParams `json:"params"`
}

type subscriptionParams struct {
	Channel  string   `json:"channel"`
	Snapshot bool     `json:"snapshot"`
	Symbol   []string `json:"symbol"`
}

func (k *Kraken) Subscribe(base, quote string) error {
	k.mu.RLock()
	defer k.mu.RUnlock()

	pair := fmt.Sprintf("%s/%s", strings.ToUpper(base), strings.ToUpper(quote))
	subMsg := subscribeMessage{
		Method: "subscribe",
		Params: subscriptionParams{
			Channel:  "trade",
			Snapshot: true,
			Symbol:   []string{pair},
		},
	}

	payload, err := json.Marshal(subMsg)
	if err != nil {
		return fmt.Errorf("error marshalling subscription message: %v", err)
	}

	for !k.isConnected {
	}

	if err := k.conn.WriteMessage(websocket.TextMessage, payload); err != nil {
		return fmt.Errorf("error writing subscription message: %v", err)
	}
	return nil
}

func (k *Kraken) Close() error {
	k.mu.Lock()
	defer k.mu.Unlock()

	conn := k.conn
	k.conn = nil

	if conn == nil {
		return nil
	}
	return conn.Close()
}

func (k *Kraken) connect() error {
	k.mu.Lock()
	defer k.mu.Unlock()

	var err error
	for {
		k.conn, _, err = k.dialer.Dial("wss://ws.kraken.com/v2", nil)
		if err != nil {
			logger.Error(err)
			time.Sleep(k.retryPeriod)
			continue
		}

		k.isConnected = true
		return nil
	}
}

func (k *Kraken) subscribe() error {
	availablePairs, err := getKrakenPairs()
	if err != nil {
		return err
	}

	marketList, err := k.marketCache.GetActive()
	if err != nil {
		return err
	}

	for _, m := range marketList {
		pair := fmt.Sprintf("%s%s", strings.ToUpper(m.BaseUnit), strings.ToUpper(m.QuoteUnit))
		if pair, ok := availablePairs[pair]; !ok || pair.Status != "online" {
			logger.Warnf("market %s doesn't exist in Kraken", m.Symbol)
			continue
		}

		if err := k.Subscribe(m.BaseUnit, m.QuoteUnit); err != nil {
			logger.Warnf("failed to subscribe to Kraken market %s: %v", m.Symbol, err)
			continue
		}

		logger.Infof("quotes service connected to Kraken %s market", m.Symbol)
		<-time.After(25 * time.Millisecond) // to cope with rate limits
	}

	return nil
}

type KrakenEvent[T KrakenStatus | KrakenTrade] struct {
	Channel string `json:"channel"`
	Type    string `json:"type"`
	Data    []T    `json:"data"`
}

type KrakenStatus struct {
	ApiVersion   string `json:"api_version"`
	ConnectionId uint64 `json:"connection_id"`
	System       string `json:"system"`
	Version      string `json:"version"`
}

type KrakenTrade struct {
	OrdType   string    `json:"ord_type"`
	Price     float64   `json:"price"`
	Qty       float64   `json:"qty"`
	Side      string    `json:"side"` // "buy" | "sell"
	Symbol    string    `json:"symbol"`
	Timestamp time.Time `json:"timestamp"`
	TradeId   int       `json:"trade_id"`
}

type KrakenResult struct {
	Method string `json:"method"`
	Result struct {
		Channel  string `json:"channel"`
		Snapshot bool   `json:"snapshot"`
		Symbol   string `json:"symbol"`
	} `json:"result"`
	Success bool      `json:"success"`
	TimeIn  time.Time `json:"time_in"`
	TimeOut time.Time `json:"time_out"`
}

func (k *Kraken) listen() {
	for {
		if !k.isConnected {
			<-time.After(k.retryPeriod)
			continue
		}

		_, rawMsg, err := k.conn.ReadMessage()
		if err != nil {
			logger.Errorf("error reading Kraken message: %v", err)
		}

		tradeEvents, err := k.parseMessage(rawMsg)
		if err != nil {
			logger.Errorf("error parsing Kraken message: %v", err)
			continue
		}
		if tradeEvents == nil {
			continue
		}

		for _, tr := range tradeEvents {
			if !k.tradeSampler.Allow(tr) {
				continue
			}

			k.outbox <- tr
			routingEvent, err := GetRoutingEvent(tr)
			if err != nil {
				logger.Warn(err)
				continue
			}
			k.output <- *routingEvent
		}
	}
}

func (k *Kraken) parseMessage(rawMsg []byte) ([]trade.Event, error) {
	var ticker KrakenEvent[KrakenTrade]
	if err := json.Unmarshal(rawMsg, &ticker); err == nil && ticker.Channel != "heartbeat" {
		return buildKrakenEvents(ticker.Data), nil
	}

	var status KrakenEvent[KrakenStatus]
	if err := json.Unmarshal(rawMsg, &status); err == nil && ticker.Channel != "heartbeat" {
		// TODO: Handle KrakenEvent[KrakenStatus]
		return nil, nil
	}

	var result KrakenResult
	if err := json.Unmarshal(rawMsg, &result); err != nil {
		return nil, err
	}
	if ticker.Channel == "heartbeat" {
		return nil, nil
	}

	if !result.Success {
		return nil, fmt.Errorf("failed to subscribe to Kraken market %s: ", result.Result.Symbol)
	}
	return nil, fmt.Errorf("unknown Kraken message: %s", string(rawMsg))
}

type krakenAssetPairs struct {
	Error  []interface{}         `json:"error"`
	Result map[string]krakenPair `json:"result"`
}

type krakenPair struct {
	Altname           string        `json:"altname"`
	Wsname            string        `json:"wsname"`
	AclassBase        string        `json:"aclass_base"`
	Base              string        `json:"base"`
	AclassQuote       string        `json:"aclass_quote"`
	Quote             string        `json:"quote"`
	Lot               string        `json:"lot"`
	CostDecimals      int           `json:"cost_decimals"`
	PairDecimals      int           `json:"pair_decimals"`
	LotDecimals       int           `json:"lot_decimals"`
	LotMultiplier     int           `json:"lot_multiplier"`
	LeverageBuy       []interface{} `json:"leverage_buy"`
	LeverageSell      []interface{} `json:"leverage_sell"`
	Fees              [][]float64   `json:"fees"`
	FeesMaker         [][]float64   `json:"fees_maker"`
	FeeVolumeCurrency string        `json:"fee_volume_currency"`
	MarginCall        int           `json:"margin_call"`
	MarginStop        int           `json:"margin_stop"`
	Ordermin          string        `json:"ordermin"`
	Costmin           string        `json:"costmin"`
	TickSize          string        `json:"tick_size"`
	Status            string        `json:"status"`
}

func getKrakenPairs() (map[string]krakenPair, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.kraken.com/0/public/AssetPairs", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pairs krakenAssetPairs
	if err := json.Unmarshal(body, &pairs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Kraken pairs response: %v", err)
	}

	return pairs.Result, nil
}

func buildKrakenEvents(trades []KrakenTrade) []trade.Event {
	var events []trade.Event
	for _, tr := range trades {
		price := decimal.NewFromFloat(tr.Price)
		amount := decimal.NewFromFloat(tr.Qty)

		takerType := trade.Buy
		if tr.Side == "sell" {
			takerType = trade.Sell
		}

		events = append(events, trade.Event{
			Market:    strings.ToLower(tr.Symbol),
			Price:     price,
			Amount:    amount,
			Total:     price.Mul(amount),
			TakerType: takerType,
			CreatedAt: tr.Timestamp.Unix(),
			Source:    "Kraken",
		})
	}

	return events
}
