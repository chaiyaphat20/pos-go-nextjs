package domain

import (
	"time"
)

type Sale struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Total       float64    `json:"total" gorm:"not null"`
	CustomerID  *uint      `json:"customer_id" gorm:"foreignKey:CustomerID;references:ID"`
	Customer    *User      `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	SaleItems   []SaleItem `json:"sale_items" gorm:"foreignKey:SaleID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type SaleItem struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	SaleID    uint     `json:"sale_id" gorm:"not null"`
	ProductID uint     `json:"product_id" gorm:"not null"`
	Product   *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Quantity  int      `json:"quantity" gorm:"not null"`
	Price     float64  `json:"price" gorm:"not null"`
	Subtotal  float64  `json:"subtotal" gorm:"not null"`
}

type Transaction interface {
	Commit() error
	Rollback() error
}

type SaleRepository interface {
	Create(sale *Sale) error
	CreateWithTx(tx Transaction, sale *Sale) error
	GetByID(id uint) (*Sale, error)
	GetAll() ([]*Sale, error)
	GetByCustomerID(customerID uint) ([]*Sale, error)
	BeginTransaction() (Transaction, error)
}

type SaleItemRepository interface {
	CreateBatch(items []SaleItem) error
	GetBySaleID(saleID uint) ([]*SaleItem, error)
}