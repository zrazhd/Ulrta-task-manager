package httpHandler

import (
	"encoding/json"
	"net/http"

	"github.com/zrazhd/Ulrta-task-manager/internal/usecase"
)

type ProjectHandler struct {
	service *usecase.ProjectService
}

func NewProjectHandler(service *usecase.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

func (handler *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := handler.service.CreateProject(r.Context(), req.Title, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (handler *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("id")

	err := handler.service.DeleteProject(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func (handler *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {

	projectID := r.PathValue("id")

	project, err := handler.service.FindByID(r.Context(), projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(project); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
