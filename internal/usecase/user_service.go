package usecase

import (
	"context"
	"fmt"
	"log"

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

func (us *UserService) RegisterUser(ctx context.Context, name, email, username, password string) (*domain.User, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Hashing password went wrong: %w", err)
	}

	user := domain.User{
		UserID:   uuid.NewString(),
		Name:     name,
		Email:    email,
		UserName: username,
		Password: string(newPassword),
		Projects: make([]domain.Project, 0),
	}
	if err = user.ValidateUser(); err != nil {
		return nil, fmt.Errorf("Invalid user: %w", err)
	}

	err = us.repo.SaveUser(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("Can't register user: %w", err)
	}

	if err = us.cache.Set(ctx, user.Email, &user); err != nil {
		log.Printf("cannot save user in redis: %s", err)
	}

	return &user, nil
}

func (userService *UserService) LoginUser(ctx context.Context, email, password string) (*domain.User, error) {

	user, err := userService.cache.Get(ctx, email)
	if err == nil && user != nil {
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, fmt.Errorf("Wrong password or email: %w", err)
		}
		log.Print("got user from cache")
		return user, nil

	}

	if err != nil {
		log.Printf("cant get user from cache: %s", err)
	}

	user, err = userService.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user not registered: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("Wrong password or email: %w", err)
	}

	if err = userService.cache.Set(ctx, email, user); err != nil {
		log.Printf("cannot save user in cache: %s", err)
	}
	log.Print("got user from db")

	return user, nil

}

func (userService *UserService) FindUserByUserName(ctx context.Context, UserName string) (*domain.User, error) {
	return userService.repo.FindByUserName(ctx, UserName)
}
