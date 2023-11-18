package usecase

import (
	"context"
	"time"

	"github.com/ritarock/bbs-app/domain"
)

type commentUsecase struct {
	commentRepo    domain.CommentRepository
	contextTimeout time.Duration
}

func NewCommentUsecase(repo domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
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

func (c *commentUsecase) GetByPostId(ctx context.Context, postId int) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepo.GetByPostId(ctx, postId)
}
