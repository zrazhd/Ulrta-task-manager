package domain

import (
	"context"
	"errors"
)

type Project struct {
	ProjectID   string
	Title       string
	Description string
}

func (p *Project) ValidateProject() error {
	if p.ProjectID == "" {
		return errors.New("There is no Project ID")
	}
	if p.Title == "" {
		return errors.New("There is no title")
	}
	if p.Description == "" {
		return errors.New("There is no description")
	}

	return nil
}

type ProjectRepo interface {
	SaveProject(ctx context.Context, p *Project) error
	DeleteProject(ctx context.Context, projectID string) error
	FindProjectByID(ctx context.Context, projectID string) (*Project, error)
	AddTaskToProject(ctx context.Context, projectID string, task *Task) (*Project, error)
	AddParticipantToProject(ctx context.Context, projectID, userName string) (*Project, error)
}
