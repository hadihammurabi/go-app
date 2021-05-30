package entity

// JWTService interface
type JWTService interface {
	Create(*User) (*Token, error)
	GetUser(string) (*User, error)
}
