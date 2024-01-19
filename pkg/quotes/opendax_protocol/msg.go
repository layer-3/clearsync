// Types and helpers for easy formatting and parsing open-finance protocol messages.
package opendax_protocol

import (
	"encoding/json"
	"errors"
)

const (
	// Request type code
	Request = 1

	// Response type code
	Response = 2

	// EventPublic is public event type code
	EventPublic = 3

	// EventPrivate is private event type code
	EventPrivate = 4

	// EventAdmin is admin event type code
	EventAdmin = 5
)

// Msg represent websocket messages, it could be either a request, a response or an event
type Msg struct {
	Type   uint8
	ReqID  uint64
	Method string
	Args   []any
}

func NewSubscribeMessage(reqID uint64, topics ...any) *Msg {
	return &Msg{
		Type:   Request,
		ReqID:  reqID,
		Method: MethodSubscribe,
		Args: []any{
			"public",
			topics,
		},
	}
}

func NewUnsubscribeMessage(reqID uint64, topics ...any) *Msg {
	return &Msg{
		Type:   Request,
		ReqID:  reqID,
		Method: MethodUnsubscribe,
		Args: []any{
			"public",
			topics,
		},
	}
}

// Encode msg into json
func (m *Msg) Encode() ([]byte, error) {
	switch m.Type {
	case Response, Request:
		return json.Marshal([]any{
			m.Type,
			m.ReqID,
			m.Method,
			m.Args,
		})

	case EventPrivate, EventPublic, EventAdmin:
		return json.Marshal([]any{
			m.Type,
			m.Method,
			m.Args,
		})

	default:
		return nil, errors.New("invalid type")
	}
}
