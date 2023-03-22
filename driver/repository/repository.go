package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository/token"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository/user"
)

// Repository struct
type Repository struct {
	User  user.UserRepository
	Token token.TokenRepository
}

// NewRepository func
func NewRepository() Repository {
	return Repository{
		User:  user.New(),
		Token: token.New(),
	}
}
