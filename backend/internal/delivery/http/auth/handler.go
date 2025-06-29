package auth

import (
	"go-clean/internal/domain/auth"
	authUsecase "go-clean/internal/usecase/auth"
	"go-clean/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	authUsecase authUsecase.UseCase
	validator   *validator.Validate
}

func NewHandler(authUsecase authUsecase.UseCase) *Handler {
	return &Handler{
		authUsecase: authUsecase,
		validator:   validator.New(),
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		response.BadRequest(c, "Validation failed", err.Error())
		return
	}

	authResponse, err := h.authUsecase.Login(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case authUsecase.ErrInvalidCredentials:
			response.Unauthorized(c, "Invalid credentials")
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.OK(c, "Login successful", authResponse)
}

func (h *Handler) Register(c *gin.Context) {
	var req auth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		response.BadRequest(c, "Validation failed", err.Error())
		return
	}

	authResponse, err := h.authUsecase.Register(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case authUsecase.ErrUserAlreadyExists:
			response.Conflict(c, "User already exists")
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.Created(c, "Registration successful", authResponse)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req auth.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		response.BadRequest(c, "Validation failed", err.Error())
		return
	}

	authResponse, err := h.authUsecase.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case authUsecase.ErrInvalidRefreshToken:
			response.Unauthorized(c, "Invalid refresh token")
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.OK(c, "Token refreshed successfully", authResponse)
}