package repository

import (
	"database/sql"

	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type ProjectRepo struct {
	db *sql.DB
}

func NewProjectRepo(db *sql.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (repo *ProjectRepo) SaveProject(p *domain.Project) error {
	return nil
}
func (repo *ProjectRepo) DeleteProject(projectID string) error {
	return nil
}
func (repo *ProjectRepo) FindProjectByID(projectID string) (*domain.Project, error) {
	return nil, nil
}
func (repo *ProjectRepo) AddTaskToProject(projectID string, task *domain.Task) (*domain.Task, error) {
	return nil, nil
}
func (repo *ProjectRepo) AddParticipantToProject(projectID, userName string) error {
	return nil
}
