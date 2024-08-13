package handler

import (
	"context"

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

	res, err := uh.UserService.RegisterUser(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

func (uh *UserHandler) GetUserById(ctx context.Context, req *Request[GetUserByIdReq]) (*Response[*GetUserByIdRes], error) {
	op := "UserHandler.GetUserById"

	res, err := uh.UserService.GetUserById(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
}

func (uh *UserHandler) Login(ctx context.Context, req *Request[LoginReq]) (*Response[*LoginRes], error) {
	op := "UserHandler.GetUserById"

	res, err := uh.UserService.Login(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return res, nil
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
	ID string `formam:"id"`
}

type GetUserByIdRes struct {
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
