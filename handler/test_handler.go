package handler

import (
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	home "github.com/Hunter-Hancock/dbproject/view/home"
)

type TestHandler struct {
	TestStore db.TestStore
}

func NewTestHandler(db db.TestStore) *TestHandler {
	return &TestHandler{TestStore: db}
}

func (t TestHandler) Home(w http.ResponseWriter, r *http.Request) {
	home.IndexPage().Render(r.Context(), w)
}

func (t TestHandler) Click(w http.ResponseWriter, r *http.Request) {
	home.Message("HTMX is working :)").Render(r.Context(), w)
}

// func (t TestHandler) Test(w http.ResponseWriter, r *http.Request) {

// 	ctx := r.Context()

// 	userIDValue := ctx.Value(mw.UserIDKey)
// 	userID, ok := userIDValue.(int)
// 	if !ok {
// 		http.Error(w, "User ID not found or not of type int", http.StatusInternalServerError)
// 		return
// 	}

// 	msg := fmt.Sprintf("User ID: %d", userID)

// 	view.IndexPage(msg).Render(r.Context(), w)
// }
