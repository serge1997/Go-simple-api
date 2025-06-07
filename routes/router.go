package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/api"
	"github.com/serge1197/go-simple-api/services"
)

func RoutesRegister(r *mux.Router) {
	fmt.Println("Server is running [:3000]")
	services.Write("Server is running [:3000]")
	r.HandleFunc("/", api.Home)

}
