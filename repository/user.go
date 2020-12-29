package repository

import (
	"belajar-go-rest-api/entity"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserRepository struct
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository func
func NewUserRepository(database *gorm.DB) entity.UserRepository {
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
	u.db.Create(&user)
	return user, nil
}

// FindByEmail func
func (u UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	u.db.Where(&entity.User{
		Email: email,
	}).First(&user)
	return user, nil
}

// FindByID func
func (u UserRepository) FindByID(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	u.db.Where("id = ?", id).First(&user)
	return user, nil
}

// ChangePassword func
func (u UserRepository) ChangePassword(id uuid.UUID, password string) (*entity.User, error) {
	user, err := u.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	u.db.Save(user)

	return user, nil
}
