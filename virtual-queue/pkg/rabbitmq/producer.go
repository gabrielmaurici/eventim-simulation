package rabbitmq

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Exchange   string
}

func NewProducer(conn *amqp.Connection, exchange, exchangeKind string) (*Producer, error) {
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

	return &Producer{
		Connection: conn,
		Channel:    ch,
		Exchange:   exchange,
	}, nil
}

func (p *Producer) Publish(msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("erro ao tratar mensagem para publicação: %w", err)
	}

	err = p.Channel.Publish(
		p.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("erro ao publicar mensagem: %w", err)
	}

	return nil
}
