package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type NotificationRabbitMQModel struct {
	Token    string `json:"token"`
	Position int64  `json:"position"`
}

type WebSocketVirtualQueueHandler struct {
	upgrader    websocket.Upgrader
	msgChan     chan []byte
	connections map[string]*websocket.Conn
	mu          sync.Mutex
}

func NewWebSocketVirtualQueueHandler(msgChan chan []byte) *WebSocketVirtualQueueHandler {
	return &WebSocketVirtualQueueHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		msgChan:     msgChan,
		connections: make(map[string]*websocket.Conn),
	}
}

func (h *WebSocketVirtualQueueHandler) NotifyPositionSocket(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token não fornecido", http.StatusBadRequest)
		return
	}

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Erro ao upgrade para WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	h.mu.Lock()
	h.connections[token] = conn
	h.mu.Unlock()

	go func() {
		defer func() {
			h.mu.Lock()
			delete(h.connections, token) // Remover conexão ao finalizar
			h.mu.Unlock()
		}()

		for msg := range h.msgChan {
			fmt.Println("chegou aqui")
			var message NotificationRabbitMQModel
			if err := json.Unmarshal(msg, &message); err != nil {
				log.Printf("Erro ao deserializar mensagem: %v", err)
				continue
			}

			responseMsg, err := json.Marshal(message)
			if err != nil {
				log.Printf("Erro ao serializar a resposta: %v", err)
				continue
			}

			h.mu.Lock()
			conn, exists := h.connections[message.Token]
			h.mu.Unlock()

			if exists {
				if err := conn.WriteMessage(websocket.TextMessage, responseMsg); err != nil {
					log.Printf("Erro ao enviar mensagem: %v", err)
					conn.Close()
					h.mu.Lock()
					delete(h.connections, message.Token)
					h.mu.Unlock()
				}
			}
		}
	}()

	// Loop para ler mensagens do cliente (opcional)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break // Sai do loop se houver erro
		}
	}
}
