package api

import "time"

type Song struct {
	ID          string     `json:"id" example:"f4ff0500-7dcd-4f28-9e3e-77859a02977"`
	Title       string     `json:"title" example:"Song Title"`
	Album       Album      `json:"album"`
	Duration    int64      `json:"duration" example:"60"`
	ReleaseDate *time.Time `json:"release_date" example:"2024-07-18T16:19:49.703231Z"`
	CreatedAt   time.Time  `json:"created_at" example:"2024-07-18T16:19:49.703231Z"`
	UpdatedAt   time.Time  `json:"updated_at" example:"2024-07-18T16:19:49.703231Z"`
}

type SongReq struct {
	Title       string     `json:"title" validate:"required" example:"Song Title"`
	AlbumID     string     `json:"album_id" validate:"required" example:"f4ff0500-7dcd-4f28-9e3e-77859a02977"`
	Duration    int64      `json:"duration" validate:"required" example:"60"`
	ReleaseDate *time.Time `json:"release_date,omitempty" example:"2024-07-18T16:19:49.703231Z"`
}
