package domain

import (
	"context"
	"time"
)

type Post struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PostedAt time.Time `json:"posted_at"`
}

//go:generate mockery --name PostUsecase
type PostUsecase interface {
	Create(ctx context.Context, post *Post) error
	GetById(ctx context.Context, id int) (Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int) error
}

//go:generate mockery --name PostRepository
type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetById(ctx context.Context, id int) (Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int) error
}
