package handler

import (
	"net/http"

	"github.com/Hunter-Hancock/dbproject/view/home"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	home.IndexPage().Render(r.Context(), w)
}
