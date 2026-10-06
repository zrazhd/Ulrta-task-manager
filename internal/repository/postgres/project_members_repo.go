package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type ProjectMembersRepo struct {
	db *pgxpool.Pool
}

func NewProjectMembersRepo(db *pgxpool.Pool) *ProjectMembersRepo {
	return &ProjectMembersRepo{db: db}
}

func (repo *ProjectMembersRepo) AddMember(ctx context.Context, pm *domain.ProjectMember) error {
	sqlStr := `INSERT INTO project_members(project_id, user_id, role, joined_at)`

	_, err := repo.db.Exec(ctx, sqlStr, pm.ProjectID, pm.UserID, pm.Role, time.Now())
	return err
}
func (repo *ProjectMembersRepo) ListMembers(ctx context.Context, projectID string) ([]domain.ProjectMember, error) {
	sqlStr := `SELECT * FROM project_members WHERE project_id = $1`

	rows, err := repo.db.Query(ctx, sqlStr, projectID)
	if err != nil {
		return []domain.ProjectMember{}, fmt.Errorf("cannot get members from db: %w", err)
	}
	defer rows.Close()
	var members []domain.ProjectMember
	for rows.Next() {
		var member domain.ProjectMember
		if err := rows.Scan(&member.ProjectID, &member.UserID, &member.Role, &member.CreatedAt); err != nil {
			return []domain.ProjectMember{}, fmt.Errorf("cannot scan project members: %w", err)
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return []domain.ProjectMember{}, fmt.Errorf("error scannig pm: %w", err)
	}
	return members, nil
}
func (repo *ProjectMembersRepo) ChangeRole(ctx context.Context, pm *domain.ProjectMember) error {
	sqlStr := `UPDATE project_members SET role = $1 WHERE project_id = $2 AND user_id = $3`

	_, err := repo.db.Exec(ctx, sqlStr, pm.Role, pm.ProjectID, pm.UserID)
	return err
}
func (repo *ProjectMembersRepo) GetRole(ctx context.Context, projectID, userID string) (string, error) {
	sqlStr := `SELECT role FROM project_members WHERE project_id = $1 AND user_id = $2`

	var role string
	err := repo.db.QueryRow(ctx, sqlStr, projectID, userID).Scan(&role)

	if err != nil {
		return "", fmt.Errorf("cannot get member's role: %w", err)
	}
	return role, nil
}
func (repo *ProjectMembersRepo) GetMember(ctx context.Context, projectID, userID string) (*domain.ProjectMember, error) {
	sqlStr := `SELECT * FROM project_members WHERE project_id = $1 AND user_id = $2`

	var member domain.ProjectMember
	err := 
	return nil, nil
}
