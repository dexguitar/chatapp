package model

import validation "github.com/go-ozzo/ozzo-validation"

type Message struct {
	Username string
	Value    string
	Receiver string
}

type SendMessageReq struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type SendMessageRes struct{}

func (r *SendMessageReq) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Message, validation.Required, validation.Length(1, 1000)),
		validation.Field(&r.Receiver, validation.Required, validation.Length(1, 1000)),
	)
}
