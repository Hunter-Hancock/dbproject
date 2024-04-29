package db

import (
	"database/sql"
	"fmt"
	"time"
)

type PaymentMethod struct {
	ID   string
	Name string
}

type OrderMethod struct {
	ID   string
	Name string
}

type Order struct {
	ID              string
	CustomerID      string
	OrderMethodID   string
	PaymentMethodID string
	OrderTime       time.Time
	Rating          int
}

type OrderDetail struct {
	ID          string
	Quantity    int
	ItemID      string
	Item        *FoodItem
	OrderID     string
	PromotionID string
	OrderTotal  float64
	Tip         float64
	OrderTime   time.Time
}

type OrderStore interface {
	CreateOrder(id, pid string) (Order, error)
	CreateDetail(id string, items []FoodItem) (OrderDetail, error)

	GetOrders(id string) ([]Order, error)
	GetOrderDetail(id []Order) ([]*OrderDetail, error)

	GetFoodByID(id string) (*FoodItem, error)
	GetPaymentMethods() ([]PaymentMethod, error)
	GetOrderMethods() ([]OrderMethod, error)
}

type SQLOrderStore struct {
	db *sql.DB
}

func NewSQLOrderStore(db *sql.DB) *SQLOrderStore {
	return &SQLOrderStore{db: db}
}

func (o *SQLOrderStore) CreateOrder(cid, pid string) (Order, error) {
	query := "INSERT INTO ORDERS OUTPUT INSERTED.ORDER_ID VALUES (@ID, @CID, 1, @PID, CURRENT_TIMESTAMP, NULL)"

	oid, _ := GenerateRandomString(2)

	var order Order
	err := o.db.QueryRow(query, sql.Named("ID", oid), sql.Named("CID", cid), sql.Named("PID", pid)).Scan(&order.ID)
	if err != nil {
		fmt.Println(err)
	}

	return order, nil
}

func (o *SQLOrderStore) CreateDetail(orderID string, items []FoodItem) (OrderDetail, error) {
	// Begin a database transaction
	tx, err := o.db.Begin()
	if err != nil {
		return OrderDetail{}, err
	}
	defer tx.Rollback() // Rollback the transaction if there's an error

	var order OrderDetail
	order.OrderID = orderID

	for _, item := range items {
		query := "INSERT INTO ORDER_DETAIL (ORDER_DETAIL_ID, ORDER_ID, FOOD_ITEM_ID, QUANTITY, RATING, PROMOTION_ID, ORDER_TOTAL, TIP) VALUES (@DetailID, @OrderID, @FoodItemID, @Quantity, NULL, NULL, @Total, 0.00)"

		// Generate a unique detailID for each item
		detailID, _ := GenerateRandomString(2)

		order.ID = detailID
		order.Item = &item

		var total float64
		for _, item := range items {
			total += item.Price * float64(item.Quantity)
		}

		// Execute the SQL query to insert the item detail into the database
		_, err := tx.Exec(query,
			sql.Named("DetailID", detailID),
			sql.Named("OrderID", orderID),
			sql.Named("FoodItemID", item.ID),
			sql.Named("Quantity", item.Quantity),
			sql.Named("Total", total),
		)
		if err != nil {
			return OrderDetail{}, err
		}
	}

	// Commit the transaction if all insertions are successful
	if err := tx.Commit(); err != nil {
		return OrderDetail{}, err
	}

	// If everything is successful, return the created order detail
	return order, nil
}

func (o *SQLOrderStore) GetOrders(id string) ([]Order, error) {
	query := "SELECT ORDER_ID, ORDER_DATETIME FROM ORDERS WHERE CUSTOMER_ID = @ID"
	rows, err := o.db.Query(query, sql.Named("ID", id))
	if err != nil {
		fmt.Println(err)
	}

	var orders []Order

	for rows.Next() {
		var order Order
		rows.Scan(&order.ID, &order.OrderTime)
		orders = append(orders, order)
	}

	return orders, nil
}

func (o *SQLOrderStore) GetOrderDetail(ods []Order) ([]*OrderDetail, error) {
	var orders []*OrderDetail

	for _, od := range ods {

		query := "SELECT ORDER_DETAIL_ID, QUANTITY, FOOD_ITEM_ID, ORDER_ID, ORDER_TOTAL FROM ORDER_DETAIL WHERE ORDER_ID = @ID"
		rows, err := o.db.Query(query, sql.Named("ID", od.ID))
		if err != nil {
			fmt.Println(err)
		}

		for rows.Next() {
			var order OrderDetail
			rows.Scan(&order.ID, &order.Quantity, &order.ItemID, &order.OrderID, &order.OrderTotal)
			order.OrderTime = od.OrderTime
			orders = append(orders, &order)
		}
	}

	return orders, nil
}

func (o *SQLOrderStore) GetPaymentMethods() ([]PaymentMethod, error) {
	query := "SELECT * FROM PAYMENT_METHOD"
	rows, err := o.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var methods []PaymentMethod

	for rows.Next() {
		var method PaymentMethod
		rows.Scan(&method.ID, &method.Name)
		methods = append(methods, method)
	}

	return methods, nil
}

func (o *SQLOrderStore) GetOrderMethods() ([]OrderMethod, error) {
	return nil, nil
}

func (o *SQLOrderStore) GetFoodByID(id string) (*FoodItem, error) {
	query := "SELECT * FROM FOOD_ITEM WHERE FOOD_ITEM_ID = @ID"

	var item FoodItem

	err := o.db.QueryRow(query, sql.Named("ID", id)).Scan(&item.ID, &item.Name, &item.Size, &item.Quantity, &item.Price, &item.SubcategoryID)
	if err != nil {
		fmt.Println(err)
		return &FoodItem{}, err
	}

	return &item, nil
}
