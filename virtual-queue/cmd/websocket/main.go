package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/internal/websocket"
	"github.com/gabrielmaurici/eventim-simulation/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitmqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(fmt.Errorf("erro ao conectar rabbitmq: %w", err))
	}
	defer rabbitmqConn.Close()

	consumer, err := rabbitmq.NewConsumer(rabbitmqConn, "virtual_queue")
	if err != nil {
		panic(fmt.Errorf("erro ao criar consumer rabbitmq: %w", err))
	}

	msgChan := make(chan []byte)
	go consumer.Consume(msgChan)

	webSocketVirtualQueueHandler := websocket.NewWebSocketVirtualQueueHandler(msgChan)

	http.HandleFunc("/ws/virtual-queue", webSocketVirtualQueueHandler.NotifyPositionSocket)

	fmt.Println("Websocket is running!")
	if err := http.ListenAndServe(":5001", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
