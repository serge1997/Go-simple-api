package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/db"
	"github.com/serge1197/go-simple-api/repository/author"
	"github.com/serge1197/go-simple-api/services"
)

var ErrEmptyAuthorRequestBody error = errors.New("please inform author body")

func StoreAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services.Write(r.Method + " " + r.URL.Path)
	params := r.Body
	var response services.HttpResponse
	if params == nil {
		w.WriteHeader(http.StatusNoContent)
		response.Code = http.StatusNoContent
		response.Message = ErrEmptyAuthorRequestBody.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response.Code = http.StatusCreated
	response.Message = "author created successfully"
	var author author.Author
	db.ConnSqlite()
	json.NewDecoder(r.Body).Decode(&author)
	author.CreatedAt = time.Now()
	author.UpdatedAt = nil
	id, err := author.Persist(db.Connection)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Code = http.StatusBadRequest
		response.Message = ErrEmptyAuthorRequestBody.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	author.Id = *id
	tobyte, err := json.MarshalIndent(&author, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Code = http.StatusBadRequest
		response.Message = ErrEmptyAuthorRequestBody.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = (*json.RawMessage)(&tobyte)
	json.NewEncoder(w).Encode(response)
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services.Write(r.Method + " " + r.URL.Path)
	var response services.HttpResponse
	params := mux.Vars(r)
	db.ConnSqlite()
	parseId, err := strconv.ParseInt(params["id"], 0, 32)
	if err != nil {
		panic(err)
	}
	id := int(parseId)
	author, err := author.Find(db.Connection, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Code = http.StatusNotFound
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(author)
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response services.HttpResponse
	services.Write(r.Method + " " + r.URL.Path)
	db.ConnSqlite()
	authors, err := author.FindAll(db.Connection)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Code = http.StatusNotFound
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := json.MarshalIndent(authors, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Code = http.StatusNotFound
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Code = http.StatusOK
	response.Message = "List of all author"
	response.Data = (*json.RawMessage)(&data)
	json.NewEncoder(w).Encode(authors)

}
