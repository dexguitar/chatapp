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
		Sender:   req.Body.Sender,
		Receiver: req.Body.Receiver,
		Content:  req.Body.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Response[SendMessageRes]{
		StatusCode: http.StatusOK,
	}, nil
}

type SendMessageReq struct {
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Content  string `json:"content"`
}

type SendMessageRes struct{}

func (r SendMessageReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Sender, validation.Required),
		validation.Field(&r.Receiver, validation.Required),
		validation.Field(&r.Content, validation.Required),
	)
}
