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
const ticketReservationKey string = "ticket_reservaiton:"

func NewTicketReservationDb(db *redis.Client) *TicketReservationDb {
	return &TicketReservationDb{
		Db: db,
	}
}

func (trb *TicketReservationDb) Reserve(userToken string, ticketId string, ctx context.Context) error {
	key := ticketReservationKey + userToken

	err := trb.Db.LPush(ctx, key, ticketId).Err()
	if err != nil {
		return err
	}

	err = trb.Db.Expire(ctx, key, expireTime).Err()
	if err != nil {
		return err
	}

	return nil
}
