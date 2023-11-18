package domain

import (
	"context"
	"errors"
	"time"
)

type Comment struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	CommentedAt time.Time `json:"commented_at"`
	PostId      int       `json:"post_id"`
}

type CommentUsecase interface {
	Create(ctx context.Context, comment *Comment) error
	GetByPostId(ctx context.Context, postId int) ([]Comment, error)
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) error
	GetByPostId(ctx context.Context, postId int) ([]Comment, error)
}

func (c *Comment) Validate() error {
	if len(c.Content) > 255 {
		return errors.New("content length must be less than equal to 30")
	}
	if c.CommentedAt.Equal(time.Time{}) {
		return errors.New("time must be set")
	}
	if c.PostId == 0 {
		return errors.New("post id must be set")
	}
	return nil
}
