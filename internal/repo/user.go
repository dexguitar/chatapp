package repo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dexguitar/chatapp/internal/model"
	"github.com/dexguitar/chatapp/internal/utils"
	"github.com/jackc/pgx/v5"
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
		return utils.NewCustomError("error creating user", http.StatusInternalServerError, fmt.Errorf("%s: %w", op, err))
	}

	return nil
}

func (u *UserRepository) FindUserByUsername(ctx context.Context, username string, db pg.Read) (*model.User, error) {
	op := "UserRepository.FindUserByUsername"

	var user model.User

	q := "SELECT id, username, email, password FROM users WHERE username = $1"

	if err := db.QueryRow(
		ctx,
		q,
		username,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, utils.NewCustomError(utils.ErrUserNotFound.Error(), http.StatusNotFound, utils.ErrUserNotFound)
		}

		return nil, utils.NewCustomError("internal server error", http.StatusInternalServerError, fmt.Errorf("%s: %w", op, err))
	}

	return &user, nil
}

func (u *UserRepository) FindUserById(ctx context.Context, id string, db pg.Read) (*model.User, error) {
	op := "UserRepository.FindUserById"

	var user model.User
	q := "SELECT id, username, email, password FROM users WHERE id = $1"

	if err := db.QueryRow(ctx, q, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, utils.NewCustomError("user not found", http.StatusNotFound, fmt.Errorf("%s: %w", op, err))
		}

		return nil, utils.NewCustomError("internal server error", http.StatusInternalServerError, fmt.Errorf("%s: %w", op, err))
	}

	return &user, nil
}
