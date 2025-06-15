package book

import (
	"database/sql"
)

func (book Book) Persist(db *sql.DB) (*int64, error) {
	stmt, err := db.Prepare("INSERT INTO books(Title, Edition, Year, AuthorId) VALUES(?,?,?,?)")
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(book.Title, book.Edition, book.Year, book.Author.Id)
	if err != nil {
		return nil, err
	}
	if id, err := result.LastInsertId(); err != nil {
		return nil, err
	} else {
		return &id, nil
	}

}

func FindAll(db *sql.DB) {

}

func Find(db *sql.DB, id int) {

}
func FindByTitle(db *sql.DB, title string) {

}

func (book Book) Delete(db *sql.DB, id int) {

}
