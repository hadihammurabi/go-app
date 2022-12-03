package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/db/table"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
	"github.com/hadihammurabi/go-ioc/ioc"

	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	All(context.Context) ([]*table.User, error)
	Create(context.Context, *table.User) (*table.User, error)
	FindByID(context.Context, uuid.UUID) (*table.User, error)
	FindByEmail(context.Context, string) (*table.User, error)
	ChangePassword(context.Context, uuid.UUID, string) (*table.User, error)
}

// userRepository struct
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository func
func NewUserRepository() UserRepository {
	db := ioc.Get(database.Database{}).DB

	repo := &userRepository{
		db: db,
	}
	return repo
}

// All func
func (u userRepository) All(c context.Context) (users []*table.User, err error) {
	users = make([]*table.User, 0)
	err = u.db.WithContext(c).Find(&users).Error
	return users, err
}

// Create func
func (u userRepository) Create(c context.Context, user *table.User) (*table.User, error) {
	err := u.db.WithContext(c).Create(&user).Error
	return user, err
}

// FindByEmail func
func (u userRepository) FindByEmail(c context.Context, email string) (*table.User, error) {
	user := &table.User{}
	err := u.db.WithContext(c).Where(&table.User{
		Email: email,
	}).First(&user).Error
	return user, err
}

// FindByID func
func (u userRepository) FindByID(c context.Context, id uuid.UUID) (*table.User, error) {
	user := &table.User{}
	err := u.db.WithContext(c).Where("id = ?", id).First(&user).Error
	return user, err
}

// ChangePassword func
func (u userRepository) ChangePassword(c context.Context, id uuid.UUID, password string) (*table.User, error) {
	user, err := u.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	err = u.db.WithContext(c).Save(user).Error

	return user, err
}
