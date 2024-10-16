package usecase

import (
	"errors"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/gateway"
)

type ReserveTicketsInputUseCaseDTO struct {
	UserToken string `json:"user_token"`
	Quantity  int8   `json:"quantity"`
}

type ReserveTicketsOutputUseCaseDTO struct {
	TicketId string `json:"ticket_id"`
}

type ReserveTicketsUseCase struct {
	TicketGateway            gateway.TicketGateway
	TicketReservationGateway gateway.TicketReservationGateway
}

func NewReserveTickets(tg gateway.TicketGateway, trg gateway.TicketReservationGateway) *ReserveTicketsUseCase {
	return &ReserveTicketsUseCase{
		TicketGateway:            tg,
		TicketReservationGateway: trg,
	}
}

func (uc *ReserveTicketsUseCase) Execute(input ReserveTicketsInputUseCaseDTO) (output []ReserveTicketsOutputUseCaseDTO, err error) {
	tickets, err := uc.TicketGateway.GetAvailableTickets(input.Quantity)
	if err != nil {
		return nil, err
	}

	if len(tickets) <= 0 {
		return nil, errors.New("ingressos indisponíveis")
	}

	for _, ticket := range tickets {
		err = uc.TicketReservationGateway.Reserve(input.UserToken, ticket.Id)
		if err != nil {
			return nil, err
		}

		ticket.UpdateToUnavailable()
		err = uc.TicketGateway.Update(ticket)
		if err != nil {
			return nil, err
		}

		ticketDto := ReserveTicketsOutputUseCaseDTO{
			TicketId: ticket.Id,
		}
		output = append(output, ticketDto)
	}

	return output, nil
}
