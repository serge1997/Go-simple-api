package models

type Book struct {
	Id       int
	Title    string
	Edition  string
	Year     string
	AuthorId uint `gorm:"column:AuthorId"`
}
