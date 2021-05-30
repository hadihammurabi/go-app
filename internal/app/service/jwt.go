package service

import (
	"errors"
	"time"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/entity"
	"github.com/sarulabs/di"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTService struct
type JWTService struct {
	Config       *config.JWTConfig
	UserService  entity.UserService
	TokenService entity.TokenService
}

// NewJWTService func
func NewJWTService(ioc di.Container) entity.JWTService {
	jwtConfig := ioc.Get("config").(*(config.Config)).JWT

	return &JWTService{
		Config:       jwtConfig,
		UserService:  NewUserService(ioc),
		TokenService: NewTokenService(ioc),
	}
}

// Create func
func (jwtService JWTService) Create(userData *entity.User) (*entity.Token, error) {
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = time.Now().Add(time.Hour * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtService.Config.Secret))
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	expToTime := time.Unix(claims.ExpiresAt, 0)
	tokenCreated, err := jwtService.TokenService.Create(&entity.Token{
		UserID:    userData.ID,
		Token:     t,
		ExpiredAt: &expToTime,
	})
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	return tokenCreated, nil
}

// GetUser func
func (jwtService JWTService) GetUser(token string) (*entity.User, error) {
	tokenData, err := jwtService.TokenService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := jwtService.UserService.FindByID(tokenData.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
