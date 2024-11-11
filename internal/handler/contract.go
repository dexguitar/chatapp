package handler

import (
	"context"

	"github.com/dexguitar/chatapp/internal/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, user *model.User) (string, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
}

type MessageService interface {
	SendMessage(ctx context.Context, message *model.Message) error
}

type Validator interface {
	Validate() error
}
