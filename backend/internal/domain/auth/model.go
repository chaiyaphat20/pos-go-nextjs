package auth

import (
	"go-clean/internal/domain/user"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string        `json:"username" validate:"required,min=3,max=50"`
	Email    string        `json:"email" validate:"required,email"`
	Password string        `json:"password" validate:"required,min=6"`
	Role     user.UserRole `json:"role"`
}

type AuthResponse struct {
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	User         *user.User `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type TokenClaims struct {
	UserID string        `json:"user_id"`
	Email  string        `json:"email"`
	Role   user.UserRole `json:"role"`
}