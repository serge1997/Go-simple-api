package models

type Book struct {
	Id      int
	Title   string
	Edition string
	Year    string
	Author  Author
}
