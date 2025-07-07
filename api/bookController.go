package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ShowBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var ErrEmptyID error = errors.New("book id missed")
	json.NewDecoder(r.Body).Decode(&book)
	if book.Id == 0 {
		services.JSONError(w, ErrEmptyID.Error(), nil, http.StatusBadRequest)
		return
	}
	db.ConnSqlite()
	repository := repository.Init(db.Connection)
	result, err := repository.UpdateBook(book)
	if err != nil {
		services.JSONError(w, err.Error(), nil, http.StatusBadRequest)
		return
	}
	dto := dto.BookResource(result)
	services.JSONSuccess(w, "book updated successfully", dto, http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	db.ConnSqlite()
	repository := repository.Init(db.Connection)
	_, err := repository.DeleteBook(id)
	if err != nil {
		services.JSONError(w, err.Error(), nil, http.StatusBadRequest)
		return
	}
	services.JSONSuccess(w, "book removed successfully", nil, http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	db.ConnSqlite()
	repository := repository.Init(db.Connection)
	result, err := repository.UpdateBook(book)
	if err != nil {
		services.JSONError(w, err.Error(), nil, http.StatusBadRequest)
		return
	}
	res := dto.BookResource(result)
	services.JSONSuccess(w, "book updated successfuly", res, http.StatusOK)
}
