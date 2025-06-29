package middleware

import (
	"strings"

	"go-clean/pkg/jwt"
	"go-clean/pkg/response"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "user_id"
	UserEmailKey        = "user_email"
	UserRoleKey         = "user_role"
)

func AuthMiddleware(jwtManager *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			response.Unauthorized(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := jwtManager.ValidateToken(token)
		if err != nil {
			switch err {
			case jwt.ErrExpiredToken:
				response.Unauthorized(c, "Token has expired")
			case jwt.ErrInvalidToken:
				response.Unauthorized(c, "Invalid token")
			default:
				response.Unauthorized(c, "Unauthorized")
			}
			c.Abort()
			return
		}

		userID, err := claims.GetUserID()
		if err != nil {
			response.Unauthorized(c, "Invalid user ID in token")
			c.Abort()
			return
		}

		c.Set(UserIDKey, userID)
		c.Set(UserEmailKey, claims.Email)
		c.Set(UserRoleKey, string(claims.Role))

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(UserRoleKey)
		if !exists {
			response.Unauthorized(c, "User role not found")
			c.Abort()
			return
		}

		if role != "admin" {
			response.Forbidden(c, "Admin access required")
			c.Abort()
			return
		}

		c.Next()
	}
}