package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type VirtualQueueDb struct {
	RedisDb *redis.Client
	Context context.Context
}

const virtualQueueKey string = "virtual_queue_key"

func NewVirtualQueueDb(r *redis.Client, ctx context.Context) *VirtualQueueDb {
	return &VirtualQueueDb{
		RedisDb: r,
		Context: ctx,
	}
}

func (db *VirtualQueueDb) Enqueue(token string) (position int64, err error) {
	position, err = db.RedisDb.RPush(db.Context, virtualQueueKey, token).Result()
	if err != nil {
		return 0, err
	}
	return position, nil
}

func (db *VirtualQueueDb) Dequeue() (token string, err error) {
	token, err = db.RedisDb.LPop(db.Context, virtualQueueKey).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (db *VirtualQueueDb) GetAll() (tokens []string, err error) {
	tokens, err = db.RedisDb.LRange(db.Context, virtualQueueKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
