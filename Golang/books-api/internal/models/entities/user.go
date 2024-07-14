package entities

import "time"

type User struct {
	ID        string
	Username  string
	Password  string // hash password with bcrypt
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
