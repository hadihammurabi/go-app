package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/db/table"
	"golang.org/x/crypto/bcrypt"
)

// User entity
type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

// IsPasswordValid func
func (u User) IsPasswordValid(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}

// ToTable func
func (u User) ToTable() *table.User {
	return &table.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

// UserFromTable func
func UserFromTable(fromTable *table.User) *User {
	return &User{
		ID:       fromTable.ID,
		Email:    fromTable.Email,
		Password: fromTable.Password,
	}
}
