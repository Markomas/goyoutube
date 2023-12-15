package security

import (
	"context"
	"errors"
	"net/http"
	"youtube/internal/db"
	"youtube/internal/response"
)

func ApiKeyMiddleware(next http.HandlerFunc, database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			response.ErrorResponse(w, errors.New("missing X-API-KEY header"), http.StatusBadRequest)
			return
		}

		user, err := database.CheckApiKey(apiKey)
		if err != nil {
			response.ErrorResponse(w, errors.New("bad X-API-KEY"), http.StatusInternalServerError)
			return
		}
		if user == nil {
			response.ErrorResponse(w, errors.New("bad X-API-KEY"), http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user", user))
		next.ServeHTTP(w, r)
	}
}
