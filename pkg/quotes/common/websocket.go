package common

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
)

type WsTransport interface {
	IsConnected() bool
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

type WsDialer interface {
	Dial(url string, header http.Header) (WsTransport, *http.Response, error)
}

type WsDialWrapper struct {
	rps int
}

func (w *WsDialWrapper) WithRateLimit(rps int) {
	w.rps = rps
}

func (w *WsDialWrapper) Dial(url string, header http.Header) (WsTransport, *http.Response, error) {
	conn, resp, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, resp, err
	}

	connWrapper := &WsConn{
		conn:    conn,
		limiter: rate.NewLimiter(rate.Limit(w.rps), w.rps),
	}
	return connWrapper, resp, err
}

type WsConn struct {
	conn    *websocket.Conn
	mu      sync.RWMutex
	limiter *rate.Limiter
}

func (w *WsConn) IsConnected() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.conn != nil
}

func (w *WsConn) ReadMessage() (messageType int, p []byte, err error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.conn.ReadMessage()
}

func (w *WsConn) WriteMessage(messageType int, data []byte) error {
	if w.limiter != nil {
		// Wait until the limiter allows sending a message
		ctx := context.Background()
		if err := w.limiter.Wait(ctx); err != nil {
			return err
		}
	}

	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.conn.WriteMessage(messageType, data)
}

func (w *WsConn) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.Close()
}
