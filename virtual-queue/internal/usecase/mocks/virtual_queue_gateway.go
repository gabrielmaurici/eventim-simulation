package mocks

import "github.com/stretchr/testify/mock"

type VirtualQueueGatewayMock struct {
	mock.Mock
}

func (m *VirtualQueueGatewayMock) Enqueue(token string) (position int64, err error) {
	args := m.Called(token)
	return int64(args.Int(0)), args.Error(1)
}

func (m *VirtualQueueGatewayMock) Dequeue() (token string, err error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *VirtualQueueGatewayMock) GetAll() (tokens []string, err error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}
