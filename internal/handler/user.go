package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dexguitar/chatapp/internal/errs"
	"github.com/dexguitar/chatapp/internal/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type UserHandler struct {
	UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) RegisterUser(ctx context.Context, req *Request[CreateUserReq]) (*Response[*CreateUserRes], error) {
	op := "UserHandler.RegisterUser"

	res, err := uh.UserService.RegisterUser(ctx, &model.User{
		Username: req.Body.Username,
		Email:    req.Body.Email,
		Password: req.Body.Password,
	})
	if err != nil {
		if errors.Is(err, errs.ErrUserExists) {
			return nil, errs.NewCustomError(errs.ErrUserExists.Error(), http.StatusConflict, fmt.Errorf("%s: %w", op, err))
		}

		return nil, err
	}

	return &Response[*CreateUserRes]{
		Body: &CreateUserRes{
			Username: res.Username,
			Email:    res.Email,
		},
		StatusCode: http.StatusCreated,
	}, nil
}

func (uh *UserHandler) GetUserById(ctx context.Context, req *Request[GetUserByIdReq]) (*Response[*GetUserByIdRes], error) {
	op := "UserHandler.GetUserById"

	user, err := uh.UserService.GetUserById(ctx, req.Params.ID)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return nil, errs.NewCustomError(errs.ErrUserNotFound.Error(), http.StatusNotFound, err)
		}
		return nil, errs.NewCustomError(errs.ErrInternal.Error(), http.StatusInternalServerError, fmt.Errorf("%s: %w", op, err))
	}

	return &Response[*GetUserByIdRes]{
		Body: &GetUserByIdRes{
			Username: user.Username,
			Email:    user.Email,
		},
		StatusCode: http.StatusOK,
	}, nil
}

func (uh *UserHandler) Login(ctx context.Context, req *Request[LoginReq]) (*Response[*LoginRes], error) {
	op := "UserHandler.GetUserById"

	u := &model.User{
		Username: req.Body.Username,
		Password: req.Body.Password,
	}

	token, err := uh.UserService.Login(ctx, u)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return &Response[*LoginRes]{}, errs.NewCustomError(
				errs.ErrUserNotFound.Error(), http.StatusNotFound, fmt.Errorf("%s: %w", op, err),
			)
		}

		return &Response[*LoginRes]{}, errs.NewCustomError(
			errs.ErrInvalidCreds.Error(), http.StatusBadRequest, fmt.Errorf("%s: %w", op, err),
		)
	}

	if token == "" {
		return nil, errs.NewCustomError(
			errs.ErrInvalidCreds.Error(), http.StatusUnauthorized, fmt.Errorf("%s: %w", op, errs.ErrInvalidCreds),
		)
	}

	return &Response[*LoginRes]{
		Body: &LoginRes{
			Token: token,
		},
		StatusCode: http.StatusOK,
	}, nil
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type GetUserByIdReq struct {
	ID string `json:"id"`
}

type GetUserByIdRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r CreateUserReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

func (r LoginReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

func (r GetUserByIdReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.ID, validation.Required),
	)
}
