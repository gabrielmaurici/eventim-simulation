package worker

import (
	"fmt"
	"time"

	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/processing_virtual_queue"
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
		w.ProcessingVirtualQueueUseCase.Execute()
		fmt.Println("Processado")
		time.Sleep(5 * time.Second)
	}
}
