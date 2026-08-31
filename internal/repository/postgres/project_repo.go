package repository

import (
	"context"
	"encoding/json"
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

	sqlStr := `INSERT INTO projects (id, title, description, owner, tasks, participants) VALUES($1, $2, $3, $4, $5, $6)`

	tasks, err := json.Marshal(p.Tasks)
	if err != nil {
		return fmt.Errorf("cant marshal task creating project: %w", err)
	}

	// participants, err := json.Marshal(p.Participants)
	// if err != nil {
	// 	return fmt.Errorf("cant marshal participants creating project: %w", err)
	// }

	_, err = repo.db.Exec(ctx, sqlStr, p.ProjectID, p.Title, p.Description, p.Owner, tasks, p.Participants)

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
	var rawTasks []byte
	// var rawParticipants []byte

	err := repo.db.QueryRow(ctx, sqlStr, projectID).Scan(&project.ProjectID, &project.Title, &project.Description, &project.Owner, &rawTasks, &project.Participants)
	if err != nil {
		return nil, fmt.Errorf("cant get project from database: %w", err)
	}

	if err = json.Unmarshal(rawTasks, &project.Tasks); err != nil {
		return nil, fmt.Errorf("cant unmurshal tasks: %w", err)
	}

	// if err = json.Unmarshal(rawParticipants, &project.Participants); err != nil {
	// 	return nil, fmt.Errorf("cant unmurshal participants: %w", err)
	// }

	return &project, nil
}
func (repo *ProjectRepo) AddTaskToProject(ctx context.Context, projectID string, task *domain.Task) (*domain.Project, error) {
	slqStr := "SELECT tasks FROM projects WHERE id = $1"
	var oldRawTasks []byte
	var tasks []domain.Task
	err := repo.db.QueryRow(ctx, slqStr, projectID).Scan(&oldRawTasks)
	if err != nil {
		return nil, fmt.Errorf("can't get tasks from database: %w", err)
	}

	if err = json.Unmarshal(oldRawTasks, &tasks); err != nil {
		return nil, fmt.Errorf("can't unmarshal tasks from projects: %w", err)
	}

	tasks = append(tasks, *task)

	newTasks, err := json.Marshal(tasks)
	if err != nil {
		return nil, fmt.Errorf("can't marhal new tasks in projects: %w", err)
	}

	sqlString := `UPDATE projects SET tasks = $1 WHERE id = $2 RETURNING *`

	var project domain.Project
	var rawTasks []byte
	// var rawParticipants []byte

	err = repo.db.QueryRow(ctx, sqlString, newTasks, projectID).Scan(&project.ProjectID, &project.Title, &project.Description, &project.Owner, &rawTasks, &project.Participants)
	if err != nil {
		return nil, fmt.Errorf("cant get project from database: %w", err)
	}

	if err = json.Unmarshal(rawTasks, &project.Tasks); err != nil {
		return nil, fmt.Errorf("cant unmurshal tasks: %w", err)
	}

	// if err = json.Unmarshal(rawParticipants, &project.Participants); err != nil {
	// 	return nil, fmt.Errorf("cant unmurshal participants: %w", err)
	// }

	return &project, nil
}
func (repo *ProjectRepo) AddParticipantToProject(ctx context.Context, projectID, userName string) (*domain.Project, error) {

	var project domain.Project
	var rawTasks []byte
	// var rawParticipants []byte

	slqStr := `UPDATE projects SET participants = array_append(participants, $1) WHERE id = $2 RETURNING *`

	err := repo.db.QueryRow(ctx, slqStr, userName, projectID).Scan(&project.ProjectID, &project.Title, &project.Description, &project.Owner, &rawTasks, &project.Participants)
	if err != nil {
		return nil, fmt.Errorf("cant get project from database: %w", err)
	}

	if err = json.Unmarshal(rawTasks, &project.Tasks); err != nil {
		return nil, fmt.Errorf("cant unmurshal tasks: %w", err)
	}

	// if err = json.Unmarshal(rawParticipants, &project.Participants); err != nil {
	// 	return nil, fmt.Errorf("cant unmurshal participants: %w", err)
	// }

	return &project, nil
}
