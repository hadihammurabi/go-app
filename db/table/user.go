package table

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"-"`
}

// BeforeCreate func
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	u.ID = id
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashedPassword)
	return
}

// BeforeSave func
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashedPassword)
	return
}
