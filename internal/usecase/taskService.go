package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type TaskRepo interface {
	CreateTask(*domain.Task) error
	FindTaskByID(taskID string) (*domain.Task, error)
	DeleteTask(taskID string) (*domain.Task, error)
	AddCommentToTask(taskID string, com *domain.Comment) (*domain.Comment, error)
	UpdateStatus(taskID, status string) error
}

type TaskService struct {
	repo TaskRepo
}

func NewTaskService(repo TaskRepo) *TaskService {
	return &TaskService{repo: repo}
}

func (ts *TaskService) CreateTask(title, description, performer string, deadline time.Time) (*domain.Task, error) {
	task := &domain.Task{
		TaskID:      uuid.NewString(),
		Title:       title,
		Description: description,
		Performer:   performer,
		Status:      "to do",
		Deadline:    deadline,
	}

	err := task.ValidateTask()
	if err != nil {
		return nil, fmt.Errorf("Wrong data: %w", err)
	}

	ts.repo.CreateTask(task)

	return task, nil
}

func (ts *TaskService) DeleteTaskByID(taskID string) (*domain.Task, error) {
	return ts.repo.DeleteTask(taskID)
}

func (ts *TaskService) FindTask(taskID string) (*domain.Task, error) {
	return ts.repo.FindTaskByID(taskID)
}

func (ts *TaskService) AddCommentToTask(taskID string, com *domain.Comment) (*domain.Comment, error) {
	return ts.repo.AddCommentToTask(taskID, com)
}

func (ts *TaskService) UpdateStatus(taskID, status string) error {
	return ts.repo.UpdateStatus(taskID, status)
}
