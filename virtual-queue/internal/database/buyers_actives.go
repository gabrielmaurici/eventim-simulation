package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type BuyersActivesDb struct {
	RedisDb *redis.Client
	Context context.Context
}

const buyersActivesCountKey string = "buyers_actives_count_key"

func NewBuyersActivesDb(r *redis.Client, ctx context.Context) *BuyersActivesDb {
	return &BuyersActivesDb{
		RedisDb: r,
		Context: ctx,
	}
}

func (db *BuyersActivesDb) GetBuyersActives() (total int64, err error) {
	expiration := fmt.Sprintf("%d", time.Now().Unix())
	_, err = db.RedisDb.ZRemRangeByScore(db.Context, buyersActivesCountKey, "-inf", expiration).Result()
	if err != nil {
		return 0, err
	}

	total, err = db.RedisDb.ZCount(db.Context, buyersActivesCountKey, expiration, "+inf").Result()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (db *BuyersActivesDb) Add(token string) error {
	expiration := time.Now().Add(30 * time.Second).Unix()
	err := db.RedisDb.ZAdd(db.Context, buyersActivesCountKey, redis.Z{
		Score:  float64(expiration),
		Member: token,
	}).Err()
	if err != nil {
		return err
	}

	return nil
}
