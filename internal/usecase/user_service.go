package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo  domain.UserRepo
	cache domain.CacheRepo[domain.User]
}

func NewUserService(repo domain.UserRepo, cache domain.CacheRepo[domain.User]) *UserService {
	return &UserService{repo: repo, cache: cache}
}

func (us *UserService) RegisterUser(ctx context.Context, name, email, password string) (*domain.User, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Hashing password went wrong: %w", err)
	}

	user := domain.User{
		UserID:   uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: string(newPassword),
		Projects: make([]domain.Project, 0),
	}
	if err = user.ValidateUser(); err != nil {
		return nil, fmt.Errorf("Invalid user: %w", err)
	}

	err = us.repo.SaveUser(&user)
	if err != nil {
		return nil, fmt.Errorf("Can't register user: %w", err)
	}
	return &user, nil

}

func (userService *UserService) LoginUser(email, password string) (*domain.User, error) {
	user, err := userService.repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not registered: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Wrong password or email: %w", err)
	}

	return user, nil

}

func (userService *UserService) FindUserByUserName(UserName string) (*domain.User, error) {
	return userService.repo.FindByUserName(UserName)
}
