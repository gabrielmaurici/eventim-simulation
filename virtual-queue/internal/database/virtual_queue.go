package database

import (
	"github.com/go-redis/redis"
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

func (db *VirtualQueueDb) Enqueue(token string) (position int64, err error) {
	position, err = db.RedisDb.RPush(virtualQueueKey, token).Result()
	if err != nil {
		return 0, err
	}
	return position, nil
}
