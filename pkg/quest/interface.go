package quest

import "context"

type Handler interface {
	Handle(ctx context.Context, userAddress string) (bool, error)
}
