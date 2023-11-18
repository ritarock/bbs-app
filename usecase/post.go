package usecase

import (
	"context"
	"time"

	"github.com/ritarock/bbs-app/domain"
)

type postUsecase struct {
	postRepo       domain.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(repo domain.PostRepository, timeout time.Duration) domain.PostUsecase {
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

func (p *postUsecase) GetById(ctx context.Context, id int) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepo.GetById(ctx, id)
}

func (p *postUsecase) Update(ctx context.Context, post *domain.Post, id int) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	_, err := p.postRepo.GetById(ctx, id)
	if err != nil {
		return domain.ErrNotFound
	}

	return p.postRepo.Update(ctx, post, id)
}

func (p *postUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	_, err := p.postRepo.GetById(ctx, id)
	if err != nil {
		return domain.ErrNotFound
	}

	return p.postRepo.Delete(ctx, id)
}
