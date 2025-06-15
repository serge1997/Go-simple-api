package dto

type BookDto struct {
	Id      int
	Titulo  string
	Edition string
	Year    string
	Author  AuthorDto
}
