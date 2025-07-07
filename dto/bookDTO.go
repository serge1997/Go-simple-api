package dto

import (
	"github.com/serge1197/go-simple-api/models"
)

type BookDto struct {
	Id      int
	Title   string
	Edition string
	Year    string
	Author  AuthorDto
}

func BookResource(book *models.Book) BookDto {
	dto := BookDto{
		Id:      book.Id,
		Title:   book.Title,
		Edition: book.Edition,
		Author:  AuthorToResource(book.Author),
		Year:    book.Year,
	}
	return dto
}

func BookCollection(books *[]models.Book) *[]BookDto {
	var booksDto []BookDto = []BookDto{}
	for _, book := range *books {
		booksDto = append(booksDto, BookDto{
			Id:      book.Id,
			Title:   book.Title,
			Year:    book.Year,
			Edition: book.Edition,
			Author:  AuthorToResource(book.Author),
		})
	}
	return &booksDto
}
