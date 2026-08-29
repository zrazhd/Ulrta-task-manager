package domain

import (
	"errors"
	"time"
)

type Task struct {
	TaskID      string
	Title       string
	Description string
	Performer   string
	Status      string
	Deadline    time.Time
	Comments    []Comment
}

func (t *Task) ValidateTask() error {

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
	CreateTask(*Task) error
	FindTaskByID(taskID string) (*Task, error)
	DeleteTask(taskID string) (*Task, error)
	AddCommentToTask(taskID string, com *Comment) (*Comment, error)
	UpdateStatus(taskID, status string) error
}
