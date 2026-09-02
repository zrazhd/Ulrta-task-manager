package domain

import "errors"

type ProjectMember struct {
	ProjectID string
	UserID    string
	Role      string
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
