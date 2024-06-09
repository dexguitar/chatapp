package service

import (
	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Conns map[string]*websocket.Conn
}

func NewHub(queue Queue) *Hub {
	return &Hub{Conns: make(map[string]*websocket.Conn, 0)}
}

func (h *Hub) AddConn(username string, conn *websocket.Conn) {
	h.Conns[username] = conn
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
