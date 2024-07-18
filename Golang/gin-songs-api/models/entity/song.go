package entity

import "time"

type Song struct {
	ID          string
	Title       string
	Album       Album
	Duration    int64
	ReleaseDate *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
