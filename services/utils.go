package services

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONSuccess(w http.ResponseWriter, message string, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	var response HttpResponse
	if code == 0 {
		response.Code = 200
	}
	response.Code = code
	response.Message = message
	response.Data = data
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
func JSONError(w http.ResponseWriter, message string, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	var response HttpResponse
	if code == 0 {
		response.Code = 501
	}
	response.Code = code
	response.Message = message
	response.Data = data
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
