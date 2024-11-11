package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func InitRouter(uh *UserHandler, mh *MessageHandler, wsh *WSHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/users", Handle(uh.RegisterUser))
	r.Post("/login", Handle(uh.Login))
	r.Get("/users", Handle(uh.GetUserById))

	r.Post("/sendMessage", Handle(mh.SendMessage))

	r.Handle("/ws", wsh)

	return r
}
