package service

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/repository/table"
	"gorm.io/gorm"
)

type TokenService struct {
	db   *gorm.DB
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(db *gorm.DB, repo *repository.Repository) TokenService {
	return TokenService{
		db,
		repo,
	}
}

// Create func
func (u TokenService) Create(token *entity.Token) (*entity.Token, error) {
	tokenFromTable := table.Token{
		Token:     token.Token,
		UserID:    token.UserID,
		ExpiredAt: token.ExpiredAt,
	}
	err := u.db.Create(&tokenFromTable).Error
	if err != nil {
		return nil, err
	}

	return tokenFromTable.ToEntity(), nil
}

// FindByUserID func
func (u TokenService) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	tokenFromTable := table.Token{}
	err := u.db.First(&tokenFromTable, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return tokenFromTable.ToEntity(), nil
}

func (u TokenService) FindByToken(token string) (*entity.Token, error) {
	tokenFromTable := table.Token{}
	err := u.db.First(&tokenFromTable, "token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return tokenFromTable.ToEntity(), nil
}
