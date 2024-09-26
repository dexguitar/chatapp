package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dexguitar/chatapp/internal/errs"
	"github.com/dexguitar/chatapp/internal/model"
	"github.com/pkg/errors"
	pg "github.com/snaffi/pg-helper"
)

const activeToken = "active_token"

type UserService struct {
	Repo
	connPool pg.DB
}

func NewUserService(repo Repo, connPool pg.DB) *UserService {
	return &UserService{repo, connPool}
}

func (u *UserService) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	op := "UserService.RegisterUser"

	result, err := u.Repo.FindUserByUsername(ctx, u.connPool.Replica(), user.Username)
	if result != nil {
		return nil, fmt.Errorf("%s: %w", op, errs.ErrUserExists)
	}
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return nil, errors.Wrap(err, op)
	}

	res, err := u.Repo.CreateUser(ctx, u.connPool, user)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return &model.User{
		Username: res.Username,
		Email:    res.Email,
	}, nil
}

func (u *UserService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	op := "UserService.GetUserById"

	user, err := u.Repo.FindUserById(ctx, u.connPool.Replica(), id)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return user, nil
}

func (u *UserService) Login(ctx context.Context, userInput *model.User) (string, error) {
	op := "UserService.Login"

	res, err := u.Repo.FindUserByUsername(ctx, u.connPool.Replica(), userInput.Username)
	if res == nil {
		return "", fmt.Errorf("%s: %w", op, errs.ErrNotFound)
	}
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return "", errors.Wrap(err, op)
	}

	if res.Password != userInput.Password {
		return "", errs.NewCustomError(errs.ErrInvalidCreds.Error(), http.StatusUnauthorized, errs.ErrInvalidCreds)
	}

	return activeToken, nil
}
