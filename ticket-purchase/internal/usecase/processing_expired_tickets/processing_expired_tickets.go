package processing_expired_tickets

import (
	"context"
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/gateway"
)

type ProcessingExpiredTicketsUseCase struct {
	TicketGateway            gateway.TicketGateway
	TicketReservationGateway gateway.TicketReservationGateway
}

func NewProcessingExpiredTicketsUseCase(
	tg gateway.TicketGateway,
	trg gateway.TicketReservationGateway) *ProcessingExpiredTicketsUseCase {
	return &ProcessingExpiredTicketsUseCase{
		TicketGateway:            tg,
		TicketReservationGateway: trg,
	}
}

func (uc *ProcessingExpiredTicketsUseCase) Execute(ctx context.Context) error {
	expiredTickets, err := uc.TicketReservationGateway.GetAndDeleteExpiredTickets(ctx)
	if err != nil {
		return err
	}

	if len(expiredTickets) <= 0 {
		return nil
	}

	for _, ticket := range expiredTickets {
		ticket, err := uc.TicketGateway.Get(ticket)
		if err != nil {
			fmt.Println("erro ao obter ingresso: %w", err)
			continue
		}

		ticket.UpdateToAvailable()
		err = uc.TicketGateway.Update(ticket)
		if err != nil {
			fmt.Println("erro ao atualizar ingresso: %w", err)
			continue
		}
	}

	return nil
}
