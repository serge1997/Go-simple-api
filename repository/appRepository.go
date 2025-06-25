package repository

import "database/sql"

type AppRepostory struct {
	db *sql.DB
}

func Init(db *sql.DB) AppRepostory {
	app := AppRepostory{db}
	return app
}
