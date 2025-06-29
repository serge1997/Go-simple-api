package routes

import (
	"net/http"

	"github.com/serge1197/go-simple-api/api"
)

var authorRoutes = []Route{
	{
		Uri:         "/author",
		Method:      http.MethodPost,
		Handle:      api.StoreAuthor,
		RequireAuth: false,
	},
	{
		Uri:         "/author",
		Method:      http.MethodGet,
		Handle:      api.ListAll,
		RequireAuth: false,
	},
	{
		Uri:         "/author/{id}",
		Method:      http.MethodGet,
		Handle:      api.Show,
		RequireAuth: false,
	},
	{
		Uri:         "/author",
		Method:      http.MethodPut,
		Handle:      api.UpdateAuthor,
		RequireAuth: false,
	},
	{
		Uri:         "/author",
		Method:      http.MethodDelete,
		Handle:      api.DeleteAuthor,
		RequireAuth: false,
	},
}
