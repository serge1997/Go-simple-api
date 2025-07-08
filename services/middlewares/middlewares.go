package middlewares

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s \n", r.Method, r.URL.Path, r.Host)
		next(w, r)
	}
}
