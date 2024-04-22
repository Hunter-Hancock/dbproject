package handler

import (
	"fmt"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/session"
	"github.com/Hunter-Hancock/dbproject/view/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthStore db.AuthStore
	Session   *session.SessionManager
}

func NewAuthHandler(db db.AuthStore, session *session.SessionManager) *AuthHandler {
	return &AuthHandler{AuthStore: db, Session: session}
}

func (a *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")

	password := r.Form.Get("password")

	firstName := r.Form.Get("firstName")
	lastName := r.Form.Get("lastName")

	customer := &db.Customer{
		FirstName: firstName,
		LastName:  lastName,
	}

	hashed, err := HashPassword(password)
	if err != nil {
		fmt.Println(err.Error())
	}

	user := &db.User{
		Email:          email,
		Password:       password,
		HashedPassword: hashed,
		Customer:       customer,
	}

	if err := a.AuthStore.CreateUser(user); err != nil {
		fmt.Println("Error creating User", err)
		return
	}

	if err2 := a.AuthStore.CreateCustomer(user); err2 != nil {
		fmt.Println("Error creating Customer", err2)
	}

	a.Session.SetAuthSession(w, r, user)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.Form.Get("email")

	password := r.Form.Get("password")

	user := &db.User{
		Email:    email,
		Password: password,
	}

	valid := a.AuthStore.ValidateUser(user)

	if valid {

		a.AuthStore.GetCustomer(user)

		err := a.Session.SetAuthSession(w, r, user)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error setting session")
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	formError := &auth.FormError{
		Message: "User Not Found",
	}

	auth.LoginPage(formError).Render(r.Context(), w)
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	s, err := a.Session.GetAuthSession(r)
	if err != nil {
		return
	}

	s.Values[session.SessionAccessTokenKey] = ""
	s.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *AuthHandler) SigupPage(w http.ResponseWriter, r *http.Request) {
	auth.SignUpPage().Render(r.Context(), w)
}

func (a *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	auth.LoginPage(&auth.FormError{}).Render(r.Context(), w)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
