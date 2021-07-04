package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/entity"
	"github.com/sarulabs/di"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	All() ([]*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	FindByID(id uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	ChangePassword(id uuid.UUID, password string) (*entity.User, error)
}

// userRepository struct
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository func
func NewUserRepository(ioc di.Container) UserRepository {
	database := getDatabase(ioc)
	repo := &userRepository{
		db: database,
	}
	return repo
}

// All func
func (u userRepository) All() (users []*entity.User, err error) {
	users = make([]*entity.User, 0)
	err = u.db.Find(&users).Error
	return users, err
}

// Create func
func (u userRepository) Create(user *entity.User) (*entity.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

// FindByEmail func
func (u userRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Where(&entity.User{
		Email: email,
	}).First(&user).Error
	return user, err
}

// FindByID func
func (u userRepository) FindByID(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Where("id = ?", id).First(&user).Error
	return user, err
}

// ChangePassword func
func (u userRepository) ChangePassword(id uuid.UUID, password string) (*entity.User, error) {
	user, err := u.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	err = u.db.Save(user).Error

	return user, err
}
