package service

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/repository"
)

// AuthService struct
type AuthService struct {
	userService entities.UserService
	jwtService  entities.JWTService
}

// NewAuthService func
func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		userService: NewUserService(repo),
		jwtService:  NewJWTService(),
	}
}

// Login func
func (a AuthService) Login(userInput *entities.User) (string, error) {
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

	return token, nil
}
