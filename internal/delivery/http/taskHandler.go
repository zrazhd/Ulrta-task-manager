package httpHandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
	"github.com/zrazhd/Ulrta-task-manager/internal/usecase"
)

type TaskHandler struct {
	service *usecase.TaskService
}

func NewTaskHandler(service *usecase.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

type CreateTaskReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Performer   string    `json:"performer"`
	Deadline    time.Time `json:"deadline"`
}

func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var odt CreateTaskReq
	if err := json.NewDecoder(r.Body).Decode(&odt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := th.service.CreateTask(odt.Title, odt.Description, odt.Performer, odt.Deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("id")

	task, err := th.service.DeleteTaskByID(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func (th *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.PathValue("id")

	task, err := th.service.FindTask(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.NewEncoder(w).Encode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (th *TaskHandler) AddCommentToTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Sender  string `json:"sender"`
		Message string `json:"message"`
	}
	taskID := r.PathValue("id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	com := domain.Comment{
		Sender:  req.Sender,
		Message: req.Message,
	}

	task, err := th.service.AddCommentToTask(taskID, &com)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (th *TaskHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Status string `json:"status"`
	}
	taskID := r.PathValue("id")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := th.service.UpdateStatus(taskID, req.Status); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
