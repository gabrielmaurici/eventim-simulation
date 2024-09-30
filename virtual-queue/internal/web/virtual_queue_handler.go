package web

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielmaurici/eventim-simulation/internal/usecase/entry_virtual_queue"
)

type WebVirtualQueueHandler struct {
	EntryVirtualQueueUseCase entry_virtual_queue.EntryVirtualQueueUseCase
}

func NewWebVirtualQueueHandler(uc entry_virtual_queue.EntryVirtualQueueUseCase) *WebVirtualQueueHandler {
	return &WebVirtualQueueHandler{
		EntryVirtualQueueUseCase: uc,
	}
}

func (h *WebVirtualQueueHandler) EntryQueue(w http.ResponseWriter, r *http.Request) {
	output, err := h.EntryVirtualQueueUseCase.Execute()
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