package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type BuyersActivesDb struct {
	RedisDb *redis.Client
}

const buyersActivesKey string = "buyers_actives:"

func NewBuyersActivesDb(r *redis.Client) *BuyersActivesDb {
	return &BuyersActivesDb{
		RedisDb: r,
	}
}

func (db *BuyersActivesDb) GetBuyersActives(ctx context.Context) (total int64, err error) {
	pattern := buyersActivesKey + "*"
	var cursor uint64
	var count int

	for {
		keys, nextCursor, err := db.RedisDb.Scan(ctx, cursor, pattern, 6).Result()
		if err != nil {
			return 0, err
		}
		count += len(keys)
		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}
	return int64(count), nil
}

func (db *BuyersActivesDb) Add(token string, ctx context.Context) error {
	expiration := 30 * time.Second
	key := buyersActivesKey + token
	return db.RedisDb.Set(ctx, key, token, expiration).Err()
}

func (db *BuyersActivesDb) Delete(token string, ctx context.Context) error {
	key := buyersActivesKey + token
	return db.RedisDb.Del(ctx, key).Err()
}
