package domain

import (
	"errors"
	"time"
)

type Comment struct {
	ID          int       `json:"id"`
	PostID      int       `json:"post_id"`
	Content     string    `json:"content"`
	CommentedAt time.Time `json:"commented_at"`
}

func (c *Comment) Validate() error {
	if c.PostID == 0 {
		return errors.New("post_id is not set")
	}
	if len(c.Content) == 0 || len(c.Content) > 255 {
		return errors.New("length of content must range from 1 to 255 characters")
	}
	if c.CommentedAt.IsZero() {
		return errors.New("time is not set")
	}

	return nil
}
