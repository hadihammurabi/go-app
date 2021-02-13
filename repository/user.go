package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/sarulabs/di"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserRepository struct
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository func
func NewUserRepository(ioc di.Container) entity.UserRepository {
	database := getDatabase(ioc)
	repo := &UserRepository{
		db: database,
	}
	return repo
}

// All func
func (u UserRepository) All() (users []*entity.User, err error) {
	users = make([]*entity.User, 0)
	err = u.db.Find(&users).Error
	return users, err
}

// Create func
func (u UserRepository) Create(user *entity.User) (*entity.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

// FindByEmail func
func (u UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Where(&entity.User{
		Email: email,
	}).First(&user).Error
	return user, err
}

// FindByID func
func (u UserRepository) FindByID(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Where("id = ?", id).First(&user).Error
	return user, err
}

// ChangePassword func
func (u UserRepository) ChangePassword(id uuid.UUID, password string) (*entity.User, error) {
	user, err := u.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	err = u.db.Save(user).Error

	return user, err
}
