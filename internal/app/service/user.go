package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/repository"
	"github.com/sarulabs/di"

	uuid "github.com/satori/go.uuid"
)

// UserService interface
type UserService interface {
	All() ([]*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	FindByID(id uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	ChangePassword(id uuid.UUID, password string) (*entity.User, error)
}

// userService struct
type userService struct {
	repo *repository.Repository
}

// NewUserService func
func NewUserService(ioc di.Container) UserService {
	repo := getRepository(ioc)
	return &userService{
		repo,
	}
}

// All func
func (u userService) All() ([]*entity.User, error) {
	return u.repo.User.All()
}

// Create func
func (u userService) Create(user *entity.User) (*entity.User, error) {
	return u.repo.User.Create(user)
}

// FindByEmail func
func (u userService) FindByEmail(email string) (*entity.User, error) {
	return u.repo.User.FindByEmail(email)
}

// FindByID func
func (u userService) FindByID(id uuid.UUID) (*entity.User, error) {
	return u.repo.User.FindByID(id)
}

// ChangePassword func
func (u userService) ChangePassword(id uuid.UUID, password string) (*entity.User, error) {
	return u.repo.User.ChangePassword(id, password)
}
