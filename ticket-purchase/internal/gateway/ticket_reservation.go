package gateway

import "context"

type TicketReservationGateway interface {
	CreateTicketReservation(userToken string, ctx context.Context) error
	RegisterTickets(userToken string, ticketsId []string, ctx context.Context) error
	GetAndDeleteExpiredTickets(ctx context.Context) (expiredTickets *[]string, err error)
	GetReservedTickets(userToken string, ctx context.Context) (reservedTickets []string, err error)
	DeleteReservedTickets(userToken string, ctx context.Context) error
}
