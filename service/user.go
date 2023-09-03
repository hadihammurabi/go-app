package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/hash"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/repository/table"
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

func (u UserService) All(c context.Context) (users []*entity.User, err error) {
	usersFromTable := []table.User{}
	err = u.db.Find(&usersFromTable).Error
	if err != nil {
		return nil, err
	}

	for _, uft := range usersFromTable {
		users = append(users, uft.ToEntity())
	}

	return users, nil
}

// Create func
func (u UserService) Create(c context.Context, user *entity.User) (*entity.User, error) {
	userFromTable := table.User{
		Email:    user.Email,
		Password: user.Password,
	}
	err := u.db.Create(&userFromTable).Error
	if err != nil {
		return nil, err
	}

	return userFromTable.ToEntity(), nil
}

// FindByEmail func
func (u UserService) FindByEmail(c context.Context, email string) (*entity.User, error) {
	userFromTable := table.User{}
	err := u.db.First(&userFromTable, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return userFromTable.ToEntity(), nil
}

// FindByID func
func (u UserService) FindByID(c context.Context, id uuid.UUID) (*entity.User, error) {
	userFromTable := table.User{}
	err := u.db.First(&userFromTable, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return userFromTable.ToEntity(), nil
}

// ChangePassword func
func (u UserService) ChangePassword(c context.Context, id uuid.UUID, password string) (*entity.User, error) {
	user, err := u.FindByID(c, id)
	if err != nil {
		return nil, err
	}
	pass := hash.PasswordHash(password, u.config.App.Key)
	user.Password = pass.Hashed
	err = u.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
