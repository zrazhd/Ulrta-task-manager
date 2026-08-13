package usecase

import (
	"errors"
	"fmt"

	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type UserRepo interface {
	FindByID(userID string) (*domain.User, error)
}

type UserService struct {
	ur UserRepo
}

func NewUserService(ur UserRepo) *UserService {
	return &UserService{ur: ur}
}

func (us *UserService) RegisterUser(u *domain.User) (*domain.User, error) {
	err := u.ValidateUser()
	if err != nil {
		return &domain.User{}, fmt.Errorf("can't register user: %w", err)
	}

	_, err = us.ur.FindByID(u.UserID)
	if err == nil {
		return &domain.User{}, errors.New("User already exists")
	}

	return &domain.User{}, nil

}

func LoginUser(email, password string) (*domain.User, error) {
	return &domain.User{}, nil

}
