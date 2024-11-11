package repo

import (
	"context"
	"fmt"

	"github.com/dexguitar/chatapp/internal/errs"
	"github.com/dexguitar/chatapp/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
	pg "github.com/snaffi/pg-helper"
)

type UserRepository struct{}

func NewUserRepo() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(ctx context.Context, db pg.DB, user *model.User) (*model.User, error) {
	op := "UserRepository.CreateUser"

	_, err := pgxutil.Insert(ctx, db, pgx.Identifier{"users"}, []map[string]any{
		{"username": user.Username, "password": user.Password, "email": user.Email},
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &model.User{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (u *UserRepository) FindUserByUsername(ctx context.Context, db pg.Read, username string) (*model.User, error) {
	op := "UserRepository.FindUserByUsername"

	q := `select id, username, email, password from users where username = $1`
	user, err := pgxutil.SelectRow(ctx, db, q, []any{username}, pgx.RowToAddrOfStructByPos[model.User])
	if err != nil {
		return handleError(err, op)
	}

	return user, nil
}

func (u *UserRepository) FindUserById(ctx context.Context, db pg.Read, id string) (*model.User, error) {
	op := "UserRepository.FindUserById"

	q := `select id, username, email, password from users where id = $1`
	user, err := pgxutil.SelectRow(ctx, db, q, []any{id}, pgx.RowToAddrOfStructByPos[model.User])
	if err != nil {
		return handleError(err, op)
	}

	return user, nil
}

func handleError(e error, op string) (*model.User, error) {
	if e == pgx.ErrNoRows {
		return nil, fmt.Errorf("%s: %w", op, errs.ErrUserNotFound)
	}

	return nil, fmt.Errorf("%s: %w", op, e)
}
