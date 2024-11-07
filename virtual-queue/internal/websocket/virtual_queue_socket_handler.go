package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type NotificationRabbitMQModel struct {
	Token             string `json:"token"`
	Position          int64  `json:"position"`
	EstimatedWaitTime int64  `json:"estimated_wait_time"`
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

func (s *WebSocketVirtualQueueHandler) NotifyPositionSocket(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token não fornecido", http.StatusBadRequest)
		return
	}

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Erro ao usar upgrade para WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	s.mu.Lock()
	s.connections[token] = conn
	s.mu.Unlock()

	go func() {
		defer func() {
			s.mu.Lock()
			delete(s.connections, token)
			s.mu.Unlock()
		}()

		for msg := range s.msgChan {
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

			s.mu.Lock()
			conn, exists := s.connections[message.Token]
			s.mu.Unlock()

			if exists {
				if err := conn.WriteMessage(websocket.TextMessage, responseMsg); err != nil {
					log.Printf("Erro ao enviar mensagem: %v", err)
					conn.Close()
					s.mu.Lock()
					delete(s.connections, message.Token)
					s.mu.Unlock()
				}
			}
		}
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
