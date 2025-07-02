package routes

import (
	"net/http"

	"github.com/serge1197/go-simple-api/api"
)

var bookRoutes = []Route{
	{
		Uri:         "/book",
		Method:      http.MethodPost,
		Handle:      api.StoreBook,
		RequireAuth: false,
	},
	{
		Uri:         "/book",
		Method:      http.MethodGet,
		Handle:      func(w http.ResponseWriter, r *http.Request) {},
		RequireAuth: false,
	},
	{
		Uri:         "/book/{id}",
		Method:      http.MethodGet,
		Handle:      func(w http.ResponseWriter, r *http.Request) {},
		RequireAuth: false,
	},
	{
		Uri:         "/book",
		Method:      http.MethodPut,
		Handle:      func(w http.ResponseWriter, r *http.Request) {},
		RequireAuth: false,
	},
	{
		Uri:         "/book/{id}",
		Method:      http.MethodDelete,
		Handle:      func(w http.ResponseWriter, r *http.Request) {},
		RequireAuth: false,
	},
}
