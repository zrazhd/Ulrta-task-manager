package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type TaskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) *TaskRepo {
	return &TaskRepo{db: db}
}

func (repo *TaskRepo) CreateTask(ctx context.Context, task *domain.Task) error {

	sqlStr := `INSERT INTO tasks(task_id, project_id, creator_id, title, decription, performer, status, deadline, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := repo.db.Exec(ctx, sqlStr, task.TaskID, task.ProjectID, task.CreatorID, task.Title, task.Description, task.Performer, task.Status, task.Deadline, time.Now())
	return err

}
func (repo *TaskRepo) FindTaskByID(ctx context.Context, taskID string) (*domain.Task, error) {
	sqlStr := `SELECT task_id, project_id, creator_id, title, decription, performer, status, deadline FROM tasks WHERE id = $1`

	var task domain.Task

	err := repo.db.QueryRow(ctx, sqlStr, taskID).Scan(&task.TaskID, &task.ProjectID, &task.CreatorID, &task.Title, &task.Description, &task.Performer, &task.Status, &task.Deadline)
	if err != nil {
		return nil, fmt.Errorf("there is no data: %w", err)
	}

	return &task, nil
}
func (repo *TaskRepo) DeleteTask(ctx context.Context, taskID string) error {
	sqlStr := `DELETE FROM tasks WHERE id = $1`

	_, err := repo.db.Exec(context.Background(), sqlStr, taskID)

	return err
}
func (repo *TaskRepo) CreateCommentToTask(ctx context.Context, com *domain.Comment) error {

	sqlStr := `INSERT INTO comments(comment_id, task_id, creator_id, message) VALUES($1, $2, $3, $4) `

	_, err := repo.db.Exec(ctx, sqlStr, com.CommentID, com.TaskID, com.CreatorID, com.Message)

	return err
}
func (repo *TaskRepo) UpdateStatus(ctx context.Context, taskID, status string) (*domain.Task, error) {
	sqlStr := `UPDATE tasks SET status = $1 WHERE id = $2 RETURNING *`

	var task domain.Task
	var rawComments []byte

	err := repo.db.QueryRow(context.Background(), sqlStr, status, taskID).Scan(&task.TaskID, &task.Title, &task.Description, &task.Performer, &task.Status, &task.Deadline, &rawComments)
	if err != nil {
		return nil, fmt.Errorf("cannot update status in task: %w", err)
	}

	if err = json.Unmarshal(rawComments, &task.Comments); err != nil {
		return nil, fmt.Errorf("cannot unmarshal comments in task updating status: %w", err)
	}

	return &task, nil
}
