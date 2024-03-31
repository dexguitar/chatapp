package service

import (
	"context"
	"fmt"

	"github.com/dexguitar/chatapp/internal/model"
	pg "github.com/snaffi/pg-helper"
)

type UserService struct {
	Repo
	connPool pg.DB
}

func NewUserService(repo Repo, connPool pg.DB) *UserService {
	return &UserService{repo, connPool}
}

func (u *UserService) RegisterUser(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error) {
	op := "UserService.RegisterUser"

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := u.Repo.CreateUser(ctx, user, u.connPool)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &model.CreateUserRes{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (u *UserService) Login(ctx context.Context, req *model.LoginReq) (*model.LoginRes, error) {
	op := "UserService.Login"

	if _, err := u.Repo.FindUserByCreds(ctx, req.Username, req.Password, u.connPool.Replica()); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &model.LoginRes{
		Token: "active_token",
	}, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *model.GetUserByIdReq) (*model.GetUserByIdRes, error) {
	op := "UserService.GetUserById"

	user, err := u.Repo.FindUserById(ctx, req.ID, u.connPool.Replica())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &model.GetUserByIdRes{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
