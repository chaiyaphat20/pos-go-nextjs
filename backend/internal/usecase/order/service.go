package order

import (
	"context"

	"go-clean/internal/domain/order"
	"go-clean/internal/domain/product"
	"go-clean/internal/domain/user"

	"github.com/google/uuid"
)

type service struct {
	orderRepo   order.Repository
	productRepo product.Repository
	userRepo    user.Repository
}

func NewService(orderRepo order.Repository, productRepo product.Repository, userRepo user.Repository) UseCase {
	return &service{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (s *service) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*order.Order, error) {
	exists, err := s.userRepo.Exists(ctx, req.UserID)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if !exists {
		return nil, ErrUserNotFound
	}

	o := &order.Order{
		ID:         uuid.New(),
		UserID:     req.UserID,
		Status:     order.OrderStatusPending,
		OrderItems: make([]order.OrderItem, 0, len(req.OrderItems)),
	}

	totalPrice := 0.0

	for _, itemReq := range req.OrderItems {
		p, err := s.productRepo.GetByID(ctx, itemReq.ProductID)
		if err != nil {
			return nil, ErrProductNotFound
		}

		if !p.IsInStock(itemReq.Quantity) {
			return nil, ErrInsufficientStock
		}

		orderItem := order.OrderItem{
			ID:        uuid.New(),
			OrderID:   o.ID,
			ProductID: itemReq.ProductID,
			Quantity:  itemReq.Quantity,
			Price:     p.Price,
			Subtotal:  float64(itemReq.Quantity) * p.Price,
		}

		o.OrderItems = append(o.OrderItems, orderItem)
		totalPrice += orderItem.Subtotal

		if err := p.DeductStock(itemReq.Quantity); err != nil {
			return nil, ErrInsufficientStock
		}

		if err := s.productRepo.Update(ctx, p); err != nil {
			return nil, ErrDatabaseError
		}
	}

	o.TotalPrice = totalPrice

	if err := s.orderRepo.Create(ctx, o); err != nil {
		return nil, ErrDatabaseError
	}

	return o, nil
}

func (s *service) GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	o, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrOrderNotFound
	}

	return o, nil
}

func (s *service) GetOrders(ctx context.Context, limit, offset int) ([]*order.Order, error) {
	orders, err := s.orderRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	return orders, nil
}

func (s *service) GetOrdersByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*order.Order, error) {
	exists, err := s.userRepo.Exists(ctx, userID)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if !exists {
		return nil, ErrUserNotFound
	}

	orders, err := s.orderRepo.GetByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	return orders, nil
}

func (s *service) GetOrdersByStatus(ctx context.Context, status order.OrderStatus, limit, offset int) ([]*order.Order, error) {
	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}

	orders, err := s.orderRepo.GetByStatus(ctx, status, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	return orders, nil
}

func (s *service) UpdateOrderStatus(ctx context.Context, id uuid.UUID, status order.OrderStatus) error {
	if !status.IsValid() {
		return ErrInvalidStatus
	}

	exists, err := s.orderRepo.Exists(ctx, id)
	if err != nil {
		return ErrDatabaseError
	}
	if !exists {
		return ErrOrderNotFound
	}

	if err := s.orderRepo.UpdateStatus(ctx, id, status); err != nil {
		return ErrDatabaseError
	}

	return nil
}

func (s *service) CancelOrder(ctx context.Context, id uuid.UUID) error {
	o, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return ErrOrderNotFound
	}

	if !o.CanBeCancelled() {
		return ErrCannotCancel
	}

	for _, item := range o.OrderItems {
		p, err := s.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return ErrProductNotFound
		}

		p.AddStock(item.Quantity)

		if err := s.productRepo.Update(ctx, p); err != nil {
			return ErrDatabaseError
		}
	}

	if err := s.orderRepo.UpdateStatus(ctx, id, order.OrderStatusCancelled); err != nil {
		return ErrDatabaseError
	}

	return nil
}