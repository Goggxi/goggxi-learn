package entities

import "time"

type Book struct {
	ID          string
	Title       string
	UserID      string
	AuthorID    string
	BookAttrsID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
