package api

import (
	"encoding/json"
	"net/http"

	"github.com/serge1197/go-simple-api/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response services.HttpResponse
	response.Code = http.StatusOK
	response.Message = "Go simple api is running"
	json.NewEncoder(w).Encode(response)
	services.Write(r.Method + " " + r.URL.Path)
}
