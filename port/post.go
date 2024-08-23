package port

import (
	"context"

	"github.com/ritarock/bbs-app/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetAll(ctx context.Context) ([]*domain.Post, error)
	GetByID(ctx context.Context, id int) (*domain.Post, error)
	Update(ctx context.Context, post *domain.Post, id int) error
	Delete(ctx context.Context, id int) error
}

type PostUsecase interface {
	Create(ctx context.Context, post *domain.Post) error
	GetAll(ctx context.Context) ([]*domain.Post, error)
	GetByID(ctx context.Context, id int) (*domain.Post, error)
	Update(ctx context.Context, post *domain.Post, id int) error
	Delete(ctx context.Context, id int) error
}
