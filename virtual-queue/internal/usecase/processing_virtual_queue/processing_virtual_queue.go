package processing_virtual_queue

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/internal/gateway"
)

type ProcessingVirtualQueueUseCase struct {
	BuyersActivesGateway gateway.BuyersActivesGateway
	VirtualQueueGateway  gateway.VirtualQueueGateway
}

const MaxBuyersActivesCapacity = 5

func NewProcessingVirtualQueueUseCase(
	b gateway.BuyersActivesGateway,
	v gateway.VirtualQueueGateway) *ProcessingVirtualQueueUseCase {
	return &ProcessingVirtualQueueUseCase{
		BuyersActivesGateway: b,
		VirtualQueueGateway:  v,
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

		fmt.Println("Usuário liberado para compra: " + token)
	}
}

func (uc *ProcessingVirtualQueueUseCase) updateAndNotificationPositionVirtualQueue() {
	tokensInQueue, err := uc.VirtualQueueGateway.GetAll()
	if err != nil {
		fmt.Println("erro ao obter tokens da fila: %w", err)
		return
	}

	for position, token := range tokensInQueue {

		message := fmt.Sprintf("Posição do usuário %s atualizada: %d", token, position+1)
		fmt.Println(message)
	}
}
