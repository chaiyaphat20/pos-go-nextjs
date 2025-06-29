package user

import (
	"context"

	"go-clean/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) user.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *postgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *postgresRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *postgresRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *postgresRepository) GetAll(ctx context.Context, limit, offset int) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *postgresRepository) GetByRole(ctx context.Context, role user.UserRole, limit, offset int) ([]*user.User, error) {
	var users []*user.User
	err := r.db.WithContext(ctx).Where("role = ?", role).Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *postgresRepository) Update(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Save(u).Error
}

func (r *postgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&user.User{}, "id = ?", id).Error
}

func (r *postgresRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func (r *postgresRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *postgresRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}