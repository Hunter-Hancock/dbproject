package mw

import (
	"context"
	"fmt"
	"net/http"
)

type MiddleWare struct{}

type contextKey string

const UserIDKey contextKey = "UserID"

func (mw *MiddleWare) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		test, err := r.Cookie("email")
		if err != nil {
			fmt.Println(err)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, test.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserName(ctx context.Context) string {
	if userName, ok := ctx.Value(UserIDKey).(string); ok {
		return userName
	}

	return "Not Logged In"
}
