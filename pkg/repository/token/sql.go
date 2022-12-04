package token

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

// tokenSQL struct
type tokenSQL struct {
	db *gorm.DB
}

// newSQL func
func newSQL(db *gorm.DB) *tokenSQL {
	return &tokenSQL{
		db: db,
	}
}

// Create func
func (u tokenSQL) Create(token *Table) (*Table, error) {
	err := u.db.Create(&token).Error
	return token, err
}

// FindByUserID func
func (u tokenSQL) FindByUserID(id uuid.UUID) (*Table, error) {
	token := &Table{}
	err := u.db.Where("id = ?", id).First(&token).Error
	return token, err
}

// FindByToken func
func (u tokenSQL) FindByToken(token string) (*Table, error) {
	tokenDB := &Table{}
	err := u.db.Where("token = ?", token).First(&tokenDB).Error
	return tokenDB, err
}
