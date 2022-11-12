package repository

// Repository struct
type Repository struct {
	User  UserRepository
	Token TokenRepository
}

// NewRepository func
func NewRepository() Repository {
	return Repository{
		User:  NewUserRepository(),
		Token: NewTokenRepository(),
	}
}
