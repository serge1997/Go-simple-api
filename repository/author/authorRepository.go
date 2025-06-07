package author

import (
	"database/sql"
	"errors"
)

func (author Author) Persist(db *sql.DB) (Author, error) {
	err := errors.New("")

	return author, err
}
