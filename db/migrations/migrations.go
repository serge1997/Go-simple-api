package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/serge1197/go-simple-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	connection, err := gorm.Open(sqlite.Open("../database.sqlite"))
	if err != nil {
		panic(err)
	}
	db = connection
}
func main() {
	RunMigrations()
	fmt.Println("Migrations done!")
}

func RunMigrations() {
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
}
