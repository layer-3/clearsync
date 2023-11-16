package quotes

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WSTransport interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

type WSDialer interface {
	Dial(url string, header http.Header) (WSTransport, *http.Response, error)
}

type WSDialWrapper struct{}

func (WSDialWrapper) Dial(url string, header http.Header) (WSTransport, *http.Response, error) {
	return websocket.DefaultDialer.Dial(url, header)
}
