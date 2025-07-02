package repository

import (
	"errors"

	"github.com/serge1197/go-simple-api/models"
)

func (app AppRepostory) PersistAuthor(author *models.Author) (*models.Author, error) {
	var ErrExists = errors.New("author already exists in database")

	exist, err := app.FindAuthorByName(author.Name)
	if err != nil {
		return nil, ErrExists
	}
	if exist.Id != 0 {
		return nil, ErrExists
	}
	if err := app.db.Create(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (app AppRepostory) FindAuthor(id int) (*models.Author, error) {
	var author models.Author
	var ErrNotFound = errors.New("author not Found")
	if err := app.db.Preload("Books").First(&author, "Id = ?", id).Error; err != nil {
		return nil, err
	}
	if author.Id == 0 {
		return nil, ErrNotFound
	}
	return &author, nil
}

func (app AppRepostory) FindAllAuthor() (*[]models.Author, error) {
	var authors []models.Author
	if err := app.db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return &authors, nil
}

func (app AppRepostory) FindAuthorByName(name string) (*models.Author, error) {
	var author models.Author
	if err := app.db.First(&author, "Name = ?", name).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (app AppRepostory) UpdateAuthor(author models.Author) (*models.Author, error) {
	finded, err := app.FindAuthor(int(author.Id))
	if err != nil {
		return nil, err
	}
	if err = app.db.Model(&finded).Updates(author).Error; err != nil {
		return nil, err
	}
	return finded, nil
}

func (app AppRepostory) DeleteAuthor(id int) (bool, error) {

	finded, err := app.FindAuthor(id)
	if err != nil {
		return false, err
	}
	if err = app.db.Delete(&finded, id).Error; err != nil {
		return false, nil
	}
	return true, nil
}
