package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	CommentedAt time.Time `json:"commented_at"`
	PostID      int       `json:"post_id"`
}

//go:generate mockery --name CommentUsecase
type CommentUsecase interface {
	Create(ctx context.Context, comment *Comment) error
	GetAllByPost(ctx context.Context, postId int) ([]Comment, error)
}

//go:generate mockery --name CommentRepository
type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) error
	GetAllByPost(ctx context.Context, postId int) ([]Comment, error)
}
