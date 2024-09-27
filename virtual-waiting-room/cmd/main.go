package main

import (
	"github.com/gabrielmaurici/eventim-simulation/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/entry_queue_usecase"
	"github.com/go-redis/redis"
)

func main() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	buyersActivesDb := database.NewBuyersActivesDb(redisDb)
	virtualQueueDb := database.NewVirtualQueueDb(redisDb)

	entryQueueUsecase := entry_queue_usecase.NewEntryQueueUseCase(buyersActivesDb, virtualQueueDb)
}
