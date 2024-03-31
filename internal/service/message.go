package service

import (
	"context"
	"fmt"

	"github.com/dexguitar/chatapp/internal/model"
)

type MessageService struct {
	Queue
}

func NewMessageService(queue Queue) *MessageService {
	return &MessageService{queue}
}

func (m *MessageService) SendMessage(ctx context.Context, req *model.SendMessageReq) (*model.SendMessageRes, error) {
	op := "MessageService.SendMessage"

	err := m.Queue.WriteMessage(ctx, &model.Message{
		Username: req.Username,
		Value:    req.Message,
		Receiver: req.Receiver,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &model.SendMessageRes{}, nil
}
