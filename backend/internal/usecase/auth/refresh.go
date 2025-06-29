package auth

import (
	"context"

	"go-clean/internal/domain/auth"
)

func (s *service) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.AuthResponse, error) {
	claims, err := s.jwtManager.ValidateToken(req.RefreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	if claims.Type != "refresh" {
		return nil, ErrInvalidRefreshToken
	}

	userID, err := claims.GetUserID()
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	u, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	if !u.IsActive {
		return nil, ErrInvalidRefreshToken
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