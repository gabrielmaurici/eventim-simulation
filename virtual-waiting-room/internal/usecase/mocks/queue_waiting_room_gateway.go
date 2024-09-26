package mocks

import "github.com/stretchr/testify/mock"

type QueueWaitingRoomGatewayMock struct {
	mock.Mock
}

func (m *QueueWaitingRoomGatewayMock) Enqueue(token string) (position int, err error) {
	args := m.Called(token)
	return args.Int(0), args.Error(1)
}
