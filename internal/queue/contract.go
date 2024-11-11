package queue

import (
	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

type Hub interface {
	AddConn(username string, conn *websocket.Conn)
	GetConns() map[string]*websocket.Conn
	RemoveConn(username string)
	Broadcast(m *sarama.ConsumerMessage)
}
