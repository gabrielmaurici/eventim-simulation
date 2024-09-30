package entry_virtual_queue

import (
	"testing"

	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/mocks"
	"github.com/gabrielmaurici/eventim-simulation/pkg/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_WhenActiveBuyersCapacityIsAvailable_Execute_ReturnsExpectedOutput(t *testing.T) {
	token, _ := token.GenerateUniqueAccessToken()
	outputExpected := &EntryVirtualQueueOutputUseCaseDTO{
		Token:    token,
		Position: 0,
	}
	buyersActivesGatewayMock := &mocks.BuyersActivesGatewayMock{}
	virtualQueueGatewayMock := &mocks.VirtualQueueGatewayMock{}
	buyersActivesGatewayMock.On("GetBuyersActives").Return(1, nil)
	buyersActivesGatewayMock.On("Add", mock.Anything).Return(nil)
	entryVirtualQueueUsecase := NewEntryQueueUseCase(buyersActivesGatewayMock, virtualQueueGatewayMock)

	output, _ := entryVirtualQueueUsecase.Execute()

	assert.Equal(t, len(outputExpected.Token), len(output.Token))
	assert.Equal(t, outputExpected.Position, output.Position)
	buyersActivesGatewayMock.AssertNumberOfCalls(t, "GetBuyersActives", 1)
	buyersActivesGatewayMock.AssertNumberOfCalls(t, "Add", 1)
	virtualQueueGatewayMock.AssertNumberOfCalls(t, "Enqueue", 0)
}

func Test_WhenActiveBuyersCapacityIsNotAvailable_Execute_ReturnsExpectedOutput(t *testing.T) {
	token, _ := token.GenerateUniqueAccessToken()
	outputExpected := &EntryVirtualQueueOutputUseCaseDTO{
		Token:    token,
		Position: 10,
	}
	buyersActivesGatewayMock := &mocks.BuyersActivesGatewayMock{}
	virtualQueueGatewayMock := &mocks.VirtualQueueGatewayMock{}
	buyersActivesGatewayMock.On("GetBuyersActives").Return(10, nil)
	virtualQueueGatewayMock.On("Enqueue", mock.Anything).Return(outputExpected.Position, nil)
	entryVirtualQueueUsecase := NewEntryQueueUseCase(buyersActivesGatewayMock, virtualQueueGatewayMock)

	output, _ := entryVirtualQueueUsecase.Execute()

	assert.Equal(t, len(outputExpected.Token), len(output.Token))
	assert.Equal(t, outputExpected.Position, output.Position)
	buyersActivesGatewayMock.AssertNumberOfCalls(t, "GetBuyersActives", 1)
	buyersActivesGatewayMock.AssertNumberOfCalls(t, "Add", 0)
	virtualQueueGatewayMock.AssertNumberOfCalls(t, "Enqueue", 1)
}