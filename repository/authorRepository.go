package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/serge1197/go-simple-api/models"
)

func (app AppRepostory) PersistAuthor(author models.Author) (*int64, error) {
	var erro error
	var ErrExists = errors.New("author already exists in database")

	exist := app.FindAuthorByName(author.Name)
	if exist != nil {
		return nil, ErrExists
	}

	smt, err := app.db.Prepare("INSERT INTO authors(Name, Website, CreatedAt) VALUES(?, ?, ?)")

	if err != nil {
		erro = err
		return nil, err
	}
	result, err := smt.Exec(author.Name, author.Website, author.CreatedAt)

	if err != nil {
		erro = err
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &id, erro
}

func (app AppRepostory) FindAuthor(id int) (*models.Author, error) {
	var author models.Author
	var ErrNotFound = errors.New("Author not Found")
	smt := app.db.QueryRow("SELECT * FROM authors WHERE Id = ?", id)
	err := smt.Scan(&author.Id, &author.Name, &author.Website, &author.CreatedAt, &author.UpdatedAt)
	if err == sql.ErrNoRows || author.Name == "" {
		return nil, ErrNotFound
	}
	return &author, nil
}

func (app AppRepostory) FindAllAuthor() (*[]models.Author, error) {
	var authors []models.Author
	rows, err := app.db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author models.Author

		rows.Scan(&author.Id, &author.Name, &author.Website, &author.CreatedAt, &author.UpdatedAt)
		authors = append(authors, author)
	}
	return &authors, nil
}

func (app AppRepostory) FindAuthorByName(name string) *models.Author {
	var author models.Author
	row := app.db.QueryRow("SELECT * FROM authors WHERE Name = ?", name)
	row.Scan(&author.Id, &author.Name, &author.Website, &author.CreatedAt, &author.UpdatedAt)
	if author.Id == 0 {
		return nil
	}
	return &author
}

func (app AppRepostory) UpdateAuthor(author models.Author) (*models.Author, error) {
	var ErrOcurred = errors.New("erro ocorred on updating author")
	finded, err := app.FindAuthor(int(author.Id))
	if err != nil {
		return nil, err
	}
	smt, err := app.db.Prepare("UPDATE authors SET Name = ?, Website = ?, UpdatedAt = ? WHERE Id = ?")
	if err != nil {
		return nil, err
	}
	result, err := smt.Exec(author.Name, author.Website, time.Now(), finded.Id)
	if err != nil {
		return nil, err
	}
	isUpdated, _ := result.RowsAffected()
	if isUpdated >= 1 {
		retrieve, _ := app.FindAuthor(int(author.Id))
		return retrieve, nil
	}
	return nil, ErrOcurred
}

func (app AppRepostory) DeleteAuthor(id int) (bool, error) {
	stmt, err := app.db.Prepare("DELETE FROM authors WHERE id = ?")

	if err != nil {
		return false, errors.New("nao foi possivel remover o registro, " + err.Error())
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return false, errors.New("nao foi possivel remover o registro, " + err.Error())
	}
	if _, err := result.RowsAffected(); err == nil {
		return true, nil
	} else {
		return false, errors.New("um erro ocorreu ao remover o registro")
	}

}
