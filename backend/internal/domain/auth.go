package domain

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type AuthService interface {
	Login(email, password string) (*LoginResponse, error)
	ValidateToken(token string) (*User, error)
}