package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dexguitar/chatapp/internal/model"
	validation "github.com/go-ozzo/ozzo-validation"
)

type MessageHandler struct {
	MessageService
}

func NewMessageHandler(messageService MessageService) *MessageHandler {
	return &MessageHandler{messageService}
}

func (mh *MessageHandler) SendMessage(ctx context.Context, req *Request[SendMessageReq]) (*Response[SendMessageRes], error) {
	op := "MessageHandler.SendMessage"
  
	err := mh.MessageService.SendMessage(ctx, &model.Message{
		Username: req.Body.Username,
		Receiver: req.Body.Receiver,
		Value:    req.Body.Value,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
  
	return &Response[SendMessageRes]{
		StatusCode: http.StatusOK,
	}, nil
}

type SendMessageReq struct {
	Username string `json:"username"`
	Receiver string `json:"receiver"`
	Value    string `json:"value"`
}

type SendMessageRes struct{}

func (r SendMessageReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Receiver, validation.Required),
		validation.Field(&r.Value, validation.Required),
	)
}
