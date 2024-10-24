package web

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/usecase/reserve_ticket"
)

type WebTicketsReservationHandler struct {
	ReserverTicketsUseCase reserve_ticket.ReserveTicketUseCase
}

func NewWebTicketsReservationHandler(uc reserve_ticket.ReserveTicketUseCase) *WebTicketsReservationHandler {
	return &WebTicketsReservationHandler{
		ReserverTicketsUseCase: uc,
	}
}

func (h *WebTicketsReservationHandler) Reserve(w http.ResponseWriter, r *http.Request) {
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

// TODO
func (h *WebTicketsReservationHandler) Purchase(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	// var input reserve_ticket.ReserveTicketInputUseCaseDTO
	// err := json.NewDecoder(r.Body).Decode(&input)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// err = h.ReserverTicketsUseCase.Execute(input, ctx)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
}
