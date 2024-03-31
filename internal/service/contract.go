package service

import (
	"context"

	"github.com/dexguitar/chatapp/internal/model"
	pg "github.com/snaffi/pg-helper"
)

type Repo interface {
	CreateUser(ctx context.Context, user *model.User, db pg.DB) error
	FindUserByCreds(ctx context.Context, username, password string, db pg.Read) (*model.User, error)
	FindUserById(ctx context.Context, id string, db pg.Read) (*model.User, error)
}
