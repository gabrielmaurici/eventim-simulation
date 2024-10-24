package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type BuyersActivesDb struct {
	RedisDb *redis.Client
}

const buyersActivesCountKey string = "buyers_actives_count_key"

func NewBuyersActivesDb(r *redis.Client) *BuyersActivesDb {
	return &BuyersActivesDb{
		RedisDb: r,
	}
}

func (db *BuyersActivesDb) GetBuyersActives(ctx context.Context) (total int64, err error) {
	expiration := fmt.Sprintf("%d", time.Now().Unix())
	_, err = db.RedisDb.ZRemRangeByScore(ctx, buyersActivesCountKey, "-inf", expiration).Result()
	if err != nil {
		return 0, err
	}

	total, err = db.RedisDb.ZCount(ctx, buyersActivesCountKey, expiration, "+inf").Result()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (db *BuyersActivesDb) Add(token string, ctx context.Context) error {
	expiration := time.Now().Add(30 * time.Second).Unix()
	err := db.RedisDb.ZAdd(ctx, buyersActivesCountKey, redis.Z{
		Score:  float64(expiration),
		Member: token,
	}).Err()
	if err != nil {
		return err
	}

	return nil
}
