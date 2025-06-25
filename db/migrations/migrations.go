package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/serge1197/go-simple-api/db/sqlmig"
)

var db *sql.DB

func init() {
	connection, err := sql.Open("sqlite3", "../database.sqlite")
	if err != nil {
		panic(err)
	}
	db = connection
}
func main() {
	err := RunMigrations()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrations done!")
}

func RunMigrations() error {
	var err error
	if db == nil {
		err = errors.New("unable to run migrations. check database connection")
		return err
	}
	sqlmig.CreateBooksTable(db)
	sqlmig.CreateAuthorsTable(db)
	return nil
}
