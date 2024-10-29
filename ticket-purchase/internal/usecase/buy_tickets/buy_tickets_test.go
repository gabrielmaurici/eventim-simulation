package buy_tickets

import (
	"context"
	"errors"
	"testing"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_WhenInputIsValid_Execute_ReturnsPurchasedTickets(t *testing.T) {
	input := BuyTicketsInputDTO{
		UserToken: "1",
	}
	tickets := []string{
		"123",
		"456",
	}
	ctx := context.Background()
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	ticketReservationGatewayMock.On("GetReservedTickets", input.UserToken, ctx).Return(tickets, nil)
	ticketReservationGatewayMock.On("DeleteReservedTickets", input.UserToken, ctx).Return(nil)
	producerMock.On("Publish", input).Return(nil)
	buyTicketsUseCase := NewBuyTicketsUseCase(ticketReservationGatewayMock, producerMock, func() bool {
		return true
	})

	output, err := buyTicketsUseCase.Execute(input, ctx)

	assert.Nil(t, err)
	assert.Equal(t, tickets[0], output.TicketsPurchased[0])
	assert.Equal(t, tickets[1], output.TicketsPurchased[1])
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetReservedTickets", 1)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "DeleteReservedTickets", 1)
	producerMock.AssertNumberOfCalls(t, "Publish", 1)
}

func Test_WhenInputIsValidButThereIsAnErrorInThePayment_Execute_ReturnsErrorPayment(t *testing.T) {
	input := BuyTicketsInputDTO{
		UserToken: "1",
	}
	tickets := []string{
		"123",
		"456",
	}
	ctx := context.Background()
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	ticketReservationGatewayMock.On("GetReservedTickets", input.UserToken, ctx).Return(tickets, nil)
	buyTicketsUseCase := NewBuyTicketsUseCase(ticketReservationGatewayMock, producerMock, func() bool {
		return false
	})

	output, err := buyTicketsUseCase.Execute(input, ctx)

	assert.Nil(t, output)
	assert.Equal(t, "erro ao realizar pagamento, tente novamente", err.Error())
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetReservedTickets", 1)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "DeleteReservedTickets", 0)
	producerMock.AssertNumberOfCalls(t, "Publish", 0)
}

func Test_WhenInputIsInvalid_Execute_ReturnsError(t *testing.T) {
	input := BuyTicketsInputDTO{
		UserToken: "",
	}
	ctx := context.Background()
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	buyTicketsUseCase := NewBuyTicketsUseCase(ticketReservationGatewayMock, producerMock, func() bool {
		return false
	})

	output, err := buyTicketsUseCase.Execute(input, ctx)

	assert.Nil(t, output)
	assert.Equal(t, "token de usuário inválido", err.Error())
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetReservedTickets", 0)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "DeleteReservedTickets", 0)
	producerMock.AssertNumberOfCalls(t, "Publish", 0)
}

func Test_WhenNotHaveReservedTickets_Execute_ReturnsError(t *testing.T) {
	input := BuyTicketsInputDTO{
		UserToken: "123",
	}
	tickets := []string{}
	ctx := context.Background()
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	ticketReservationGatewayMock.On("GetReservedTickets", input.UserToken, ctx).Return(tickets, nil)
	buyTicketsUseCase := NewBuyTicketsUseCase(ticketReservationGatewayMock, producerMock, func() bool {
		return false
	})

	output, err := buyTicketsUseCase.Execute(input, ctx)

	assert.Nil(t, output)
	assert.Equal(t, "reserva expirada ou inexistente", err.Error())
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetReservedTickets", 1)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "DeleteReservedTickets", 0)
	producerMock.AssertNumberOfCalls(t, "Publish", 0)
}

func Test_WhenThereIsAnErrorWhenDeletingReservations_Execute_ReturnsError(t *testing.T) {
	input := BuyTicketsInputDTO{
		UserToken: "123",
	}
	tickets := []string{
		"123",
		"456",
	}
	ctx := context.Background()
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	ticketReservationGatewayMock.On("GetReservedTickets", input.UserToken, ctx).Return(tickets, nil)
	ticketReservationGatewayMock.On("DeleteReservedTickets", input.UserToken, ctx).Return(errors.New("teste"))
	buyTicketsUseCase := NewBuyTicketsUseCase(ticketReservationGatewayMock, producerMock, func() bool {
		return true
	})

	output, err := buyTicketsUseCase.Execute(input, ctx)

	assert.Nil(t, output)
	assert.Equal(t, "erro ao remover a reserva: teste", err.Error())
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetReservedTickets", 1)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "DeleteReservedTickets", 1)
	producerMock.AssertNumberOfCalls(t, "Publish", 0)
}
