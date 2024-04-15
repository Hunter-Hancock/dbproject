package main

//
// import (
// 	"database/sql"
// 	"fmt"
//
// 	_ "github.com/denisenkom/go-mssqldb"
// )
//
// type DBConfig struct {
// 	user     string
// 	dbname   string
// 	password string
// 	dbtype   string
// 	host     string
// }
//
// func NewDBConfig(host, dbtype, dbname, user, password string) *DBConfig {
// 	return &DBConfig{user: user, dbname: dbname, password: password, dbtype: dbtype, host: host}
// }
//
// type DB interface {
// 	GetAllOrders() ([]*Order, error)
// 	CreateOrder(string) (*Order, error)
// 	DeleteOrder(int) error
//
// 	CreateEmployee(first_name, last_name, street, city, state, zip, role string) (*Employee, error)
// }
//
// type Store struct {
// 	db *sql.DB
// }
//
// func InitDB(config *DBConfig) (*Store, error) {
// 	connectionString := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", config.host, config.dbname, config.user, config.password)
// 	db, err := sql.Open(config.dbtype, connectionString)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}
//
// 	store := &Store{db: db}
//
// 	return store, nil
// }
//
// func (s *Store) CreateEmployee(first_name, last_name, street, city, state, zip, role string) (*Employee, error) {
// 	query := `INSERT INTO Employee (first_name, last_name, street, city, state, zip, role) OUTPUT INSERTED.ID VALUES (@first_name, @last_name, @street, @city, @state, @zip, @role)`
//
// 	var createdId int
// 	err := s.db.QueryRow(query, sql.Named("first_name", first_name), sql.Named("last_name", last_name), sql.Named("street", street), sql.Named("city", city), sql.Named("state", state), sql.Named("zip", zip), sql.Named("role", role)).Scan(&createdId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	employee := &Employee{Id: createdId, FirstName: first_name, LastName: last_name, Street: street, City: city, State: state, Zip: zip, Role: role}
//
// 	return employee, nil
// }
//
// func (s *Store) GetAllOrders() ([]*Order, error) {
// 	orders := []*Order{}
//
// 	rows, err := s.db.Query("SELECT * FROM Orders")
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for rows.Next() {
// 		order := new(Order)
// 		if err := rows.Scan(&order.Id, &order.Method); err != nil {
// 			return nil, fmt.Errorf("Error scanning: %s", err)
// 		}
//
// 		orders = append(orders, order)
// 	}
//
// 	return orders, nil
// }
//
// func (s *Store) CreateOrder(method string) (*Order, error) {
// 	query := "INSERT INTO Orders (Method) OUTPUT INSERTED.ID VALUES (@method)"
//
// 	var createdID int
// 	err := s.db.QueryRow(query, sql.Named("Method", method)).Scan(&createdID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	order := &Order{Id: createdID, Method: method}
//
// 	return order, nil
// }
//
// func (s *Store) DeleteOrder(id int) error {
// 	_, err := s.db.Exec("DELETE FROM Orders WHERE id = @id", sql.Named("id", id))
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
