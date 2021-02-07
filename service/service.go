package service

import (
	"belajar-go-rest-api/config"
	"belajar-go-rest-api/entity"
	"belajar-go-rest-api/repository"
)

// Service struct
type Service struct {
	Auth  entity.AuthService
	User  entity.UserService
	Token entity.TokenService
	JWT   entity.JWTService
}

// NewService func
func NewService(repo *repository.Repository) (service *Service) {
	service = &Service{
		Auth:  NewAuthService(repo),
		User:  NewUserService(repo),
		Token: NewTokenService(repo),
		JWT:   NewJWTService(config.ConfigureJWT()),
	}
	return
}
