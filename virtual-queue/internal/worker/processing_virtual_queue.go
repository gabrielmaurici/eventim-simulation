package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/usecase/processing_virtual_queue"
)

type ProcessingVirtualQueueWorker struct {
	ProcessingVirtualQueueUseCase processing_virtual_queue.ProcessingVirtualQueueUseCase
}

func NewProcessingVirtualQueueWorker(uc processing_virtual_queue.ProcessingVirtualQueueUseCase) *ProcessingVirtualQueueWorker {
	return &ProcessingVirtualQueueWorker{
		ProcessingVirtualQueueUseCase: uc,
	}
}

func (w *ProcessingVirtualQueueWorker) Start() {
	for {
		ctx := context.Background()
		w.ProcessingVirtualQueueUseCase.Execute(ctx)
		fmt.Println("Fila virtual processada")
		time.Sleep(5 * time.Second)
	}
}
