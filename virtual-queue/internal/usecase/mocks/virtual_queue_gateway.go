package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type VirtualQueueGatewayMock struct {
	mock.Mock
}

func (m *VirtualQueueGatewayMock) Enqueue(token string, ctx context.Context) (position int64, err error) {
	args := m.Called(token, ctx)
	return args.Get(0).(int64), args.Error(1)
}

func (m *VirtualQueueGatewayMock) Dequeue(ctx context.Context) (token string, err error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *VirtualQueueGatewayMock) GetAll(ctx context.Context) (tokens []string, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}
