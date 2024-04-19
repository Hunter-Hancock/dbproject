package handler

import (
	"net/http"

	"github.com/Hunter-Hancock/dbproject/view/auth"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) SigupPage(w http.ResponseWriter, r *http.Request) {
	auth.SignUpPage().Render(r.Context(), w)
}
