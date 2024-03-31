package handler

import (
	"context"

	"github.com/dexguitar/chatapp/internal/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error)
	Login(ctx context.Context, req *model.LoginReq) (*model.LoginRes, error)
	GetUserById(ctx context.Context, req *model.GetUserByIdReq) (*model.GetUserByIdRes, error)
}

type Validator interface {
	Validate() error
}
