package service

import (
	"context"

	"github.com/dexguitar/chatapp/internal/model"
)

type MessageService struct {
	Queue Queue
}

func NewMessageService(queue Queue) *MessageService {
	return &MessageService{Queue: queue}
}

func (ms *MessageService) SendMessage(ctx context.Context, message *model.Message) error {
	// DB write here
	// .............

	return ms.Queue.WriteMessage(ctx, message)
}
