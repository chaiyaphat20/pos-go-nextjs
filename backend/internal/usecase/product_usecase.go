package usecase

import (
	"go-clean/internal/domain"
)

type ProductUsecase interface {
	CreateProduct(name, description string, price float64, stock int) (*domain.Product, error)
	GetProduct(id uint) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint) error
	UpdateStock(id uint, stock int) error
}

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(productRepo domain.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: productRepo}
}

func (u *productUsecase) CreateProduct(name, description string, price float64, stock int) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
	err := u.productRepo.Create(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (u *productUsecase) GetProduct(id uint) (*domain.Product, error) {
	return u.productRepo.GetByID(id)
}

func (u *productUsecase) GetAllProducts() ([]*domain.Product, error) {
	return u.productRepo.GetAll()
}

func (u *productUsecase) UpdateProduct(product *domain.Product) error {
	return u.productRepo.Update(product)
}

func (u *productUsecase) DeleteProduct(id uint) error {
	return u.productRepo.Delete(id)
}

func (u *productUsecase) UpdateStock(id uint, stock int) error {
	return u.productRepo.UpdateStock(id, stock)
}