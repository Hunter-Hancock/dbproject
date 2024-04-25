package handler

import (
	"net/http"

	"github.com/Hunter-Hancock/dbproject/view/error_page"
	"github.com/Hunter-Hancock/dbproject/view/home"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	home.IndexPage().Render(r.Context(), w)
}

func (h *Handler) ErrorPage(w http.ResponseWriter, r *http.Request) {
	eMsg := r.URL.Query().Get("error")
	e := error_page.PError{
		Message: eMsg,
	}
	error_page.ErrorPage(e).Render(r.Context(), w)
}
