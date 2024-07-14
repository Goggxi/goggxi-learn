package entities

import "time"

type Author struct {
	ID        string
	Name      string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
