package service

import (
	"sync"

	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Conns map[string]*websocket.Conn
	mu    sync.Mutex
}

func NewHub() *Hub {
	return &Hub{Conns: make(map[string]*websocket.Conn)}
}

func (h *Hub) AddConn(username string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Conns[username] = conn
}

func (h *Hub) GetConns() map[string]*websocket.Conn {
	return h.Conns
}

func (h *Hub) RemoveConn(username string) {
	delete(h.Conns, username)
}

func (h *Hub) Broadcast(m *sarama.ConsumerMessage) {
	for username, conn := range h.Conns {
		if username == getHeaderValue(m, "receiver") {
			conn.WriteMessage(websocket.TextMessage, m.Value)
		}
	}
}

func getHeaderValue(m *sarama.ConsumerMessage, key string) string {
	for _, header := range m.Headers {
		if string(header.Key) == key {
			return string(header.Value)
		}
	}
	return ""
}
