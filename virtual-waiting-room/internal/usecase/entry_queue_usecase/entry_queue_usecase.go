package entry_queue_usecase

import (
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/internal/gateway"
	"github.com/gabrielmaurici/eventim-simulation/pkg/token"
)

type EntryQueueOutputUseCaseDTO struct {
	Token    string `json:"token"`
	Position int64  `json:"position"`
}

type EntryQueueUseCase struct {
	BuyersActivesGateway gateway.BuyersActivesGateway
	VirtualQueueGateway  gateway.VirtualQueueGateway
}

const MaxBuyersActivesCapacity = 5

func NewEntryQueueUseCase(bg gateway.BuyersActivesGateway, vq gateway.VirtualQueueGateway) *EntryQueueUseCase {
	return &EntryQueueUseCase{
		BuyersActivesGateway: bg,
		VirtualQueueGateway:  vq,
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

	var position int64
	if buyersActivesTotal < MaxBuyersActivesCapacity {
		err = uc.BuyersActivesGateway.Add(token)
		if err != nil {
			return nil, fmt.Errorf("erro ao adicionar usuário ao grupo de compradores ativos: %w", err)
		}
	} else {
		position, err = uc.VirtualQueueGateway.Enqueue(token)
		if err != nil {
			return nil, fmt.Errorf("erro ao adicionar usuário a fila de espera: %w", err)
		}
	}

	return &EntryQueueOutputUseCaseDTO{
		Token:    token,
		Position: position,
	}, nil
}
