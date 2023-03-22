package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/driver/util/jwt"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
)

// JWTService interface
type JWTService interface {
	Create(*entity.User) (*entity.Token, error)
	GetUser(context.Context, string) (*entity.User, error)
}

// jwtService struct
type jwtService struct {
	Config       *gowok.Config
	UserService  UserService
	TokenService TokenService
	// Cache        *cache.Redis
}

// NewJWTService func
func NewJWTService() JWTService {
	config := ioc.Get(gowok.Config{})

	return jwtService{
		Config:       config,
		UserService:  NewUserService(),
		TokenService: NewTokenService(),
		// Cache:        config.Redis,
	}
}

// Create func
func (s jwtService) Create(userData *entity.User) (*entity.Token, error) {
	claims := &entity.JWTClaims{
		UserID: userData.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}
	t, err := jwtUtil.CreateJWTWithClaims(s.Config.Security.Secret, claims)
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	// if s.Cache != nil {
	// 	userData.CreatedAt = nil
	// 	s.Cache.Set(stringUtil.ToCacheKey("auth", "token", t), userData, 3*time.Hour)
	// }

	return &entity.Token{
		Token: t,
	}, nil
}

// GetUser func
func (s jwtService) GetUser(c context.Context, token string) (*entity.User, error) {
	tokenData, err := s.TokenService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.UserService.FindByID(c, tokenData.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
