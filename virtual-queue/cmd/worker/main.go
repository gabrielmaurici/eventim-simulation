package main

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/processing_virtual_queue"
	"github.com/gabrielmaurici/eventim-simulation/internal/worker"
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
	processingVirtualQueueUseCase := processing_virtual_queue.NewProcessingVirtualQueueUseCase(buyersActivesDb, virtualQueueDb)
	processingVirtualQueueWorker := worker.NewProcessingVirtualQueueWorker(*processingVirtualQueueUseCase)
	fmt.Println("Worker is running!")

	processingVirtualQueueWorker.Start()
}
