package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRouter(uh *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/register", Handle(uh.UserService.RegisterUser))
	r.Post("/login", Handle(uh.UserService.Login))
	r.Get("/getUserById", Handle(uh.UserService.GetUserById))

	return r
}
