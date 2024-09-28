package mocks

import "github.com/stretchr/testify/mock"

type VirtualQueueGatewayMock struct {
	mock.Mock
}

func (m *VirtualQueueGatewayMock) Enqueue(token string) (position int64, err error) {
	args := m.Called(token)
	return int64(args.Int(0)), args.Error(1)
}
