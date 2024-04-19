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
	categoryId := chi.URLParam(r, "id")
	subcategory, err := f.FoodStore.GetSubCategory(categoryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	items, err := f.FoodStore.GetFoodItemsBySubID(subcategory.ID)
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

	category.Category(c).Render(r.Context(), w)
}
