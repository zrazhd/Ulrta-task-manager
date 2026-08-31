package repository

import (
	"context"
	"encoding/json"
	"fmt"

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

	comments, err := json.Marshal(task.Comments)
	if err != nil {
		return fmt.Errorf("cant convert comments to json")
	}

	sqlStr := `INSERT INTO tasks(id, title, decription, performer, status, deadline, comments) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = repo.db.Exec(ctx, sqlStr, task.TaskID, task.Title, task.Description, task.Performer, task.Status, task.Deadline, comments)
	return err

}
func (repo *TaskRepo) FindTaskByID(ctx context.Context, taskID string) (*domain.Task, error) {
	var rawComments []byte
	sqlStr := `SELECT * FROM tasks WHERE id = $1`

	var task domain.Task

	err := repo.db.QueryRow(ctx, sqlStr, taskID).Scan(&task.TaskID, &task.Title, &task.Description, &task.Performer, &task.Status, &task.Deadline, &rawComments)
	if err != nil {
		return nil, fmt.Errorf("there is no data: %w", err)
	}

	if err = json.Unmarshal(rawComments, &task.Comments); err != nil {
		return nil, fmt.Errorf("cant unmarshal comments: %w", err)
	}

	return &task, nil
}
func (repo *TaskRepo) DeleteTask(ctx context.Context, taskID string) error {
	sqlStr := `DELETE FROM tasks WHERE id = $1`

	_, err := repo.db.Exec(context.Background(), sqlStr, taskID)

	return err
}
func (repo *TaskRepo) AddCommentToTask(ctx context.Context, taskID string, com *domain.Comment) (*domain.Task, error) {

	sqlStr := `SELECT comments FROM tasks WHERE id = $1`

	var rawComments []byte
	var comment []domain.Comment

	err := repo.db.QueryRow(context.Background(), sqlStr, taskID).Scan(&rawComments)
	if err != nil {
		return nil, fmt.Errorf("cannot get data: %w", err)
	}

	err = json.Unmarshal(rawComments, &comment)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal comments: %w", err)
	}

	comment = append(comment, *com)
	newComment, err := json.Marshal(comment)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal comments: %w", err)
	}

	sqlStr = `UPDATE tasks SET comments = $1 WHERE id = $2 RETURNING *`

	var task domain.Task

	err = repo.db.QueryRow(context.Background(), sqlStr, newComment, taskID).Scan(&task.TaskID, &task.Title, &task.Description, &task.Performer, &task.Status, &task.Deadline, &rawComments)
	if err != nil {
		return nil, fmt.Errorf("Cannot set new comments: %w", err)
	}

	if err = json.Unmarshal(rawComments, &task.Comments); err != nil {
		return nil, fmt.Errorf("Cannot unmarshal comments adding it in task: %w", err)
	}

	return &task, nil
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
