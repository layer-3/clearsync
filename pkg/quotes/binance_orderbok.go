package quotes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

// gapThreshold is the maximum gap size (in update IDs) before we refetch the snapshot.
const gapThreshold = 25

// ErrInvalidSnapshot is returned if the snapshot from REST is empty.
var ErrInvalidSnapshot = errors.New("invalid snapshot received")

// BinanceDepthSnapshot is the REST snapshot.
type BinanceDepthSnapshot struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// BinanceDepthEvent is the WebSocket update event.
type BinanceDepthEvent struct {
	FirstUpdateID int64      `json:"U"`
	LastUpdateID  int64      `json:"u"`
	Bids          [][]string `json:"b"`
	Asks          [][]string `json:"a"`
}

// BinanceOrderBookLevel represents a single level in the order book.
type BinanceOrderBookLevel struct {
	Price  decimal.Decimal
	Amount decimal.Decimal
}

// BinanceOrderBookOutboxEvent is sent on each update.
type BinanceOrderBookOutboxEvent struct {
	Bids []BinanceOrderBookLevel
	Asks []BinanceOrderBookLevel
}

// BinanceOrderBook holds the current order book state.
type BinanceOrderBook struct {
	mu           sync.Mutex
	LastUpdateID int64
	// Both Bids and Asks are maintained as maps using a normalized price string as key.
	// The value is a normalized string representation of the quantity.
	Bids map[string]string
	Asks map[string]string

	Outbox    chan<- BinanceOrderBookOutboxEvent
	TopLevels int
}

// NewBinanceOrderBook creates a new order book. The Outbox channel receives the top-level updates.
// The provided context cancels the routine.
func NewBinanceOrderBook(ctx context.Context, market Market, topLevels int, outbox chan<- BinanceOrderBookOutboxEvent) (*BinanceOrderBook, error) {
	if topLevels <= 0 {
		return nil, errors.New("topLevels must be greater than 0")
	}
	if outbox == nil {
		return nil, errors.New("outbox channel is required")
	}

	ob := &BinanceOrderBook{
		Bids:      make(map[string]string),
		Asks:      make(map[string]string),
		Outbox:    outbox,
		TopLevels: topLevels,
	}

	// Run the order book update loop.
	go ob.run(ctx, market)
	return ob, nil
}

// run connects to the WebSocket feed and processes updates in a loop.
func (ob *BinanceOrderBook) run(ctx context.Context, market Market) {
	for {
		if ctx.Err() != nil {
			return
		}
		err := ob.connectWebSocket(ctx, market)
		if err != nil {
			fmt.Printf("WebSocket error: %v\n", err)
		}
		// Pause before attempting to reconnect.
		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			return
		}
	}
}

// connectWebSocket fetches a snapshot and then processes WebSocket updates.
func (ob *BinanceOrderBook) connectWebSocket(ctx context.Context, market Market) error {
	// Build the WebSocket URL.
	symbol := strings.ReplaceAll(strings.ToLower(market.String()), "/", "")
	wsURL := fmt.Sprintf("wss://stream.binance.com:9443/ws/%s@depth@100ms", symbol)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		return fmt.Errorf("failed to dial WebSocket: %w", err)
	}
	defer conn.Close()

	// Simple ping/pong to keep the connection alive.
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
				return
			}
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Fetch the initial snapshot.
	snapshot, err := ob.fetchSnapshot(ctx, market)
	if err != nil {
		return fmt.Errorf("failed to fetch snapshot: %w", err)
	}
	ob.applySnapshot(snapshot)
	fmt.Printf("Snapshot applied: LastUpdateID=%d\n", snapshot.LastUpdateID)
	nextUpdateID := snapshot.LastUpdateID + 1

	// Process incoming updates.
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}

		var event BinanceDepthEvent
		if err := json.Unmarshal(message, &event); err != nil {
			fmt.Printf("Unmarshal error: %v\n", err)
			continue
		}

		// Skip events that are entirely outdated.
		if event.LastUpdateID < nextUpdateID {
			continue
		}

		// According to Binance, we must have:
		//    event.FirstUpdateID <= nextUpdateID <= event.LastUpdateID
		if event.FirstUpdateID <= nextUpdateID && nextUpdateID <= event.LastUpdateID {
			ob.applyUpdate(event)
			nextUpdateID = event.LastUpdateID + 1
			ob.sendUpdate()
		} else {
			// A gap or an unexpected range – refetch the snapshot.
			fmt.Printf("Gap detected (expected %d, got [%d, %d]). Refetching snapshot...\n",
				nextUpdateID, event.FirstUpdateID, event.LastUpdateID)
			snapshot, err = ob.fetchSnapshot(ctx, market)
			if err != nil {
				return fmt.Errorf("failed to refetch snapshot: %w", err)
			}
			ob.applySnapshot(snapshot)
			nextUpdateID = snapshot.LastUpdateID + 1
			fmt.Printf("Snapshot reapplied: LastUpdateID=%d\n", snapshot.LastUpdateID)
		}
	}
}

// fetchSnapshot retrieves the current order book snapshot via REST.
func (ob *BinanceOrderBook) fetchSnapshot(ctx context.Context, market Market) (BinanceDepthSnapshot, error) {
	symbol := strings.ReplaceAll(strings.ToUpper(market.String()), "/", "")
	url := fmt.Sprintf("https://api.binance.com/api/v3/depth?symbol=%s&limit=1000", symbol)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return BinanceDepthSnapshot{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return BinanceDepthSnapshot{}, err
	}
	defer resp.Body.Close()

	var snapshot BinanceDepthSnapshot
	if err := json.NewDecoder(resp.Body).Decode(&snapshot); err != nil {
		return BinanceDepthSnapshot{}, err
	}
	if snapshot.LastUpdateID == 0 {
		return BinanceDepthSnapshot{}, ErrInvalidSnapshot
	}
	return snapshot, nil
}

// normalizePrice takes a price string, parses it to a decimal, and returns its normalized string.
func normalizePrice(priceStr string) (string, error) {
	d, err := decimal.NewFromString(priceStr)
	if err != nil {
		return "", err
	}
	return d.String(), nil
}

// normalizeQuantity takes a quantity string, parses it to a decimal, and returns its normalized string.
func normalizeQuantity(qtyStr string) (string, error) {
	d, err := decimal.NewFromString(qtyStr)
	if err != nil {
		return "", err
	}
	return d.String(), nil
}

// applySnapshot completely replaces the order book with the snapshot data.
// Prices and quantities are normalized so that keys match across snapshot and updates.
func (ob *BinanceOrderBook) applySnapshot(snapshot BinanceDepthSnapshot) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	ob.LastUpdateID = snapshot.LastUpdateID
	ob.Bids = make(map[string]string)
	ob.Asks = make(map[string]string)

	for _, bid := range snapshot.Bids {
		if len(bid) < 2 {
			continue
		}
		normPrice, err := normalizePrice(bid[0])
		if err != nil {
			continue
		}
		normQty, err := normalizeQuantity(bid[1])
		if err != nil {
			continue
		}
		// Only store if the quantity is nonzero.
		qtyDec, _ := decimal.NewFromString(normQty)
		if qtyDec.IsZero() {
			continue
		}
		ob.Bids[normPrice] = normQty
	}

	for _, ask := range snapshot.Asks {
		if len(ask) < 2 {
			continue
		}
		normPrice, err := normalizePrice(ask[0])
		if err != nil {
			continue
		}
		normQty, err := normalizeQuantity(ask[1])
		if err != nil {
			continue
		}
		qtyDec, _ := decimal.NewFromString(normQty)
		if qtyDec.IsZero() {
			continue
		}
		ob.Asks[normPrice] = normQty
	}
}

// applyUpdate applies an incremental update event to the order book.
// Prices and quantities are normalized so that update keys match those from the snapshot.
func (ob *BinanceOrderBook) applyUpdate(event BinanceDepthEvent) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	for _, bid := range event.Bids {
		if len(bid) < 2 {
			continue
		}
		normPrice, err := normalizePrice(bid[0])
		if err != nil {
			continue
		}
		normQty, err := normalizeQuantity(bid[1])
		if err != nil {
			continue
		}
		qtyDec, _ := decimal.NewFromString(normQty)
		if qtyDec.IsZero() {
			delete(ob.Bids, normPrice)
		} else {
			ob.Bids[normPrice] = normQty
		}
	}

	for _, ask := range event.Asks {
		if len(ask) < 2 {
			continue
		}
		normPrice, err := normalizePrice(ask[0])
		if err != nil {
			continue
		}
		normQty, err := normalizeQuantity(ask[1])
		if err != nil {
			continue
		}
		qtyDec, _ := decimal.NewFromString(normQty)
		if qtyDec.IsZero() {
			delete(ob.Asks, normPrice)
		} else {
			ob.Asks[normPrice] = normQty
		}
	}

	ob.LastUpdateID = event.LastUpdateID
}

// sendUpdate builds the top-level order book levels and pushes an update on Outbox.
func (ob *BinanceOrderBook) sendUpdate() {
	bids := ob.getSortedLevels(ob.Bids, true)
	asks := ob.getSortedLevels(ob.Asks, false)
	update := BinanceOrderBookOutboxEvent{
		Bids: bids,
		Asks: asks,
	}
	// Nonblocking send.
	select {
	case ob.Outbox <- update:
	default:
	}
}

// getSortedLevels returns the top levels from the given side (bids or asks).
// Bids are sorted descending, asks ascending.
func (ob *BinanceOrderBook) getSortedLevels(side map[string]string, isBid bool) []BinanceOrderBookLevel {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	var levels []BinanceOrderBookLevel
	for priceStr, qtyStr := range side {
		price, err := decimal.NewFromString(priceStr)
		if err != nil {
			continue
		}
		qty, err := decimal.NewFromString(qtyStr)
		if err != nil {
			continue
		}
		levels = append(levels, BinanceOrderBookLevel{
			Price:  price,
			Amount: qty,
		})
	}

	sort.Slice(levels, func(i, j int) bool {
		if isBid {
			return levels[i].Price.GreaterThan(levels[j].Price)
		}
		return levels[i].Price.LessThan(levels[j].Price)
	})

	if len(levels) > ob.TopLevels {
		levels = levels[:ob.TopLevels]
	}
	return levels
}
