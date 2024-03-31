package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
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
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *CreateUserReq) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

func (r *LoginReq) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

func (r *GetUserByIdReq) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.ID, validation.Required),
	)
}
