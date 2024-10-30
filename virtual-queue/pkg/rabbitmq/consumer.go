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

func NewConsumer(conn *amqp.Connection, queueName, exchange, routingKey, exchangeKind string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir canal: %w", err)
	}

	err = ch.ExchangeDeclare(
		exchange,
		exchangeKind,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao declarar exchange: %w", err)
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao declarar fila: %w", err)
	}

	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao vincular fila à exchange: %w", err)
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
