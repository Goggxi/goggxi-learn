package api

import "time"

type Song struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Album       Album     `json:"album"`
	Duration    float64   `json:"duration"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SongReq struct {
	Name        string    `json:"name" validate:"required"`
	AlbumID     string    `json:"album_id" validate:"required"`
	Duration    float64   `json:"duration" validate:"required"`
	ReleaseDate time.Time `json:"release_date"`
}
