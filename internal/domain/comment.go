package domain

import "errors"

type Comment struct {
	CommentID string
	TaskID    string
	CreatorID string
	Message   string
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
