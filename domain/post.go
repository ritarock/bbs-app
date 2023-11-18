package domain

import (
	"context"
	"errors"
	"time"
)

type Post struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PostedAt time.Time `json:"posted_at"`
}

type PostUsecase interface {
	Create(ctx context.Context, post *Post) error
	GetAll(ctx context.Context) ([]Post, error)
	GetById(ctx context.Context, id int) (*Post, error)
	Update(ctx context.Context, post *Post, id int) error
	Delete(ctx context.Context, id int) error
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetAll(ctx context.Context) ([]Post, error)
	GetById(ctx context.Context, id int) (*Post, error)
	Update(ctx context.Context, post *Post, id int) error
	Delete(ctx context.Context, id int) error
}

func (p *Post) Validate() error {
	if len(p.Title) > 30 {
		return errors.New("title length must be less than equal to 30")
	}
	if len(p.Content) > 255 {
		return errors.New("content length must be less than equal to 30")
	}
	if p.PostedAt.Equal(time.Time{}) {
		return errors.New("time must be set")
	}
	return nil
}
