package usecase

import (
	"context"
	"ritarock/bbs-app/domain"
	"time"
)

type postUsecase struct {
	postRepo       domain.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(p domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &postUsecase{
		postRepo:       p,
		contextTimeout: timeout,
	}
}

func (p *postUsecase) Create(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepo.Create(ctx, post)
}

func (p *postUsecase) GetById(ctx context.Context, id int) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	res, err := p.postRepo.GetById(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	return res, nil
}

func (p *postUsecase) GetAll(ctx context.Context) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	res, err := p.postRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *postUsecase) Update(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepo.Update(ctx, post)
}

func (p *postUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	existedPost, err := p.postRepo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if existedPost == (domain.Post{}) {
		return domain.ErrNotFound
	}

	return p.postRepo.Delete(ctx, id)
}
