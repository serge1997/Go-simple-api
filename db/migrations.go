package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Connection *sql.DB

func ConnSqlite() {
	conn, err := sql.Open("sqlite3", "./db/database.sqlite")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")
	Connection = conn
}
