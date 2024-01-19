package opendax_protocol

import (
	"encoding/json"
	"errors"
	"fmt"
)

const ParseError = "unexpected message"

func ParseRaw(msg []byte) (*Msg, error) {
	req := Msg{}
	var v []interface{}
	if err := json.Unmarshal(msg, &v); err != nil {
		return nil, fmt.Errorf("could not parse message: %w", err)
	}

	if len(v) < 3 {
		return nil, errors.New("message is too small")
	}

	it := NewArgIterator(v)
	tInt := it.NextUint64()
	if it.Err() != nil {
		return nil, fmt.Errorf("failed to parse type: %w", it.Err())
	}

	t := uint8(tInt)

	var reqID uint64
	var method string
	var args []interface{}

	switch t {
	case Request, Response:
		if len(v) != 4 {
			return nil, errors.New("message must contain 4 elements")
		}

		reqID = it.NextUint64()
		if it.Err() != nil {
			return nil, fmt.Errorf("failed to parse request ID: %w", it.Err())
		}

		method = it.NextString()
		if it.Err() != nil {
			return nil, fmt.Errorf("failed to parse method: %w", it.Err())
		}

		args = it.NextSlice()
		if it.Err() != nil {
			return nil, fmt.Errorf("failed to parse arguments: %w", it.Err())
		}

	case EventPrivate, EventPublic:
		if len(v) != 3 {
			return nil, errors.New("message must contain 3 elements")
		}

		method = it.NextString()
		if it.Err() != nil {
			return nil, fmt.Errorf("failed to parse method: %w", it.Err())
		}

		args = it.NextSlice()
		if it.Err() != nil {
			return nil, fmt.Errorf("failed to parse arguments: %w", it.Err())
		}
	default:
		return nil, errors.New("message type must be 1, 2, 3 or 4")
	}

	req.Type = t
	req.ReqID = reqID
	req.Method = method
	req.Args = args

	return &req, it.Err()
}
