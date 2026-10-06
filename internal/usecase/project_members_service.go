package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/zrazhd/Ulrta-task-manager/internal/domain"
)

type ProjectMemberService struct {
	repo  domain.ProjectMemberRepo
	cache domain.CacheRepo[domain.ProjectMember]
}

func NewProjectMemverService(repo domain.ProjectMemberRepo, cache domain.CacheRepo[domain.ProjectMember]) *ProjectMemberService {
	return &ProjectMemberService{repo: repo, cache: cache}
}

func (pm *ProjectMemberService) AddMember(ctx context.Context, projectID, userID, role string) error {

	member := domain.ProjectMember{
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
	}

	if err := member.Validate(); err != nil {
		return fmt.Errorf("Invalid Project member: %w", err)
	}

	if err := pm.repo.AddMember(ctx, &member); err != nil {
		return fmt.Errorf("cannot save member: %w", err)
	}

	if err := pm.cache.Set(ctx, projectID, &member); err != nil {
		log.Printf("cannot save member in cache: %s", err)
	}

	return nil
}

func (pm *ProjectMemberService) ListMembers(ctx context.Context, projectID string) ([]domain.ProjectMember, error) {
	members, err := pm.repo.ListMembers(ctx, projectID)
	if err != nil {
		return []domain.ProjectMember{}, fmt.Errorf("cannot get members: %w", err)
	}

	return members, nil
}
func (pm *ProjectMemberService) ChangeRole(ctx context.Context, member *domain.ProjectMember) error {
	err := pm.repo.ChangeRole(ctx, member)
	if err != nil {
		return fmt.Errorf("cannot change member's role: %w", err)
	}

	return nil
}
func (pm *ProjectMemberService) GetRole(ctx context.Context, projectID, userID string) (string, error) {
	role, err := pm.repo.GetRole(ctx, projectID, userID)
	if err != nil {
		return "", fmt.Errorf("cannot get member's role: %w", err)
	}
	return role, nil
}
func (pm *ProjectMemberService) GetMember(ctx context.Context, projectID, userID string) (*domain.ProjectMember, error) {
	member, err := pm.repo.GetMember(ctx, projectID, userID)
	if err != nil {
		return nil, fmt.Errorf("cannot get member: %w", err)
	}
	return member, nil
}
