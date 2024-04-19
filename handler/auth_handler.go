package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hunter-Hancock/dbproject/view/auth"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	fmt.Println(email)

	password := r.Form.Get("password")
	fmt.Println(password)

	cookie := &http.Cookie{
		Name:    "email",
		Value:   email,
		Expires: time.Now().AddDate(0, 0, 1),
	}

	http.SetCookie(w, cookie)

	w.Header().Set("HX-Push-Url", "/")
}

func (a *AuthHandler) SigupPage(w http.ResponseWriter, r *http.Request) {
	auth.SignUpPage().Render(r.Context(), w)
}
