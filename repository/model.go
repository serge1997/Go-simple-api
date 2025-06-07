package repository

import "time"

type Model struct {
	Id        int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
