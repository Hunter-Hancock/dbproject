package db

import (
	"database/sql"
)

type FoodItem struct {
	ID            string
	Name          string
	Size          string
	Quantity      int
	Price         float32
	SubcategoryID string
}

type Category struct {
	ID   string
	Name string
}

type FoodStore interface {
	GetAllFoodItems() ([]FoodItem, error)

	GetAllCategories() ([]Category, error)
}

type SQLFoodStore struct {
	db *sql.DB
}

func NewSQLFoodStore(db *sql.DB) *SQLFoodStore {
	return &SQLFoodStore{db: db}
}

func (f *SQLFoodStore) GetAllFoodItems() ([]FoodItem, error) {
	query := "SELECT * FROM FOOD_ITEM"
	rows, err := f.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []FoodItem

	for rows.Next() {
		var item FoodItem
		rows.Scan(&item.ID, &item.Name, &item.Size, &item.Quantity, &item.Price, &item.SubcategoryID)
		items = append(items, item)
	}
	return items, nil
}

func (f *SQLFoodStore) GetAllCategories() ([]Category, error) {
	query := "SELECT * FROM FOOD_CATEGORY"
	rows, err := f.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}

	return categories, nil
}
