package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/buy_tickets"
	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/reserve_ticket"
)

type WebTicketsHandler struct {
	ReserverTicketsUseCase reserve_ticket.ReserveTicketUseCase
	BuyTicketsUseCase      buy_tickets.BuyTicketsUseCase
}

func NewWebTicketsHandler(
	ucr reserve_ticket.ReserveTicketUseCase,
	ucb buy_tickets.BuyTicketsUseCase) *WebTicketsHandler {
	return &WebTicketsHandler{
		ReserverTicketsUseCase: ucr,
		BuyTicketsUseCase:      ucb,
	}
}

func (h *WebTicketsHandler) Reserve(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var input reserve_ticket.ReserveTicketInputUseCaseDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.ReserverTicketsUseCase.Execute(input, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *WebTicketsHandler) Purchase(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var input buy_tickets.BuyTicketsInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("erro")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("passsou " + input.UserToken)

	output, err := h.BuyTicketsUseCase.Execute(input, ctx)
	if err != nil {
		fmt.Println("erro 2 " + err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(output)
}
