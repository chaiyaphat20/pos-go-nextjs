package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null;default:0"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) BeforeCreate() error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (p *Product) IsInStock(quantity int) bool {
	return p.Stock >= quantity
}

func (p *Product) DeductStock(quantity int) error {
	if !p.IsInStock(quantity) {
		return ErrInsufficientStock
	}
	p.Stock -= quantity
	return nil
}

func (p *Product) AddStock(quantity int) {
	p.Stock += quantity
}