package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/websocket"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@rabbitmq-virtual-queue:5672/")
	if err != nil {
		panic(fmt.Errorf("erro ao conectar rabbitmq: %w", err))
	}
	defer rabbitmqConn.Close()

	consumer, err := rabbitmq.NewConsumer(rabbitmqConn, "", "virtual_queue_exchange", "fanout")
	if err != nil {
		panic(fmt.Errorf("erro ao criar consumer rabbitmq: %w", err))
	}
	msgChan := make(chan []byte)
	go consumer.Consume(msgChan)

	webSocketVirtualQueueHandler := websocket.NewWebSocketVirtualQueueHandler(msgChan)
	http.HandleFunc("/ws/virtual-queue", webSocketVirtualQueueHandler.NotifyPositionSocket)
	fmt.Println("Websocket is running!")
	http.ListenAndServe(":5001", nil)
}
