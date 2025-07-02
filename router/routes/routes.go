package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri         string
	Method      string
	Handle      func(w http.ResponseWriter, r *http.Request)
	RequireAuth bool
}

func RoutesRegister(r *mux.Router) {
	var appRoutes = authorRoutes
	appRoutes = append(appRoutes, homeRoute)
	appRoutes = append(appRoutes, bookRoutes...)
	for _, route := range appRoutes {
		r.HandleFunc(route.Uri, route.Handle).Methods(route.Method)
	}
}
