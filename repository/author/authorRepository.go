package author

import (
	"database/sql"
	"errors"
)

func (author Author) Persist(db *sql.DB) (*int64, error) {
	var erro error
	smt, err := db.Prepare("INSERT INTO authors(Name, Website, CreatedAt) VALUES(?, ?, ?)")

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

func Find(db *sql.DB, id int) (*Author, error) {
	var author Author
	var ErrNotFound = errors.New("Author not Found")
	smt := db.QueryRow("SELECT * FROM authors WHERE Id = ?", id)
	err := smt.Scan(&author.Id, &author.Name, &author.Website, &author.CreatedAt, &author.UpdatedAt)
	if err == sql.ErrNoRows || author.Name == "" {
		return nil, ErrNotFound
	}
	return &author, nil
}

func FindAll(db *sql.DB) (*[]Author, error) {
	var authors []Author
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author Author

		rows.Scan(&author.Id, &author.Name, &author.Website, &author.CreatedAt, &author.UpdatedAt)
		authors = append(authors, author)
	}
	return &authors, nil
}
