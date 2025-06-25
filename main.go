package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/routes"
)

func main() {
	//setup router
	router := mux.NewRouter()
	routes.RoutesRegister(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
