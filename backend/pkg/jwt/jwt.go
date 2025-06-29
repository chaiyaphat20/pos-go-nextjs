package jwt

import (
	"errors"
	"time"

	"go-clean/internal/domain/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type JWTManager struct {
	secretKey           string
	issuer              string
	accessTokenExpiry   time.Duration
	refreshTokenExpiry  time.Duration
}

func NewJWTManager(secretKey, issuer string, accessTokenExpiry, refreshTokenExpiry time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:           secretKey,
		issuer:              issuer,
		accessTokenExpiry:   accessTokenExpiry,
		refreshTokenExpiry:  refreshTokenExpiry,
	}
}

func (j *JWTManager) GenerateAccessToken(user *user.User) (string, error) {
	claims := &Claims{
		UserID: user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) GenerateRefreshToken(user *user.User) (string, error) {
	claims := &Claims{
		UserID: user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) GenerateTokens(user *user.User) (accessToken, refreshToken string, err error) {
	accessToken, err = j.GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = j.GenerateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (j *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, ErrExpiredToken
	}

	return claims, nil
}

type Claims struct {
	UserID string        `json:"user_id"`
	Email  string        `json:"email"`
	Role   user.UserRole `json:"role"`
	Type   string        `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

func (c *Claims) GetUserID() (uuid.UUID, error) {
	return uuid.Parse(c.UserID)
}