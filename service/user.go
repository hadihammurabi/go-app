package service

import (
	"belajar-go-rest-api/entity"
	"belajar-go-rest-api/repository"

	uuid "github.com/satori/go.uuid"
)

// UserService struct
type UserService struct {
	repo *repository.Repository
}

// NewUserService func
func NewUserService(repo *repository.Repository) entity.UserService {
	return &UserService{
		repo: repo,
	}
}

// All func
func (u UserService) All() ([]*entity.User, error) {
	return u.repo.User.All()
}

// Create func
func (u UserService) Create(user *entity.User) (*entity.User, error) {
	return u.repo.User.Create(user)
}

// FindByEmail func
func (u UserService) FindByEmail(email string) (*entity.User, error) {
	return u.repo.User.FindByEmail(email)
}

// FindByID func
func (u UserService) FindByID(id uuid.UUID) (*entity.User, error) {
	return u.repo.User.FindByID(id)
}

// ChangePassword func
func (u UserService) ChangePassword(id uuid.UUID, password string) (*entity.User, error) {
	return u.repo.User.ChangePassword(id, password)
}
