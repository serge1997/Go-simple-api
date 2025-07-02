package repository

import (
	"gorm.io/gorm"
)

type AppRepostory struct {
	db *gorm.DB
}

func Init(db *gorm.DB) AppRepostory {
	app := AppRepostory{db}
	return app
}
