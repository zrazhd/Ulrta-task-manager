package domain

import (
	"errors"
)

type Project struct {
	ProjectID    string
	Title        string
	Description  string
	Owner        string
	Tasks        []Task
	Participants []string
}

func (p *Project) ValidateProject() error {
	if p.Title == "" {
		return errors.New("There is no title")
	}
	if p.Description == "" {
		return errors.New("There is no description")
	}
	if p.Owner == "" {
		return errors.New("There is no owner")
	}

	return nil
}

type ProjectRepo interface {
	SaveProject(p *Project) error
	DeleteProject(projectID string) (*Project, error)
	FindProjectByID(projectID string) (*Project, error)
	AddTaskToProject(projectID string, task *Task) (*Task, error)
	AddParticipantToProject(projectID, userName string) error
}
