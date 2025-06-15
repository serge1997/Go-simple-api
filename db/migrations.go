package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Connection *sql.DB
var ErrNoDbConnection error = errors.New("no connection provided")

func ConnSqlite() {
	conn, err := sql.Open("sqlite3", "./db/database.sqlite")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")
	Connection = conn
}
func RunMigrations() *error {
	if Connection == nil {
		return &ErrNoDbConnection
	}
	createBooksTable()
	createAuthorsTable()
	return nil
}

func createAuthorsTable() {
	smt, err := Connection.Prepare(`
		CREATE TABLE IF NOT EXISTS authors(
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name VARCHAR(60),
			Website VARCHAR(30) NULL,
			CreatedAt DATETIME,
			UpdatedAt DATETIME NULL
		)
	`)
	if err != nil {
		panic(err)
	}
	smt.Exec()
}

func createBooksTable() {
	stmt, err := Connection.Prepare(`
		CREATE TABLE IF NOT EXISTS books(
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Title VARCHAR(60),
			Edition VARCHAR(45),
			Year INTEGER,
			AuthorId INTEGER,
			CreatedAt DATETIME,
			UpdatedAt DATETIME NULL,
			FOREIGN KEY (AuthorId) REFERENCES authors(Id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}
