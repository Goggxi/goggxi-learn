package api

import "time"

type Artist struct {
	ID        string    `json:"id" example:"f4ff0500-7dcd-4f28-9e3e-77859a02977"`
	Name      string    `json:"name" example:"John Doe"`
	Bio       string    `json:"bio" example:"Lorem ipsum dolor sit amet"`
	CreatedAt time.Time `json:"created_at" example:"2021-08-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-08-01T00:00:00Z"`
}

type ArtistReq struct {
	Name string `json:"name" validate:"required" example:"John Doe"`
	Bio  string `json:"bio" example:"Lorem ipsum dolor sit amet"`
}
