package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type ProjectService struct {
	repo  domain.ProjectRepo
	cache domain.CacheRepo[domain.Project]
}

func NewProjectService(repo domain.ProjectRepo, cache domain.CacheRepo[domain.Project]) *ProjectService {
	return &ProjectService{repo: repo, cache: cache}
}

func (ps *ProjectService) CreateProject(ctx context.Context, title, description, owner string) (*domain.Project, error) {
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

	err := ps.repo.SaveProject(ctx, project)
	if err != nil {
		return &domain.Project{}, fmt.Errorf("Can't save project: %w", err)
	}
	err = ps.cache.Set(ctx, project.ProjectID, project)
	if err != nil {
		log.Printf("cant set project to cache: %s", err)
	}

	return project, nil
}

func (ps *ProjectService) DeleteProject(ctx context.Context, projectID string) error {
	err := ps.cache.Del(context.Background(), projectID)
	if err != nil {
		return fmt.Errorf("can't delete project from cache: %w", err)
	}
	return ps.repo.DeleteProject(ctx, projectID)
}

func (ps *ProjectService) FindByID(ctx context.Context, projectID string) (*domain.Project, error) {
	project, err := ps.cache.Get(context.Background(), projectID)
	if err != nil {
		log.Printf("something wrong with getting project in cache: %s", err)
	}
	if err == nil && project == nil {
		return ps.repo.FindProjectByID(ctx, projectID)
	} else {
		return project, nil
	}

}

func (ps *ProjectService) AddTaskToProject(ctx context.Context, projectID string, task *domain.Task) (*domain.Project, error) {
	project, err := ps.repo.AddTaskToProject(ctx, projectID, task)
	if err != nil {
		return nil, fmt.Errorf("Error adding task to project: %w", err)
	}
	err = ps.cache.Set(context.Background(), projectID, project)
	if err != nil {
		return nil, fmt.Errorf("cant set project in cache after adding task: %w", err)
	}

	return project, nil
}

func (ps *ProjectService) AddPersonToProject(ctx context.Context, projectID, userName string) (*domain.Project, error) {

	project, err := ps.repo.AddParticipantToProject(ctx, projectID, userName)
	if err != nil {
		return nil, fmt.Errorf("Error adding participant to project: %w", err)
	}
	err = ps.cache.Set(context.Background(), projectID, project)
	if err != nil {
		return nil, fmt.Errorf("cant set project in cache after adding task: %s", err)
	}

	return project, nil
}
