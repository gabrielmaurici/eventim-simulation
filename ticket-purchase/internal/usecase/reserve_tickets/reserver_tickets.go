package usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/gateway"
)

type ReserverTicketsInputUseCaseDTO struct {
	UserToken string `json:"user_token"`
	Quantity  int8   `json:"quantity"`
}

type ReserverTicketsOutputUseCaseDTO struct {
	TicketId string `json:"ticket_id"`
}

type ReserverTicketsUseCase struct {
	TicketGateway            gateway.TicketGateway
	TicketReservationGateway gateway.TicketReservationGateway
}

func NewReserverTickets(tg gateway.TicketGateway, trg gateway.TicketReservationGateway) *ReserverTicketsUseCase {
	return &ReserverTicketsUseCase{
		TicketGateway:            tg,
		TicketReservationGateway: trg,
	}
}

func (uc *ReserverTicketsUseCase) Execute(input ReserverTicketsInputUseCaseDTO) (output []ReserverTicketsOutputUseCaseDTO, err error) {
	err = inputValidation(input)
	if err != nil {
		return nil, err
	}

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

		ticketDto := ReserverTicketsOutputUseCaseDTO{
			TicketId: ticket.Id,
		}
		output = append(output, ticketDto)
	}

	return output, nil
}

func inputValidation(input ReserverTicketsInputUseCaseDTO) error {
	var errs []string
	if input.Quantity <= 0 {
		errs = append(errs, "quantidade de ingressos inválido")
	}

	if input.UserToken == "" {
		errs = append(errs, "token de usuário inválido")
	}

	if len(errs) > 0 {
		return fmt.Errorf("falha ao reservar ingressos: %v", strings.Join(errs, ". "))
	}

	return nil
}
