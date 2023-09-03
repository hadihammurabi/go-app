package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/hash"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"gorm.io/gorm"
)

type UserService struct {
	config *gowok.Config
	db     *gorm.DB
	repo   *repository.Repository
}

// NewUserService func
func NewUserService(config *gowok.Config, db *gorm.DB, repo *repository.Repository) UserService {
	return UserService{
		config,
		db,
		repo,
	}
}

// All func
func (u UserService) All(c context.Context) (users []*entity.User, err error) {
	usersFromTable, err := u.repo.User.All(c)
	if err != nil {
		return nil, err
	}

	for _, uft := range usersFromTable {
		users = append(users, uft)
	}

	return users, nil
}

// Create func
func (u UserService) Create(c context.Context, user *entity.User) (*entity.User, error) {
	userFromTable, err := u.repo.User.Create(c, user)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}

// FindByEmail func
func (u UserService) FindByEmail(c context.Context, email string) (*entity.User, error) {
	userFromTable, err := u.repo.User.FindByEmail(c, email)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}

// FindByID func
func (u UserService) FindByID(c context.Context, id uuid.UUID) (*entity.User, error) {
	userFromTable, err := u.repo.User.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}

// ChangePassword func
func (u UserService) ChangePassword(c context.Context, id uuid.UUID, password string) (*entity.User, error) {
	pass := hash.PasswordHash(password, u.config.App.Key)
	userFromTable, err := u.repo.User.ChangePassword(c, id, pass.Hashed)
	if err != nil {
		return nil, err
	}

	return userFromTable, nil
}
