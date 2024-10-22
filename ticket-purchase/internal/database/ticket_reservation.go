package database

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type TicketReservationDb struct {
	Db *redis.Client
}

const expireTime time.Duration = 1 * time.Minute
const ticketReservationKey string = "ticket_reservation:"
const listTicketsReservationKey string = "ticket_reservation_tickets:"

func NewTicketReservationDb(db *redis.Client) *TicketReservationDb {
	return &TicketReservationDb{
		Db: db,
	}
}

func (trb *TicketReservationDb) CreateTicketReservation(userToken string, ctx context.Context) error {
	reservationKey := ticketReservationKey + userToken
	reservationFields := []string{
		"dateReservation", time.Now().Local().UTC().Format("02/01/2006 15:04:05"),
		"expire", expireTime.String(),
	}
	err := trb.Db.HSet(ctx, reservationKey, reservationFields).Err()
	if err != nil {
		return err
	}

	errExpire := trb.Db.HExpire(ctx, reservationKey, expireTime, "expire").Err()
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

func (trb *TicketReservationDb) GetAndDeleteExpiredTickets(ctx context.Context) (expiredTickets *[]string, err error) {
	expiredTickets = &[]string{}
	reservationKey := ticketReservationKey + "*"
	reservations, err := trb.Db.Keys(ctx, reservationKey).Result()
	if err != nil {
		return nil, err
	}

	for _, key := range reservations {
		expireField, err := trb.Db.HExists(ctx, key, "expire").Result()
		if err != nil {
			continue
		}
		if expireField {
			continue
		}

		userToken := strings.Split(key, ":")[1]
		ticketsReservationKey := listTicketsReservationKey + userToken
		expiredTicketsUser, err := trb.Db.LRange(ctx, ticketsReservationKey, 0, -1).Result()
		if err != nil {
			continue
		}
		if len(expiredTicketsUser) > 0 {
			*expiredTickets = append(*expiredTickets, expiredTicketsUser...)
		}

		trb.Db.Del(ctx, key, ticketsReservationKey)
	}

	return expiredTickets, nil
}
