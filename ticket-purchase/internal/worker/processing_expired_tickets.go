package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/processing_expired_tickets"
)

type ProcessingExpiredTicketsWorker struct {
	ProcessingExpiredTicketsUseCase processing_expired_tickets.ProcessingExpiredTicketsUseCase
}

func NewProcessingExpiredTicketsWorker(uc processing_expired_tickets.ProcessingExpiredTicketsUseCase) *ProcessingExpiredTicketsWorker {
	return &ProcessingExpiredTicketsWorker{
		ProcessingExpiredTicketsUseCase: uc,
	}
}

func (w *ProcessingExpiredTicketsWorker) Start() {
	for {
		ctx := context.Background()
		err := w.ProcessingExpiredTicketsUseCase.Execute(ctx)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("ingressos expirados processados")
		time.Sleep(5 * time.Second)
	}
}
