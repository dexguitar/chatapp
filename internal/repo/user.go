package repo

import (
	"context"
	"fmt"

	"github.com/dexguitar/chatapp/internal/model"
	pg "github.com/snaffi/pg-helper"
)

type UserRepository struct{}

func NewUserRepo() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *model.User, db pg.DB) error {
	op := "UserRepository.CreateUser"

	q := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"

	_, err := db.Exec(ctx, q, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (u *UserRepository) FindUserByCreds(ctx context.Context, username, password string, db pg.Read) (*model.User, error) {
	op := "UserRepository.FindUserByCreds"

	var user model.User

	q := "SELECT * FROM users WHERE username = $1 AND password = $2"

	if err := db.QueryRow(
		ctx,
		q,
		username,
		password,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (u *UserRepository) FindUserById(ctx context.Context, id string, db pg.Read) (*model.User, error) {
	op := "UserRepository.FindUserById"

	var user model.User

	q := "SELECT * FROM users WHERE id = $1"

	if err := db.QueryRow(ctx, q, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
