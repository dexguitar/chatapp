package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/dexguitar/chatapp/internal/utils"
	"github.com/monoculum/formam"
)

type Request[T Validator] struct {
	Body T
}

type Response[T any] struct {
	StatusCode int
	Body       T
}

type HandleFunc[In Validator, Out any] func(context.Context, *Request[In]) (*Response[Out], error)

func Handle[In Validator, Out any](f HandleFunc[In, Out]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &Request[In]{}
		response := &Response[Out]{}

		switch r.Method {
		case http.MethodPost:
			err := json.NewDecoder(r.Body).Decode(&request.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case http.MethodGet:
			r.ParseForm()
			decoder := formam.NewDecoder(&formam.DecoderOptions{})
			err := decoder.Decode(r.Form, &request.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		err := request.Body.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()

		response, err = f(ctx, request)
		if err != nil {
			var customErr utils.CustomError
			if errors.As(err, &customErr) {
				slog.Error(err.Error())
				http.Error(w, customErr.Message, customErr.StatusCode)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response.Body)
	}
}
