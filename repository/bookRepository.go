package repository

import (
	"github.com/serge1197/go-simple-api/models"
)

func (app AppRepostory) PersistBook(book models.Book) (*int64, error) {
	stmt, err := app.db.Prepare("INSERT INTO books(Title, Edition, Year, AuthorId) VALUES(?,?,?,?)")
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

func (app AppRepostory) FindAllBook() {

}

func (app AppRepostory) FindBook(id int) {

}
func (app AppRepostory) FindByTitle(title string) {

}

func (app AppRepostory) DeleteBook(id int) {

}
