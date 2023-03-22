package entity

import (
	"github.com/google/uuid"
)

// User entity
type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}
