package sqlmig

import (
	"database/sql"
	"log"
)

func CreateBooksTable(Connection *sql.DB) {
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

func CreateAuthorsTable(Connection *sql.DB) {
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
