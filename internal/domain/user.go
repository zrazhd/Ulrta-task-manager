package domain

import (
	"errors"
)

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
	Projects []Project
}

func (u *User) ValidateUser() error {
	if u.Name == "" {
		return errors.New("There is no name")
	}
	if u.Email == "" {
		return errors.New("There is no email")
	}
	if u.Password == "" {
		return errors.New("There is no password")
	}
	return nil
}