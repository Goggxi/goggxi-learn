package entity

import "time"

type Album struct {
	ID          string
	Title       string
	Genre       string
	ArtistID    string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
