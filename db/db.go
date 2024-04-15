package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/joho/godotenv/autoload"
)

func Open() (*sql.DB, error) {
	host := os.Getenv("DBHOST")
	dbtype := os.Getenv("DBTYPE")
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")

	connectionString := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", host, dbname, dbuser, dbpass)
	db, err := sql.Open(dbtype, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
