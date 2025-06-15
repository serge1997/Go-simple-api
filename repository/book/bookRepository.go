package book

import "database/sql"

func (book Book) Persist(db *sql.DB) {

}

func FindAll(db *sql.DB) {

}

func Find(db *sql.DB, id int) {

}
func FindByTitle(db *sql.DB, title string) {

}

func (book Book) Delete(db *sql.DB, id int) {

}
