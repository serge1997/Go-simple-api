package db

import (
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB
var ErrNoDbConnection error = errors.New("no connection provided")

func init() {
	ConnSqlite()
}
func ConnSqlite() {
	conn, err := gorm.Open(sqlite.Open("./db/database.sqlite"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")
	Connection = conn
}

func Close() {
	db, _ := Connection.DB()
	db.Close()
}
