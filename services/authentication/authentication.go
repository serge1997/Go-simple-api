package authentication

import (
	"net/http"

	"github.com/serge1197/go-simple-api/models"
)

func GenerateToken(user *models.User) (string, error) {
	return "", nil
}

func extractToken(r *http.Request) (string, error) {
	return "", nil
}
