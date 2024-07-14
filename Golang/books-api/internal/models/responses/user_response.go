package responses

import "time"

type UserResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Username  string    `json:"username" example:"John Doe"`
	FullName  string    `json:"full_name" example:"johndoe"`
	CreatedAt time.Time `json:"created_at" example:"2023-07-14T12:34:56Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-07-14T12:34:56Z"`
}

type UserTokenResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
