package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serge1197/go-simple-api/db"
	"github.com/serge1197/go-simple-api/dto"
	"github.com/serge1197/go-simple-api/models"
	"github.com/serge1197/go-simple-api/repository"
	"github.com/serge1197/go-simple-api/services"
)

func StoreBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services.Write(r.Method + " " + r.URL.Path)
	params := r.Body
	fmt.Println(params)
	var book models.Book

	db.ConnSqlite()
	defer db.Close()
	json.NewDecoder(r.Body).Decode(&book)
	repository := repository.Init(db.Connection)
	result, err := repository.PersistBook(&book)
	if err != nil {
		services.JSONError(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	response := dto.BookResource(result)
	services.JSONSuccess(w, "book created", response, http.StatusCreated)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db.ConnSqlite()
	repository := repository.Init(db.Connection)
	books, err := repository.FindAllBook()
	if err != nil {
		services.JSONError(w, err.Error(), nil, http.StatusInternalServerError)
	}
	collection := dto.BookCollection(books)
	services.JSONSuccess(w, "list of all book", collection, http.StatusOK)
}
