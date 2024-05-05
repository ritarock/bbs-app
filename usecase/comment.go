package usecase

import (
	"context"
	"time"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type commentUsecase struct {
	commentRepo    port.CommentRepository
	contextTimeout time.Duration
}

func NewCommentUsecase(repo port.CommentRepository, timeout time.Duration) port.CommentUsecase {
	return &commentUsecase{
		commentRepo:    repo,
		contextTimeout: timeout,
	}
}

func (c *commentUsecase) Create(ctx context.Context, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepo.Create(ctx, comment)
}

func (c *commentUsecase) GetAll(ctx context.Context) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepo.GetAll(ctx)
}

func (c *commentUsecase) GetByPostID(ctx context.Context, postID int) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepo.GetByPostID(ctx, postID)
}
