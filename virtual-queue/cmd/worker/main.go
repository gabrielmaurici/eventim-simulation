package main

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/usecase/processing_virtual_queue"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/worker"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "redis-virtual-queue:6379",
		Password: "",
		DB:       0,
	})

	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@rabbitmq-virtual-queue:5672/")
	if err != nil {
		panic(fmt.Errorf("erro ao conectar rabbitmq: %w", err))
	}
	defer rabbitmqConn.Close()

	producer, err := rabbitmq.NewProducer(rabbitmqConn, "virtual_queue_exchange", "fanout")
	if err != nil {
		panic(fmt.Errorf("erro ao criar produtor rabbitmq: %w", err))
	}

	buyersActivesDb := database.NewBuyersActivesDb(redisDb)
	virtualQueueDb := database.NewVirtualQueueDb(redisDb)
	processingVirtualQueueUseCase := processing_virtual_queue.NewProcessingVirtualQueueUseCase(buyersActivesDb, virtualQueueDb, producer)
	processingVirtualQueueWorker := worker.NewProcessingVirtualQueueWorker(*processingVirtualQueueUseCase)
	fmt.Println("Worker is running!")

	processingVirtualQueueWorker.Start()
}
