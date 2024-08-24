package port

import (
	"context"

	"github.com/ritarock/bbs-app/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	GetAll(ctx context.Context, postID int) ([]*domain.Comment, error)
}

type CommentUsecase interface {
	Create(ctx context.Context, comment *domain.Comment) error
	GetAll(ctx context.Context, postID int) ([]*domain.Comment, error)
}
