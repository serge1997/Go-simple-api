package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/db"
	"github.com/serge1197/go-simple-api/routes"
)

func main() {
	//setup db actions
	db.ConnSqlite()
	migrationErro := db.RunMigrations()
	if migrationErro != nil {
		panic(migrationErro)
	}
	fmt.Println("Migrations done !")
	defer db.Connection.Close()

	//setup router
	router := mux.NewRouter()
	routes.RoutesRegister(router)
	http.ListenAndServe(":3000", router)
}
