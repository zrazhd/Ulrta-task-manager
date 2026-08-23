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

func (repo *TaskRepo) CreateTask(task *domain.Task) error {

	comments, err := json.Marshal(task.Comments)
	if err != nil {
		return fmt.Errorf("cant convert comments to json")
	}

	sqlStr := `INSERT INTO tasks(id, title, decription, performer, status, deadline, comments) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = repo.db.Exec(context.Background(), sqlStr, task.TaskID, task.Title, task.Description, task.Performer, task.Status, task.Deadline, comments)
	return err

}
func (repo *TaskRepo) FindTaskByID(taskID string) (*domain.Task, error) {
	var rawComments []byte
	sqlStr := `SELECT * FROM tasks WHERE id = $1`

	var task domain.Task

	err := repo.db.QueryRow(context.Background(), sqlStr, taskID).Scan(&task.TaskID, &task.Title, &task.Description, &task.Performer, &task.Status, &task.Deadline, &rawComments)
	if err != nil {
		return nil, fmt.Errorf("there is no data: %w", err)
	}

	if err = json.Unmarshal(rawComments, &task.Comments); err != nil {
		return nil, fmt.Errorf("cant unmarshal comments: %w", err)
	}

	return &task, nil
}
func (repo *TaskRepo) DeleteTask(taskID string) (*domain.Task, error) {
	sqlStr := `DELETE FROM tasks WHERE id = $1 RETURNING *`

	var task domain.Task
	var rawComments []byte

	err := repo.db.QueryRow(context.Background(), sqlStr, taskID).Scan(&task.TaskID, &task.Title, &task.Description, &task.Performer, &task.Status, &task.Deadline, &rawComments)
	if err != nil {
		return nil, fmt.Errorf("cannot delete data: %w", err)
	}

	err = json.Unmarshal(rawComments, &task.Comments)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal comments: %w", err)
	}

	return &task, nil
}
func (repo *TaskRepo) AddCommentToTask(taskID string, com *domain.Comment) (*domain.Comment, error) {

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

	sqlStr = `UPDATE tasks SET comments = $1 WHERE id = $2`
	_, err = repo.db.Exec(context.Background(), sqlStr, newComment, taskID)
	if err != nil {
		return nil, fmt.Errorf("Cannot set new comments: %w", err)
	}

	return com, nil
}
func (repo *TaskRepo) UpdateStatus(taskID, status string) error {
	sqlStr := `UPDATE tasks SET status = $1 WHERE id = $2`

	_, err := repo.db.Exec(context.Background(), sqlStr, status, taskID)

	return err
}
