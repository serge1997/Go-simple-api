package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/api"
	"github.com/serge1197/go-simple-api/router/routes"
	"github.com/serge1197/go-simple-api/services"
)

func Congiguration(r *mux.Router) *mux.Router {
	fmt.Println("Server is running [:3000]")
	services.Write("Server is running [:3000]")
	r.HandleFunc("/", api.Home).Methods("GET")
	routes.RoutesRegister(r)
	return r
}
