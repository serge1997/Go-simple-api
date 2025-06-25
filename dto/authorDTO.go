package dto

import "github.com/serge1197/go-simple-api/repository/author"

type AuthorDto struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Website   *string `json:"website"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"upated_at"`
}

func AuthorsCollection(authors []author.Author) []AuthorDto {
	var result []AuthorDto
	for _, author := range authors {
		create_at := author.CreatedAt.Format("02-01-2006 15:04:15")
		var updated_at string
		if author.UpdatedAt != nil {
			updated_at = author.UpdatedAt.Format("02-01-2006 15:04:15")
		}
		var dto = AuthorDto{Id: author.Id, Name: author.Name, Website: author.Website, CreatedAt: create_at, UpdatedAt: &updated_at}
		result = append(result, dto)
	}
	return result
}

func AuthorToResource(author author.Author) AuthorDto {
	create_at := author.CreatedAt.Format("02-01-2006 15:04:15")
	var updated_at string
	if author.UpdatedAt != nil {
		updated_at = author.UpdatedAt.Format("02-01-2006 15:04:15")
	}
	return AuthorDto{
		Id:        author.Id,
		Name:      author.Name,
		Website:   author.Website,
		CreatedAt: create_at,
		UpdatedAt: &updated_at,
	}
}
