package service

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/repository"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// User service
type User struct {
	userRepository *repository.User
}

// NewUser func
func NewUser(database *gorm.DB) *User {
	return &User{
		userRepository: repository.NewUser(),
	}
}

// All func
func (u User) All() []entities.User {
	return u.userRepository.All()
}

// Create func
func (u User) Create(user *entities.User) *entities.User {
	return u.userRepository.Create(user)
}

// FindByEmail func
func (u User) FindByEmail(email string) (*entities.User, error) {
	return u.userRepository.FindByEmail(email)
}

// FindByID func
func (u User) FindByID(id uuid.UUID) (*entities.User, error) {
	return u.userRepository.FindByID(id)
}

// ChangePassword func
func (u User) ChangePassword(id uuid.UUID, password string) (*entities.User, error) {
	return u.userRepository.ChangePassword(id, password)
}
