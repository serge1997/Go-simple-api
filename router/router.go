package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/config"
	"github.com/serge1197/go-simple-api/router/routes"
	"github.com/serge1197/go-simple-api/services"
)

func Congiguration(r *mux.Router) *mux.Router {
	port := config.APP_PORT
	s := fmt.Sprintf("Server is running [:%s]", port)
	fmt.Println(s)
	services.Write(s)
	routes.RoutesRegister(r)
	return r
}
