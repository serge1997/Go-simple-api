package models

import (
	"errors"
	"strings"
)

type User struct {
	Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (user User) ValidateUser() error {
	var ErrEmptyName = errors.New("name is required")
	var ErrEmptyEmail = errors.New("email is required")
	var ErrEmptyPassword = errors.New("password is required")
	if strings.TrimSpace(user.Name) == "" {
		return ErrEmptyName
	}
	if strings.TrimSpace(user.Password) == "" {
		return ErrEmptyPassword
	}
	if strings.TrimSpace(user.Email) == "" {
		return ErrEmptyEmail
	}
	return nil
}
