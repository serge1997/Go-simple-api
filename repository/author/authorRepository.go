package author

import (
	"database/sql"
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
