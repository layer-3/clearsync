package main

import (
	"context"
	"fmt"

	"github.com/layer-3/clearsync/pkg/userop"
)

func main() {
	client, err := userop.NewClient(config)
	if err != nil {
		panic(fmt.Errorf("failed to create userop client: %w", err))
	}

	op, err := client.NewUserOp(context.Background(), sender, receiver, token, amount)
	if err != nil {
		panic(fmt.Errorf("failed to build userop: %w", err))
	}

	callback := func() {}
	if err := client.SendUserOp(context.Background(), op, callback); err != nil {
		panic(fmt.Errorf("failed to send userop: %w", err))
	}
}
