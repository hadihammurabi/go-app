package service

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
)

type TokenService struct {
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(repo *repository.Repository) TokenService {
	return TokenService{
		repo,
	}
}

// Create func
func (u TokenService) Create(token *entity.Token) (*entity.Token, error) {
	tokenFromTable, err := u.repo.Token.Create(token)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}

// FindByUserID func
func (u TokenService) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	tokenFromTable, err := u.repo.Token.FindByUserID(id)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}

// FindByToken func
func (u TokenService) FindByToken(token string) (*entity.Token, error) {
	tokenFromTable, err := u.repo.Token.FindByToken(token)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}
