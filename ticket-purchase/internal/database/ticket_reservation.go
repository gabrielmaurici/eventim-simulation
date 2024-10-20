package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TicketReservationDb struct {
	Db *redis.Client
}

const expireTime time.Duration = 1 * time.Minute
const ticketReservationKey string = "ticket_reservation:"
const listTicketsReservationKey string = "ticket_reservation:tickets:"

func NewTicketReservationDb(db *redis.Client) *TicketReservationDb {
	return &TicketReservationDb{
		Db: db,
	}
}

func (trb *TicketReservationDb) HasReservation(userToken string, ctx context.Context) (Has bool, err error) {
	reservationKey := ticketReservationKey + userToken
	reservation, err := trb.Db.HGetAll(ctx, reservationKey).Result()
	if err != nil {
		return false, err
	}
	if len(reservation) == 0 {
		return false, nil
	}

	return true, nil
}

func (trb *TicketReservationDb) CreateTicketReservation(userToken string, ctx context.Context) error {
	reservationKey := ticketReservationKey + userToken
	reservationFields := []string{
		"teste", "teste1",
		"dateReservation", time.Now().String(),
	}
	err := trb.Db.HSet(ctx, reservationKey, reservationFields).Err()
	if err != nil {
		return err
	}

	errExpire := trb.Db.HExpire(ctx, reservationKey, expireTime, "dateReservation").Err()
	if errExpire != nil {
		return errExpire
	}

	return nil
}

func (trb *TicketReservationDb) RegisterTickets(userToken string, ticketsId []string, ctx context.Context) error {
	ticketsReservationKey := listTicketsReservationKey + userToken
	err := trb.Db.LPush(ctx, ticketsReservationKey, ticketsId).Err()
	if err != nil {
		return err
	}

	return nil
}
