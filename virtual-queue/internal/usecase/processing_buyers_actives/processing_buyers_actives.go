package processing_buyers_actives

import (
	"context"
	"errors"
	"fmt"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/gateway"
)

type ProcessingBuyersActivesInputDTO struct {
	UserToken string `json:"user_token"`
}

type ProcessingBuyersActivesUseCase struct {
	BuyersActivesGateway gateway.BuyersActivesGateway
}

func NewProcessingBuyersActivesUseCase(bg gateway.BuyersActivesGateway) *ProcessingBuyersActivesUseCase {
	return &ProcessingBuyersActivesUseCase{
		BuyersActivesGateway: bg,
	}
}

func (uc *ProcessingBuyersActivesUseCase) Execute(input ProcessingBuyersActivesInputDTO, ctx context.Context) error {
	if input.UserToken == "" {
		return errors.New("token do usuário é obrigatório")
	}

	err := uc.BuyersActivesGateway.Delete(input.UserToken, ctx)
	if err != nil {
		return fmt.Errorf("erro ao deletar comprador ativo: %w", err)
	}

	return nil
}
