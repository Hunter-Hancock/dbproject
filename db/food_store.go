package db

import (
	"database/sql"
	"fmt"
)

type FoodItem struct {
	ID            string
	Name          string
	Size          string
	Quantity      int
	Price         float64
	SubcategoryID string
}

type Category struct {
	ID   string
	Name string
}

type Subcategory struct {
	ID         string
	Name       string
	CategoryID string
}

type FoodStore interface {
	GetAllFoodItems() ([]FoodItem, error)
	GetFoodItemsBySubID(id string) ([]FoodItem, error)

	GetAllCategories() ([]Category, error)
	GetCategory(name string) (*Category, error)
	GetSubCategories(id string) ([]Subcategory, error)
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

func (f *SQLFoodStore) GetFoodItemsBySubID(id string) ([]FoodItem, error) {
	query := "SELECT * FROM FOOD_ITEM WHERE SUBCATEGORY_ID = @ID"
	rows, err := f.db.Query(query, sql.Named("ID", id))
	if err != nil {
		return nil, err
	}

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

func (f *SQLFoodStore) GetCategory(name string) (*Category, error) {
	query := "SELECT * FROM FOOD_CATEGORY WHERE CATEGORY_NAME = @Name"
	rows, err := f.db.Query(query, sql.Named("Name", name))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var category Category

	for rows.Next() {
		rows.Scan(&category.ID, &category.Name)
	}

	return &category, nil
}

func (f *SQLFoodStore) GetSubCategories(id string) ([]Subcategory, error) {
	query := "SELECT * FROM FOOD_SUBCATEGORY WHERE CATEGORY_ID = @ID"
	rows, err := f.db.Query(query, sql.Named("ID", id))
	if err != nil {
		return nil, err
	}

	var subcategories []Subcategory

	for rows.Next() {
		var subcategory Subcategory
		rows.Scan(&subcategory.ID, &subcategory.Name, &subcategory.CategoryID)
		subcategories = append(subcategories, subcategory)
	}

	return subcategories, err
}
