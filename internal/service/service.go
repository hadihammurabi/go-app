package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/db/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/util/di"
)

// Service struct
type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
	JWT   JWTService
}

// NewService func
func NewService(ioc di.IOC) (service Service) {
	service = Service{
		Auth:  NewAuthService(ioc),
		User:  NewUserService(ioc),
		Token: NewTokenService(ioc),
		JWT:   NewJWTService(ioc),
	}
	return
}

func getConfig(ioc di.IOC) config.Config {
	return ioc[di.DI_CONFIG].(config.Config)
}

func getRepository(ioc di.IOC) repository.Repository {
	return ioc[di.DI_REPOSITORY].(repository.Repository)
}
