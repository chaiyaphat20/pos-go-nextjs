package admin

import (
	"go-clean/internal/domain/user"
	userUsecase "go-clean/internal/usecase/user"
	"go-clean/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase userUsecase.UseCase
}

func NewHandler(userUsecase userUsecase.UseCase) *Handler {
	return &Handler{
		userUsecase: userUsecase,
	}
}

func (h *Handler) GetSystemStats(c *gin.Context) {
	// Get all users to calculate stats
	users, err := h.userUsecase.GetUsers(c.Request.Context(), 1000, 0)
	if err != nil {
		response.InternalServerError(c, "Failed to get system stats")
		return
	}

	var adminCount, userCount, activeCount int
	for _, u := range users {
		if u.Role == user.UserRoleAdmin {
			adminCount++
		} else {
			userCount++
		}
		if u.IsActive {
			activeCount++
		}
	}

	stats := gin.H{
		"total_users":  len(users),
		"admin_count":  adminCount,
		"user_count":   userCount,
		"active_count": activeCount,
		"system_info": gin.H{
			"version": "1.0.0",
			"status":  "healthy",
		},
	}

	response.OK(c, "System stats retrieved successfully", stats)
}

func (h *Handler) GetAllUsersDetailed(c *gin.Context) {
	users, err := h.userUsecase.GetUsers(c.Request.Context(), 1000, 0)
	if err != nil {
		response.InternalServerError(c, "Failed to get users")
		return
	}

	// Return detailed user information (admin only)
	result := gin.H{
		"users": users,
		"count": len(users),
	}

	response.OK(c, "Users retrieved successfully", result)
}