package userop

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserOperationClient struct {
	mock.Mock
}

func (c *MockUserOperationClient) NewUserOp(ctx context.Context, calls []Call) (UserOperation, error) {
	args := c.Called(ctx, calls)
	return args.Get(0).(UserOperation), args.Error(1)
}

func (c *MockUserOperationClient) SendUserOp(ctx context.Context, op UserOperation, callback func()) error {
	args := c.Called(ctx, op, callback)
	return args.Error(0)
}
