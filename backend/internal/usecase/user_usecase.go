package usecase

import (
	"go-clean/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(name, email string) (*domain.User, error)
	CreateUserWithPassword(name, email, password string) (*domain.User, error)
	CreateUserWithRole(name, email, password, role string) (*domain.User, error)
	GetUser(id uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uint) error
}

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		Name:  name,
		Email: email,
	}
	err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) CreateUserWithPassword(name, email, password string) (*domain.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user", // Default role เป็น user
	}
	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) CreateUserWithRole(name, email, password, role string) (*domain.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Validate role
	if role != "user" && role != "admin" {
		role = "user" // Default to user if invalid role
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}
	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUser(id uint) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return u.userRepo.GetByEmail(email)
}

func (u *userUsecase) GetAllUsers() ([]*domain.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUsecase) UpdateUser(user *domain.User) error {
	return u.userRepo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}