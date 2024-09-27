package service

import (
	"context"

	"github.com/dexguitar/chatapp/internal/model"
	pg "github.com/snaffi/pg-helper"
)

type Repo interface {
	CreateUser(ctx context.Context, db pg.DB, user *model.User) (*model.User, error)
	FindUserByUsername(ctx context.Context, db pg.Read, username string) (*model.User, error)
	FindUserById(ctx context.Context, db pg.Read, id string) (*model.User, error)
}

type Queue interface {
	WriteMessage(ctx context.Context, message *model.Message) error
}
