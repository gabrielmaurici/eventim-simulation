package processing_virtual_queue

import (
	"context"
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/gateway"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/pkg/rabbitmq"
)

type NotificationPositionRabbitMQ struct {
	Token             string `json:"token"`
	Position          int64  `json:"position"`
	EstimatedWaitTime int64  `json:"estimated_wait_time"`
}

type ProcessingVirtualQueueUseCase struct {
	BuyersActivesGateway gateway.BuyersActivesGateway
	VirtualQueueGateway  gateway.VirtualQueueGateway
	Producer             rabbitmq.ProducerInterface
}

const MaxBuyersActivesCapacity = 5

func NewProcessingVirtualQueueUseCase(
	b gateway.BuyersActivesGateway,
	v gateway.VirtualQueueGateway,
	p rabbitmq.ProducerInterface) *ProcessingVirtualQueueUseCase {
	return &ProcessingVirtualQueueUseCase{
		BuyersActivesGateway: b,
		VirtualQueueGateway:  v,
		Producer:             p,
	}
}

func (uc *ProcessingVirtualQueueUseCase) Execute(ctx context.Context) {
	totalBuyersActives, err := uc.BuyersActivesGateway.GetBuyersActives(ctx)
	if err != nil {
		fmt.Println("erro ao obter compradores ativos: %w", err)
	}

	if totalBuyersActives == MaxBuyersActivesCapacity {
		return
	}

	quantityNextBuyersActives := MaxBuyersActivesCapacity - totalBuyersActives
	uc.updateAndNotificationNextBuyersActives(int(quantityNextBuyersActives), ctx)
	uc.updateAndNotificationPositionVirtualQueue(ctx)
}

func (uc *ProcessingVirtualQueueUseCase) updateAndNotificationNextBuyersActives(quantityNextBuyersActives int, ctx context.Context) {
	for i := 0; i < quantityNextBuyersActives; i++ {
		token, err := uc.VirtualQueueGateway.Dequeue(ctx)
		if err != nil {
			fmt.Println("erro ao remover token da fila: %w", err)
			continue
		}
		uc.BuyersActivesGateway.Add(token, ctx)
		err = uc.Producer.Publish(NotificationPositionRabbitMQ{
			Token:    token,
			Position: 0,
		})
		if err != nil {
			fmt.Println("erro ao publicar no rabbitmq: %w", err)
		}
	}
}

func (uc *ProcessingVirtualQueueUseCase) updateAndNotificationPositionVirtualQueue(ctx context.Context) {
	tokensInQueue, err := uc.VirtualQueueGateway.GetAll(ctx)
	if err != nil {
		fmt.Println("erro ao obter tokens da fila: %w", err)
		return
	}

	for index, token := range tokensInQueue {
		position := index + 1
		avarageWaitTimeInSeconds := 30
		estimatedWaitTime := position * avarageWaitTimeInSeconds
		uc.Producer.Publish(NotificationPositionRabbitMQ{
			Token:             token,
			Position:          int64(position),
			EstimatedWaitTime: int64(estimatedWaitTime),
		})
	}
}
