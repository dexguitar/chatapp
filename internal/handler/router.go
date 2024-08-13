package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func InitRouter(uh *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/users", Handle(uh.RegisterUser))
	r.Post("/login", Handle(uh.Login))
	r.Get("/users", Handle(uh.GetUserById))

	return r
}
