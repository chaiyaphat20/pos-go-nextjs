package order

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrOrderExists       = errors.New("order already exists")
	ErrInvalidInput      = errors.New("invalid input")
	ErrDatabaseError     = errors.New("database error")
	ErrInvalidStatus     = errors.New("invalid order status")
	ErrCannotCancel      = errors.New("order cannot be cancelled")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type Repository interface {
	Create(ctx context.Context, order *Order) error
	GetByID(ctx context.Context, id uuid.UUID) (*Order, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*Order, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Order, error)
	GetByStatus(ctx context.Context, status OrderStatus, limit, offset int) ([]*Order, error)
	Update(ctx context.Context, order *Order) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status OrderStatus) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

type OrderItemRepository interface {
	Create(ctx context.Context, orderItem *OrderItem) error
	GetByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderItem, error)
	Update(ctx context.Context, orderItem *OrderItem) error
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByOrderID(ctx context.Context, orderID uuid.UUID) error
}