package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge1197/go-simple-api/services/middlewares"
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
		r.HandleFunc(route.Uri, middlewares.LoggerMiddleware(route.Handle)).Methods(route.Method)
	}
}
