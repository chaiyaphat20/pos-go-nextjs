package database

import (
	"go-clean/internal/domain"

	"gorm.io/gorm"
)

type saleRepository struct {
	db *gorm.DB
}

type saleItemRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) domain.SaleRepository {
	return &saleRepository{db: db}
}

func NewSaleItemRepository(db *gorm.DB) domain.SaleItemRepository {
	return &saleItemRepository{db: db}
}

func (r *saleRepository) Create(sale *domain.Sale) error {
	return r.db.Create(sale).Error
}

func (r *saleRepository) CreateWithTx(tx domain.Transaction, sale *domain.Sale) error {
	gormTx := tx.(*gormTransaction).tx
	return gormTx.Create(sale).Error
}

func (r *saleRepository) BeginTransaction() (domain.Transaction, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return newGormTransaction(tx), nil
}

func (r *saleRepository) GetByID(id uint) (*domain.Sale, error) {
	var sale domain.Sale
	err := r.db.Preload("Customer").Preload("SaleItems.Product").First(&sale, id).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *saleRepository) GetAll() ([]*domain.Sale, error) {
	var sales []*domain.Sale
	err := r.db.Preload("Customer").Preload("SaleItems.Product").Find(&sales).Error
	return sales, err
}

func (r *saleRepository) GetByCustomerID(customerID uint) ([]*domain.Sale, error) {
	var sales []*domain.Sale
	err := r.db.Preload("Customer").Preload("SaleItems.Product").Where("customer_id = ?", customerID).Find(&sales).Error
	return sales, err
}

func (r *saleItemRepository) CreateBatch(items []domain.SaleItem) error {
	return r.db.Create(&items).Error
}

func (r *saleItemRepository) GetBySaleID(saleID uint) ([]*domain.SaleItem, error) {
	var items []*domain.SaleItem
	err := r.db.Preload("Product").Where("sale_id = ?", saleID).Find(&items).Error
	return items, err
}