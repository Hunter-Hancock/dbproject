package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pressly/goose/v3"
)

const (
	dialect     = "sqlserver"
	fmtDBString = "Server=%s;Database=%s;User Id=%s;Password=%s;"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}
	command := args[0]

	host := os.Getenv("DBHOST")
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")

	dbstring := fmt.Sprintf(fmtDBString, host, dbname, dbuser, dbpass)

	db, err := goose.OpenDBWithDriver(dialect, dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
