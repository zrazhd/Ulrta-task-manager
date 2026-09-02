package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type TaskService struct {
	repo  domain.TaskRepo
	cache domain.CacheRepo[domain.Task]
}

func NewTaskService(repo domain.TaskRepo, cache domain.CacheRepo[domain.Task]) *TaskService {
	return &TaskService{repo: repo, cache: cache}
}

func (ts *TaskService) CreateTask(ctx context.Context, projectID, creatorID, title, description, performer string, deadline time.Time) (*domain.Task, error) {
	task := &domain.Task{
		TaskID:      uuid.NewString(),
		ProjectID:   projectID,
		CreatorID:   creatorID,
		Title:       title,
		Description: description,
		Performer:   performer,
		Status:      "Not Started",
		Deadline:    deadline,
	}

	err := task.ValidateTask()
	if err != nil {
		return nil, fmt.Errorf("Wrong data: %w", err)
	}

	ts.repo.CreateTask(ctx, task)

	err = ts.cache.Set(ctx, task.TaskID, task)
	if err != nil {
		log.Printf("cant save task to cache: %s", err)
	}
	return task, nil
}

func (ts *TaskService) DeleteTaskByID(ctx context.Context, taskID string) error {
	err := ts.repo.DeleteTask(ctx, taskID)
	if err != nil {
		return fmt.Errorf("cant delete task from db: %w", err)
	}

	err = ts.cache.Del(ctx, taskID)
	if err != nil {
		return fmt.Errorf("cant delete task from cache: %w", err)
	}

	return nil

}

func (ts *TaskService) FindTask(ctx context.Context, taskID string) (*domain.Task, error) {
	task, err := ts.cache.Get(ctx, taskID)
	if err == nil && task != nil {
		return task, nil
	}
	if err != nil {
		log.Printf("cannot get task from cache: %s", err)
	}
	return ts.repo.FindTaskByID(ctx, taskID)
}
func (ts *TaskService) AddCommentToTask(ctx context.Context, taskID string, com *domain.Comment) (*domain.Task, error) {
	task, err := ts.repo.AddCommentToTask(ctx, taskID, com)
	if err != nil {
		return nil, fmt.Errorf("cannot add comment to task: %w", err)
	}

	if err = ts.cache.Set(ctx, taskID, task); err != nil {
		return nil, fmt.Errorf("cannot add comment to task: %w", err)
	}

	return task, nil
}

func (ts *TaskService) UpdateStatus(ctx context.Context, taskID, status string) (*domain.Task, error) {
	task, err := ts.repo.UpdateStatus(ctx, taskID, status)
	if err != nil {
		return nil, fmt.Errorf("cannot update status in task: %w", err)
	}

	if err = ts.cache.Set(ctx, taskID, task); err != nil {
		return nil, fmt.Errorf("cannot update status in task cache: %w", err)
	}

	return task, nil
}
