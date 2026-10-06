package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type ProjectRepo struct {
	db *pgxpool.Pool
}

func NewProjectRepo(db *pgxpool.Pool) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (repo *ProjectRepo) SaveProject(ctx context.Context, p *domain.Project) error {

	sqlStr := `INSERT INTO projects (project_id, title, description) VALUES($1, $2, $3)`

	_, err := repo.db.Exec(ctx, sqlStr, p.ProjectID, p.Title, p.Description)

	return err
}
func (repo *ProjectRepo) DeleteProject(ctx context.Context, projectID string) error {

	sqlStr := `DELETE FROM projects WHERE id = $1`

	_, err := repo.db.Exec(ctx, sqlStr, projectID)

	return err
}
func (repo *ProjectRepo) FindProjectByID(ctx context.Context, projectID string) (*domain.Project, error) {
	sqlStr := `SELECT * FROM projects WHERE id = $1`

	var project domain.Project

	err := repo.db.QueryRow(ctx, sqlStr, projectID).Scan(&project.ProjectID, &project.Title, &project.Description)
	if err != nil {
		return nil, fmt.Errorf("cant get project from database: %w", err)
	}

	return &project, nil
}
