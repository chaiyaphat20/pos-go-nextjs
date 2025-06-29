package order

import (
	"context"

	"go-clean/internal/domain/order"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) order.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, o *order.Order) error {
	return r.db.WithContext(ctx).Create(o).Error
}

func (r *postgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	var o order.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Where("id = ?", id).
		First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *postgresRepository) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*order.Order, error) {
	var orders []*order.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&orders).Error
	return orders, err
}

func (r *postgresRepository) GetAll(ctx context.Context, limit, offset int) ([]*order.Order, error) {
	var orders []*order.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Limit(limit).
		Offset(offset).
		Find(&orders).Error
	return orders, err
}

func (r *postgresRepository) GetByStatus(ctx context.Context, status order.OrderStatus, limit, offset int) ([]*order.Order, error) {
	var orders []*order.Order
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Where("status = ?", status).
		Limit(limit).
		Offset(offset).
		Find(&orders).Error
	return orders, err
}

func (r *postgresRepository) Update(ctx context.Context, o *order.Order) error {
	return r.db.WithContext(ctx).Save(o).Error
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status order.OrderStatus) error {
	return r.db.WithContext(ctx).Model(&order.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *postgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&order.Order{}, "id = ?", id).Error
}

func (r *postgresRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&order.Order{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}