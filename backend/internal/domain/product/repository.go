package product

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrInsufficientStock  = errors.New("insufficient stock")
	ErrProductExists      = errors.New("product already exists")
	ErrInvalidInput       = errors.New("invalid input")
	ErrDatabaseError      = errors.New("database error")
)

type Repository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
	GetByName(ctx context.Context, name string) (*Product, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	UpdateStock(ctx context.Context, id uuid.UUID, quantity int) error
	SearchByName(ctx context.Context, name string, limit, offset int) ([]*Product, error)
	GetLowStock(ctx context.Context, threshold int) ([]*Product, error)
}