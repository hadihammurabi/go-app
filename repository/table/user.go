package table

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"-"`
}

// UserName func
func (u *User) TableName(tx *gorm.DB) string {
	return "users"
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

// ToEntity func
func (t User) ToEntity() *entity.User {
	return &entity.User{
		ID:       t.ID,
		Email:    t.Email,
		Password: t.Password,
	}
}
