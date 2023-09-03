package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"gorm.io/gorm"
)

type JwtService struct {
	config       *gowok.Config
	userService  UserService
	tokenService TokenService
	// Cache        *cache.Redis
}

// NewJWTService func
func NewJWTService(config *gowok.Config, db *gorm.DB, repo *repository.Repository) JwtService {
	return JwtService{
		config:       config,
		userService:  NewUserService(config, db, repo),
		tokenService: NewTokenService(db, repo),
		// Cache:        config.Redis,
	}
}

// Create func
func (s JwtService) Create(userData *entity.User) (*entity.Token, error) {
	claims := &entity.JWTClaims{
		UserID: userData.ID,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	jwk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := jwk.SignedString([]byte(s.config.Security.Secret))
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
func (s JwtService) GetUser(c context.Context, token string) (*entity.User, error) {
	tokenData, err := s.tokenService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.FindByID(c, tokenData.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
