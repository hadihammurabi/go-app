package repository

import (
	"belajar-go-rest-api/entities"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserRepository struct
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository func
func NewUserRepository(database *gorm.DB) (repo *UserRepository) {
	repo = &UserRepository{
		db: database,
	}
	return
}

// All func
func (u UserRepository) All() (users []*entities.User, err error) {
	users = make([]*entities.User, 0)
	err = u.db.Find(&users).Error
	return users, err
}

// Create func
func (u UserRepository) Create(user *entities.User) *entities.User {
	u.db.Create(&user)
	return user
}

// FindByEmail func
func (u UserRepository) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	u.db.Where(&entities.User{
		Email: email,
	}).First(&user)
	return user, nil
}

// FindByID func
func (u UserRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	user := &entities.User{}
	u.db.Where("id = ?", id).First(&user)
	return user, nil
}

// ChangePassword func
func (u UserRepository) ChangePassword(id uuid.UUID, password string) (*entities.User, error) {
	user, err := u.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	u.db.Save(user)

	return user, nil
}
