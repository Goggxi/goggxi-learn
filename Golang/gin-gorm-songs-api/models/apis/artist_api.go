package apis

import (
	"time"
)

type ArtistRes struct {
	ID        string    `json:"id" example:"uuid"`
	Name      string    `json:"name" example:"John Doe"`
	Bio       string    `json:"bio" example:"Lorem ipsum dolor sit amet"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01T00:00:00Z"`
}

type ArtistReq struct {
	Name string `json:"name" example:"John Doe"`
	Bio  string `json:"bio" example:"Lorem ipsum dolor sit amet"`
}
