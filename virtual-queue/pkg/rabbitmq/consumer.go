package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
}

func NewConsumer(conn *amqp.Connection, queueName string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir canal: %w", err)
	}

	return &Consumer{
		Connection: conn,
		Channel:    ch,
		QueueName:  queueName,
	}, nil
}

func (c *Consumer) Consume(msgChan chan []byte) error {
	msgs, err := c.Channel.Consume(
		c.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("falha ao receber mensagens: %w", err)
	}

	for msg := range msgs {
		msgChan <- msg.Body
	}

	return nil
}
