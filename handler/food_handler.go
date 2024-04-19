package handler

import (
	"context"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/view/category"
	"github.com/Hunter-Hancock/dbproject/view/home"
	"github.com/go-chi/chi/v5"
)

type FoodHandler struct {
	FoodStore db.FoodStore
}

func NewFoodHandler(db db.FoodStore) *FoodHandler {
	return &FoodHandler{FoodStore: db}
}

func (f *FoodHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	// foodItems, err := f.FoodStore.GetAllFoodItems()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// home.Products(foodItems).Render(r.Context(), w)
}

func (f *FoodHandler) HandleGetProductsBySubID(w http.ResponseWriter, r *http.Request) {
	subcatId := chi.URLParam(r, "id")

	items, err := f.FoodStore.GetFoodItemsBySubID(subcatId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	category.Products(items).Render(r.Context(), w)
}

func (f *FoodHandler) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := f.FoodStore.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	type contextKey string

	var userContextKey contextKey = "userID"

	ctx := context.WithValue(context.Background(), userContextKey, "DOC")

	home.Categories(categories).Render(ctx, w)
}

func (f *FoodHandler) HandleGetCategory(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	c, err := f.FoodStore.GetCategory(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	subcats, err2 := f.FoodStore.GetSubCategories(c.ID)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusNotFound)
	}

	category.Subcategories(subcats).Render(r.Context(), w)
}
