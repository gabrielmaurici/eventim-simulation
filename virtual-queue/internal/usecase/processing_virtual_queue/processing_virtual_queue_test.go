package processing_virtual_queue

import (
	"context"
	"testing"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_WhenActiveBuyersCapacityIsAvailable_Execute_NotifyWebsocket(t *testing.T) {
	buyersActivesMock := &mocks.BuyersActivesGatewayMock{}
	virtualQueueMock := &mocks.VirtualQueueGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	buyersActivesMock.On("GetBuyersActives", mock.Anything).Return(2, nil)
	buyersActivesMock.On("Add", "some-token", mock.Anything).Return(nil)
	virtualQueueMock.On("Dequeue", mock.Anything).Return("some-token", nil)
	virtualQueueMock.On("GetAll", mock.Anything).Return([]string{"some-token1", "some-token-2"}, nil)
	producerMock.On("Publish", mock.Anything).Return(nil)
	processingVirtualQueueUseCase := NewProcessingVirtualQueueUseCase(
		buyersActivesMock,
		virtualQueueMock,
		producerMock,
	)

	processingVirtualQueueUseCase.Execute(context.Background())

	assert.NoError(t, nil)
	buyersActivesMock.AssertNumberOfCalls(t, "GetBuyersActives", 1)
	buyersActivesMock.AssertNumberOfCalls(t, "Add", 3)
	virtualQueueMock.AssertNumberOfCalls(t, "Dequeue", 3)
	virtualQueueMock.AssertNumberOfCalls(t, "GetAll", 1)
	producerMock.AssertNumberOfCalls(t, "Publish", 5)
}

func Test_WhenActiveBuyersCapacityIsNotAvailable_Execute_DontNotifyWebsocket(t *testing.T) {
	buyersActivesMock := &mocks.BuyersActivesGatewayMock{}
	virtualQueueMock := &mocks.VirtualQueueGatewayMock{}
	producerMock := &mocks.ProducerMock{}
	buyersActivesMock.On("GetBuyersActives", mock.Anything).Return(5, nil)
	processingVirtualQueueUseCase := NewProcessingVirtualQueueUseCase(
		buyersActivesMock,
		virtualQueueMock,
		producerMock,
	)

	processingVirtualQueueUseCase.Execute(context.Background())

	assert.NoError(t, nil)
	buyersActivesMock.AssertNumberOfCalls(t, "GetBuyersActives", 1)
	buyersActivesMock.AssertNumberOfCalls(t, "Add", 0)
	virtualQueueMock.AssertNumberOfCalls(t, "Dequeue", 0)
	virtualQueueMock.AssertNumberOfCalls(t, "GetAll", 0)
	producerMock.AssertNumberOfCalls(t, "Publish", 0)
}
