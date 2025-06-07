package main

import "github.com/serge1197/go-simple-api/db"

func main() {
	db.ConnSqlite()
	defer db.Connection.Close()
}
