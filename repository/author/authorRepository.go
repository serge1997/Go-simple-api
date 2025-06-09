package author

import (
	"database/sql"
	"errors"
	"time"
)

func (author Author) Persist(db *sql.DB) (*int64, error) {
	var erro error
	var ErrExists = errors.New("author already exists in database")
	exist := FindByName(db, author.Name)
	if exist != nil {
		return nil, ErrExists
	}
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

func FindByName(db *sql.DB, name string) *Author {
	var author Author
	row := db.QueryRow("SELECT * FROM authors WHERE Name = ?", name)
	row.Scan(&author.Id, &author.Name, &author.Website, &author.CreatedAt, &author.UpdatedAt)
	if author.Id == 0 {
		return nil
	}
	return &author
}

func (author Author) Update(db *sql.DB) (*Author, error) {
	var ErrOcurred = errors.New("erro ocorred on updating author")
	finded, err := Find(db, int(author.Id))
	if err != nil {
		return nil, err
	}
	smt, err := db.Prepare("UPDATE authors SET Name = ?, Website = ?, UpdatedAt = ? WHERE Id = ?")
	if err != nil {
		return nil, err
	}
	result, err := smt.Exec(author.Name, author.Website, time.Now(), finded.Id)
	if err != nil {
		return nil, err
	}
	isUpdated, _ := result.RowsAffected()
	if isUpdated >= 1 {
		retrieve, _ := Find(db, int(author.Id))
		return retrieve, nil
	}
	return nil, ErrOcurred
}
