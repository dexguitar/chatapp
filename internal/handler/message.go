package handler

type MessageHandler struct {
	MessageService
}

func NewMessageHandler(msgService MessageService) *MessageHandler {
	return &MessageHandler{msgService}
}
