package responses

import "time"

type BookResponse struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	UserID    string            `json:"user_id"`
	Author    AuthorResponse    `json:"author"`
	BookAttrs BookAttrsResponse `json:"book_attrs"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
