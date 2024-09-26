package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/dexguitar/chatapp/internal/errs"
	"github.com/gorilla/schema"
)

type Request[T Validator] struct {
	r      *http.Request
	Params T
	Body   T
}

type Response[T any] struct {
	StatusCode int
	Body       T
}

func (req *Request[T]) parseQuery() error {
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&req.Params, req.r.URL.Query())
	if err != nil {
		return err
	}

	err = req.Params.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (req *Request[T]) parseBody() error {
	err := json.NewDecoder(req.r.Body).Decode(&req.Body)
	if err != nil {
		return err
	}

	err = req.Body.Validate()
	if err != nil {
		return err
	}

	return nil
}

type HandleFunc[In Validator, Out any] func(context.Context, *Request[In]) (*Response[Out], error)

func Handle[In Validator, Out any](f HandleFunc[In, Out]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &Request[In]{r: r}
		response := &Response[Out]{}

		switch r.Method {
		case http.MethodPost:
			err := request.parseBody()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		case http.MethodGet:
			err := request.parseQuery()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		response, err := f(r.Context(), request)
		if err != nil {
			var customErr errs.CustomError
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
