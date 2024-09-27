package database

import (
	"github.com/go-redis/redis"
)

type BuyersActivesDb struct {
	RedisDb *redis.Client
}

const buyersActivesDb string = "buyers_actives_key"

func NewBuyersActivesDb(r *redis.Client) *BuyersActivesDb {
	return &BuyersActivesDb{
		RedisDb: r,
	}
}

func (db *BuyersActivesDb) GetBuyersActives() (total int64, err error) {
	total, err = db.RedisDb.LLen(buyersActivesDb).Result()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (db *BuyersActivesDb) Add(token string) error {
	_, err := db.RedisDb.LPush(buyersActivesDb, token).Result()
	return err
}
