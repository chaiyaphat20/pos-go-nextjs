package domain

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null;default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id uint) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id uint) error
	UpdateStock(id uint, stock int) error
	UpdateStockWithTx(tx Transaction, id uint, stock int) error
}