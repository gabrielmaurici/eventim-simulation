package entry_virtual_queue

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/gateway"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/pkg/token"
)

type EntryVirtualQueueOutputUseCaseDTO struct {
	Token    string `json:"token"`
	Position int64  `json:"position"`
}

type EntryVirtualQueueUseCase struct {
	BuyersActivesGateway gateway.BuyersActivesGateway
	VirtualQueueGateway  gateway.VirtualQueueGateway
}

const MaxBuyersActivesCapacity = 5

func NewEntryQueueUseCase(b gateway.BuyersActivesGateway, v gateway.VirtualQueueGateway) *EntryVirtualQueueUseCase {
	return &EntryVirtualQueueUseCase{
		BuyersActivesGateway: b,
		VirtualQueueGateway:  v,
	}
}

func (uc *EntryVirtualQueueUseCase) Execute() (output *EntryVirtualQueueOutputUseCaseDTO, err error) {
	token, err := token.GenerateUniqueAccessToken()
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar token de acesso único: %w", err)
	}

	totalBuyersActives, err := uc.BuyersActivesGateway.GetBuyersActives()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter compradores ativos: %w", err)
	}

	var position int64
	if totalBuyersActives < MaxBuyersActivesCapacity {
		err = uc.BuyersActivesGateway.Add(token)
		if err != nil {
			return nil, fmt.Errorf("erro ao adicionar usuário ao grupo de compradores ativos: %w", err)
		}

		return &EntryVirtualQueueOutputUseCaseDTO{
			Token:    token,
			Position: 0,
		}, nil
	}

	position, err = uc.VirtualQueueGateway.Enqueue(token)
	if err != nil {
		return nil, fmt.Errorf("erro ao adicionar usuário a fila de espera: %w", err)
	}

	return &EntryVirtualQueueOutputUseCaseDTO{
		Token:    token,
		Position: position,
	}, nil
}
