package service

import (
	"belajar-go-rest-api/repository"
)

// Service struct
type Service struct {
	Auth *AuthService
	User *UserService
	JWT  *JWTService
}

// NewService func
func NewService(repo *repository.Repository) (service *Service) {
	service = &Service{
		Auth: NewAuthService(repo),
		User: NewUserService(repo),
		JWT:  NewJWTService(repo),
	}
	return
}
