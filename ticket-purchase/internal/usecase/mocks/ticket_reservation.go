package mocks

import "github.com/stretchr/testify/mock"

type TicketReservationGatewayMock struct {
	mock.Mock
}

func (m *TicketReservationGatewayMock) Reserve(userToken string, ticketId string) error {
	args := m.Called(userToken, ticketId)
	return args.Error(0)
}
