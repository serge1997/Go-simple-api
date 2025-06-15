package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serge1197/go-simple-api/repository/book"
	"github.com/serge1197/go-simple-api/services"
)

func StoreBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services.Write(r.Method + " " + r.URL.Path)
	params := r.Body
	fmt.Println(params)
	var response services.HttpResponse
	var book book.Book

	json.NewDecoder(r.Body).Decode(&book)
	if book.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "Informe book body data"
		response.Code = http.StatusBadRequest
		json.NewEncoder(w).Encode(response)
		return
	}
	//id, _ := strconv.ParseInt(params["authorId"], 0, 64)
	//author, err := author.Find()
}
