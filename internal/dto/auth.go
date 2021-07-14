package dto

// UserLoginRequest struct
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserChangePasswordRequest struct
type UserChangePasswordRequest struct {
	Password string `json:"password"`
}
