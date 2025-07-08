package dto

import (
	"time"

	"github.com/serge1197/go-simple-api/models"
)

type UserDTO struct {
	Id        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func UserResource(user *models.User) *UserDTO {
	return nil
}

func UserCollection(users *[]models.User) *[]models.User {
	return nil
}
