package book

import "go-rest-api/internal/author"

type Book struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	AuthorID author.Author `json:"author"`
}
