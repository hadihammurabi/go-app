package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/sarulabs/di"
)

// AuthService struct
type AuthService struct {
	userService  entity.UserService
	tokenService entity.TokenService
	jwtService   entity.JWTService
}

// NewAuthService func
func NewAuthService(ioc di.Container) entity.AuthService {
	return &AuthService{
		userService:  NewUserService(ioc),
		tokenService: NewTokenService(ioc),
		jwtService:   NewJWTService(ioc),
	}
}

// Login func
func (a AuthService) Login(userInput *entity.User) (string, error) {
	user, err := a.userService.FindByEmail(userInput.Email)
	if err != nil {
		return "", err
	}

	err = user.IsPasswordValid(userInput.Password)
	if err != nil {
		return "", err
	}

	token, err := a.jwtService.Create(user)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
