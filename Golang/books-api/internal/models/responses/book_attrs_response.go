package responses

import "time"

type BookAttrsResponse struct {
	ID          string    `json:"id"`
	Publisher   string    `json:"publisher"`
	Pages       int       `json:"pages"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
