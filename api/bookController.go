package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serge1197/go-simple-api/db"
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

	json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	repository := repository.Init(db.Connection)
	result, err := repository.PersistBook(&book)
	if err != nil {
		services.JSONError(w, err.Error(), nil, http.StatusInternalServerError)
		return
	}
	services.JSONSuccess(w, "book created", result, http.StatusCreated)
	//id, _ := strconv.ParseInt(params["authorId"], 0, 64)
	//author, err := author.Find()
}
