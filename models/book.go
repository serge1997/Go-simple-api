package models

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Edition  string `json:"edition"`
	Year     string `json:"year"`
	AuthorId uint   `gorm:"column:AuthorId" json:"author_id"`
	Author   Author `gorm:"foreignKey:AuthorId" json:"author"`
}
