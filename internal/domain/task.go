package domain

import (
	"errors"
	"time"
)

type Task struct {
	TaksID      string
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
