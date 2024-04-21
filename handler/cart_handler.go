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

	message := fmt.Sprintf("Added to cart: %s item: %s", user.Customer.FirstName, id)

	io.WriteString(w, message)
}

func (ch *CartHandler) HandleCartTotal(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(mw.UserIDKey).(*db.User)
	if !ok {
		return
	}

	total, err := ch.CartStore.GetTotal(user.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, fmt.Sprintf("Cart: %d", total))
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
