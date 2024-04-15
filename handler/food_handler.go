package handler

import (
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/view/home"
)

type FoodHandler struct {
	FoodStore db.FoodStore
}

func NewFoodHandler(db db.FoodStore) *FoodHandler {
	return &FoodHandler{FoodStore: db}
}

func (f *FoodHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	foodItems, err := f.FoodStore.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	home.Products(foodItems).Render(r.Context(), w)
}
