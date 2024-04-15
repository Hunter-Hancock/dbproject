package db

import "database/sql"

type TestStore interface {
	test()
}

type SQLTestStore struct {
	db *sql.DB
}

func NewSQLTestStore(db *sql.DB) *SQLTestStore {
	return &SQLTestStore{db: db}
}

func (t *SQLTestStore) test() {
}
