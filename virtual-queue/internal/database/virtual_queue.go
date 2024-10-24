package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type VirtualQueueDb struct {
	RedisDb *redis.Client
}

const virtualQueueKey string = "virtual_queue_key"

func NewVirtualQueueDb(r *redis.Client) *VirtualQueueDb {
	return &VirtualQueueDb{
		RedisDb: r,
	}
}

func (db *VirtualQueueDb) Enqueue(token string, ctx context.Context) (position int64, err error) {
	position, err = db.RedisDb.RPush(ctx, virtualQueueKey, token).Result()
	if err != nil {
		return 0, err
	}
	return position, nil
}

func (db *VirtualQueueDb) Dequeue(ctx context.Context) (token string, err error) {
	token, err = db.RedisDb.LPop(ctx, virtualQueueKey).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (db *VirtualQueueDb) GetAll(ctx context.Context) (tokens []string, err error) {
	tokens, err = db.RedisDb.LRange(ctx, virtualQueueKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
