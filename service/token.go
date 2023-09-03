package service

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
)

// TokenService interface
type TokenService interface {
	Create(*entity.Token) (*entity.Token, error)
	FindByUserID(uuid.UUID) (*entity.Token, error)
	FindByToken(string) (*entity.Token, error)
}

// tokenService struct
type tokenService struct {
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(repo *repository.Repository) TokenService {
	return tokenService{
		repo,
	}
}

// Create func
func (u tokenService) Create(token *entity.Token) (*entity.Token, error) {
	tokenFromTable, err := u.repo.Token.Create(token)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}

// FindByUserID func
func (u tokenService) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	tokenFromTable, err := u.repo.Token.FindByUserID(id)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}

// FindByToken func
func (u tokenService) FindByToken(token string) (*entity.Token, error) {
	tokenFromTable, err := u.repo.Token.FindByToken(token)
	if err != nil {
		return nil, err
	}

	return tokenFromTable, nil
}
