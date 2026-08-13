package domain

import "errors"

type Comment struct {
	Sender  string
	Message string
}

func (c *Comment) ValidateComment() error {
	if c.Message == "" {
		return errors.New("Comment body is empty")
	}

	return nil
}
