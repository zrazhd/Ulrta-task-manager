package domain

import (
	"context"
	"errors"
	"time"
)

type Task struct {
	TaskID      string
	ProjectID   string
	CreatorID   string
	Title       string
	Description string
	Performer   string
	Status      string
	Deadline    time.Time
}

func (t *Task) ValidateTask() error {
	if t.TaskID == "" {
		return errors.New("There is no Task ID")
	}
	if t.ProjectID == "" {
		return errors.New("There is no Project ID")
	}
	if t.CreatorID == "" {
		return errors.New("There is no Creator ID")
	}

	if t.Title == "" {
		return errors.New("There is no title")
	}
	if t.Description == "" {
		return errors.New("There is no description")
	}
	if t.Performer == "" {
		return errors.New("There is no performer")
	}
	if t.Status != "to do" && t.Status != "in progress" && t.Status != "done" {
		return errors.New("Invalid status")
	}
	if t.Deadline.Before(time.Now()) {
		return errors.New("Deadline must be in future")
	}

	return nil
}

type TaskRepo interface {
	CreateTask(ctx context.Context, t *Task) error
	FindTaskByID(ctx context.Context, taskID string) (*Task, error)
	DeleteTask(ctx context.Context, taskID string) error
	AddCommentToTask(ctx context.Context, taskID string, com *Comment) (*Task, error)
	UpdateStatus(ctx context.Context, taskID, status string) (*Task, error)
}
