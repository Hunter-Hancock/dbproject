package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	mw "github.com/Hunter-Hancock/dbproject/middleware"
	"github.com/Hunter-Hancock/dbproject/view/cart"
	"github.com/go-chi/chi/v5"
)

type CartHandler struct {
	CartStore db.CartStore
}

func NewCartHandler(db db.CartStore) *CartHandler {
	return &CartHandler{CartStore: db}
}

func (ch *CartHandler) HandleAdd(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(mw.UserIDKey).(*db.User)

	id := chi.URLParam(r, "id")

	if err2 := ch.CartStore.Add(user, id); err2 != nil {
		fmt.Println(err2)
	}

	io.WriteString(w, "Added to Cart!")
}

func (ch *CartHandler) HandleRemove(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(mw.UserIDKey).(*db.User)

	id := chi.URLParam(r, "id")
	if err := ch.CartStore.Remove(user, id); err != nil {
		fmt.Println(err)
	}

}

func (ch *CartHandler) HandleCartQuantity(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(mw.UserIDKey).(*db.User)
	if !ok {
		return
	}

	quantity, err := ch.CartStore.GetCartQuantity(user.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, fmt.Sprintf("%d", quantity))
}

func (ch *CartHandler) HandleCartTotal(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(mw.UserIDKey).(*db.User)
	if !ok {
		return
	}

	total, err := ch.CartStore.GetCartTotal(user.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, fmt.Sprintf("%.2f", total))
}

func (ch *CartHandler) ShowCart(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(mw.UserIDKey).(*db.User)
	if !ok {
		return
	}
	items, err := ch.CartStore.GetItems(user.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}

	cart.CartPage(items).Render(r.Context(), w)
}
