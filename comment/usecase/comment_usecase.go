package usecase

import (
	"context"
	"ritarock/bbs-app/domain"
	"time"
)

type commentUsecase struct {
	commentRepo    domain.CommentRepository
	contextTimeout time.Duration
}

func NewCommentUsecase(c domain.CommentRepository, timeout time.Duration) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo:    c,
		contextTimeout: timeout,
	}
}

func (c *commentUsecase) Create(ctx context.Context, comment *domain.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepo.Create(ctx, comment)
}

func (c *commentUsecase) GetAllByPost(ctx context.Context, postId int) ([]domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	res, err := c.commentRepo.GetAllByPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
