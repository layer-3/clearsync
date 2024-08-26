package quest

import "context"

type Handler interface {
	Handle(ctx context.Context, userAddress string) (Result, error)
}

type Result struct {
	Valid bool                   `json:"valid"` // Verification result
	Data  map[string]interface{} `json:"data"`  // Optional additional data
}
