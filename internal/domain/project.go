package domain

import "errors"

type Project struct {
	ProjectID    string
	Title        string
	Description  string
	Owner        string
	Tasks        []Task
	Participants []User
}

func (p *Project) ValidateProject() error {
	if p.Title == "" {
		return errors.New("There is no title")
	}
	if p.Description == "" {
		return errors.New("There is no description")
	}
	if p.Owner == "" {
		return errors.New("There is no owner")
	}

	return nil
}
