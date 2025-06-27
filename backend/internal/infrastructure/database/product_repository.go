package database

import (
	"go-clean/internal/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}

func (r *productRepository) UpdateStock(id uint, stock int) error {
	return r.db.Model(&domain.Product{}).Where("id = ?", id).Update("stock", stock).Error
}

func (r *productRepository) UpdateStockWithTx(tx domain.Transaction, id uint, stock int) error {
	gormTx := tx.(*gormTransaction).tx
	return gormTx.Model(&domain.Product{}).Where("id = ?", id).Update("stock", stock).Error
}