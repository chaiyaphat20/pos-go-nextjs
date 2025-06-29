package order

import (
	"time"

	"go-clean/internal/domain/product"
	"go-clean/internal/domain/user"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID   `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID     uuid.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	User       *user.User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	TotalPrice float64     `json:"total_price" gorm:"not null"`
	Status     OrderStatus `json:"status" gorm:"type:varchar(20);not null;default:'pending'"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderItem struct {
	ID        uuid.UUID        `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrderID   uuid.UUID        `json:"order_id" gorm:"type:uuid;not null"`
	ProductID uuid.UUID        `json:"product_id" gorm:"type:uuid;not null"`
	Product   *product.Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	Quantity  int              `json:"quantity" gorm:"not null"`
	Price     float64          `json:"price" gorm:"not null"`
	Subtotal  float64          `json:"subtotal" gorm:"not null"`
	CreatedAt time.Time        `json:"created_at" gorm:"autoCreateTime"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusPending, OrderStatusConfirmed, OrderStatusShipped, OrderStatusDelivered, OrderStatusCancelled:
		return true
	default:
		return false
	}
}

func (Order) TableName() string {
	return "orders"
}

func (OrderItem) TableName() string {
	return "order_items"
}

func (o *Order) BeforeCreate() error {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate() error {
	if oi.ID == uuid.Nil {
		oi.ID = uuid.New()
	}
	oi.Subtotal = float64(oi.Quantity) * oi.Price
	return nil
}

func (o *Order) CalculateTotal() {
	total := 0.0
	for _, item := range o.OrderItems {
		total += item.Subtotal
	}
	o.TotalPrice = total
}

func (o *Order) CanBeCancelled() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusConfirmed
}