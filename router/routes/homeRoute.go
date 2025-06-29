package routes

import (
	"net/http"

	"github.com/serge1197/go-simple-api/api"
)

var homeRoute = Route{
	Uri:    "/",
	Method: http.MethodGet,
	Handle: api.Home,
}
