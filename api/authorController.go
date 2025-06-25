package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/db"
	"github.com/serge1197/go-simple-api/dto"
	"github.com/serge1197/go-simple-api/repository/author"
	"github.com/serge1197/go-simple-api/services"
)

var ErrEmptyAuthorRequestBody error = errors.New("please inform author body")

func StoreAuthor(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	var author author.Author
	db.ConnSqlite()
	json.NewDecoder(r.Body).Decode(&author)
	author.CreatedAt = time.Now()
	author.UpdatedAt = nil
	id, err := author.Persist(db.Connection)
	if err != nil {
		msg := err.Error()
		services.JSONError(w, msg, nil, http.StatusBadRequest)
		return
	}
	author.Id = *id
	data := dto.AuthorToResource(author)

	message := "author created successfully"
	services.JSONSuccess(w, message, data, http.StatusCreated)
}

func Show(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	params := mux.Vars(r)
	db.ConnSqlite()
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		er := err.Error()
		services.JSONError(w, er, nil, http.StatusBadRequest)
		return
	}
	author, err := author.Find(db.Connection, id)
	if err != nil {
		er := err.Error()
		services.JSONError(w, er, nil, http.StatusNotFound)
		return
	}
	authorDto := dto.AuthorToResource(*author)
	message := "Showing an author"
	services.JSONSuccess(w, message, authorDto, http.StatusOK)
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	db.ConnSqlite()
	defer db.Connection.Close()
	authors, err := author.FindAll(db.Connection)
	if err != nil {
		msg := err.Error()
		services.JSONSuccess(w, msg, nil, 404)
		return
	}

	collection := dto.AuthorsCollection(*authors)
	message := "List of all author"
	services.JSONSuccess(w, message, collection, 200)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	var author author.Author
	db.ConnSqlite()
	json.NewDecoder(r.Body).Decode(&author)
	result, err := author.Update(db.Connection)
	if err != nil {
		message := err.Error()
		services.JSONError(w, message, nil, 501)
		return
	}
	dto := dto.AuthorToResource(*result)
	message := "Author updated successfully"
	services.JSONSuccess(w, message, dto, 200)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	params := mux.Vars(r)
	parseId, err := strconv.Atoi(params["id"])
	if err != nil {
		message := fmt.Sprintf("Parametro %s Invalido: %v", params["id"], err.Error())
		services.JSONError(w, message, nil, http.StatusBadRequest)
		return
	}
	db.ConnSqlite()
	id := int(parseId)
	author, err := author.Find(db.Connection, id)
	if err != nil {
		message := err.Error()
		services.JSONError(w, message, nil, http.StatusNotFound)
		return
	}
	_, errr := author.Delete(db.Connection)
	if errr != nil {
		message := err.Error()
		services.JSONError(w, message, nil, http.StatusNotFound)
		return
	}
	message := "author removed successfully"
	services.JSONSuccess(w, message, nil, 200)
}
