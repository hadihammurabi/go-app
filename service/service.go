package service

import (
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/exception"
	"github.com/gowok/gowok/singleton"
	"github.com/gowok/plugins/gorm"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
)

// Service struct
type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
	// JWT   JWTService
}

// NewService func
func NewService() Service {
	config := gowok.Get().Config
	sql := gorm.DB().OrPanic(exception.ErrNoDatabaseFound)
	repo := repository.Get()

	service := Service{
		Auth:  NewAuthService(config, sql, repo),
		User:  NewUserService(config, sql, repo),
		Token: NewTokenService(sql, repo),
		// JWT:   NewJWTService(),
	}

	return service
}

var Get = singleton.New(NewService)
