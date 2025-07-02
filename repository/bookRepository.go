package repository

import (
	"errors"

	"github.com/serge1197/go-simple-api/models"
)

func (app AppRepostory) PersistBook(book *models.Book) (*models.Book, error) {
	if err := app.db.Create(&book).Error; err != nil {
		return nil, err
	}
	if book.Id == 0 {
		var err = errors.New("erro when trying to create book")
		return nil, err
	}
	return book, nil
}

func (app AppRepostory) FindAllBook() {

}

func (app AppRepostory) FindBook(id int) {

}
func (app AppRepostory) FindByTitle(title string) {

}

func (app AppRepostory) DeleteBook(id int) {

}
