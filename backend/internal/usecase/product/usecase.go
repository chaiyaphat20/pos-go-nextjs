package product

import (
	"context"
	"errors"

	"go-clean/internal/domain/product"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrProductExists      = errors.New("product already exists")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInsufficientStock  = errors.New("insufficient stock")
	ErrDatabaseError      = errors.New("database error")
	ErrInternalServer     = errors.New("internal server error")
)

type UseCase interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest) (*product.Product, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (*product.Product, error)
	GetProducts(ctx context.Context, limit, offset int) ([]*product.Product, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, req *UpdateProductRequest) (*product.Product, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	UpdateStock(ctx context.Context, id uuid.UUID, quantity int) error
	SearchProducts(ctx context.Context, name string, limit, offset int) ([]*product.Product, error)
	GetLowStockProducts(ctx context.Context, threshold int) ([]*product.Product, error)
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Description string  `json:"description" validate:"max=1000"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"gte=0"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty" validate:"omitempty,min=1,max=255"`
	Description *string  `json:"description,omitempty" validate:"omitempty,max=1000"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Stock       *int     `json:"stock,omitempty" validate:"omitempty,gte=0"`
}