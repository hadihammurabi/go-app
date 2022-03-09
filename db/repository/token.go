package repository

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/db/table"
	"github.com/hadihammurabi/belajar-go-rest-api/util/di"

	"gorm.io/gorm"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*table.Token) (*table.Token, error)
	FindByUserID(uuid.UUID) (*table.Token, error)
	FindByToken(string) (*table.Token, error)
}

// tokenRepository struct
type tokenRepository struct {
	db *gorm.DB
}

// NewTokenRepository func
func NewTokenRepository(ioc di.IOC) TokenRepository {
	database := getDatabase(ioc)

	return &tokenRepository{
		db: database,
	}
}

// Create func
func (u tokenRepository) Create(token *table.Token) (*table.Token, error) {
	err := u.db.Create(&token).Error
	return token, err
}

// FindByUserID func
func (u tokenRepository) FindByUserID(id uuid.UUID) (*table.Token, error) {
	token := &table.Token{}
	err := u.db.Where("id = ?", id).First(&token).Error
	return token, err
}

// FindByToken func
func (u tokenRepository) FindByToken(token string) (*table.Token, error) {
	tokenDB := &table.Token{}
	err := u.db.Where("token = ?", token).First(&tokenDB).Error
	return tokenDB, err
}
