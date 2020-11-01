package service

import (
	"belajar-go-rest-api/entities"

	"gorm.io/gorm"
)

// Auth service
type Auth struct {
	userService *User
	jwtService  *JWT
}

// NewAuth func
func NewAuth(database *gorm.DB) *Auth {
	return &Auth{
		userService: NewUser(database),
		jwtService:  NewJWT(database),
	}
}

// Login func
func (a Auth) Login(userInput *entities.User) (string, error) {
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
