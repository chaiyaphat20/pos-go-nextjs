package user

import (
	"context"
	"errors"

	"go-clean/internal/domain/user"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidInput      = errors.New("invalid input")
	ErrInvalidUserRole   = errors.New("invalid user role")
	ErrDatabaseError     = errors.New("database error")
	ErrInternalServer    = errors.New("internal server error")
)

type UseCase interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*user.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUsers(ctx context.Context, limit, offset int) ([]*user.User, error)
	GetUsersByRole(ctx context.Context, role user.UserRole, limit, offset int) ([]*user.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*user.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type CreateUserRequest struct {
	Username string        `json:"username" validate:"required,min=3,max=50"`
	Email    string        `json:"email" validate:"required,email"`
	Password string        `json:"password" validate:"required,min=6"`
	Role     user.UserRole `json:"role"`
}

type UpdateUserRequest struct {
	Username *string        `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Email    *string        `json:"email,omitempty" validate:"omitempty,email"`
	Role     *user.UserRole `json:"role,omitempty"`
	IsActive *bool          `json:"is_active,omitempty"`
}
