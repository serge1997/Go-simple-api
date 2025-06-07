package author

import "github.com/serge1197/go-simple-api/repository"

type Author struct {
	repository.Model
	Name    string
	Website *string
}
