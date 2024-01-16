package common

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
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

type WsDialWrapper struct{}

func (WsDialWrapper) Dial(url string, header http.Header) (WsTransport, *http.Response, error) {
	conn, resp, err := websocket.DefaultDialer.Dial(url, header)
	return &WsConn{conn: conn}, resp, err
}

type WsConn struct {
	conn *websocket.Conn
	mu   *sync.Mutex
}

func (w *WsConn) IsConnected() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn != nil
}

func (w *WsConn) ReadMessage() (messageType int, p []byte, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.ReadMessage()
}

func (w *WsConn) WriteMessage(messageType int, data []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.WriteMessage(messageType, data)
}

func (w *WsConn) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.conn.Close()
}
