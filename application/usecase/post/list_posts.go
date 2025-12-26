package post

import (
	"context"

	"github.com/ritarock/bbs-app/application/dto"
	"github.com/ritarock/bbs-app/domain/repository"
)

type ListPostUsecase struct {
	postRepo repository.PostRepository
}

func NewListPostUsecase(postRepo repository.PostRepository) *ListPostUsecase {
	return &ListPostUsecase{
		postRepo: postRepo,
	}
}

func (u *ListPostUsecase) Execute(ctx context.Context) (*dto.ListPostOutput, error) {
	posts, err := u.postRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]dto.PostItem, len(posts))
	for i, post := range posts {
		items[i] = dto.PostItem{
			ID:        post.ID().Int(),
			Title:     post.Title().String(),
			Content:   post.Content().String(),
			CreatedAt: post.CreatedAt(),
		}
	}

	return &dto.ListPostOutput{Posts: items}, nil
}
