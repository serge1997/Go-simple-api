package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/config"
	"github.com/serge1197/go-simple-api/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	//setup router
	r := mux.NewRouter()
	r = router.Congiguration(r)
	port := fmt.Sprintf(":%s", config.APP_PORT)
	log.Fatal(http.ListenAndServe(port, r))
}
