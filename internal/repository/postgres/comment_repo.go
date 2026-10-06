package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type CommentRepo struct {
	db *pgxpool.Pool
}

func NewCommentRepo(db *pgxpool.Pool) *CommentRepo {
	return &CommentRepo{db: db}
}

func (repo *CommentRepo) CreateComment(ctx context.Context, com *domain.Comment) error {
	sqlStr := `INSERT INTO comments(comment_id, task_id, creator_id, message, created_at) VALUES($1, $2, $3, $4, $5)`

	_, err := repo.db.Exec(ctx, sqlStr, com.CommentID, com.TaskID, com.CreatorID, com.Message, time.Now())

	return err
}
func (repo *CommentRepo) CommentsByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error) {
	sqlStr := `SELECT * FROM comments WHERE task_id = $1`
	var comments []domain.Comment

	rows, err := repo.db.Query(ctx, sqlStr, taskID)
	if err != nil {
		return []domain.Comment{}, fmt.Errorf("cannot get comments from task: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var com domain.Comment
		if err := rows.Scan(&com.CommentID, &com.TaskID, &com.CreatorID, &com.Message); err != nil {
			return []domain.Comment{}, fmt.Errorf("cannot scan data from comments: %w", err)
		}
		comments = append(comments, com)
	}

	if err := rows.Err(); err != nil {
		return []domain.Comment{}, err
	}
	return comments, nil
}
