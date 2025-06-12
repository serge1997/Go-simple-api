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

	var author author.Author
	db.ConnSqlite()
	json.NewDecoder(r.Body).Decode(&author)
	author.CreatedAt = time.Now()
	author.UpdatedAt = nil
	id, err := author.Persist(db.Connection)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	author.Id = *id
	dto := dto.AuthorToResource(author)
	tobyte, err := json.MarshalIndent(&dto, "", "\t")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Code = http.StatusBadRequest
		response.Message = ErrEmptyAuthorRequestBody.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response.Code = http.StatusCreated
	response.Message = "author created successfully"
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
	authorDto := dto.AuthorToResource(*author)
	data, err := json.MarshalIndent(authorDto, "", "\t")
	if err != nil {
		panic(err)
	}
	response.Message = "Showing an author"
	response.Code = http.StatusOK
	response.Data = (*json.RawMessage)(&data)
	json.NewEncoder(w).Encode(response)
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
	collection := dto.AuthorsCollection(*authors)
	data, err := json.MarshalIndent(collection, "", "\t")
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
	json.NewEncoder(w).Encode(response)

}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	var response services.HttpResponse
	var author author.Author

	params := r.Body
	if params == nil {
		w.WriteHeader(http.StatusNoContent)
		response.Code = http.StatusNoContent
		response.Message = "Please informe a author data"
		json.NewEncoder(w).Encode(response)
		return
	}
	db.ConnSqlite()
	json.NewDecoder(r.Body).Decode(&author)
	result, err := author.Update(db.Connection)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		response.Code = http.StatusConflict
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	dto := dto.AuthorToResource(*result)
	data, _ := json.MarshalIndent(dto, "", "\t")
	w.WriteHeader(http.StatusOK)
	response.Code = http.StatusOK
	response.Message = "Author updated successfully"
	response.Data = (*json.RawMessage)(&data)

	json.NewEncoder(w).Encode(response)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	services.Write(r.Method + " " + r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	var response services.HttpResponse
	params := mux.Vars(r)
	parseId, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = fmt.Sprintf("Parametro %s Invalido: %v", params["id"], err.Error())
		response.Code = http.StatusBadRequest
		json.NewEncoder(w).Encode(response)
		return
	}
	db.ConnSqlite()
	id := int(parseId)
	author, err := author.Find(db.Connection, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = fmt.Sprintf("Erro : %s", err.Error())
		response.Code = http.StatusBadRequest
		json.NewEncoder(w).Encode(response)
		return
	}
	_, errr := author.Delete(db.Connection)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = fmt.Sprintf("Erro : %s", errr.Error())
		response.Code = http.StatusBadRequest
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response.Message = fmt.Sprintf("%s removido com successo", author.Name)
	response.Code = http.StatusOK
	json.NewEncoder(w).Encode(response)
}
