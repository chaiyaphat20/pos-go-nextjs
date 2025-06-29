package order

import (
	"context"
	"errors"

	"go-clean/internal/domain/order"

	"github.com/google/uuid"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrInvalidInput      = errors.New("invalid input")
	ErrInvalidStatus     = errors.New("invalid order status")
	ErrCannotCancel      = errors.New("order cannot be cancelled")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrDatabaseError     = errors.New("database error")
	ErrInternalServer    = errors.New("internal server error")
	ErrUserNotFound      = errors.New("user not found")
	ErrProductNotFound   = errors.New("product not found")
)

type UseCase interface {
	CreateOrder(ctx context.Context, req *CreateOrderRequest) (*order.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error)
	GetOrders(ctx context.Context, limit, offset int) ([]*order.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*order.Order, error)
	GetOrdersByStatus(ctx context.Context, status order.OrderStatus, limit, offset int) ([]*order.Order, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, status order.OrderStatus) error
	CancelOrder(ctx context.Context, id uuid.UUID) error
}

type CreateOrderRequest struct {
	UserID     uuid.UUID           `json:"user_id" validate:"required"`
	OrderItems []CreateOrderItemRequest `json:"order_items" validate:"required,min=1"`
}

type CreateOrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
}