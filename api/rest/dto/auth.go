package dto

// UserLoginRequest struct
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLoginResponse struct
type UserLoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type,omitempty"`
}

// UserChangePasswordRequest struct
type UserChangePasswordRequest struct {
	Password string `json:"password"`
}
