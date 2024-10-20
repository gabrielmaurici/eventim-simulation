package reserve_ticket

import (
	"context"
	"errors"
	"fmt"
	"strings"

	entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/gateway"
)

type ReserveTicketInputUseCaseDTO struct {
	UserToken string `json:"user_token"`
	Quantity  int8   `json:"quantity"`
}

type ReserveTicketUseCase struct {
	TicketGateway            gateway.TicketGateway
	TicketReservationGateway gateway.TicketReservationGateway
}

func NewReserveTicket(tg gateway.TicketGateway, trg gateway.TicketReservationGateway) *ReserveTicketUseCase {
	return &ReserveTicketUseCase{
		TicketGateway:            tg,
		TicketReservationGateway: trg,
	}
}

func (uc *ReserveTicketUseCase) Execute(input ReserveTicketInputUseCaseDTO, ctx context.Context) error {
	err := validateInput(input)
	if err != nil {
		return err
	}

	tickets, err := uc.GetAvailableTickets(input.Quantity)
	if err != nil {
		return err
	}

	reservation, err := uc.TicketReservationGateway.HasReservation(input.UserToken, ctx)
	if err != nil {
		return fmt.Errorf("erro ao verificar se já possui reserva. %w", err)
	}

	if !reservation {
		if err = uc.TicketReservationGateway.CreateTicketReservation(input.UserToken, ctx); err != nil {
			return fmt.Errorf("erro ao gerar registro da reserva. %w", err)
		}
	}

	var ticketsIds []string
	for _, ticket := range *tickets {
		ticketsIds = append(ticketsIds, ticket.Id)

		ticket.UpdateToUnavailable()
		err = uc.TicketGateway.Update(&ticket)
		if err != nil {
			return err
		}
	}

	err = uc.TicketReservationGateway.RegisterTickets(input.UserToken, ticketsIds, ctx)
	if err != nil {
		return fmt.Errorf("erro ao gravar ingressos para a reserva. %w", err)
	}

	return nil
}

func validateInput(input ReserveTicketInputUseCaseDTO) error {
	var err []string
	if input.UserToken == "" {
		err = append(err, "token de usuário inválido")
	}
	if input.Quantity <= 0 {
		err = append(err, "quantidade inválida")
	}
	if input.Quantity > 10 {
		err = append(err, "não é possível reservar mais que 10 ingressos")
	}

	if len(err) > 0 {
		return errors.New(strings.Join(err, ". "))
	}

	return nil
}

func (uc *ReserveTicketUseCase) GetAvailableTickets(quantity int8) (*[]entity.Ticket, error) {
	tickets, err := uc.TicketGateway.GetAvailableTickets(quantity)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter ingresso disponível. %w", err)
	}
	if len(*tickets) <= 0 {
		return nil, errors.New("nenhum ingresso disponível")
	}

	return tickets, nil
}
