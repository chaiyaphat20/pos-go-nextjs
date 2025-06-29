package product

import (
	"context"

	"go-clean/internal/domain/product"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) product.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, p *product.Product) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *postgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*product.Product, error) {
	var p product.Product
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postgresRepository) GetByName(ctx context.Context, name string) (*product.Product, error) {
	var p product.Product
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postgresRepository) GetAll(ctx context.Context, limit, offset int) ([]*product.Product, error) {
	var products []*product.Product
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *postgresRepository) Update(ctx context.Context, p *product.Product) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *postgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&product.Product{}, "id = ?", id).Error
}

func (r *postgresRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&product.Product{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *postgresRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&product.Product{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

func (r *postgresRepository) UpdateStock(ctx context.Context, id uuid.UUID, quantity int) error {
	return r.db.WithContext(ctx).Model(&product.Product{}).Where("id = ?", id).Update("stock", quantity).Error
}

func (r *postgresRepository) SearchByName(ctx context.Context, name string, limit, offset int) ([]*product.Product, error) {
	var products []*product.Product
	err := r.db.WithContext(ctx).Where("name ILIKE ?", "%"+name+"%").Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *postgresRepository) GetLowStock(ctx context.Context, threshold int) ([]*product.Product, error) {
	var products []*product.Product
	err := r.db.WithContext(ctx).Where("stock <= ?", threshold).Find(&products).Error
	return products, err
}