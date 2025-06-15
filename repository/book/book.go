package book

import "github.com/serge1197/go-simple-api/repository/author"

type Book struct {
	Id      int
	Title   string
	Edition string
	Year    string
	Author  author.Author
}
