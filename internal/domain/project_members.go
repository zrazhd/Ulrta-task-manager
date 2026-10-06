package domain

import (
	"context"
	"errors"
	"time"
)

type ProjectMember struct {
	ProjectID string
	UserID    string
	Role      string
	CreatedAt time.Time
}

func (pm *ProjectMember) Validate() error {
	if pm.ProjectID == "" {
		return errors.New("There is no Project ID")
	}
	if pm.UserID == "" {
		return errors.New("There is no User ID")
	}
	if pm.Role == "" || (pm.Role != "owner" && pm.Role != "member" && pm.Role != "viewer") {
		return errors.New("Wrong Role")
	}
	return nil
}

type ProjectMemberRepo interface {
	AddMember(ctx context.Context, pm *ProjectMember) error
	ListMembers(ctx context.Context, projectID string) ([]ProjectMember, error)
	ChangeRole(ctx context.Context, pm *ProjectMember) error
	GetRole(ctx context.Context, projectID, userID string) (string, error)
	GetMember(ctx context.Context, projectID, userID string) (*ProjectMember, error)
}
