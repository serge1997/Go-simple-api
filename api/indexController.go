package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serge1197/go-simple-api/config"
	"github.com/serge1197/go-simple-api/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response services.HttpResponse
	response.Code = http.StatusOK
	response.Message = fmt.Sprintf("%s api is running on p[:%s]", config.APP_NAME, config.APP_PORT)
	json.NewEncoder(w).Encode(response)
	services.Write(r.Method + " " + r.URL.Path)
}
