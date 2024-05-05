package usecase

import (
	"context"
	"time"

	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type postUsecase struct {
	postRepo       port.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(repo port.PostRepository, timeout time.Duration) port.PostUsecase {
	return &postUsecase{
		postRepo:       repo,
		contextTimeout: timeout,
	}
}

func (p *postUsecase) Create(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepo.Create(ctx, post)
}

func (p *postUsecase) GetAll(ctx context.Context) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepo.GetAll(ctx)
}

func (p *postUsecase) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepo.GetByID(ctx, id)
}

func (p *postUsecase) Update(ctx context.Context, postID int, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	_, err := p.postRepo.GetByID(ctx, postID)
	if err != nil {
		return domain.ErrNotFound
	}

	return p.postRepo.Update(ctx, postID, post)
}

func (p *postUsecase) Delete(ctx context.Context, postID int) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	_, err := p.postRepo.GetByID(ctx, postID)
	if err != nil {
		return domain.ErrNotFound
	}

	return p.postRepo.Delete(ctx, postID)
}
