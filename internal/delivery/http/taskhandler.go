package http

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

func NewProjectHandler(service usecase.TaskService) *TaskHandler {
	return &TaskHandler{service: &service}
}

type CreateTaskReq struct {
	Title       string
	Description string
	Performer   string
	Deadline    time.Time
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
	var req struct {
		taskID string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := th.service.DeleteTaskByID(req.taskID)
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

func (th *TaskHandler) FindTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		taskID string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := th.service.FindTask(req.taskID)
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
		TaskID  string
		Sender  string
		Message string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	com := domain.Comment{
		Sender:  req.Sender,
		Message: req.Message,
	}

	task, err := th.service.AddCommentToTask(req.TaskID, &com)
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
		TaskID string
		Status string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := th.service.UpdateStatus(req.TaskID, req.Status); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
