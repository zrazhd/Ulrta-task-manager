package domain

import (
	"context"
	"errors"
)

type User struct {
	UserID   string
	Name     string
	UserName string
	Email    string
	Password string
	Projects []Project
}

func (u *User) ValidateUser() error {
	if u.Name == "" {
		return errors.New("There is no name")
	}
	if u.UserName == "" {
		return errors.New("There is no UserName")
	}
	if u.Email == "" {
		return errors.New("There is no email")
	}
	if u.Password == "" {
		return errors.New("There is no password")
	}
	return nil
}

type UserRepo interface {
	SaveUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserName(ctx context.Context, UserName string) (*User, error)
}
