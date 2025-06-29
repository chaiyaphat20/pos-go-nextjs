package auth

import (
	"context"

	"go-clean/internal/domain/auth"
	"go-clean/internal/domain/user"
	"go-clean/pkg/jwt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo   user.Repository
	jwtManager *jwt.JWTManager
}

func NewService(userRepo user.Repository, jwtManager *jwt.JWTManager) UseCase {
	return &service{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *service) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	u, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	if !u.IsActive {
		return nil, ErrInvalidCredentials
	}

	accessToken, refreshToken, err := s.jwtManager.GenerateTokens(u)
	if err != nil {
		return nil, ErrInternalServer
	}

	u.Password = ""
	return &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         u,
	}, nil
}

func (s *service) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.AuthResponse, error) {
	if req.Role == "" {
		req.Role = user.UserRoleUser
	}

	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInternalServer
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	exists, err = s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, ErrInternalServer
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternalServer
	}

	u := &user.User{
		ID:       uuid.New(),
		Username: &req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
		IsActive: true,
	}

	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, ErrInternalServer
	}

	accessToken, refreshToken, err := s.jwtManager.GenerateTokens(u)
	if err != nil {
		return nil, ErrInternalServer
	}

	u.Password = ""
	return &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         u,
	}, nil
}