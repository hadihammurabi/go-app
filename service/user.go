package service

import (
	"belajar-go-rest-api/model"
	"belajar-go-rest-api/repository"

	uuid "github.com/satori/go.uuid"
)

// User service
type User struct {
	userRepository *repository.User
}

// NewUser func
func NewUser() *User {
	return &User{
		userRepository: repository.NewUser(),
	}
}

// All func
func (u User) All() []model.User {
	return u.userRepository.All()
}

// Create func
func (u User) Create(user *model.User) *model.User {
	return u.userRepository.Create(user)
}

// FindByEmail func
func (u User) FindByEmail(email string) (*model.User, error) {
	return u.userRepository.FindByEmail(email)
}

// FindByID func
func (u User) FindByID(id uuid.UUID) (*model.User, error) {
	return u.userRepository.FindByID(id)
}
