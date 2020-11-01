package repository

import (
	"belajar-go-rest-api/config/database"
	"belajar-go-rest-api/entities"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// User repository
type User struct {
	db *gorm.DB
}

// NewUser func
func NewUser() *User {
	db, _ := database.ConfigureDatabase()
	return &User{
		db: db,
	}
}

// All func
func (u User) All() []entities.User {
	users := []entities.User{}
	u.db.Find(&users)
	return users
}

// Create func
func (u User) Create(user *entities.User) *entities.User {
	u.db.Create(&user)
	return user
}

// FindByEmail func
func (u User) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	u.db.Where(&entities.User{
		Email: email,
	}).First(&user)
	return user, nil
}

// FindByID func
func (u User) FindByID(id uuid.UUID) (*entities.User, error) {
	user := &entities.User{}
	u.db.Where("id = ?", id).First(&user)
	return user, nil
}

// ChangePassword func
func (u User) ChangePassword(id uuid.UUID, password string) (*entities.User, error) {
	user, err := u.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	u.db.Save(user)

	return user, nil
}
