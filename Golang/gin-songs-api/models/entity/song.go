package entity

import "time"

type Song struct {
	ID          string
	Name        string
	AlbumID     string
	Duration    float64
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
