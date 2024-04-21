package db

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             mssql.UniqueIdentifier
	Email          string
	Password       string
	HashedPassword string
	LoggedIn       bool
	Customer       *Customer
}

type Customer struct {
	ID         string
	FirstName  string
	LastName   string
	ClubCardID string
	UserID     mssql.UniqueIdentifier
}

type AuthStore interface {
	CreateUser(user *User) error

	GetCustomer(user *User) error
	CreateCustomer(user *User) error
	ValidateUser(user *User) bool
}

type SQLAuthStore struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) *SQLAuthStore {
	return &SQLAuthStore{db: db}
}

func (a *SQLAuthStore) CreateUser(user *User) error {
	query := "INSERT INTO USERS (Email, PasswordHash) OUTPUT INSERTED.ID VALUES (@Email, @Hash)"
	var userID mssql.UniqueIdentifier
	err := a.db.QueryRow(query, sql.Named("Email", user.Email), sql.Named("Hash", user.HashedPassword)).Scan(&userID)

	user.ID = userID

	return err
}

func (a *SQLAuthStore) CreateCustomer(user *User) error {
	rid, _ := generateRandomString(2)
	query := "INSERT INTO CLUB_CARD VALUES (@ID, 0)"

	a.db.Exec(query, sql.Named("ID", rid))

	rid2, _ := generateRandomString(2)

	user.Customer.ID = rid2
	user.Customer.ClubCardID = rid

	query2 := "INSERT INTO CUSTOMER VALUES (@ID, @FName, @LName, '123 Address Rd', 'Somewhere', 'Alabama', '99999', '1999-06-01', @CCID, @UID)"

	if _, err := a.db.Exec(query2, sql.Named("ID", rid2), sql.Named("FName", user.Customer.FirstName), sql.Named("LName", user.Customer.LastName), sql.Named("CCID", rid), sql.Named("UID", user.ID)); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (a *SQLAuthStore) GetCustomer(user *User) error {
	query := "SELECT CUSTOMER_ID, F_NAME, L_NAME, CLUB_CARD_ID FROM CUSTOMER WHERE USER_ID = @ID"
	row := a.db.QueryRow(query, sql.Named("ID", user.ID))

	c := &Customer{}

	if err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.ClubCardID); err != nil {
		fmt.Println("db error", err)
	}

	user.Customer = c

	return nil
}

func (a *SQLAuthStore) ValidateUser(user *User) bool {
	query := "SELECT * FROM USERS WHERE EMAIL = @Email"
	row := a.db.QueryRow(query, sql.Named("Email", user.Email))

	if err := row.Scan(&user.ID, &user.Email, &user.HashedPassword); err != nil {
		fmt.Println(err)
	}

	return CheckPasswordHash(user.Password, user.HashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	randomString := hex.EncodeToString(bytes)

	return randomString[:length], nil
}
