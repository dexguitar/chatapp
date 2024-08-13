package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dexguitar/chatapp/internal/handler"
	"github.com/dexguitar/chatapp/internal/model"
	"github.com/dexguitar/chatapp/internal/utils"
	"github.com/pkg/errors"
	pg "github.com/snaffi/pg-helper"
)

type UserService struct {
	Repo
	connPool pg.DB
}

func NewUserService(repo Repo, connPool pg.DB) *UserService {
	return &UserService{repo, connPool}
}

func (u *UserService) RegisterUser(ctx context.Context, req *handler.Request[handler.CreateUserReq]) (*handler.Response[*handler.CreateUserRes], error) {
	op := "UserService.RegisterUser"

	user := &model.User{
		Username: req.Body.Username,
		Email:    req.Body.Email,
		Password: req.Body.Password,
	}

	result, err := u.Repo.FindUserByUsername(ctx, user.Username, u.connPool.Replica())
	if result != nil {
		return nil, utils.NewCustomError("user already exists", http.StatusConflict, fmt.Errorf("%s: %w", op, errors.New("attempt to create existing user")))
	}

	var customError utils.CustomError
	if errors.As(err, &customError) {
		if customError.Err != utils.ErrUserNotFound {
			return nil, errors.Wrap(err, op)
		}
	}

	err = u.Repo.CreateUser(ctx, user, u.connPool)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return &handler.Response[*handler.CreateUserRes]{
		Body: &handler.CreateUserRes{
			Username: user.Username,
			Email:    user.Email,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *handler.Request[handler.GetUserByIdReq]) (*handler.Response[*handler.GetUserByIdRes], error) {
	op := "UserService.GetUserById"

	user, err := u.Repo.FindUserById(ctx, req.Body.ID, u.connPool.Replica())
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return &handler.Response[*handler.GetUserByIdRes]{
		Body: &handler.GetUserByIdRes{
			Username: user.Username,
			Email:    user.Email,
		},
		StatusCode: http.StatusOK,
	}, nil
}

func (u *UserService) Login(ctx context.Context, req *handler.Request[handler.LoginReq]) (*handler.Response[*handler.LoginRes], error) {
	op := "UserService.Login"

	user, err := u.Repo.FindUserByUsername(ctx, req.Body.Username, u.connPool.Replica())
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if user.Password != req.Body.Password {
		return nil, utils.NewCustomError(utils.ErrInvalidCreds.Error(), http.StatusUnauthorized, utils.ErrInvalidCreds)
	}

	return &handler.Response[*handler.LoginRes]{
		Body: &handler.LoginRes{
			Token: "active_token",
		},
		StatusCode: http.StatusOK,
	}, nil
}
