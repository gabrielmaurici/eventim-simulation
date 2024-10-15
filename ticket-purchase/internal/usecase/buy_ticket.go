package usecase

import (
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/gateway"
)

type BuyTicketUseCase struct {
	TicketGateway gateway.TicketGateway
}
