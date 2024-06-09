package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRouter(uh *UserHandler, mh *MessageHandler, wsh *WSHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/register", Handle(uh.UserService.RegisterUser))
	r.Post("/login", Handle(uh.UserService.Login))
	r.Post("/sendMessage", Handle(mh.MessageService.SendMessage))
	r.Get("/getUserById", Handle(uh.UserService.GetUserById))
	r.Handle("/ws", wsh)

	return r
}
