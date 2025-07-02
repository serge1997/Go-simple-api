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

func (app AppRepostory) FindAllBook() (*[]models.Book, error) {
	var books []models.Book
	if err := app.db.Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}
	return &books, nil
}

func (app AppRepostory) FindBook(id int) (*models.Book, error) {
	var book models.Book
	if err := app.db.Preload("Author").First(&book, "Id = ?", id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
func (app AppRepostory) FindByTitle(title string) (*models.Book, error) {
	var book models.Book
	if err := app.db.First(&book, "Title = ?", title).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (app AppRepostory) DeleteBook(id int) (bool, error) {
	book, err := app.FindBook(id)
	if err != nil {
		return false, nil
	}
	if err = app.db.Delete(&book, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
