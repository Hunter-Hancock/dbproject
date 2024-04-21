package mw

import (
	"context"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/session"
)

type MiddleWare struct {
	AuthStore db.AuthStore
	Session   *session.SessionManager
}

type contextKey string

const UserIDKey contextKey = "UserID"

func (mw *MiddleWare) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, err := mw.Session.GetAuthSession(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user, ok := s.Values[session.SessionAccessTokenKey].(*db.User)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUser(ctx context.Context) *db.User {
	if user, ok := ctx.Value(UserIDKey).(*db.User); ok {

		return user
	}

	return nil
}

func GetCustomer(user *db.User) *db.Customer {
	return user.Customer
}

func GetUserName(ctx context.Context) string {
	if userName, ok := ctx.Value(UserIDKey).(string); ok {
		return userName
	}

	return "Not Logged In"
}
