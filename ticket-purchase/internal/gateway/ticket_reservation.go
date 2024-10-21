package gateway

import "context"

type TicketReservationGateway interface {
	CreateTicketReservation(userToken string, ctx context.Context) error
	RegisterTickets(userToken string, ticketsId []string, ctx context.Context) error
	DeleteExpiredReservations(ctx context.Context) error
}
