package entities

import "time"

type BookAttrs struct {
	ID          string
	Publisher   string
	Pages       int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
