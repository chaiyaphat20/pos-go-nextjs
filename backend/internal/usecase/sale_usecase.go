package usecase

import (
	"errors"
	"go-clean/internal/domain"
)

type CreateSaleRequest struct {
	CustomerID *uint             `json:"customer_id"`
	Items      []SaleItemRequest `json:"items"`
}

type SaleItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type SaleUsecase interface {
	CreateSale(req CreateSaleRequest) (*domain.Sale, error)
	GetSale(id uint) (*domain.Sale, error)
	GetAllSales() ([]*domain.Sale, error)
	GetSalesByCustomer(customerID uint) ([]*domain.Sale, error)
}

type saleUsecase struct {
	saleRepo     domain.SaleRepository
	saleItemRepo domain.SaleItemRepository
	productRepo  domain.ProductRepository
}

func NewSaleUsecase(saleRepo domain.SaleRepository, saleItemRepo domain.SaleItemRepository, productRepo domain.ProductRepository) SaleUsecase {
	return &saleUsecase{
		saleRepo:     saleRepo,
		saleItemRepo: saleItemRepo,
		productRepo:  productRepo,
	}
}

func (u *saleUsecase) CreateSale(req CreateSaleRequest) (*domain.Sale, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("sale items cannot be empty")
	}

	// Validate products and calculate total first
	var total float64
	var saleItems []domain.SaleItem
	var stockUpdates []struct {
		ProductID uint
		NewStock  int
	}

	for _, item := range req.Items {
		product, err := u.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		subtotal := product.Price * float64(item.Quantity)
		total += subtotal

		saleItems = append(saleItems, domain.SaleItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			Subtotal:  subtotal,
		})

		stockUpdates = append(stockUpdates, struct {
			ProductID uint
			NewStock  int
		}{
			ProductID: item.ProductID,
			NewStock:  product.Stock - item.Quantity,
		})
	}

	// Start transaction
	tx, err := u.saleRepo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Create sale with items
	sale := &domain.Sale{
		CustomerID: req.CustomerID,
		Total:      total,
		SaleItems:  saleItems,
	}

	err = u.saleRepo.CreateWithTx(tx, sale)
	if err != nil {
		return nil, err
	}

	// Update stock for all products
	for _, update := range stockUpdates {
		err = u.productRepo.UpdateStockWithTx(tx, update.ProductID, update.NewStock)
		if err != nil {
			return nil, err
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return u.saleRepo.GetByID(sale.ID)
}

func (u *saleUsecase) GetSale(id uint) (*domain.Sale, error) {
	return u.saleRepo.GetByID(id)
}

func (u *saleUsecase) GetAllSales() ([]*domain.Sale, error) {
	return u.saleRepo.GetAll()
}

func (u *saleUsecase) GetSalesByCustomer(customerID uint) ([]*domain.Sale, error) {
	return u.saleRepo.GetByCustomerID(customerID)
}