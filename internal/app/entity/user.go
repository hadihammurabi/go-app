package entity

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"-"`
}

// BeforeCreate func
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashedPassword)
	return
}

// BeforeSave func
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashedPassword)
	return
}

// IsPasswordValid func
func (u User) IsPasswordValid(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.New("Invalid credentials")
	}

	return nil
}

// UserLoginDTO struct
type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserChangePasswordDTO struct
type UserChangePasswordDTO struct {
	Password string `json:"password"`
}

// UserService interface
type UserService interface {
	All() ([]*User, error)
	Create(user *User) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
	ChangePassword(id uuid.UUID, password string) (*User, error)
}

// UserRepository interface
type UserRepository interface {
	All() ([]*User, error)
	Create(user *User) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
	ChangePassword(id uuid.UUID, password string) (*User, error)
}
