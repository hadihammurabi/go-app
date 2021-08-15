package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"

	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	All(context.Context) ([]*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	FindByID(context.Context, uuid.UUID) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	ChangePassword(context.Context, uuid.UUID, string) (*model.User, error)
}

// userRepository struct
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository func
func NewUserRepository(ioc di.IOC) UserRepository {
	database := getDatabase(ioc)
	repo := &userRepository{
		db: database,
	}
	return repo
}

// All func
func (u userRepository) All(c context.Context) (users []*model.User, err error) {
	users = make([]*model.User, 0)
	err = u.db.WithContext(c).Find(&users).Error
	return users, err
}

// Create func
func (u userRepository) Create(c context.Context, user *model.User) (*model.User, error) {
	err := u.db.WithContext(c).Create(&user).Error
	return user, err
}

// FindByEmail func
func (u userRepository) FindByEmail(c context.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := u.db.WithContext(c).Where(&model.User{
		Email: email,
	}).First(&user).Error
	return user, err
}

// FindByID func
func (u userRepository) FindByID(c context.Context, id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := u.db.WithContext(c).Where("id = ?", id).First(&user).Error
	return user, err
}

// ChangePassword func
func (u userRepository) ChangePassword(c context.Context, id uuid.UUID, password string) (*model.User, error) {
	user, err := u.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	err = u.db.WithContext(c).Save(user).Error

	return user, err
}
