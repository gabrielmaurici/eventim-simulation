package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/internal/usecase/processing_buyers_actives"
	"github.com/gabrielmaurici/eventim-simulation/virtual-queue/pkg/rabbitmq"
)

type ProcessingBuyersActivesWorker struct {
	Consumer                       rabbitmq.Consumer
	ProcessingBuyersActivesUseCase processing_buyers_actives.ProcessingBuyersActivesUseCase
}

func NewProcessingBuyersActivesWorker(
	c rabbitmq.Consumer,
	uc processing_buyers_actives.ProcessingBuyersActivesUseCase) *ProcessingBuyersActivesWorker {
	return &ProcessingBuyersActivesWorker{
		Consumer:                       c,
		ProcessingBuyersActivesUseCase: uc,
	}
}

func (w *ProcessingBuyersActivesWorker) Start() {
	ctx := context.Background()

	msgChan := make(chan []byte)
	go w.Consumer.Consume(msgChan)

	go func() {
		for msg := range msgChan {
			var message processing_buyers_actives.ProcessingBuyersActivesInputDTO
			if err := json.Unmarshal(msg, &message); err != nil {
				log.Printf("Erro ao deserializar mensagem: %v", err)
				continue
			}

			fmt.Println("removendo comprador ativo: " + message.UserToken)

			err := w.ProcessingBuyersActivesUseCase.Execute(message, ctx)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	select {}
}
