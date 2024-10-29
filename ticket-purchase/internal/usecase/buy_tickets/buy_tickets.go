package buy_tickets

import (
	"context"
	"errors"
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/gateway"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/pkg/rabbitmq"
)

type BuyTicketsInputDTO struct {
	UserToken string `json:"user_token"`
}

type BuyTicketsOutputDTO struct {
	TicketsPurchased []string `json:"tickets_purchased"`
}

type BuyTicketsUseCase struct {
	TicketReservationGateway gateway.TicketReservationGateway
	Producer                 rabbitmq.ProducerInterface
	SimulatePaymentFunc      func() bool
}

func NewBuyTicketsUseCase(
	trg gateway.TicketReservationGateway,
	p rabbitmq.ProducerInterface,
	spf func() bool) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{
		TicketReservationGateway: trg,
		Producer:                 p,
		SimulatePaymentFunc:      spf,
	}
}

func (uc *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO, ctx context.Context) (output *BuyTicketsOutputDTO, err error) {
	if input.UserToken == "" {
		return nil, errors.New("token de usuário inválido")
	}

	reservedTickets, err := uc.TicketReservationGateway.GetReservedTickets(input.UserToken, ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter a reserva dos ingressos: %w", err)
	}

	if len(reservedTickets) == 0 {
		return nil, errors.New("reserva expirada ou inexistente")
	}

	payment := uc.SimulatePaymentFunc()
	if !payment {
		return nil, errors.New("erro ao realizar pagamento, tente novamente")
	}

	err = uc.TicketReservationGateway.DeleteReservedTickets(input.UserToken, ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao remover a reserva: %w", err)
	}

	uc.Producer.Publish(input)

	return &BuyTicketsOutputDTO{
		TicketsPurchased: reservedTickets,
	}, nil
}
