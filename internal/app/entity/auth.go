package entity

// AuthService interface
type AuthService interface {
	Login(userInput *User) (string, error)
}
