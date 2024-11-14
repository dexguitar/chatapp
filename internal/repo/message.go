package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/dexguitar/chatapp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
	pg "github.com/snaffi/pg-helper"
)

type MessageRepository struct{}

func NewMessageRepo() *MessageRepository {
	return &MessageRepository{}
}

func (m *MessageRepository) StoreMessage(ctx context.Context, db pg.DB, message *model.Message) error {
	op := "MessageRepository.StoreMessage"

	_, err := pgxutil.Insert(ctx, db, pgx.Identifier{"messages"}, []map[string]any{
		{"timestamp": time.Now(), "sender_id": message.Username, "receiver_id": message.Receiver, "content": message.Value},
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
