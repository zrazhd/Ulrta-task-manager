package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type ProjectRepo interface {
	SaveProject(p *domain.Project) error
	DeleteProject(projectID string) (*domain.Project, error)
	FindProjectByID(projectID string) (*domain.Project, error)
	AddTaskToProject(projectID string, task *domain.Task) (*domain.Task, error)
	AddParticipantToProject(projectID, userName string) error
}

type ProjectService struct {
	repo ProjectRepo
}

func NewProjectService(repo ProjectRepo) *ProjectService {
	return &ProjectService{repo: repo}
}

func (ps *ProjectService) CreateProject(title, description, owner string) (*domain.Project, error) {
	project := &domain.Project{
		ProjectID:    uuid.NewString(),
		Title:        title,
		Description:  description,
		Owner:        owner,
		Tasks:        make([]domain.Task, 0),
		Participants: make([]string, 0),
	}

	if err := project.ValidateProject(); err != nil {
		return nil, fmt.Errorf("Invalid project: %w", err)
	}

	err := ps.repo.SaveProject(project)
	if err != nil {
		return &domain.Project{}, fmt.Errorf("Can't save project: %w", err)
	}

	return project, nil
}

func (ps *ProjectService) DeleteProject(projectID string) (*domain.Project, error) {
	return ps.repo.DeleteProject(projectID)
}

func (ps *ProjectService) FindByID(projectID string) (*domain.Project, error) {
	return ps.repo.FindProjectByID(projectID)
}

func (ps *ProjectService) AddTaskToProject(projectID string, task *domain.Task) (*domain.Task, error) {
	return ps.repo.AddTaskToProject(projectID, task)
}

func (ps *ProjectService) AddPersonToProject(projectID, userName string) error {
	return ps.repo.AddParticipantToProject(projectID, userName)
}
