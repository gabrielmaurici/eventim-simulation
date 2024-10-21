package reserve_ticket

import (
	"context"
	"testing"

	entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_WhenInputIsValid_Execute_ReturnsReservedTicket(t *testing.T) {
	input := ReserveTicketInputUseCaseDTO{
		UserToken: "1",
		Quantity:  2,
	}
	tickets := &[]entity.Ticket{
		{
			Id:        "1",
			Available: true,
		},
		{
			Id:        "2",
			Available: true,
		},
	}
	ctx := context.Background()
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketGatewayMock.On("GetAvailableTickets", input.Quantity).Return(tickets, nil)
	ticketReservationGatewayMock.On("CreateTicketReservation", input.UserToken, ctx).Return(nil)
	ticketGatewayMock.On("Update", mock.AnythingOfType("*entity.Ticket")).Return(nil)
	ticketGatewayMock.On("Update", mock.AnythingOfType("*entity.Ticket")).Return(nil)
	ticketReservationGatewayMock.On("RegisterTickets", input.UserToken, mock.AnythingOfType("[]string"), ctx).Return(nil)
	reserveTicketsUseCase := NewReserveTicket(ticketGatewayMock, ticketReservationGatewayMock)

	err := reserveTicketsUseCase.Execute(input, ctx)

	assert.Nil(t, err)
	ticketGatewayMock.AssertNumberOfCalls(t, "GetAvailableTickets", 1)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 2)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "CreateTicketReservation", 1)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "RegisterTickets", 1)
}

func Test_WhenInvalidInput_Execute_ReturnsError(t *testing.T) {
	input := ReserveTicketInputUseCaseDTO{
		UserToken: "",
		Quantity:  0,
	}
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	reserveTicketsUseCase := NewReserveTicket(ticketGatewayMock, ticketReservationGatewayMock)

	err := reserveTicketsUseCase.Execute(input, context.Background())

	assert.Error(t, err)
	assert.Equal(t, "token de usuário inválido. quantidade inválida", err.Error())
	ticketGatewayMock.AssertNumberOfCalls(t, "GetAvailableTickets", 0)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "CreateTicketReservation", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "RegisterTickets", 0)
}

func Test_WhenQuantityIsGreaterThanTen_Execute_ReturnsError(t *testing.T) {
	input := ReserveTicketInputUseCaseDTO{
		UserToken: "1",
		Quantity:  11,
	}
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	reserveTicketsUseCase := NewReserveTicket(ticketGatewayMock, ticketReservationGatewayMock)

	err := reserveTicketsUseCase.Execute(input, context.Background())

	assert.Error(t, err)
	assert.Equal(t, "não é possível reservar mais que 10 ingressos", err.Error())
	ticketGatewayMock.AssertNumberOfCalls(t, "GetAvailableTickets", 0)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "CreateTicketReservation", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "RegisterTickets", 0)
}

func Test_WhenTicketsUnavailable_Execute_ReturnsError(t *testing.T) {
	input := ReserveTicketInputUseCaseDTO{
		UserToken: "1",
		Quantity:  3,
	}
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketGatewayMock.On("GetAvailableTickets", input.Quantity).Return(&[]entity.Ticket{}, nil)
	reserverTicketsUseCase := NewReserveTicket(ticketGatewayMock, ticketReservationGatewayMock)

	err := reserverTicketsUseCase.Execute(input, context.Background())

	assert.Error(t, err)
	assert.Equal(t, "nenhum ingresso disponível", err.Error())
	ticketGatewayMock.AssertNumberOfCalls(t, "GetAvailableTickets", 1)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "CreateTicketReservation", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "RegisterTickets", 0)
}
