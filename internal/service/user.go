package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

// UserService interface
type UserService interface {
	All(context.Context) ([]*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	FindByID(context.Context, uuid.UUID) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	ChangePassword(context.Context, uuid.UUID, string) (*model.User, error)
}

// userService struct
type userService struct {
	repo repository.Repository
}

// NewUserService func
func NewUserService(ioc di.IOC) UserService {
	repo := getRepository(ioc)
	return userService{
		repo,
	}
}

// All func
func (u userService) All(c context.Context) ([]*model.User, error) {
	return u.repo.User.All(c)
}

// Create func
func (u userService) Create(c context.Context, user *model.User) (*model.User, error) {
	return u.repo.User.Create(c, user)
}

// FindByEmail func
func (u userService) FindByEmail(c context.Context, email string) (*model.User, error) {
	return u.repo.User.FindByEmail(c, email)
}

// FindByID func
func (u userService) FindByID(c context.Context, id uuid.UUID) (*model.User, error) {
	return u.repo.User.FindByID(c, id)
}

// ChangePassword func
func (u userService) ChangePassword(c context.Context, id uuid.UUID, password string) (*model.User, error) {
	return u.repo.User.ChangePassword(c, id, password)
}
