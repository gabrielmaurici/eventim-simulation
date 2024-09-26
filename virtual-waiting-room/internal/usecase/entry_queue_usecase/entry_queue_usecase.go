package entry_queue_usecase

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/internal/gateway"
	"github.com/gabrielmaurici/eventim-simulation/pkg/token"
)

type EntryQueueOutputUseCaseDTO struct {
	Token    string `json:"token"`
	Position int    `json:"position"`
}

type EntryQueueUseCase struct {
	BuyersActivesGateway    gateway.BuyersActivesGateway
	QueueWaitingRoomGateway gateway.QueueWaitingRoomGateway
}

const MaxBuyersActivesCapacity = 5

func NewEntryQueueUseCase(bg gateway.BuyersActivesGateway, qg gateway.QueueWaitingRoomGateway) *EntryQueueUseCase {
	return &EntryQueueUseCase{
		BuyersActivesGateway:    bg,
		QueueWaitingRoomGateway: qg,
	}
}

func (uc *EntryQueueUseCase) Execute() (output *EntryQueueOutputUseCaseDTO, err error) {
	token, err := token.GenerateUniqueAccessToken()
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token de acesso único: %w", err)
	}

	buyersActivesTotal, err := uc.BuyersActivesGateway.GetBuyersActives()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter compradores ativos: %w", err)
	}

	var position int
	if buyersActivesTotal < MaxBuyersActivesCapacity {
		err = uc.BuyersActivesGateway.Add(token)
		if err != nil {
			return nil, fmt.Errorf("erro ao adicionar usuário ao grupo de compradores ativos: %w", err)
		}
	} else {
		position, err = uc.QueueWaitingRoomGateway.Enqueue(token)
		if err != nil {
			return nil, fmt.Errorf("erro ao adicionar usuário a fila de espera: %w", err)
		}
	}

	return &EntryQueueOutputUseCaseDTO{
		Token:    token,
		Position: position,
	}, nil
}
