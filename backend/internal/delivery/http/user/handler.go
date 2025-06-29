package user

import (
	"strconv"

	"go-clean/internal/usecase/user"
	"go-clean/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	userUseCase user.UseCase
}

func NewHandler(userUseCase user.UseCase) *Handler {
	return &Handler{
		userUseCase: userUseCase,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req user.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	u, err := h.userUseCase.CreateUser(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case user.ErrUserAlreadyExists:
			response.Conflict(c, "User already exists")
		case user.ErrInvalidInput, user.ErrInvalidUserRole:
			response.BadRequest(c, "Invalid input", err.Error())
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.Created(c, "User created successfully", u)
}

func (h *Handler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	u, err := h.userUseCase.GetUserByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			response.NotFound(c, "User not found")
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.OK(c, "User retrieved successfully", u)
}

func (h *Handler) GetUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.BadRequest(c, "Invalid limit parameter", err.Error())
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		response.BadRequest(c, "Invalid offset parameter", err.Error())
		return
	}

	users, err := h.userUseCase.GetUsers(c.Request.Context(), limit, offset)
	if err != nil {
		response.InternalServerError(c, "Internal server error")
		return
	}

	response.OK(c, "Users retrieved successfully", users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	var req user.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	u, err := h.userUseCase.UpdateUser(c.Request.Context(), id, &req)
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			response.NotFound(c, "User not found")
		case user.ErrInvalidInput, user.ErrInvalidUserRole:
			response.BadRequest(c, "Invalid input", err.Error())
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.OK(c, "User updated successfully", u)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid user ID", err.Error())
		return
	}

	err = h.userUseCase.DeleteUser(c.Request.Context(), id)
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			response.NotFound(c, "User not found")
		default:
			response.InternalServerError(c, "Internal server error")
		}
		return
	}

	response.OK(c, "User deleted successfully", nil)
}