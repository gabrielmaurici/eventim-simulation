package mocks

import (
	entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"
	"github.com/stretchr/testify/mock"
)

type TicketGatewayMock struct {
	mock.Mock
}

func (m *TicketGatewayMock) GetAvailableTickets(quantity int8) ([]entity.Ticket, error) {
	args := m.Called(quantity)
	return args.Get(0).([]entity.Ticket), args.Error(1)
}

func (m *TicketGatewayMock) Update(ticket entity.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}
