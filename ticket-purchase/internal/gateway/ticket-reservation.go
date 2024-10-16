package gateway

type TicketReservationGateway interface {
	Reserve(userToken string, ticketId string) error
}
