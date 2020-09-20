package repository

import (
	"belajar-go-rest-api/config"
	"belajar-go-rest-api/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// User repository
type User struct {
	db *gorm.DB
}

// NewUser func
func NewUser() *User {
	db, _ := config.ConfigureDatabase()
	return &User{
		db: db,
	}
}

// All func
func (u User) All() []model.User {
	users := []model.User{}
	u.db.Find(&users)
	return users
}

// Create func
func (u User) Create(user *model.User) *model.User {
	u.db.Create(&user)
	return user
}

// FindByEmail func
func (u User) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	u.db.Where(&model.User{
		Email: email,
	}).First(&user)
	return user, nil
}

// FindByID func
func (u User) FindByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	u.db.Where("id = ?", id).First(&user)
	return user, nil
}

// ChangePassword func
func (u User) ChangePassword(id uuid.UUID, password string) (*model.User, error) {
	user, err := u.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	u.db.Save(user)

	return user, nil
}
