package api

import "time"

type Artist struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArtistReq struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio"`
}
