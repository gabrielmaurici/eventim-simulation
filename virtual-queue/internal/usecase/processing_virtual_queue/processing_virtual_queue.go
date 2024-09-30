package processing_virtual_queue

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/internal/gateway"
	"github.com/gabrielmaurici/eventim-simulation/pkg/rabbitmq"
)

type NotificationPositionRabbitMQ struct {
	Token    string `json:"token"`
	Position int64  `json:"position"`
}

type ProcessingVirtualQueueUseCase struct {
	BuyersActivesGateway gateway.BuyersActivesGateway
	VirtualQueueGateway  gateway.VirtualQueueGateway
	Producer             rabbitmq.Producer
}

const MaxBuyersActivesCapacity = 5

func NewProcessingVirtualQueueUseCase(
	b gateway.BuyersActivesGateway,
	v gateway.VirtualQueueGateway,
	p rabbitmq.Producer) *ProcessingVirtualQueueUseCase {
	return &ProcessingVirtualQueueUseCase{
		BuyersActivesGateway: b,
		VirtualQueueGateway:  v,
		Producer:             p,
	}
}

func (uc *ProcessingVirtualQueueUseCase) Execute() {
	totalBuyersActives, err := uc.BuyersActivesGateway.GetBuyersActives()
	if err != nil {
		fmt.Println("erro ao obter compradores ativos: %w", err)
	}

	if totalBuyersActives == MaxBuyersActivesCapacity {
		return
	}

	quantityNextBuyersActives := MaxBuyersActivesCapacity - totalBuyersActives
	uc.updateAndNotificationNextBuyersActives(int(quantityNextBuyersActives))
	uc.updateAndNotificationPositionVirtualQueue()
}

func (uc *ProcessingVirtualQueueUseCase) updateAndNotificationNextBuyersActives(quantityNextBuyersActives int) {
	for i := 0; i < quantityNextBuyersActives; i++ {
		token, err := uc.VirtualQueueGateway.Dequeue()
		if err != nil {
			fmt.Println("erro ao remover token da fila: %w", err)
			continue
		}
		uc.BuyersActivesGateway.Add(token)
		uc.Producer.Publish(NotificationPositionRabbitMQ{
			Token:    token,
			Position: 0,
		})
	}
}

func (uc *ProcessingVirtualQueueUseCase) updateAndNotificationPositionVirtualQueue() {
	tokensInQueue, err := uc.VirtualQueueGateway.GetAll()
	if err != nil {
		fmt.Println("erro ao obter tokens da fila: %w", err)
		return
	}

	for index, token := range tokensInQueue {
		position := index + 1
		uc.Producer.Publish(NotificationPositionRabbitMQ{
			Token:    token,
			Position: int64(position),
		})
	}
}
