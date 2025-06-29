package order

import (
	"strconv"

	"go-clean/internal/domain/order"
	orderUseCase "go-clean/internal/usecase/order"
	"go-clean/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	orderUseCase orderUseCase.UseCase
}

func NewHandler(orderUseCase orderUseCase.UseCase) *Handler {
	return &Handler{
		orderUseCase: orderUseCase,
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req orderUseCase.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request format", err.Error())
		return
	}

	// Get user ID from JWT token in middleware
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User ID not found in token")
		return
	}
	
	// Override the user_id from request with the one from JWT token
	req.UserID = userID.(uuid.UUID)

	o, err := h.orderUseCase.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case orderUseCase.ErrUserNotFound:
			response.NotFound(c, "User not found")
		case orderUseCase.ErrProductNotFound:
			response.NotFound(c, "Product not found")
		case orderUseCase.ErrInsufficientStock:
			response.BadRequest(c, "Insufficient stock", err.Error())
		case orderUseCase.ErrInvalidInput:
			response.BadRequest(c, "Invalid input", err.Error())
		default:
			response.InternalServerError(c, "Failed to create order")
		}
		return
	}

	response.Created(c, "Order created successfully", o)
}

func (h *Handler) GetOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid order ID", err.Error())
		return
	}

	o, err := h.orderUseCase.GetOrderByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case orderUseCase.ErrOrderNotFound:
			response.NotFound(c, "Order not found")
		default:
			response.InternalServerError(c, "Failed to get order")
		}
		return
	}

	response.OK(c, "Order retrieved successfully", o)
}

func (h *Handler) GetOrders(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	userIDStr := c.Query("user_id")
	statusStr := c.Query("status")

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

	var orders []*order.Order

	if userIDStr != "" {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			response.BadRequest(c, "Invalid user ID", err.Error())
			return
		}
		orders, err = h.orderUseCase.GetOrdersByUserID(c.Request.Context(), userID, limit, offset)
	} else if statusStr != "" {
		status := order.OrderStatus(statusStr)
		orders, err = h.orderUseCase.GetOrdersByStatus(c.Request.Context(), status, limit, offset)
	} else {
		orders, err = h.orderUseCase.GetOrders(c.Request.Context(), limit, offset)
	}

	if err != nil {
		switch err {
		case orderUseCase.ErrUserNotFound:
			response.NotFound(c, "User not found")
		case orderUseCase.ErrInvalidStatus:
			response.BadRequest(c, "Invalid status", err.Error())
		default:
			response.InternalServerError(c, "Failed to get orders")
		}
		return
	}

	response.OK(c, "Orders retrieved successfully", orders)
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid order ID", err.Error())
		return
	}

	var req struct {
		Status order.OrderStatus `json:"status" validate:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request format", err.Error())
		return
	}

	err = h.orderUseCase.UpdateOrderStatus(c.Request.Context(), id, req.Status)
	if err != nil {
		switch err {
		case orderUseCase.ErrOrderNotFound:
			response.NotFound(c, "Order not found")
		case orderUseCase.ErrInvalidStatus:
			response.BadRequest(c, "Invalid status", err.Error())
		default:
			response.InternalServerError(c, "Failed to update order status")
		}
		return
	}

	response.OK(c, "Order status updated successfully", nil)
}

func (h *Handler) CancelOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid order ID", err.Error())
		return
	}

	err = h.orderUseCase.CancelOrder(c.Request.Context(), id)
	if err != nil {
		switch err {
		case orderUseCase.ErrOrderNotFound:
			response.NotFound(c, "Order not found")
		case orderUseCase.ErrCannotCancel:
			response.BadRequest(c, "Cannot cancel order", err.Error())
		default:
			response.InternalServerError(c, "Failed to cancel order")
		}
		return
	}

	response.OK(c, "Order cancelled successfully", nil)
}