package service

import (
	"belajar-go-rest-api/model"
	"belajar-go-rest-api/repository"
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
