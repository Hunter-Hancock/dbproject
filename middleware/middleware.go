package mw

import (
	"context"
	"net/http"
)

type MiddleWare struct{}

type contextKey string

const UserIDKey contextKey = "UserID"

func (mw *MiddleWare) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), UserIDKey, 123)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
