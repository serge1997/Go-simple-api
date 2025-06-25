package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Connection *sql.DB
var ErrNoDbConnection error = errors.New("no connection provided")

func init() {
	ConnSqlite()
}
func ConnSqlite() {
	conn, err := sql.Open("sqlite3", "./db/database.sqlite")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")
	Connection = conn
}
