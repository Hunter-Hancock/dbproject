package db

import "database/sql"

type FoodItem struct {
	ID            string
	Name          string
	Size          string
	Quantity      float32
	Price         float32
	SubcategoryID string
}

type FoodStore interface {
	GetAll() ([]FoodItem, error)
}

type SQLFoodStore struct {
	db *sql.DB
}

func NewSQLFoodStore(db *sql.DB) *SQLFoodStore {
	return &SQLFoodStore{db: db}
}

func (f *SQLFoodStore) GetAll() ([]FoodItem, error) {
	query := "SELECT * FROM FOOD_ITEM"
	rows, err := f.db.Query(query)
	if err != nil {
		return nil, err
	}

	var item FoodItem
	var items []FoodItem

	for rows.Next() {
		rows.Scan(&item.ID, &item.Name, &item.Size, &item.Quantity, &item.SubcategoryID, &item.Price)
		items = append(items, item)
	}
	return items, nil
}
