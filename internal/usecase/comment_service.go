package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type CommentService struct {
	repo  domain.CommentRepo
	cache domain.CacheRepo[[]domain.Comment]
}

func NewCommentService(repo domain.CommentRepo, cache domain.CacheRepo[[]domain.Comment]) *CommentService {
	return &CommentService{repo: repo, cache: cache}
}

func (cs *CommentService) AddComment(ctx context.Context, taskID, creatorID, message string) error {
	comment := domain.Comment{
		CommentID: uuid.NewString(),
		TaskID:    taskID,
		CreatorID: creatorID,
		Message:   message,
	}

	if err := comment.ValidateComment(); err != nil {
		return fmt.Errorf("Invalid comment: %w", err)
	}

	if err := cs.repo.CreateComment(ctx, &comment); err != nil {
		return fmt.Errorf("cannot save comment in db: %w", err)
	}

	if err := cs.cache.Del(ctx, taskID); err != nil {
		log.Printf("cannot save comment in cache: %s", err)
	}

	return nil
}

func (cs *CommentService) GetCommentsByTaskID(ctx context.Context, taskID string) ([]domain.Comment, error) {
	cachedComments, err := cs.cache.Get(ctx, taskID)
	if err != nil {
		log.Printf("can't get comments from cache: %s", err)
	}
	if err == nil && cachedComments != nil {
		return *cachedComments, nil
	}

	comments, err := cs.repo.CommentsByTaskID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("cannot get comments from task db: %w", err)
	}
	if err = cs.cache.Set(ctx, taskID, &comments); err != nil {
		log.Printf("cannot save comments in cache: %s", err)
	}

	return comments, nil

}
