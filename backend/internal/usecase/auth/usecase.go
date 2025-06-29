package auth

import (
	"context"
	"errors"

	"go-clean/internal/domain/auth"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInternalServer     = errors.New("internal server error")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type UseCase interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.AuthResponse, error)
	RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.AuthResponse, error)
}