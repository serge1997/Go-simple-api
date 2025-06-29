package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/router"
)

func main() {
	//setup router
	r := mux.NewRouter()
	r = router.Congiguration(r)
	log.Fatal(http.ListenAndServe(":3000", r))
}
