package utils

import "github.com/google/uuid"

func GenerateID() string {
	ids := uuid.New()
	return ids.String()
}
