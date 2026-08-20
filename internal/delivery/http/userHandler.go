package httpHandler

import "github.com/zrazhd/Ulrta-task-manager/internal/usecase"

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
	return &UserHandler{service: service}
}
