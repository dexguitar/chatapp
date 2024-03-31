package service

import (
	"context"

	pg "github.com/snaffi/pg-helper"
	"github.com/dexguitar/chatapp/internal/model"
)

type Repo interface {
	CreateUser(ctx context.Context, user *model.User, db pg.DB) error
	FindUserByCreds(ctx context.Context, username, password string, db pg.Read) (*model.User, error)
	FindUserById(ctx context.Context, id string, db pg.Read) (*model.User, error)
}

type Queue interface {
	WriteMessage(ctx context.Context, message *model.Message) error
}
