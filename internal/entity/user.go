package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/repository/user"
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
func (u User) ToTable() *user.Table {
	return &user.Table{
		Email:    u.Email,
		Password: u.Password,
	}
}

// UserFromTable func
func UserFromTable(fromTable *user.Table) *User {
	return &User{
		ID:       fromTable.ID,
		Email:    fromTable.Email,
		Password: fromTable.Password,
	}
}
