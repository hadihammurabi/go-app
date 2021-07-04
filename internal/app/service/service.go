package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/repository"
	"github.com/sarulabs/di"
)

// Service struct
type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
	JWT   JWTService
}

// NewService func
func NewService(ioc di.Container) (service *Service) {
	service = &Service{
		Auth:  NewAuthService(ioc),
		User:  NewUserService(ioc),
		Token: NewTokenService(ioc),
		JWT:   NewJWTService(ioc),
	}
	return
}

func getRepository(ioc di.Container) *repository.Repository {
	return ioc.Get("repository").(*repository.Repository)
}
