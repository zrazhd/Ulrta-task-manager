package domain

import (
	"context"
	"errors"
	"time"
)

type Comment struct {
	CommentID string
	TaskID    string
	CreatorID string
	Message   string
	CreatedAt time.Time
}

func (c *Comment) ValidateComment() error {
	if c.CommentID == "" {
		return errors.New("There is no Comment ID")
	}
	if c.TaskID == "" {
		return errors.New("There is no Task ID")
	}
	if c.CreatorID == "" {
		return errors.New("There is no Creator ID")
	}
	if c.Message == "" {
		return errors.New("Comment body is empty")
	}

	return nil
}

type CommentRepo interface {
	CreateComment(ctx context.Context, com *Comment) error
	CommentsByTaskID(ctx context.Context, taskID string) ([]Comment, error)
}
