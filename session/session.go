package session

import (
	"encoding/gob"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/gorilla/sessions"
)

const (
	SessionUserKey        = "user"
	SessionAccessTokenKey = "accessToken"
)

type SessionManager struct {
	Store *sessions.CookieStore
}

func NewSessionManager(store *sessions.CookieStore) *SessionManager {
	return &SessionManager{Store: store}
}

func (sm *SessionManager) SetAuthSession(w http.ResponseWriter, r *http.Request, user *db.User) error {
	session, _ := sm.Store.Get(r, SessionUserKey)
	gob.Register(user)
	session.Values[SessionAccessTokenKey] = user
	return session.Save(r, w)
}

func (sm *SessionManager) GetAuthSession(r *http.Request) (*sessions.Session, error) {
	return sm.Store.Get(r, SessionUserKey)
}
