package api

import "time"

type Album struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	Artist      Artist    `json:"artist"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AlbumReq struct {
	Title       string    `json:"title" validate:"required"`
	Genre       string    `json:"genre" validate:"required"`
	ArtistID    string    `json:"artist_id" validate:"required"`
	ReleaseDate time.Time `json:"release_date"`
}
