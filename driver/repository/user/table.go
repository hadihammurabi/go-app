package user

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository/base"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Table model
type Table struct {
	base.Table
	Email    string `json:"email"`
	Password string `json:"-"`
}

// TableName func
func (u *Table) TableName(tx *gorm.DB) string {
	return "users"
}

// BeforeCreate func
func (u *Table) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	u.ID = id
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashedPassword)
	return
}

// BeforeSave func
func (u *Table) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hashedPassword)
	return
}

// ToEntity func
func (t Table) ToEntity() *entity.User {
	return &entity.User{
		ID:       t.ID,
		Email:    t.Email,
		Password: t.Password,
	}
}

// FromEntity func
func FromEntity(e *entity.User) *Table {
	return &Table{
		Email:    e.Email,
		Password: e.Password,
	}
}
