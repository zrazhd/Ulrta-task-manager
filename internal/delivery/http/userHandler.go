package httpHandler

import (
	"net/http"

	"github.com/zrazhd/Ulrta-task-manager/internal/usecase"
)

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (service *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:password`
	}
}

func (service *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (service *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}
