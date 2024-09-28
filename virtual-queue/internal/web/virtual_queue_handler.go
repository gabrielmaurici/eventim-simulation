package web

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/entry_queue"
)

type WebVirtualQueueHandler struct {
	EntryQueueUseCase entry_queue.EntryQueueUseCase
}

func NewWebVirtualQueueHandler(uc entry_queue.EntryQueueUseCase) *WebVirtualQueueHandler {
	return &WebVirtualQueueHandler{
		EntryQueueUseCase: uc,
	}
}

func (h *WebVirtualQueueHandler) EntryQueue(w http.ResponseWriter, r *http.Request) {
	output, err := h.EntryQueueUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao retornar resposta: " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
