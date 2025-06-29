package user

import (
	"context"

	"go-clean/internal/domain/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo user.Repository
}

func NewService(userRepo user.Repository) UseCase {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*user.User, error) {
	if req.Role == "" {
		req.Role = user.UserRoleUser
	}
	if !req.Role.IsValid() {
		return nil, ErrInvalidUserRole
	}

	exists, err := s.userRepo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	exists, err = s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternalServer
	}

	u := &user.User{
		ID:       uuid.New(),
		Username: &req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
		IsActive: true,
	}

	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, ErrDatabaseError
	}

	u.Password = ""
	return u, nil
}

func (s *service) GetUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	u, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	u.Password = ""
	return u, nil
}

func (s *service) GetUsers(ctx context.Context, limit, offset int) ([]*user.User, error) {
	users, err := s.userRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	for _, u := range users {
		u.Password = ""
	}

	return users, nil
}

func (s *service) GetUsersByRole(ctx context.Context, role user.UserRole, limit, offset int) ([]*user.User, error) {
	if !role.IsValid() {
		return nil, ErrInvalidUserRole
	}

	users, err := s.userRepo.GetByRole(ctx, role, limit, offset)
	if err != nil {
		return nil, ErrDatabaseError
	}

	for _, u := range users {
		u.Password = ""
	}

	return users, nil
}

func (s *service) UpdateUser(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*user.User, error) {
	exists, err := s.userRepo.Exists(ctx, id)
	if err != nil {
		return nil, ErrDatabaseError
	}
	if !exists {
		return nil, ErrUserNotFound
	}

	u, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if req.Username != nil {
		u.Username = req.Username
	}
	if req.Email != nil {
		u.Email = *req.Email
	}
	if req.Role != nil {
		if !req.Role.IsValid() {
			return nil, ErrInvalidUserRole
		}
		u.Role = *req.Role
	}
	if req.IsActive != nil {
		u.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(ctx, u); err != nil {
		return nil, ErrDatabaseError
	}

	u.Password = ""
	return u, nil
}

func (s *service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	exists, err := s.userRepo.Exists(ctx, id)
	if err != nil {
		return ErrDatabaseError
	}
	if !exists {
		return ErrUserNotFound
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return ErrDatabaseError
	}

	return nil
}
