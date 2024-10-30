package main

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/database"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/usecase/processing_buyers_actives"
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

	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(fmt.Errorf("erro ao conectar rabbitmq: %w", err))
	}
	defer rabbitmqConn.Close()

	consumer, err := rabbitmq.NewConsumer(rabbitmqConn, "buy_tickets_queue", "buy_tickets_exchange", "buy_tickets_routing_key", "direct")
	if err != nil {
		panic(fmt.Errorf("erro ao criar consumer rabbitmq: %w", err))
	}

	buyersActivesDb := database.NewBuyersActivesDb(redisDb)
	processingBuyersActivesUseCase := processing_buyers_actives.NewProcessingBuyersActivesUseCase(buyersActivesDb)
	processingBuyersActivesWorker := worker.NewProcessingBuyersActivesWorker(*consumer, *processingBuyersActivesUseCase)

	fmt.Println("Worker is running!")
	processingBuyersActivesWorker.Start()
}
