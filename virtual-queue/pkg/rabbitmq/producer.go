package rabbitmq

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	QueueName  string
}

func NewProducer(conn *amqp.Connection, queueName string) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir canal: %w", err)
	}

	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao declarar fila: %w", err)
	}

	return &Producer{
		Connection: conn,
		Channel:    ch,
		QueueName:  queueName,
	}, nil
}

func (p *Producer) Publish(msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("erro ao tratar mensagem para publicação: %w", err)
	}

	err = p.Channel.Publish(
		"",          // exchange
		p.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
