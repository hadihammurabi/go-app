package service

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/repository"

	uuid "github.com/satori/go.uuid"
)

// User service
type User struct {
	repo *repository.Repository
}

// NewUser func
func NewUser(repo *repository.Repository) *User {
	return &User{
		repo: repo,
	}
}

// All func
func (u User) All() []entities.User {
	return u.repo.User.All()
}

// Create func
func (u User) Create(user *entities.User) *entities.User {
	return u.repo.User.Create(user)
}

// FindByEmail func
func (u User) FindByEmail(email string) (*entities.User, error) {
	return u.repo.User.FindByEmail(email)
}

// FindByID func
func (u User) FindByID(id uuid.UUID) (*entities.User, error) {
	return u.repo.User.FindByID(id)
}

// ChangePassword func
func (u User) ChangePassword(id uuid.UUID, password string) (*entities.User, error) {
	return u.repo.User.ChangePassword(id, password)
}
