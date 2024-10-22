package gateway

import entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"

type TicketGateway interface {
	Get(id string) (*entity.Ticket, error)
	GetAvailableTickets(quantity int8) (*[]entity.Ticket, error)
	Update(ticket *entity.Ticket) error
}
