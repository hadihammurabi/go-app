package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
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

func getConfig(ioc di.Container) *config.Config {
	return ioc.Get("config").(*config.Config)
}

func getRepository(ioc di.Container) *repository.Repository {
	return ioc.Get("repository").(*repository.Repository)
}
