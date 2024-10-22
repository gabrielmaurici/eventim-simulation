package processing_expired_tickets

import (
	"context"
	"errors"
	"testing"

	entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_WhenTicketsExpire_Execute_DeleteReservation(t *testing.T) {
	expiredTickets := &[]string{"1", "2"}
	ticket := &entity.Ticket{
		Id:        "1",
		Available: false,
	}
	ctx := context.Background()
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketReservationGatewayMock.On("GetAndDeleteExpiredTickets", ctx).Return(expiredTickets, nil)
	ticketGatewayMock.On("Get", mock.AnythingOfType("string")).Return(ticket, nil)
	ticketGatewayMock.On("Update", mock.AnythingOfType("*entity.Ticket")).Return(nil)
	ticketGatewayMock.On("Update", mock.AnythingOfType("*entity.Ticket")).Return(nil)
	processingExpiredTicketsUseCase := NewProcessingExpiredTicketsUseCase(ticketGatewayMock, ticketReservationGatewayMock)

	err := processingExpiredTicketsUseCase.Execute(ctx)

	assert.Nil(t, err)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetAndDeleteExpiredTickets", 1)
	ticketGatewayMock.AssertNumberOfCalls(t, "Get", 2)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 2)
}

func Test_WhenGetAndDeleteExpiredTicketsReturnsError_Execute_RetunsError(t *testing.T) {
	expiredTickets := &[]string{}
	ctx := context.Background()
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketReservationGatewayMock.On("GetAndDeleteExpiredTickets", ctx).Return(expiredTickets, errors.New("teste"))
	processingExpiredTicketsUseCase := NewProcessingExpiredTicketsUseCase(ticketGatewayMock, ticketReservationGatewayMock)

	err := processingExpiredTicketsUseCase.Execute(ctx)

	assert.Error(t, err)
	assert.Equal(t, "erro ao obter e deletar ingressos expirados: teste", err.Error())
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetAndDeleteExpiredTickets", 1)
	ticketGatewayMock.AssertNumberOfCalls(t, "Get", 0)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 0)
}

func Test_WhenTicketsNotExpired_Execute_Retun(t *testing.T) {
	expiredTickets := &[]string{}
	ctx := context.Background()
	ticketGatewayMock := &mocks.TicketGatewayMock{}
	ticketReservationGatewayMock := &mocks.TicketReservationGatewayMock{}
	ticketReservationGatewayMock.On("GetAndDeleteExpiredTickets", ctx).Return(expiredTickets, nil)
	processingExpiredTicketsUseCase := NewProcessingExpiredTicketsUseCase(ticketGatewayMock, ticketReservationGatewayMock)

	err := processingExpiredTicketsUseCase.Execute(ctx)

	assert.Nil(t, err)
	ticketReservationGatewayMock.AssertNumberOfCalls(t, "GetAndDeleteExpiredTickets", 1)
	ticketGatewayMock.AssertNumberOfCalls(t, "Get", 0)
	ticketGatewayMock.AssertNumberOfCalls(t, "Update", 0)
}
