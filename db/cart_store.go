package db

import (
	"database/sql"
	"fmt"
)

type CartStore interface {
	Add(user *User, item string) error

	GetItems(id string) ([]FoodItem, error)
	GetTotal(id string) (int, error)
}

type SQLCartStore struct {
	db *sql.DB
}

func NewSQLCartStore(db *sql.DB) *SQLCartStore {
	return &SQLCartStore{db: db}
}

func (s *SQLCartStore) Add(user *User, item string) error {

	var existingQuantity int
	err := s.db.QueryRow("SELECT QUANTITY FROM CART WHERE CUSTOMER_ID = @ID AND FOOD_ITEM_ID = @FID",
		sql.Named("ID", user.Customer.ID), sql.Named("FID", item)).Scan(&existingQuantity)

	if err != nil && err != sql.ErrNoRows {
		// Handle error
		fmt.Println("Error checking existing quantity:", err)
		return err
	}

	// If the item already exists, update the quantity
	if err == nil {
		_, err := s.db.Exec("UPDATE CART SET QUANTITY = QUANTITY + 1 WHERE CUSTOMER_ID = @ID AND FOOD_ITEM_ID = @FID",
			sql.Named("ID", user.Customer.ID), sql.Named("FID", item))
		if err != nil {
			fmt.Println("Error updating quantity:", err)
			return err
		}
	} else {
		// If the item doesn't exist, insert a new row
		query := "INSERT INTO CART VALUES (@ID, @FID, 1)"
		if _, err := s.db.Exec(query, sql.Named("ID", user.Customer.ID), sql.Named("FID", item)); err != nil {
			fmt.Println("Error inserting new row:", err)
			return err
		}
	}

	return nil

}

func (s *SQLCartStore) GetTotal(id string) (int, error) {
	var total int
	query := "SELECT SUM(QUANTITY) AS TOTAL FROM CART WHERE CUSTOMER_ID = @ID"
	row := s.db.QueryRow(query, sql.Named("ID", id))
	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (s *SQLCartStore) GetItems(id string) ([]FoodItem, error) {
	query := "SELECT f.FOOD_ITEM_ID, f.FOOD_NAME, f.FOOD_SIZE, c.QUANTITY, f.PRICE FROM CART c JOIN FOOD_ITEM f ON c.FOOD_ITEM_ID = f.FOOD_ITEM_ID WHERE CUSTOMER_ID = @ID"
	rows, err := s.db.Query(query, sql.Named("ID", id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var items []FoodItem

	for rows.Next() {
		var item FoodItem
		rows.Scan(&item.ID, &item.Name, &item.Size, &item.Quantity, &item.Price)
		items = append(items, item)
	}

	return items, nil
}
