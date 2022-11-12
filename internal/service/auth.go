package service

import (
	"context"

	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
)

// AuthService interface
type AuthService interface {
	Login(context.Context, *entity.User) (string, error)
}

// authService struct
type authService struct {
	userService  UserService
	tokenService TokenService
	jwtService   JWTService
}

// NewAuthService func
func NewAuthService() AuthService {
	return authService{
		userService:  NewUserService(),
		tokenService: NewTokenService(),
		jwtService:   NewJWTService(),
	}
}

// Login func
func (a authService) Login(c context.Context, userInput *entity.User) (string, error) {
	user, err := a.userService.FindByEmail(c, userInput.Email)
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
