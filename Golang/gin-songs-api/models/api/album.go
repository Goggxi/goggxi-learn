package api

import "time"

type Album struct {
	ID          string     `json:"id" example:"f4ff0500-7dcd-4f28-9e3e-77859a02977"`
	Title       string     `json:"title" example:"Album Title"`
	Genre       string     `json:"genre" example:"Pop"`
	Artist      Artist     `json:"artist"`
	ReleaseDate *time.Time `json:"release_date" example:"2021-08-01T00:00:00Z"`
	CreatedAt   time.Time  `json:"created_at" example:"2024-07-18T16:19:49.703231Z"`
	UpdatedAt   time.Time  `json:"updated_at" example:"2024-07-18T16:19:49.703231Z"`
}

type AlbumReq struct {
	Title       string     `json:"title" validate:"required" example:"Album Title"`
	Genre       string     `json:"genre" validate:"required" example:"Pop"`
	ArtistID    string     `json:"artist_id" validate:"required" example:"f4ff0500-7dcd-4f28-9e3e-77859a02977"`
	ReleaseDate *time.Time `json:"release_date,omitempty" example:"2024-07-18T16:19:49.703231Z"`
}
