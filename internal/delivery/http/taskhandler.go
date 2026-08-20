package http

import (
	"net/http"

	"github.com/zrazhd/Ulrta-task-manager/internal/usecase"
)

type ProjectHandler struct {
	service *usecase.ProjectService
}

func NewProjectHandler(service usecase.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: &service}
}

type CreateTaskODT struct {
}

func (ph *ProjectHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

}
