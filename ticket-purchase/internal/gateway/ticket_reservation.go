package gateway

import "context"

type TicketReservationGateway interface {
	HasReservation(userToken string, ctx context.Context) (Has bool, err error)
	CreateTicketReservation(userToken string, ctx context.Context) error
	RegisterTickets(userToken string, ticketsId []string, ctx context.Context) error
}
