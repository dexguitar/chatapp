package service

import (
	"context"

	"github.com/dexguitar/chatapp/internal/model"
	"github.com/pkg/errors"
	pg "github.com/snaffi/pg-helper"
)

type MessageService struct {
	MessageRepo
	connPool pg.DB
	queue    Queue
}

func NewMessageService(repo MessageRepo, connPool pg.DB, queue Queue) *MessageService {
	return &MessageService{MessageRepo: repo, connPool: connPool, queue: queue}
}

func (ms *MessageService) SendMessage(ctx context.Context, message *model.Message) error {
	op := "MessageService.SendMessage"

	err := ms.MessageRepo.StoreMessage(ctx, ms.connPool, message)
	if err != nil {
		return errors.Wrap(err, op)
	}

	err = ms.queue.WriteMessage(ctx, message)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
