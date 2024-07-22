package tools

import "github.com/google/uuid"

func CreatedId() string {
	return uuid.NewString()
}
