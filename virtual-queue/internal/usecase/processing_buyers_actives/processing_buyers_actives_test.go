package processing_buyers_actives

import (
	"context"
	"errors"
	"testing"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_WhenConsumerRecevedOneUserToken_Execute_DeleteBuyerActive(t *testing.T) {
	input := ProcessingBuyersActivesInputDTO{
		UserToken: "some_token",
	}
	ctx := context.Background()
	buyersActivesMock := &mocks.BuyersActivesGatewayMock{}
	buyersActivesMock.On("Delete", input.UserToken, ctx).Return(nil)
	processingBuyersActivesUseCase := NewProcessingBuyersActivesUseCase(
		buyersActivesMock,
	)

	err := processingBuyersActivesUseCase.Execute(input, ctx)

	assert.NoError(t, err)
	buyersActivesMock.AssertNumberOfCalls(t, "Delete", 1)
}

func Test_WhenConsumerRecevedOneUserTokenEmpty_Execute_ReturnsError(t *testing.T) {
	input := ProcessingBuyersActivesInputDTO{
		UserToken: "",
	}
	ctx := context.Background()
	buyersActivesMock := &mocks.BuyersActivesGatewayMock{}
	processingBuyersActivesUseCase := NewProcessingBuyersActivesUseCase(
		buyersActivesMock,
	)

	err := processingBuyersActivesUseCase.Execute(input, ctx)

	assert.Error(t, err)
	assert.Equal(t, "token do usuário é obrigatório", err.Error())
	buyersActivesMock.AssertNumberOfCalls(t, "Delete", 0)
}

func Test_WhenItGivesAnErrorWhenDeleting_Execute_ReturnsError(t *testing.T) {
	input := ProcessingBuyersActivesInputDTO{
		UserToken: "some_token",
	}
	ctx := context.Background()
	buyersActivesMock := &mocks.BuyersActivesGatewayMock{}
	buyersActivesMock.On("Delete", input.UserToken, ctx).Return(errors.New("teste"))
	processingBuyersActivesUseCase := NewProcessingBuyersActivesUseCase(
		buyersActivesMock,
	)

	err := processingBuyersActivesUseCase.Execute(input, ctx)

	assert.Error(t, err)
	assert.Equal(t, "erro ao deletar comprador ativo: teste", err.Error())
	buyersActivesMock.AssertNumberOfCalls(t, "Delete", 1)
}
