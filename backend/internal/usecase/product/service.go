package product

import (
	"context"

	"go-clean/internal/domain/product"

	"github.com/google/uuid"
)

type service struct {
	productRepo product.Repository
}

func NewService(productRepo product.Repository) UseCase {
	return &service{
		productRepo: productRepo,
	}
}

func (s *service) CreateProduct(ctx context.Context, req *CreateProductRequest) (*product.Product, error) {
	exists, err := s.productRepo.ExistsByName(ctx, req.Name)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if exists {
		return nil, ErrProductExists
	}

	p := &product.Product{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := s.productRepo.Create(ctx, p); err != nil {
		return nil, ErrDatabaseError
	}

	return p, nil
}

func (s *service) GetProductByID(ctx context.Context, id uuid.UUID) (*product.Product, error) {
	p, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return p, nil
}

func (s *service) GetProducts(ctx context.Context, limit, offset int) ([]*product.Product, error) {
	products, err := s.productRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	return products, nil
}

func (s *service) UpdateProduct(ctx context.Context, id uuid.UUID, req *UpdateProductRequest) (*product.Product, error) {
	exists, err := s.productRepo.Exists(ctx, id)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if !exists {
		return nil, ErrProductNotFound
	}

	p, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	if req.Name != nil {
		if *req.Name != p.Name {
			nameExists, err := s.productRepo.ExistsByName(ctx, *req.Name)
			if err != nil {
				return nil, ErrDatabaseError
			}
			if nameExists {
				return nil, ErrProductExists
			}
		}
		p.Name = *req.Name
	}
	if req.Description != nil {
		p.Description = *req.Description
	}
	if req.Price != nil {
		p.Price = *req.Price
	}
	if req.Stock != nil {
		p.Stock = *req.Stock
	}

	if err := s.productRepo.Update(ctx, p); err != nil {
		return nil, ErrDatabaseError
	}

	return p, nil
}

func (s *service) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	exists, err := s.productRepo.Exists(ctx, id)
	if err != nil {
		return ErrDatabaseError
	}
	if !exists {
		return ErrProductNotFound
	}

	if err := s.productRepo.Delete(ctx, id); err != nil {
		return ErrDatabaseError
	}

	return nil
}

func (s *service) UpdateStock(ctx context.Context, id uuid.UUID, quantity int) error {
	exists, err := s.productRepo.Exists(ctx, id)
	if err != nil {
		return ErrDatabaseError
	}
	if !exists {
		return ErrProductNotFound
	}

	if err := s.productRepo.UpdateStock(ctx, id, quantity); err != nil {
		return ErrDatabaseError
	}

	return nil
}

func (s *service) SearchProducts(ctx context.Context, name string, limit, offset int) ([]*product.Product, error) {
	products, err := s.productRepo.SearchByName(ctx, name, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	return products, nil
}

func (s *service) GetLowStockProducts(ctx context.Context, threshold int) ([]*product.Product, error) {
	products, err := s.productRepo.GetLowStock(ctx, threshold)
	if err != nil {
		return nil, ErrDatabaseError
	}

	return products, nil
}