package usecase

import (
	"testing"

	entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_WhenInputIsValid_Execute_ReturnsReservedTickets(t *testing.T) {
	input := ReserverTicketsInputUseCaseDTO{
		UserToken: "1",
		Quantity:  2,
	}
	tickets := []entity.Ticket{
		*entity.NewTicket("1", true),
		*entity.NewTicket("2", true),
	}
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketGatewayMock.On("GetAvailableTickets", input.Quantity).Return(tickets, nil)
	ticketReservationGatewayMock.On("Reserve", input.UserToken, tickets[0].Id).Return(nil)
	ticketReservationGatewayMock.On("Reserve", input.UserToken, tickets[1].Id).Return(nil)
	ticketGatewayMock.On("Update", mock.AnythingOfType("entity.Ticket")).Return(nil)
	ticketGatewayMock.On("Update", mock.AnythingOfType("entity.Ticket")).Return(nil)
	reserverTicketsUseCase := NewReserverTickets(ticketGatewayMock, ticketReservationGatewayMock)

	output, _ := reserverTicketsUseCase.Execute(input)

	assert.Equal(t, len(tickets), len(output))
	assert.Equal(t, tickets[0].Id, output[0].TicketId)
	assert.Equal(t, tickets[1].Id, output[1].TicketId)
}

func Test_WhenInvalidInputs_Execute_ReturnsError(t *testing.T) {
	input := ReserverTicketsInputUseCaseDTO{
		UserToken: "",
		Quantity:  0,
	}
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	reserverTicketsUseCase := NewReserverTickets(ticketGatewayMock, ticketReservationGatewayMock)

	output, err := reserverTicketsUseCase.Execute(input)

	assert.Nil(t, output)
	assert.Error(t, err)
	assert.Equal(t, "falha ao reservar ingressos: quantidade de ingressos inválido. token de usuário inválido", err.Error())
}

func Test_WhenTicketsUnavailable_Execute_ReturnsError(t *testing.T) {
	input := ReserverTicketsInputUseCaseDTO{
		UserToken: "1",
		Quantity:  2,
	}
	tickets := []entity.Ticket{}
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketGatewayMock.On("GetAvailableTickets", input.Quantity).Return(tickets, nil)
	reserverTicketsUseCase := NewReserverTickets(ticketGatewayMock, ticketReservationGatewayMock)

	output, err := reserverTicketsUseCase.Execute(input)

	assert.Nil(t, output)
	assert.Error(t, err)
	assert.Equal(t, "ingressos indisponíveis", err.Error())
}
