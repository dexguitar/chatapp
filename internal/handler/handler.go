package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var ErrUserNotFound = errors.New("user not found")

type TargetFunc[In Validator, Out any] func(context.Context, In) (Out, error)

func Handle[In Validator, Out any](f TargetFunc[In, Out]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in In
		var out Out

		switch r.Method {
		case http.MethodGet:
			params := r.URL.Query()

			jsonString := "{"
			for key, values := range params {
				jsonString += "\"" + key + "\":\"" + values[0] + "\","
			}
			jsonString = jsonString[:len(jsonString)-1] + "}"

			if err := json.Unmarshal([]byte(jsonString), &in); err != nil {
				http.Error(w, "please fill all required parameters", http.StatusBadRequest)
				return
			}

			err := in.Validate()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if out, err = f(r.Context(), in); err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					http.Error(w, ErrUserNotFound.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(out)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case http.MethodPost:
			err := json.NewDecoder(r.Body).Decode(&in)
			if err != nil {
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}

			err = in.Validate()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if out, err = f(r.Context(), in); err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					http.Error(w, ErrUserNotFound.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(out)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
