package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type TicketReservationGatewayMock struct {
	mock.Mock
}

func (m *TicketReservationGatewayMock) HasReservation(userToken string, ctx context.Context) (Has bool, err error) {
	args := m.Called(userToken, ctx)
	return args.Bool(0), args.Error(1)
}

func (m *TicketReservationGatewayMock) CreateTicketReservation(userToken string, ctx context.Context) error {
	args := m.Called(userToken, ctx)
	return args.Error(0)
}

func (m *TicketReservationGatewayMock) RegisterTickets(userToken string, ticketsId []string, ctx context.Context) error {
	args := m.Called(userToken, ticketsId, ctx)
	return args.Error(0)
}
