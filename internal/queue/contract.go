package queue

import (
	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

type Hub interface {
	AddConn(username string, conn *websocket.Conn)
	Broadcast(m *sarama.ConsumerMessage)
}
