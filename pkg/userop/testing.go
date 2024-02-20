package userop

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

type MockUserOperationClient struct {
	mock.Mock
}

func (c *MockUserOperationClient) NewUserOp(ctx context.Context, sender common.Address, calls []Call) (UserOperation, error) {
	args := c.Called(ctx, sender, calls)
	return args.Get(0).(UserOperation), args.Error(1)
}

func (c *MockUserOperationClient) SendUserOp(ctx context.Context, op UserOperation) (<-chan struct{}, error) {
	args := c.Called(ctx, op)
	return args.Get(0).(<-chan struct{}), args.Error(1)
}
